package main

const (
	polkadotbeatConfigTpl = `polkadotbeat:
    period: 30s
    polkadot_host: "{{ .ContainerName "polkadot" }}"
    polkadot_port: "9933"
fields:
    info:
        system: BPM
        node_xid: {{ .NodeGID }}
        project: development
        protocol_type: POLKADOT
        network_type: public
        network_xid: {{ .BlockchainGID }}
        user_id: TODO
        environment: {{ .Environment }}
fields_under_root: true
output:
    logstash:
        hosts:
        - "{{ .Logstash.Host }}"
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
        system: BPM
        node_xid: {{ .NodeGID }}
        project: development
        protocol_type: POLKADOT
        network_type: public
        network_xid: {{ .BlockchainGID }}
        user_id: TODO
        environment: {{ .Environment }}
fields_under_root: true
output:
    logstash:
        hosts:
        - "{{ .Logstash.Host }}"
        ssl:
            certificate: /etc/ssl/beats/beat.crt
            certificate_authorities:
            - /etc/ssl/beats/ca.crt
            key: /etc/ssl/beats/beat.key
`
)
