This repository contains a WIP Polkadot plugin for BPM.

It currently supports:

- Watcher nodes
- Experimental validator nodes
- Sending blockheight, best peer blockheight and peer count to the backend
- Sending logs to the backend

# How to build it

- Go 1.12+ is required to build.

Run `make` which will automatically create versioned binaries for Linux and OSX
Usage.

Once built, the plugin can be used directly by calling `./bin/polkadot-<version>-<os>-amd64` or, once the plugin repository is implemented it can be uploaded to the repository and used with bpm.

# How to run it

Requirments to run:

1. Ubuntu 18.04
2. Docker-CE 19.03  
or
1. macOS recent version
2. Docker for Mac 

Your user must be in the docker group.  

To use the plugin directly:

1. Build the binary or download a pre-build binary from https://runner-test.sfo2.digitaloceanspaces.com/polkadot-0.5.0-linux-amd64
2. Create the node directory: `mkdir -p ~/.blockdaemon/nodes/polkadot-ms2/`
3. Copy the node configuration file: `cp node_example.json ~/.blockdaemon/nodes/polkadot-ms2/node.json`
4. Copy the certificates and keys directory: `cp -r beats ~/.blockdaemon/beats`
5. Run through the plugin lifecycle:

```
./polkadot-0.5.0-linux-amd64 create-configurations polkadot-ms2
./polkadot-0.5.0-linux-amd64 start polkadot-ms2
```

This will create the secrets, configuration and finally start the docker container with the blockchain node. 

You should now see a running multiple docker container: `docker ps`

# To remove it

Note: If you use the purge flag your configuration files will alse be removed, if you want to keep your configuration files omit this flag.
```
./polkadot-0.5.0-linux-amd64 remove polkadot-ms2 --purge
```

# Dependencies

* https://github.com/Blockdaemon/polkadotbeat - Contains an elasticbeat that collects blockchain information and sends them to the backend

## Credits
Thanks to [Chevdor](https://github.com/chevdor) for his great docker container: https://hub.docker.com/r/chevdor/polkadot

 
