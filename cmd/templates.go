package main

const (
	collectorEnvTpl = `SERVICE_PORT=9933
SERVICE_HOST={{ .Node.NamePrefix }}` + polkadotContainerName

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
