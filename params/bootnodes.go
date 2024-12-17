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
        "enode://94987177ceafa234fb2914471a822d6facb5b6ed560786222e65608efea3c371d6b4a48c4c79d783c757f2d7c7a5a352d4c721f6e2429db575601fd8d59e76e1@162.19.230.47:30304",
        "enode://0fe8a8bf36698fdf7cd848b5a38bf7f6ca272d7cba6c38cd0aacf1d8fe213b5502196679cc1a17a0bd5239c0eea7dfc64d095f14fc0e8a00c352d133af05f524@54.37.204.74:30304",
	"enode://bc80b5e48f9353020da879546407bd0f01d7e21cffd5826811f415731f87e0690a78722ef540caf876f4792308a5d9ed78a3c021a609b9e62f6707296353bc66@94.203.255.70:30304",
	"enode://88e1451b4d544f927fb6a0d5ce409c98195a98a87b491c4662f76701d1039413287d28d25142c5f4d4169e3314e2473a66f9879bb3d2eb6e8d010395d7cc09bb@45.61.160.54:30304",
        "enode://a8990014830cf131acad4aa18fbd7be5db24980d6c6941c93384b8091e1d309ba5081fdeb39b5ab186f1ffe1e90ba589c68e96c8a782a77c870c389284267ec9@178.72.89.199:58938",
        "enode://42405a45ff11bf3938c7734a8bffb3e08794ac5fb50875573d46862951595a9777843d72114411c06f9a7ee926f40eeb2d2e15f65a054f4d1d5cd3fb4830e2bf@194.67.196.93:30304",
        "enode://a0d5f597a8ddcc5de07ba71e596b7bb5141c07ab32e15503a5d20bebb4382146132e7cb2b003bff84d852b7466f7bb001060d938d88087cef47ffe56c93c9196@148.113.173.115:30304",
        "enode://fd3b31010a55e0a8d8975a8699c06ee1b9b8c7bc855846206c606d639f5cbc87f84a29b4c02437569ed8149b88a9fabb6058425709fd7e54acfa2317c835834e@107.172.157.49:30304",
        "enode://e80d4d38f4776f27cda0e0604003da1f0868a7610e27e5b43d5de889c6fb22b02bf3351a458da718c8825bc2ead99e5b9783c74871408248e234ad9e2d28f5c2@82.65.240.208:30304",
        "enode://44258e70f5d4fa2c0cd1a466222fc1b484fbcd7ad33b092b07b6c3b2d141aa121a2f666b163578c6b1a0db76ba22fa3f328ce5419fe3d90bd98638fddaa55eb2@91.126.138.190:30304",
        "enode://911ea9bd2be92fa11caa546f910cfe72d92dc15eb3f86a85b13465928bac37be5d238ff8394ed1735f9531ffb7a8b42d4d5ec00896ce4542685a1002975f3573@57.180.26.174:30304",
        "enode://d6090b4127a608d185a81717f1722d58bf4f07fe3eba9a90ba524e07de1ac75468794061ea4c6ecbfb8bf4b21c84511cd542167ef428a3128571c7fff01ee869@146.59.239.222:30304",	

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
