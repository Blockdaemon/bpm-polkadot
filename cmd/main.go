package main

import (
	"github.com/Blockdaemon/bpm-sdk/pkg/plugin"
	"github.com/Blockdaemon/bpm-sdk/pkg/docker"
)

var version string

const (
	polkadotContainerImage = "docker.io/chevdor/polkadot:0.4.4"
	polkadotContainerName  = "polkadot"
	polkadotDataVolumeName = "polkadot-data"
	polkadotCmdFile        = "polkadot.dockercmd"

	polkadotbeatContainerImage = "docker.io/blockdaemon/polkadotbeat:1.0.0"
	polkadotbeatContainerName  = "polkadotbeat"
	polkadotbeatConfigFile     = "polkadotbeat.yml"

	filebeatContainerImage = "docker.elastic.co/beats/filebeat:7.3.1"
	filebeatContainerName  = "filebeat"
	filebeatConfigFile     = "filebeat.yml"

	networkName = "polkadot"

)

func main() {
	polkadotContainer := docker.Container{
		Name:      polkadotContainerName,
		Image:     polkadotContainerImage,
		CmdFile:   polkadotCmdFile,
		NetworkID: networkName,
		Mounts: []docker.Mount{
			{
				Type: "volume",
				From: polkadotDataVolumeName,
				To:   "/data",
			},
		},
		Ports: []docker.Port{
			{
				HostIP:        "0.0.0.0",
				HostPort:      "30333",
				ContainerPort: "30333",
				Protocol:      "tcp",
			},
			{
				HostIP:        "127.0.0.1",
				HostPort:      "9933",
				ContainerPort: "9933",
				Protocol:      "tcp",
			},
		},
	}

	polkadotbeatContainer := docker.Container{
		Name:      polkadotbeatContainerName,
		Image:     polkadotbeatContainerImage,
		Cmd:       []string{"-e", "-strict.perms=false"},
		NetworkID: networkName,
		Mounts: []docker.Mount{
			{
				Type: "bind",
				From: polkadotbeatConfigFile,
				To:   "/usr/share/polkadotbeat/polkadotbeat.yml",
			},
		},
	}

	filebeatContainer := docker.Container{
		Name:      filebeatContainerName,
		Image:     filebeatContainerImage,
		Cmd:       []string{"-e", "-strict.perms=false"},
		NetworkID: networkName,
		Mounts: []docker.Mount{
			{
				Type: "bind",
				From: filebeatConfigFile,
				To:   "/usr/share/filebeat/filebeat.yml",
			},
		},
		User: "root",
	}

	plugin.Initialize(plugin.NewDockerPlugin(
		"polkadot",
		"A polkadot plugin",
		version,
		[]docker.Container{polkadotContainer, polkadotbeatContainer, filebeatContainer},
		map[string]string{
			polkadotCmdFile: polkadotCmdTpl,
			polkadotbeatConfigFile: polkadotbeatConfigTpl,
			filebeatConfigFile: filebeatConfigTpl,
		},
	))
}
