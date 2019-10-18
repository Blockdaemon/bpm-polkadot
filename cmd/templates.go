package main

const (
	polkadotbeatConfigTpl = `polkadotbeat:
    period: 30s
    polkadot_host: "{{ .Node.NamePrefix }}polkadot"
    polkadot_port: "9933"
fields:
    info:
        launch_type: bpm
        node_xid: {{ .Node.ID }}
        project: development
        protocol_type: POLKADOT
        network_type: public
        user_id: TODO
        environment: {{ .Node.Environment }}
fields_under_root: true
output:
    logstash:
        hosts:
        - "{{ .Node.Collection.Host }}"
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
{{ .Node.Config.name }}
--chain
{{ .Node.Environment }}
{{ if eq .Node.Subtype "validator" }}
--validator
--key
{{ .Node.Config.key }}
{{ end }}
{{ if .Node.Config.in_peers }}
--in-peers
{{ .Node.Config.in_peers }}
{{ end }}
{{ if .Node.Config.out_peers }}
--out-peers
{{ .Node.Config.out_peers }}
{{ end }}
`
)
