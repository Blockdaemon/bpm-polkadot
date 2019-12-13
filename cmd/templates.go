package main

const (
	polkadotbeatConfigTpl = `polkadotbeat:
    period: 30s
    polkadot_host: "{{ .Node.NamePrefix }}polkadot"
    polkadot_port: "9933"
fields:
    node:
        launch_type: bpm
        xid: {{ .Node.ID }}
        plugin: {{ .Node.PluginName }}
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
alexander
{{ if eq .Node.StrParameters.subtype "validator" }}
--validator
--key {{ index .Node.StrParameters "validator-key" }}
{{ end }}
`
)
