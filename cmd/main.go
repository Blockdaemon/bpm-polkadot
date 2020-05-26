package main

import (
	"go.blockdaemon.com/bpm/sdk/pkg/docker"
	"go.blockdaemon.com/bpm/sdk/pkg/node"
	"go.blockdaemon.com/bpm/sdk/pkg/plugin"
)

var version string

const (
	polkadotContainerImage = "docker.io/chevdor/polkadot:0.4.4"
	polkadotContainerName  = "polkadot"
	polkadotCmdFile        = "polkadot.dockercmd"

	collectorContainerName = "collector"
	collectorImage         = "docker.io/blockdaemon/polkadot-collector:0.5.0"
	collectorEnvFile       = "configs/collector.env"
)

type PolkadotTester struct{}

func (d PolkadotTester) Test(currentNode node.Node) (bool, error) {
	if err := runAllTests(); err != nil {
		return false, err
	}
	return true, nil
}

func main() {
	templates := map[string]string{
		polkadotCmdFile:  polkadotCmdTpl,
		collectorEnvFile: collectorEnvTpl,
	}

	parameters := []plugin.Parameter{
		{
			Name:        "subtype",
			Type:        plugin.ParameterTypeString,
			Description: "The type of node. Must be either `watcher` or `validator`",
			Mandatory:   false,
			Default:     "watcher",
		},
		{
			Name:        "validator-key",
			Type:        plugin.ParameterTypeString,
			Description: "The key used for a validator (required if subtype = validator)",
			Mandatory:   false,
		},
	}

	containers := []docker.Container{
		{
			Name:    polkadotContainerName,
			Image:   polkadotContainerImage,
			CmdFile: polkadotCmdFile,
			Mounts: []docker.Mount{
				{
					Type: "bind",
					From: "{{ index .Node.StrParameters \"data-dir\" }}",
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
		},
		{
			Name:        collectorContainerName,
			Image:       collectorImage,
			EnvFilename: collectorEnvFile,
			Mounts: []docker.Mount{
				{
					Type: "bind",
					From: "logs",
					To:   "/data/nodestate",
				},
			},
			CollectLogs: true,
		},
	}

	description := "A polkadot package"

	polkadotPlugin := plugin.NewDockerPlugin("polkadot", version, description, parameters, templates, containers)
	polkadotPlugin.ParameterValidator = NewpolkadotParameterValidator(parameters)
	polkadotPlugin.Tester = PolkadotTester{}

	plugin.Initialize(polkadotPlugin)
}
