// Copyright 2018 The go-ccmchain Authors
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
	"strings"
	"testing"
)

const mccmoddata = `
[
	{"type": "function", "name": "balance", "constant": true },
	{"type": "function", "name": "send", "constant": false, "inputs": [{ "name": "amount", "type": "uint256" }]},
	{"type": "function", "name": "transfer", "constant": false, "inputs": [{"name": "from", "type": "address"}, {"name": "to", "type": "address"}, {"name": "value", "type": "uint256"}], "outputs": [{"name": "success", "type": "bool"}]},
	{"constant":false,"inputs":[{"components":[{"name":"x","type":"uint256"},{"name":"y","type":"uint256"}],"name":"a","type":"tuple"}],"name":"tuple","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},
	{"constant":false,"inputs":[{"components":[{"name":"x","type":"uint256"},{"name":"y","type":"uint256"}],"name":"a","type":"tuple[]"}],"name":"tupleSlice","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},
	{"constant":false,"inputs":[{"components":[{"name":"x","type":"uint256"},{"name":"y","type":"uint256"}],"name":"a","type":"tuple[5]"}],"name":"tupleArray","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},
	{"constant":false,"inputs":[{"components":[{"name":"x","type":"uint256"},{"name":"y","type":"uint256"}],"name":"a","type":"tuple[5][]"}],"name":"complexTuple","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}
]`

func TestMccmodString(t *testing.T) {
	var table = []struct {
		mccmod      string
		expectation string
	}{
		{
			mccmod:      "balance",
			expectation: "function balance() constant returns()",
		},
		{
			mccmod:      "send",
			expectation: "function send(uint256 amount) returns()",
		},
		{
			mccmod:      "transfer",
			expectation: "function transfer(address from, address to, uint256 value) returns(bool success)",
		},
		{
			mccmod:      "tuple",
			expectation: "function tuple((uint256,uint256) a) returns()",
		},
		{
			mccmod:      "tupleArray",
			expectation: "function tupleArray((uint256,uint256)[5] a) returns()",
		},
		{
			mccmod:      "tupleSlice",
			expectation: "function tupleSlice((uint256,uint256)[] a) returns()",
		},
		{
			mccmod:      "complexTuple",
			expectation: "function complexTuple((uint256,uint256)[5][] a) returns()",
		},
	}

	abi, err := JSON(strings.NewReader(mccmoddata))
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range table {
		got := abi.Mccmods[test.mccmod].String()
		if got != test.expectation {
			t.Errorf("expected string to be %s, got %s", test.expectation, got)
		}
	}
}

func TestMccmodSig(t *testing.T) {
	var cases = []struct {
		mccmod string
		expect string
	}{
		{
			mccmod: "balance",
			expect: "balance()",
		},
		{
			mccmod: "send",
			expect: "send(uint256)",
		},
		{
			mccmod: "transfer",
			expect: "transfer(address,address,uint256)",
		},
		{
			mccmod: "tuple",
			expect: "tuple((uint256,uint256))",
		},
		{
			mccmod: "tupleArray",
			expect: "tupleArray((uint256,uint256)[5])",
		},
		{
			mccmod: "tupleSlice",
			expect: "tupleSlice((uint256,uint256)[])",
		},
		{
			mccmod: "complexTuple",
			expect: "complexTuple((uint256,uint256)[5][])",
		},
	}
	abi, err := JSON(strings.NewReader(mccmoddata))
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range cases {
		got := abi.Mccmods[test.mccmod].Sig()
		if got != test.expect {
			t.Errorf("expected string to be %s, got %s", test.expect, got)
		}
	}
}
