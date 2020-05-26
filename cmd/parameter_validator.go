package main

import (
	"fmt"

	"go.blockdaemon.com/bpm/sdk/pkg/node"
	"go.blockdaemon.com/bpm/sdk/pkg/plugin"
)

// polkadotParameterValidator validates parameters
type polkadotParameterValidator struct {
	plugin.SimpleParameterValidator
}

// ValidateParameters uses SimpleParameterValidator but also check is the network is correct
func (t polkadotParameterValidator) ValidateParameters(currentNode node.Node) error {
	// Call SimpleParameterValidator to check if all parameters actually have values
	if err := t.SimpleParameterValidator.ValidateParameters(currentNode); err != nil {
		return err
	}

	network := currentNode.StrParameters["network"]
	if network != "mainnet" && network != "testnet" {
		return fmt.Errorf("unknown network: %q", network)
	}

	subtype := currentNode.StrParameters["subtype"]
	if subtype != "watcher" && network != "validator" {
		return fmt.Errorf("unknown subtype: %q", subtype)
	}

	return nil
}

// NewpolkadotParameterValidator creates a new instance of polkadotParameterValidator
func NewpolkadotParameterValidator(pluginParameters []plugin.Parameter) polkadotParameterValidator {
	return polkadotParameterValidator{
		SimpleParameterValidator: plugin.NewSimpleParameterValidator(pluginParameters),
	}
}
