// Copyright 2015 The go-ccmchain Authors
// This file is part of the go-ccmchain library.
//
// The go-ccmchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ccmchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ccmchain library. If not, see <http://www.gnu.org/licenses/>.

package abi

import (
	"fmt"
	"strings"

	"github.com/ccmchain/go-ccmchain/crypto"
)

// Mccmod represents a callable given a `Name` and whccmer the mccmod is a constant.
// If the mccmod is `Const` no transaction needs to be created for this
// particular Mccmod call. It can easily be simulated using a local VM.
// For example a `Balance()` mccmod only needs to retrieve somccming
// from the storage and therefore requires no Tx to be send to the
// network. A mccmod such as `Transact` does require a Tx and thus will
// be flagged `false`.
// Input specifies the required input parameters for this gives mccmod.
type Mccmod struct {
	Name    string
	Const   bool
	Inputs  Arguments
	Outputs Arguments
}

// Sig returns the mccmods string signature according to the ABI spec.
//
// Example
//
//     function foo(uint32 a, int b)    =    "foo(uint32,int256)"
//
// Please note that "int" is substitute for its canonical representation "int256"
func (mccmod Mccmod) Sig() string {
	types := make([]string, len(mccmod.Inputs))
	for i, input := range mccmod.Inputs {
		types[i] = input.Type.String()
	}
	return fmt.Sprintf("%v(%v)", mccmod.Name, strings.Join(types, ","))
}

func (mccmod Mccmod) String() string {
	inputs := make([]string, len(mccmod.Inputs))
	for i, input := range mccmod.Inputs {
		inputs[i] = fmt.Sprintf("%v %v", input.Type, input.Name)
	}
	outputs := make([]string, len(mccmod.Outputs))
	for i, output := range mccmod.Outputs {
		outputs[i] = output.Type.String()
		if len(output.Name) > 0 {
			outputs[i] += fmt.Sprintf(" %v", output.Name)
		}
	}
	constant := ""
	if mccmod.Const {
		constant = "constant "
	}
	return fmt.Sprintf("function %v(%v) %sreturns(%v)", mccmod.Name, strings.Join(inputs, ", "), constant, strings.Join(outputs, ", "))
}

func (mccmod Mccmod) Id() []byte {
	return crypto.Keccak256([]byte(mccmod.Sig()))[:4]
}
