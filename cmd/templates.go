package main

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

    polkadotCmdTpl = `polkadot
--base-path
/data
--rpc-external
--name
{{ .Config.name }}
--chain
{{ .Environment }}
{{ if eq .Subtype "validator" }}
--validator
--key
{{ .Config.key }}
{{ end }}
{{ if .Config.in_peers }}
--in-peers
{{ .Config.in_peers }}
{{ end }}
{{ if .Config.out_peers }}
--out-peers
{{ .Config.out_peers }}
{{ end }}
`

)
