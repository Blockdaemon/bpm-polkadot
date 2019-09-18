package main

import (
	"github.com/Blockdaemon/bpm-sdk/pkg/node"
	"github.com/Blockdaemon/bpm-sdk/pkg/template"
)

const (
	polkadotbeatConfigTpl = `polkadotbeat:
    period: 30s
    polkadot_host: "{{ .ContainerName "polkadot" }}"
    polkadot_port: "9933"
fields:
    info:
        launch_type: bpm
        node_xid: {{ .ID }}
        project: development
        protocol_type: POLKADOT
        network_type: public
        user_id: TODO
        environment: {{ .Environment }}
fields_under_root: true
output:
    logstash:
        hosts:
        - "{{ .Collection.Host }}"
        ssl:
            certificate: /etc/ssl/beats/beat.crt
            certificate_authorities:
            - /etc/ssl/beats/ca.crt
            key: /etc/ssl/beats/beat.key
`

	filebeatConfigTpl = `filebeat.inputs:
- type: docker
  containers.ids:
  - '*'

filebeat.config:
  modules:
    path: ${path.config}/modules.d/*.yml
    reload.enabled: false

fields:
    info:
        launch_type: bpm
        node_xid: {{ .ID }}
        project: development
        protocol_type: POLKADOT
        network_type: public
        user_id: TODO
        environment: {{ .Environment }}
fields_under_root: true
output:
    logstash:
        hosts:
        - "{{ .Collection.Host }}"
        ssl:
            certificate: /etc/ssl/beats/beat.crt
            certificate_authorities:
            - /etc/ssl/beats/ca.crt
            key: /etc/ssl/beats/beat.key
`
)

func createConfigs(currentNode node.Node) error {
	if err := template.ConfigFileRendered(filebeatConfigFile, filebeatConfigTpl, currentNode); err != nil {
		return err
	}

	return template.ConfigFileRendered(polkadotbeatConfigFile, polkadotbeatConfigTpl, currentNode)
}
