// Copyright 2025 nutscloud authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package net

import (
	"encoding/binary"
	"unsafe"
)

// NativeEndian is the ByteOrder of the current system.
var NativeEndian binary.ByteOrder

func init() {
	// Examine the memory layout of an int16 to determine system
	// endianness.
	var one int16 = 1
	b := (*byte)(unsafe.Pointer(&one))
	if *b == 0 {
		NativeEndian = binary.BigEndian
	} else {
		NativeEndian = binary.LittleEndian
	}
}

func NativelyLittle() bool {
	return NativeEndian == binary.LittleEndian
}

func Htons(host uint16) uint16 {
	return (host&0xff)<<8 | (host >> 8)
}

func Htonl(host uint32) uint32 {
	return (host&0xff)<<24 | (host&0xff00)<<8 | (host&0xff0000)>>8 | (host >> 24)
}

func Htonll(host uint64) uint64 {
	return (host&0xff)<<56 | (host&0xff00)<<40 | (host&0xff0000)<<24 | (host&0xff000000)<<8 | (host&0xff00000000)>>8 | (host&0xff0000000000)>>24 | (host&0xff000000000000)>>40 | (host&0xff00000000000000)>>56
}

func Ntohs(net uint16) uint16 {
	return Htons(net)
}

func Ntohl(net uint32) uint32 {
	return Htonl(net)
}

func Ntohll(net uint64) uint64 {
	return Htonll(net)
}
