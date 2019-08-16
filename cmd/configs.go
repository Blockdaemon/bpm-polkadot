package main

import (
	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm-sdk/pkg/template"
)

func createConfigs(currentNode node.Node) error {
    if err := template.ConfigFileRendered(filebeatConfigFile, filebeatConfigTpl, currentNode); err != nil {
    	return err
    }
    return template.ConfigFileRendered(polkadotbeatConfigFile, polkadotbeatConfigTpl, currentNode)
}

