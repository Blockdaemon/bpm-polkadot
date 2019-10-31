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
        protocol_type: {{ .Node.Subtype }}
        network_type: {{ .Node.NetworkType }}
        environment: {{ .Node.Environment }}
fields_under_root: true
output:
{{- if .Node.Collection.Host }}
    logstash:
        hosts:
        - "{{ .Node.Collection.Host }}"
        ssl:
            certificate: /etc/ssl/beats/beat.crt
            certificate_authorities:
            - /etc/ssl/beats/ca.crt
            key: /etc/ssl/beats/beat.key
{{- else }}
    console:
        pretty: true
{{- end }}
`

	polkadotCmdTpl = `polkadot
--base-path
/data
--rpc-external
--name
{{ .Node.ID }}
--chain
{{ .Node.Environment }}
{{ if eq .Node.Subtype "validator" }}
--validator
--key {% ADD NODE KEY HERE %}
{{ end }}
`
)
