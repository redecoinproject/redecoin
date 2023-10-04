## Faster geth start node sync
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

## script install node

REDEV2 Network Node Installation

This script allows you to install a node for the REDEV2 network on your Ubuntu system. Choose the appropriate version for your Ubuntu distribution and run the corresponding command below.

Ubuntu 18.04
```shell
curl -sSL https://github.com/redecoinproject/node/raw/master/ubuntu18.sh | bash
```
Ubuntu 20.04
```shell
curl -sSL https://github.com/redecoinproject/node/raw/master/ubuntu20.sh | bash
```
Ubuntu 23.04
```shell
curl -sSL https://github.com/redecoinproject/node/raw/master/ubuntu23.04.sh | bash
```
Raspberry Pi 4 arm
```shell
curl -sSL https://github.com/redecoinproject/node/raw/master/arm.sh | bash
```

Make sure to run the script as an administrator if required by your system.

License
This project is licensed under the MIT License.

Contributing
If you'd like to contribute to this project, please follow our contribution guidelines.

Issues
If you encounter any issues or have questions, please open an issue on our GitHub repository.

Happy node installation!
