// Copyright 2019 The go-ccmchain Authors
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

package core

import (
	"sync/atomic"

	"github.com/ccmchain/go-ccmchain/common"
	"github.com/ccmchain/go-ccmchain/consensus"
	"github.com/ccmchain/go-ccmchain/core/state"
	"github.com/ccmchain/go-ccmchain/core/types"
	"github.com/ccmchain/go-ccmchain/core/vm"
	"github.com/ccmchain/go-ccmchain/params"
)

// statePrefetcher is a basic Prefetcher, which blindly executes a block on top
// of an arbitrary state with the goal of prefetching potentially useful state
// data from disk before the main block processor start executing.
type statePrefetcher struct {
	config *params.ChainConfig // Chain configuration options
	bc     *BlockChain         // Canonical block chain
	engine consensus.Engine    // Consensus engine used for block rewards
}

// newStatePrefetcher initialises a new statePrefetcher.
func newStatePrefetcher(config *params.ChainConfig, bc *BlockChain, engine consensus.Engine) *statePrefetcher {
	return &statePrefetcher{
		config: config,
		bc:     bc,
		engine: engine,
	}
}

// Prefetch processes the state changes according to the Ccmchain rules by running
// the transaction messages using the statedb, but any changes are discarded. The
// only goal is to pre-cache transaction signatures and state trie nodes.
func (p *statePrefetcher) Prefetch(block *types.Block, statedb *state.StateDB, cfg vm.Config, interrupt *uint32) {
	var (
		header  = block.Header()
		gaspool = new(GasPool).AddGas(block.GasLimit())
	)
	// Iterate over and process the individual transactions
	for i, tx := range block.Transactions() {
		// If block precaching was interrupted, abort
		if interrupt != nil && atomic.LoadUint32(interrupt) == 1 {
			return
		}
		// Block precaching permitted to continue, execute the transaction
		statedb.Prepare(tx.Hash(), block.Hash(), i)
		if err := precacheTransaction(p.config, p.bc, nil, gaspool, statedb, header, tx, cfg); err != nil {
			return // Ugh, somccming went horribly wrong, bail out
		}
	}
}

// precacheTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. The goal is not to execute
// the transaction successfully, rather to warm up touched data slots.
func precacheTransaction(config *params.ChainConfig, bc ChainContext, author *common.Address, gaspool *GasPool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, cfg vm.Config) error {
	// Convert the transaction into an executable message and pre-cache its sender
	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return err
	}
	// Create the EVM and execute the transaction
	context := NewEVMContext(msg, header, bc, author)
	vm := vm.NewEVM(context, statedb, config, cfg)

	_, _, _, err = ApplyMessage(vm, msg, gaspool)
	return err
}
