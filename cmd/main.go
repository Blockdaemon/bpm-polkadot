package main

import (
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

	filebeatContainerImage = "docker.elastic.co/beats/filebeat:7.3.1"
	filebeatContainerName  = "filebeat"
	filebeatConfigFile     = "filebeat.yml"
)

func main() {
	plugin.Initialize(plugin.Plugin{
		Name:          "polkadot",
		Description:   "A polkadot plugin",
		Version:       version,
		CreateSecrets: plugin.DefaultCreateSecrets,
		CreateConfigs: createConfigs,
		Start:         start,
		Status:        status,
		Stop:          plugin.DefaultStop,
		Upgrade:       plugin.DefaultUpgrade,
	})
}
