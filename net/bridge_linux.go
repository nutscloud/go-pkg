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
	"os"
	"path/filepath"
)

const (
	SYSCLASSNET = "/sys/class/net"
)

type linuxBridge struct{}

// GetBridgeInterfaces returns all interfaces of linuxBridge.
func (_ *linuxBridge) GetBridgeInterfaces(br string) ([]string, error) {
	ifaces := make([]string, 0)

	brIfPath := filepath.Join(SYSCLASSNET, br, "brif")
	ifs, err := os.ReadDir(brIfPath)
	if err != nil {
		return ifaces, err
	}
	for _, i := range ifs {
		ifaces = append(ifaces, i.Name())
	}

	return ifaces, nil
}

func (_ *linuxBridge) GetAllBridge() ([]string, error) {
	bridge := make([]string, 0)

	devs, err := os.ReadDir(SYSCLASSNET)
	if err != nil {
		return bridge, err
	}

	for _, d := range devs {
		bridgePath := filepath.Join(SYSCLASSNET, d.Name(), "bridge")
		stat, err := os.Stat(bridgePath)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return bridge, err
		}

		if stat.IsDir() {
			bridge = append(bridge, d.Name())
		}
	}

	return bridge, nil
}
