package main

import (
	"github.com/Blockdaemon/bpm-sdk/pkg/docker"
	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"gitlab.com/Blockdaemon/untyped"

	"context"
	"fmt"
	"path"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

func start(currentNode node.Node) error {
	cmd := []string{"polkadot", "-d /data", "--rpc-external"}

	name, err := untyped.GetString(currentNode.Config, "name")
	if err != nil {
		return err
	}
	cmd = append(cmd, "--name")
	cmd = append(cmd, name)

	cmd = append(cmd, "--chain")
	switch currentNode.Environment {
	case "kusama":
		cmd = append(cmd, "kusama")
	case "alexander":
		cmd = append(cmd, "alexander")
	default:
		return fmt.Errorf("Unknown environment: %s", currentNode.Environment)
	}

	switch currentNode.NodeSubtype {
	case "validator":
		cmd = append(cmd, "--validator")
		cmd = append(cmd, "--key")

		key, err := untyped.GetString(currentNode.Config, "key")
		if err != nil {
			return err
		}
		cmd = append(cmd, key)
	case "watcher":
		break
	default:
		return fmt.Errorf("Unknown environment: %s", currentNode.Environment)
	}

	client, err := docker.NewBasicManager()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// First, create the docker network if it doesn't exist yet
	err = client.NetworkExists(ctx, currentNode.DockerNetworkName())
	if err != nil {
		return err
	}

	// Configure the containers
	polkadotContainer := docker.Container{
		Name:      currentNode.ContainerName(polkadotContainerName),
		Image:     polkadotContainerImage,
		Cmd:       cmd,
		NetworkID: currentNode.DockerNetworkName(),
		Mounts: []docker.Mount{
			docker.Mount{
				Type: "volume",
				From: currentNode.VolumeName(polkadotDataVolumeName),
				To:   "/data",
			},
		},
		Ports: []docker.Port{
			docker.Port{
				HostIP:        "0.0.0.0",
				HostPort:      "30333",
				ContainerPort: "30333",
				Protocol:      "tcp",
			},
			docker.Port{
				HostIP:        "127.0.0.1",
				HostPort:      "9933",
				ContainerPort: "9933",
				Protocol:      "tcp",
			},
		},
	}

	// TODO: This is just temporarily until we have a proper authentication system we stick the certs in ~/.blockdaemon
	currentNode.Logstash.Key, err = homedir.Expand(currentNode.Logstash.Key)
	if err != nil {
		return err
	}
	currentNode.Logstash.Certificate, err = homedir.Expand(currentNode.Logstash.Certificate)
	if err != nil {
		return err
	}
	currentNode.Logstash.CertificateAuthorities, err = homedir.Expand(currentNode.Logstash.CertificateAuthorities)
	if err != nil {
		return err
	}
	// TODO end

	polkadotbeatContainer := docker.Container{
		Name:      currentNode.ContainerName(polkadotbeatContainerName),
		Image:     polkadotbeatContainerImage,
		Cmd: 	   []string{"-e", "-strict.perms=false"},
		NetworkID: currentNode.DockerNetworkName(),
		Mounts: []docker.Mount{
			docker.Mount{
				Type: "bind",
				From: path.Join(currentNode.ConfigsDirectory(), polkadotbeatConfigFile),
				To:   "/usr/share/polkadotbeat/polkadotbeat.yml",
			},
			docker.Mount{
				Type: "bind",
				From: currentNode.Logstash.CertificateAuthorities,
				To:   "/etc/ssl/beats/ca.crt",
			},
			docker.Mount{
				Type: "bind",
				From: currentNode.Logstash.Certificate,
				To:   "/etc/ssl/beats/beat.crt",
			},
			docker.Mount{
				Type: "bind",
				From: currentNode.Logstash.Key,
				To:   "/etc/ssl/beats/beat.key",
			},
		},
	}

	filebeatContainer := docker.Container{
		Name:      currentNode.ContainerName(filebeatContainerName),
		Image:     filebeatContainerImage,
		Cmd: 	   []string{"-e", "-strict.perms=false"},
		NetworkID: currentNode.DockerNetworkName(),
		Mounts: []docker.Mount{
			docker.Mount{
				Type: "bind",
				From: path.Join(currentNode.ConfigsDirectory(), filebeatConfigFile),
				To:   "/usr/share/filebeat/filebeat.yml",
			},
			docker.Mount{
				Type: "bind",
				From: "/var/lib/docker/containers",
				To:   "/var/lib/docker/containers",
			},
			docker.Mount{
				Type: "bind",
				From: currentNode.Logstash.CertificateAuthorities,
				To:   "/etc/ssl/beats/ca.crt",
			},
			docker.Mount{
				Type: "bind",
				From: currentNode.Logstash.Certificate,
				To:   "/etc/ssl/beats/beat.crt",
			},
			docker.Mount{
				Type: "bind",
				From: currentNode.Logstash.Key,
				To:   "/etc/ssl/beats/beat.key",
			},
		},
		User: "root",
	}

	// Next, start the containers
	if err := client.ContainerRuns(ctx, polkadotbeatContainer); err != nil {
		return err
	}
	if err := client.ContainerRuns(ctx, filebeatContainer); err != nil {
		return err
	}

	return client.ContainerRuns(ctx, polkadotContainer)
}
