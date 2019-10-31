# BPM README

## Requirements and pre-requisites

- Linux or OSX
- [Docker](https://www.docker.com/) needs to be installed
- [wget](https://www.gnu.org/software/wget/) needs to be installed

### Installing `wget`

`wget` must be installed to follow these instructions. To install it on Ubuntu run:

```bash
sudo apt install wget
```

and on MacOS (using the [homebrew package manager](https://brew.sh/)):

```bash
brew install wget
```

## Download bpm binary

```bash
wget https://runner-test.sfo2.digitaloceanspaces.com/bpm-<VERSION>-<OS>-amd64
sudo cp bpm-<VERSION>-<OS>-amd64 /usr/local/bin/bpm
sudo chmod 755 /usr/local/bin/bpm
```

Replace `<VERSION>` with the actual version of BPM (e.g. `0.6.0`) and `<OS>` with either `linux` or `darwin`.

## Tutorial

### Getting help

First run `bpm --help` to see a list of available commands:

```
Blockchain Package Manager (BPM) manages blockchain nodes on your own infrastructure.

Usage:
  bpm [command]

Available Commands:
  configure   Configure a new blockchain node
  help        Help about any command
  info        Show information about a package
  install     Installs or upgrades a package to a specific version or latest if no version is specified
  list        List installed packages
  remove      Remove blockchain node data and configuration
  search      Search available packages
  show        Print a resource to stdout
  start       Start a blockchain node
  status      Display statuses of configured nodes
  stop        Stops a running blockchain node
  test        Tests a running blockchain node
  uninstall   Uninstall a package. Data and configuration will not be removed.
  version     Print the version

Flags:
      --base-dir string           The directory plugins and configuration are stored (default "~/.bpm/")
      --debug                     Enable debug output
  -h, --help                      help for bpm
      --package-registry string   The package registry provides packages to install (default "https://dev.registry.blockdaemon.com")

Use "bpm [command] --help" for more information about a command.
```

### Installing a package

In order to install a package we first need to know it's name. The best way to find available plugins is to search for them.

Run `bpm search polkadot`. The result should look like this:

```
    NAME   | PROTOCOL |          DESCRIPTION
+----------+----------+--------------------------------+
  polkadot | Polkadot | A package to run a polkadot
           |          | full node or validator on the
           |          | alexander testnet
```

To finally install the package, run `bpm install polkadot`. This will install the latest version. Optionally you can specify a version number but we recommend to always run the latest version of a package.

After waiting for a few minutes, while bpm is downloading the package, you should see the following message:

```
The package "polkadot" has been installed.
```

To verify which packages have been installed at any time, run `bpm list`:

```
    NAME   | INSTALLED VERSION | AVAILABLE VERSION
+----------+-------------------+-------------------+
  polkadot | 0.6.0             | 0.6.0
```

To get more detailed information about a package, run `bpm info polkadot` which results in:

```
Name:         polkadot
Description:  A package to run a polkadot full node or validator on the alexander testnet
Protocol:     polkadot
Network:      alexander
Network Type: public
Subtype:      watcher, validator
Versions:     0.3.0
              0.7.0
              0.6.0
```

Be sure to have any version of the polkadot package installed before running `bpm info polkadot`, or will result in an error.

### (Optional) Install Blockdaemon monitoring credentials

Optionally you can use the Blockdaemon monitoring infrastructure. This will instruct bpm to send node logs and metrics to Blockdaemon. In the future, Blockdaemon will make this monitoring data and alerts available via it's dashboard.

To use the Blockdaemon monitoring service, obtain the credentials from Blockdaemon and unpack them into the `~/.bpm/beats` directory:

```
mkdir -p ~/.bpm/beats
mv credentials.zip ~/.bpm/beats/
cd ~/.bpm/beats
unzip credentials.zip
```

As an alternative, you can change the config to use your own monitoring endpoint, or run the bpm nodes without dedicated monitoring backend.

### Configure a polkadot watcher

Run `bpm configure polkadot` which should output something similar to:

```
Nothing to do here, skipping create-secrets
Writing file '/Users/blockdaemon/.bpm/nodes/bmtgrr5h5s5km0mm9n2g/configs/polkadot.dockercmd'
Writing file '/Users/blockdaemon/.bpm/nodes/bmtgrr5h5s5km0mm9n2g/configs/polkadotbeat.yml'
Writing file '/Users/blockdaemon/.bpm/nodes/bmtgrr5h5s5km0mm9n2g/configs/filebeat.yml'

Node with id "bmtgrr5h5s5km0mm9n2g" has been initialized.

To change the configuration, modify the files here:
    /Users/blockdaemon/.bpm/nodes/bmtgrr5h5s5km0mm9n2g/configs
To start the node, run:
    bpm start bmtgrr5h5s5km0mm9n2g
To see the status of configured nodes, run:
    bpm status
```

Before starting the node you now have the opportunity to customize the generate configuration files by editing them in the `~/.bpm/nodes/<id>/configs/` directory.

### (Optional) Configure a custom monitoring endpoint

If you haven't installed the Blockdaemon monitoring credentials you can now add you own monitoring endpoints. To do this edit the files `~/.bpm/nodes/<id>/configs/polkadotbeat.yml` and `~/.bpm/nodes/<id>/configs/filebeat.yml` and add a configure an output as described here: https://www.elastic.co/guide/en/beats/filebeat/current/configuring-output.html

### Start a polkadot watcher

Now that the node is configured, let's check it's status by running `bpm status`.

```
        NODE ID        | PLUGIN  |  STATUS  | SECRETS
+----------------------+---------+----------+---------+
  bmtgrr5h5s5km0mm9n2g | stopped | polkadot |       0
```

It is currently stopped. To start the node, run `bpm start bmtgrr5h5s5km0mm9n2g`.

```
Creating network 'bpm-bmtgrr5h5s5km0mm9n2g-polkadot'
Creating container 'bpm-bmtgrr5h5s5km0mm9n2g-polkadot'
Starting container 'bpm-bmtgrr5h5s5km0mm9n2g-polkadot'
Creating container 'bpm-bmtgrr5h5s5km0mm9n2g-polkadotbeat'
Starting container 'bpm-bmtgrr5h5s5km0mm9n2g-polkadotbeat'
Creating container 'bpm-bmtgrr5h5s5km0mm9n2g-filebeat'
Starting container 'bpm-bmtgrr5h5s5km0mm9n2g-filebeat'
The node "bmtgrr5h5s5km0mm9n2g" has been started.
```

If we run `bpm status` again we can see the the node is now running:

```
        NODE ID        | PLUGIN  |  STATUS  | SECRETS
+----------------------+---------+----------+---------+
  bmtgrr5h5s5km0mm9n2g | running | polkadot |       0
```

Running `docker ps` will show you the containers started by bpm:

```
CONTAINER ID        IMAGE                                    COMMAND                  CREATED             STATUS              PORTS                                                          NAMES
28cdeaf0ef8d        docker.elastic.co/beats/filebeat:7.3.1   "/usr/local/bin/dock…"   6 minutes ago       Up 5 minutes                                                                       bpm-bmtgrr5h5s5km0mm9n2g-filebeat
a94ae5122a63        blockdaemon/polkadotbeat:1.0.0           "/usr/local/bin/dock…"   6 minutes ago       Up 6 minutes                                                                       bpm-bmtgrr5h5s5km0mm9n2g-polkadotbeat
26caeb0369b2        chevdor/polkadot:0.4.4                   "polkadot --base-pat…"   6 minutes ago       Up 6 minutes        127.0.0.1:9933->9933/tcp, 9944/tcp, 0.0.0.0:30333->30333/tcp   bpm-bmtgrr5h5s5km0mm9n2g-polkadot
```

### Testing the node

To verify that the node is working as expected, each bpm package comes with a set of tests. Run them with: `bpm test bmtgrr5h5s5km0mm9n2g`.

If everything is ok you should see an output like this:

```
2019/10/31 16:55:57 PASSED: chain_getBlock 200
2019/10/31 16:55:57 PASSED: chain_getBlockHash 200
2019/10/31 16:55:57 PASSED: chain_getFinalizedHead 200
2019/10/31 16:55:57 PASSED: chain_getHeader 200
2019/10/31 16:55:57 PASSED: system_chain Alexander
2019/10/31 16:55:57 PASSED: system_health Syncing: true Peers: 8 Should Have Peers: true
2019/10/31 16:55:57 PASSED: system_name parity-polkadot
2019/10/31 16:55:57 PASSED: system_networkState average DL/UL 16330/15231
2019/10/31 16:55:57 PASSED: system_peers 8 peer(s)
2019/10/31 16:55:57 PASSED: system_properties token decimals: 15 token symbol: DOT
2019/10/31 16:55:57 PASSED: system_version 0.4.4
```

### Stop the node

In order to remove the node we first need to stop it: `bpm stop bmtgrr5h5s5km0mm9n2g`.

```
Stopping container 'polkadot'
Removing container 'bpm-bmtgrr5h5s5km0mm9n2g-polkadot'
Stopping container 'polkadotbeat'
Removing container 'bpm-bmtgrr5h5s5km0mm9n2g-polkadotbeat'
Stopping container 'filebeat'
Removing container 'bpm-bmtgrr5h5s5km0mm9n2g-filebeat'
The node "bmtgrr5h5s5km0mm9n2g" has been stopped.
```

This only stops the node, it doesn't delete any of the already synced blockchain data or any of the configuration. Let's check the status again: `bpm status`.

```
        NODE ID        | PLUGIN  |  STATUS  | SECRETS
+----------------------+---------+----------+---------+
  bmtgrr5h5s5km0mm9n2g | stopped | polkadot |       0
```

To double check that the configs are still there, let's view them: `bpm show config bmtgrr5h5s5km0mm9n2g`. It should print the contents of all available configuration files for this particular node.

### Remove the node

In order to finally remove the node, including the configuration and the blockchain data, run `bpm remove bmtgrr5h5s5km0mm9n2g --all`.

```
Removing file '/Users/blockdaemon/.bpm/nodes/bmtgrr5h5s5km0mm9n2g/configs/polkadot.dockercmd'
Removing file '/Users/blockdaemon/.bpm/nodes/bmtgrr5h5s5km0mm9n2g/configs/polkadotbeat.yml'
Removing file '/Users/blockdaemon/.bpm/nodes/bmtgrr5h5s5km0mm9n2g/configs/filebeat.yml'
Container 'bpm-bmtgrr5h5s5km0mm9n2g-polkadot' is not running, skipping stop
Cannot find container 'bpm-bmtgrr5h5s5km0mm9n2g-polkadot', skipping removel
Container 'bpm-bmtgrr5h5s5km0mm9n2g-polkadotbeat' is not running, skipping stop
Cannot find container 'bpm-bmtgrr5h5s5km0mm9n2g-polkadotbeat', skipping removel
Container 'bpm-bmtgrr5h5s5km0mm9n2g-filebeat' is not running, skipping stop
Cannot find container 'bpm-bmtgrr5h5s5km0mm9n2g-filebeat', skipping removel
Removing volume 'bpm-bmtgrr5h5s5km0mm9n2g-polkadot-data'
Removing network 'bpm-bmtgrr5h5s5km0mm9n2g-polkadot'
Cannot find network 'bpm-bmtgrr5h5s5km0mm9n2g-polkadot', skipping removal
Cannot find network 'bpm-bmtgrr5h5s5km0mm9n2g-polkadot', skipping removal
Removing directory "/Users/blockdaemon/.bpm/nodes/bmtgrr5h5s5km0mm9n2g"
```

The `Cannot find container` message can safely be ignored. Stopping the node in the previous step allready removed the containers so the `remove` command skips this step.

Thre is also the possibility to specify exactly what should be removed. Run `bpm remove --help` to see all available flags. Those flags can come in hadny to reset the configuration or the blockchain during development or troubleshooting.

### Start a polkadot validator

Now that we removed the watcher node, let's do something more advanced and start a validator. By now you are familiar with most commands so we'll skip the output. This section also requires understanding of how to generate the necessary accounts and validator key as described here: https://wiki.polkadot.network/docs/en/maintain-guides-how-to-validate-alexander

Configure the validator node: `bpm configure polkadot --subtype validator` 

Edit the configuration file: `nano /Users/blockdaemon/.bpm/nodes/bmthj6lh5s5licpj01sg/configs/polkadot.dockercmd` and replace the string `{% ADD NODE KEY HERE %}` with the actual node key.

Start the validator node: `bpm start bmthj6lh5s5licpj01sg`

### Uninstall a package

To uninstall a package run `bpm uninstall polkadot`

