### Requirements

- [Docker](https://www.docker.com/)

### Download bpm binary

`wget` must be installed to follow these instructions. The simplest way to install it on Ubuntu is:

```bash
sudo apt install wget
```

and on MacOS (using the [homebrew package manager](https://brew.sh/)):

```bash
brew install wget
```

Please see the official [wget homepage](https://www.gnu.org/software/wget/) for further details.

```bash
wget <BPM_URL>/bpm-0.5.0-<OS>-amd64
sudo cp bpm-0.5.0-<OS>-amd64 /usr/local/bin/bpm
sudo chmod 755 /usr/local/bin/bpm
```

Replace `<OS>` with either `linux` or `darwin`. Note that the version can change, currently we are running `0.5.0`.

## Usage

Set the current registry using the `BPM_REGISTRY_URL` env var.

```bash
Usage:
  bpm [command]

Available Commands:
  configure   Configure a new blockchain node
  help        Help about any command
  install     Installs or upgrades a package
  list        List available and installed blockchain protocols
  show        Print a resource to stdout
  start       Start a blockchain node
  status      Display statuses of configured nodes
  stop        Removes a running blockchain client. Data and configuration will not be removed.
  uninstall   Uninstall a package. Data and configuration will not be removed.
  version     Print the version

Flags:
      --base-dir string   The directory plugins and configuration are stored (default "~/.bpm/")
  -h, --help              help for bpm

Use "bpm [command] --help" for more information about a command.
```

### Example (Polkadot)

Install the polkadot package:

```bash
export BPM_REGISTRY_URL= <REGISTRY_URL>
bpm install polkadot 0.6.0
```

Configure the node and optionally pass additional fields:

```bash
bpm configure polkadot --field name=polkadot
Node with id "bm0lmirmvbaj4is78gtg" has been initialized, add your configuration (node.json) and secrets here:
...
```

Add your configs and secrets to the directory above then:

```bash
bpm start polkadot bm0lmirmvbaj4is78gtg
```

Check for running nodes with the status command:

```bash
bpm status
        NODE ID        | STATUS  |  PLUGIN  | SECRETS
+----------------------+---------+----------+---------+
  bm0lmirmvbaj4is78gtg | running | polkadot |       0
```

To stop the node run:

```bash
bpm stop polkadot bm0lmirmvbaj4is78gtg
```

Please note that the above does not remove data volumes or configurations. To force the removal of all data use:

```bash
bpm stop polkadot bm0lmirmvbaj4is78gtg --purge
```

> Be careful with the `--purge` parameters. If you purge an already fully synced blockchain you loose all data and have to re-sync from scratch. Any manual customisations to the configuration files will be lost as well.

To uninstall the package use:

```bash
bpm uninstall polkadot
```

