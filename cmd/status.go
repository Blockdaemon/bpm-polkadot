package main

import (
	"github.com/Blockdaemon/bpm-sdk/pkg/docker"
	"github.com/Blockdaemon/bpm-sdk/pkg/node"

	"context"
	"time"
)

func status(currentNode node.Node) (string, error) {
	client, err := docker.NewBasicManager()
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	polkadotContainerRunning, err := client.IsContainerRunning(ctx, currentNode.ContainerName(polkadotContainerName))
	if err != nil {
		return "", err
	}
	polkadotbeatContainerRunning, err := client.IsContainerRunning(ctx, currentNode.ContainerName(polkadotbeatContainerName))
	if err != nil {
		return "", err
	}
	filebeatCotainerRunning, err := client.IsContainerRunning(ctx, currentNode.ContainerName(filebeatContainerName))
	if err != nil {
		return "", err
	}

	if polkadotContainerRunning && polkadotbeatContainerRunning && filebeatCotainerRunning {
		return "running", nil
	} else if polkadotContainerRunning || polkadotbeatContainerRunning || filebeatCotainerRunning {
		return "incomplete", nil
	}

	return "stopped", nil
}
