## Faster start node sync
```
./geth --syncmode full console
```
## Building the source

For prerequisites and detailed build instructions please read the

Building `geth` requires both a Go (version 1.16 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make geth
```

or, to build the full suite of utilities:

```shell
make all
```

## Running `geth`

Minimum:

* CPU with 2+ cores
* 4GB RAM
* High Performance SSD
* 8 MBit/sec download Internet service

Recommended:

* Fast CPU with 4+ cores
* 16GB+ RAM
* High Performance SSD
* 25+ MBit/sec download Internet service

