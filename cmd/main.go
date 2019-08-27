package main

import (
	"github.com/Blockdaemon/bpm-sdk/pkg/plugin"
)

var pluginVersion string

const (
	pokadotAlexanderTag    = "0.4.4"
	polkadotKusamaTag      = "0.5.1"
	polkadotContainerImage = "docker.io/chevdor/polkadot"
	polkadotContainerName  = "polkadot"
	polkadotDataVolumeName = "polkadot-data"

	polkadotbeatContainerImage = "docker.io/blockdaemon/polkadotbeat:1.0.0"
	polkadotbeatContainerName  = "polkadotbeat"
	polkadotbeatConfigFile     = "polkadotbeat.yml"

	filebeatContainerImage = "docker.elastic.co/beats/filebeat:7.3.1"
	filebeatContainerName  = "filebeat"
	filebeatConfigFile     = "filebeat.yml"
)

func main() {
	plugin.Initialize(plugin.Plugin{
		Name:          "polkadot",
		Description:   "A polkadot plugin",
		Version:       pluginVersion,
		CreateSecrets: plugin.DefaultCreateSecrets,
		CreateConfigs: createConfigs,
		Start:         start,
		Remove:        plugin.DefaultRemove,
		Upgrade:       plugin.DefaultUpgrade,
	})
}
