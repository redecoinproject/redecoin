// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

var MainnetBootnodes = []string{
"enode://3eb5c62e61743442e392a3de56fda8cee541f56ef2f43cae032521f3ac715fe7393a97e08aea2ef6da06be73cd21396f8b2ed066e29181a1ff5b0519cb19eee2@37.187.131.105:30304",
"enode://e3eae0e252f687ffd0b0e07afc20346a140bc105a0015a96796b3c1e4a0e065a380a14292eb0e2907e8512a546054fd34ba2f1194aa6e92dd930720b1ef93ea5@194.180.19.36:46878",
"enode://f529ee3cd1d0b673779438b08dad82d38713e9d296bd58ad87c2d5cc2ebbbf66ecd8de7139cd4c4d939cb41ab2970c76e531eb07e40a2f2607b3cd624694be67@107.172.157.49:50896",
"enode://912f9bae91eeea4588b2ddbf9c14ca9045a4ee9da860bed8e61a925a9cfbf7391a208f605a6ef9a5b318022a544a86c61fe176fafdfe7d9cb5c77c9bc3367681@109.108.71.107:33555",
"enode://641943713a285e47c47aad7b14a669033d89a8953641b8d39a4a45e6c1180adae7f7ec66724924af31892721f9979e8585cfe70a0f281b49b257f5bd460b0bdd@95.143.216.223:54168",
"enode://2658b892f5cbe23205c7b7aebf87b5de481d8e700f457c42adc76499f5c6b452fa69eb4c6ac37ba49ddec1770d40d66b6cd990757dcbb849c52615beb20876d8@31.42.10.1:47744",
"enode://e4be4d0fd9085e502a6303dfc04e85cd3c44a9ac9a8cdb000b537807b7b3de8504dddf88b40a582bb39d2b161d92a485f05a481ae8229713452660ae77c0d8f0@146.59.18.4:40102",
"enode://16fb54e17f54586c3996628783c749f5e5949b5c92d29fd01b4a41297b5a67256be34aad8a260341b7f293b5af84cb361c2ce3286ebc4f5c93058f91a08e5880@95.143.216.223:45802",
"enode://e80d4d38f4776f27cda0e0604003da1f0868a7610e27e5b43d5de889c6fb22b02bf3351a458da718c8825bc2ead99e5b9783c74871408248e234ad9e2d28f5c2@82.65.240.208:46114",
"enode://bc80b5e48f9353020da879546407bd0f01d7e21cffd5826811f415731f87e0690a78722ef540caf876f4792308a5d9ed78a3c021a609b9e62f6707296353bc66@94.203.255.70:57744",
"enode://186baa1d91ffd2622e91267569d0089e24e6323ea1cd66434b2f11554c429554dd4001a0362865e42066cb8c76250f6dce915553cbd956a70eb47ea72960e1ba@188.165.213.69:55144",
"enode://7d14ca4ce343bd7a140b9eb19b3e4772bb2d1a920da49380f1e56d50a662ebc91adc2324cd09b0f438e809d11b026e61b18e96fa0d7afadca94691831e6bcc92@77.37.144.223:51192",
"enode://b49fbcabfb532e9753c123076d0759acb08ba9a4e49ddd2c1f81d1b0d289b5af221d67827b010001b99d0751e5ad11b8f0dfcfee493899535c03fba1442e27b5@191.101.234.48:46640",
"enode://f3cb6bbb4244e6ca950943526a14a4965e5852c1f1c95d5268a744ee63ec4e499f10bf144b3c4cf53dacb04128d6b0756b4ee8dbcb89e5b046563f8e10ca82f4@37.187.76.37:38336",
"enode://7b77c484f29888e0e5a4845e4158b3ff8a276f8794b844e5969ee0319bed4cec2dc8dbaf1b3e11bd386b8ad36d19d9fe94b9188c1735be2941f91543eeb41b43@95.143.216.223:54752",
"enode://0fe8a8bf36698fdf7cd848b5a38bf7f6ca272d7cba6c38cd0aacf1d8fe213b5502196679cc1a17a0bd5239c0eea7dfc64d095f14fc0e8a00c352d133af05f524@54.37.204.74:33124",
"enode://a05f74d5d0afa2aa0a74ee8d85c85472e24530a1ba432f702e157bbf278c555a883d60d27c79bcd0740b3a010e009a3bc6a02bb661b07ea54f50ef0f1c10c74d@89.33.193.194:57170",
"enode://477f86518ebe1f078d50dbd53bed2a4d722303c54dd373460c2e9d88b2da78430a937bd6ae4109ee596196e67169f21ccfa9039916c2034110d6b9bb3e88707a@31.42.10.1:34660",
"enode://b1d731502ab83228dbd87c02f304d64a4f526bafd35965f463bea96b54d0555c8dcca4f2801d3c16f86d0b5a0842760086872d6e992e27f672025d2c07a9daff@27.254.39.27:30304",
"enode://1cb9b707a8a2a7df290bbe2ba2f8eaec8df869a92f4437d0eace381bfda5ab7f84009e775496293bab055a0d2fb3de1492480047431b03c3dd246b993625bed1@178.72.89.199:46044",
"enode://b235cf1d0dfec0f0755f70755660b47687bc122ccd22530acb22d197e09bd0ab8bb0f2d602d4c01ce34ebd676089f6d05a8f86f77444e1842af37b398b006fa5@31.220.84.74:50966",
"enode://3fc51bbae633f941b5630ddec4ddcad5462636499d4ff849bb94ab3e72d5ea883f3ab913ddaa56a1522ada22906a5a7ddd1311e5e51a953e42e074dc18313270@202.178.115.11:35160",
"enode://27cc30f5ac099bfca272359f47aec5d86363afb7f1bb39344ce1b941abda374f63a7a34c284a49651964be05a73e37c83e0e691c5825fb1481b57ff34c053ef3@79.132.239.243:58951",
"enode://45f0cd6d3ec3efad1e4c6392455ade7064f9cc065d6c2265db0e4055b37c757badc002b5d9d265bdb498d41673b1c167e614a00a6ea22d4f322ba9a3ce636dca@146.59.32.81:34854",
}
var RopstenBootnodes = []string{
}
var SepoliaBootnodes = []string{
}
var RinkebyBootnodes = []string{
}
var GoerliBootnodes = []string{
}
var KilnBootnodes = []string{
}
var V5Bootnodes = []string{
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	case SepoliaGenesisHash:
		net = "sepolia"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
