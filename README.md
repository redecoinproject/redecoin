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

Ubuntu
```shell
curl -sSL https://github.com/redecoinproject/node/raw/master/ubuntu.sh | bash
```


Make sure to run the script as an administrator if required by your system.

License
This project is licensed under the MIT License.

Contributing
If you'd like to contribute to this project, please follow our contribution guidelines.

Issues
If you encounter any issues or have questions, please open an issue on our GitHub repository.

Happy node installation!

---
enode list 10.02.2026r
```shell
admin.addPeer("enode://5dd9bb009459497091937a78e6fd475232f039df06223c115f6ffed6b58f926d7482860926b8fc80ce18ca4451516dcae12d01d21169e5a9fd3a9dc9dee63629@2.139.240.114:30304");
admin.addPeer("enode://41816ffa95040e52266da95a9e87d7f2a8779415aa93ad0d4aec6ff60ccb6a5c735431fa44608311f1d71ca4fa1b7bced3d63757200cb29507689f8466431c69@152.70.184.19:30304");
admin.addPeer("enode://204ad03f979cac76d6ff478b81072f965e6d13292f532bd932cfe6c80d0a5b3f4142a1a8f2a0eae9024bc68ec8b437fb881afc557e31147f328375b8adb254cd@195.26.240.202:30304");
admin.addPeer("enode://99ad09d0efef37fc32859c804fe43e9c5aed7b91ecac16cf2db2ab151868576071bf42ef3f2a27147259dd9fdaec2cd04994f86d427bae6f630510c9f83f9bac@194.67.204.135:36070");
admin.addPeer("enode://fad5d2682a9e922c20251e41bfc7000d3525a78a6d941c3dc75d8fbd946aa3baa451888c414b4ce2ccfa8c34cb16ad94f85aae452fe5058d88b28ef600b25135@43.130.35.12:30304");
admin.addPeer("enode://b46b5e7352c348ab4f492a68cb9839397da2d94bd4cb5616dbf63f6d7108ddad31be7e0fc8a45b024ea6d292d016a0840f6520424d5c87b89a1732d6c52184dc@167.86.120.52:30304");
admin.addPeer("enode://177729ee0981f4c56cea74fe852d7c39d96c3f479e5b6fe1255c032cf01b86a56ee6f4c5e31dd7ca10ac58193acce385b903d8d3f9c62b36b8bee425ebc46187@167.86.70.237:41106");
admin.addPeer("enode://a065ddb9b11188ebc3a5ad8b7810d523c2a85b7a5fa4cc980146bb208bb8e5e959a88ea8eb0f5e160ffb455b5fb31eeb19a8d1d62268c166a42f38bfc5286685@107.172.157.49:30304");

