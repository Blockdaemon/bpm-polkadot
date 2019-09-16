package main

import (
	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm-sdk/pkg/template"
)

func createConfigs(currentNode node.Node) error {
	return template.ConfigFilesRendered(map[string]string{
		filebeatConfigFile:     filebeatConfigTpl,
		polkadotCmdFile:        polkadotCmdTpl,
		polkadotbeatConfigFile: polkadotbeatConfigTpl,
	}, currentNode)
}
