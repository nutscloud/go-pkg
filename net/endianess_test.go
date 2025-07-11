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

import "testing"

func TestNtohs(t *testing.T) {
	tests := []struct {
		name string
		net  uint16
		want uint16
	}{
		{"LittleEndian", 0x1234, 0x3412},
		{"BigEndian", 0x3412, 0x1234},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ntohs(tt.net); got != tt.want {
				t.Errorf("Ntohs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNtohl(t *testing.T) {
	tests := []struct {
		name string
		net  uint32
		want uint32
	}{
		{"LittleEndian", 0x12345678, 0x78563412},
		{"BigEndian", 0x78563412, 0x12345678},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ntohl(tt.net); got != tt.want {
				t.Errorf("Ntohl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNtohll(t *testing.T) {
	tests := []struct {
		name string
		net  uint64
		want uint64
	}{
		{"LittleEndian", 0x1234567890abcdef, 0xefcdab9078563412},
		{"BigEndian", 0xefcdab9078563412, 0x1234567890abcdef},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ntohll(tt.net); got != tt.want {
				t.Errorf("Ntohll() = %v, want %v", got, tt.want)
			}
		})
	}
}
