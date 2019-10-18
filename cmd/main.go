package main

import (
	"github.com/Blockdaemon/bpm-sdk/pkg/docker"
	"github.com/Blockdaemon/bpm-sdk/pkg/plugin"
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
		CollectLogs: true,
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
		CollectLogs: true,
	}

	plugin.Initialize(plugin.NewDockerPlugin(
		"polkadot",
		"A polkadot plugin",
		version,
		[]docker.Container{polkadotContainer, polkadotbeatContainer},
		map[string]string{
			polkadotCmdFile:        polkadotCmdTpl,
			polkadotbeatConfigFile: polkadotbeatConfigTpl,
		},
	))
}
