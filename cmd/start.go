package main

import (
	"github.com/Blockdaemon/bpm-sdk/pkg/docker"
	"github.com/Blockdaemon/bpm-sdk/pkg/node"

	"context"
	"path"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

func start(currentNode node.Node) error {
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
		CmdFile:   path.Join(currentNode.ConfigsDirectory(), polkadotCmdFile),
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
	currentNode.Collection.Key, err = homedir.Expand(currentNode.Collection.Key)
	if err != nil {
		return err
	}
	currentNode.Collection.Cert, err = homedir.Expand(currentNode.Collection.Cert)
	if err != nil {
		return err
	}
	currentNode.Collection.CA, err = homedir.Expand(currentNode.Collection.CA)
	if err != nil {
		return err
	}
	// TODO end

	polkadotbeatContainer := docker.Container{
		Name:      currentNode.ContainerName(polkadotbeatContainerName),
		Image:     polkadotbeatContainerImage,
		Cmd:       []string{"-e", "-strict.perms=false"},
		NetworkID: currentNode.DockerNetworkName(),
		Mounts: []docker.Mount{
			docker.Mount{
				Type: "bind",
				From: path.Join(currentNode.ConfigsDirectory(), polkadotbeatConfigFile),
				To:   "/usr/share/polkadotbeat/polkadotbeat.yml",
			},
			docker.Mount{
				Type: "bind",
				From: currentNode.Collection.CA,
				To:   "/etc/ssl/beats/ca.crt",
			},
			docker.Mount{
				Type: "bind",
				From: currentNode.Collection.Cert,
				To:   "/etc/ssl/beats/beat.crt",
			},
			docker.Mount{
				Type: "bind",
				From: currentNode.Collection.Key,
				To:   "/etc/ssl/beats/beat.key",
			},
		},
	}

	filebeatContainer := docker.Container{
		Name:      currentNode.ContainerName(filebeatContainerName),
		Image:     filebeatContainerImage,
		Cmd:       []string{"-e", "-strict.perms=false"},
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
				From: currentNode.Collection.CA,
				To:   "/etc/ssl/beats/ca.crt",
			},
			docker.Mount{
				Type: "bind",
				From: currentNode.Collection.Cert,
				To:   "/etc/ssl/beats/beat.crt",
			},
			docker.Mount{
				Type: "bind",
				From: currentNode.Collection.Key,
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
