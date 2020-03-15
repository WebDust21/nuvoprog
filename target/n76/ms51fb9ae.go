// Copyright © 2019 Erin Shepherd
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
package n76

import (
//	"errors"

	//"github.com/erincandescent/nuvoprog/protocol"
	//"github.com/erincandescent/nuvoprog/target"
	"../../protocol"
	"../..//target"
)

var MS51FB9AE = &target.Definition{
	Name:        "MS51FB9AE",
	Family:      protocol.ChipFamilyN76E003,
	DeviceID:    protocol.DeviceMS51FB9AE,
	ProgMemSize: 12 * 1024,
	LDROMOffset: 0x3000,
	Config: target.ConfigSpace{
		IHexOffset: 0x30000,
		MinSize:    4,
		ReadSize:   8,
		WriteSize:  32,
		NewConfig:  func() target.Config { return new(N76E003Config) },
	},
}

func init() {
	target.Register(MS51FB9AE)
}
