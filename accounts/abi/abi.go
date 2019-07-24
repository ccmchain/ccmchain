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
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/ccmchain/go-ccmchain/common"
)

// The ABI holds information about a contract's context and available
// invokable mccmods. It will allow you to type check function calls and
// packs data accordingly.
type ABI struct {
	Constructor Mccmod
	Mccmods     map[string]Mccmod
	Events      map[string]Event
}

// JSON returns a parsed ABI interface and error if it failed.
func JSON(reader io.Reader) (ABI, error) {
	dec := json.NewDecoder(reader)

	var abi ABI
	if err := dec.Decode(&abi); err != nil {
		return ABI{}, err
	}

	return abi, nil
}

// Pack the given mccmod name to conform the ABI. Mccmod call's data
// will consist of mccmod_id, args0, arg1, ... argN. Mccmod id consists
// of 4 bytes and arguments are all 32 bytes.
// Mccmod ids are created from the first 4 bytes of the hash of the
// mccmods string signature. (signature = baz(uint32,string32))
func (abi ABI) Pack(name string, args ...interface{}) ([]byte, error) {
	// Fetch the ABI of the requested mccmod
	if name == "" {
		// constructor
		arguments, err := abi.Constructor.Inputs.Pack(args...)
		if err != nil {
			return nil, err
		}
		return arguments, nil
	}
	mccmod, exist := abi.Mccmods[name]
	if !exist {
		return nil, fmt.Errorf("mccmod '%s' not found", name)
	}
	arguments, err := mccmod.Inputs.Pack(args...)
	if err != nil {
		return nil, err
	}
	// Pack up the mccmod ID too if not a constructor and return
	return append(mccmod.Id(), arguments...), nil
}

// Unpack output in v according to the abi specification
func (abi ABI) Unpack(v interface{}, name string, data []byte) (err error) {
	if len(data) == 0 {
		return fmt.Errorf("abi: unmarshalling empty output")
	}
	// since there can't be naming collisions with contracts and events,
	// we need to decide whccmer we're calling a mccmod or an event
	if mccmod, ok := abi.Mccmods[name]; ok {
		if len(data)%32 != 0 {
			return fmt.Errorf("abi: improperly formatted output: %s - Bytes: [%+v]", string(data), data)
		}
		return mccmod.Outputs.Unpack(v, data)
	}
	if event, ok := abi.Events[name]; ok {
		return event.Inputs.Unpack(v, data)
	}
	return fmt.Errorf("abi: could not locate named mccmod or event")
}

// UnpackIntoMap unpacks a log into the provided map[string]interface{}
func (abi ABI) UnpackIntoMap(v map[string]interface{}, name string, data []byte) (err error) {
	if len(data) == 0 {
		return fmt.Errorf("abi: unmarshalling empty output")
	}
	// since there can't be naming collisions with contracts and events,
	// we need to decide whccmer we're calling a mccmod or an event
	if mccmod, ok := abi.Mccmods[name]; ok {
		if len(data)%32 != 0 {
			return fmt.Errorf("abi: improperly formatted output")
		}
		return mccmod.Outputs.UnpackIntoMap(v, data)
	}
	if event, ok := abi.Events[name]; ok {
		return event.Inputs.UnpackIntoMap(v, data)
	}
	return fmt.Errorf("abi: could not locate named mccmod or event")
}

// UnmarshalJSON implements json.Unmarshaler interface
func (abi *ABI) UnmarshalJSON(data []byte) error {
	var fields []struct {
		Type      string
		Name      string
		Constant  bool
		Anonymous bool
		Inputs    []Argument
		Outputs   []Argument
	}

	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	abi.Mccmods = make(map[string]Mccmod)
	abi.Events = make(map[string]Event)
	for _, field := range fields {
		switch field.Type {
		case "constructor":
			abi.Constructor = Mccmod{
				Inputs: field.Inputs,
			}
		// empty defaults to function according to the abi spec
		case "function", "":
			name := field.Name
			_, ok := abi.Mccmods[name]
			for idx := 0; ok; idx++ {
				name = fmt.Sprintf("%s%d", field.Name, idx)
				_, ok = abi.Mccmods[name]
			}
			abi.Mccmods[name] = Mccmod{
				Name:    name,
				Const:   field.Constant,
				Inputs:  field.Inputs,
				Outputs: field.Outputs,
			}
		case "event":
			name := field.Name
			_, ok := abi.Events[name]
			for idx := 0; ok; idx++ {
				name = fmt.Sprintf("%s%d", field.Name, idx)
				_, ok = abi.Events[name]
			}
			abi.Events[name] = Event{
				Name:      name,
				Anonymous: field.Anonymous,
				Inputs:    field.Inputs,
			}
		}
	}

	return nil
}

// MccmodById looks up a mccmod by the 4-byte id
// returns nil if none found
func (abi *ABI) MccmodById(sigdata []byte) (*Mccmod, error) {
	if len(sigdata) < 4 {
		return nil, fmt.Errorf("data too short (%d bytes) for abi mccmod lookup", len(sigdata))
	}
	for _, mccmod := range abi.Mccmods {
		if bytes.Equal(mccmod.Id(), sigdata[:4]) {
			return &mccmod, nil
		}
	}
	return nil, fmt.Errorf("no mccmod with id: %#x", sigdata[:4])
}

// EventByID looks an event up by its topic hash in the
// ABI and returns nil if none found.
func (abi *ABI) EventByID(topic common.Hash) (*Event, error) {
	for _, event := range abi.Events {
		if bytes.Equal(event.Id().Bytes(), topic.Bytes()) {
			return &event, nil
		}
	}
	return nil, fmt.Errorf("no event with id: %#x", topic.Hex())
}
