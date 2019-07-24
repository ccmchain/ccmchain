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

// package web3ext contains gccm specific web3.js extensions.
package web3ext

var Modules = map[string]string{
	"accounting": AccountingJs,
	"admin":      AdminJs,
	"chequebook": ChequebookJs,
	"clique":     CliqueJs,
	"ccmash":     EthashJs,
	"debug":      DebugJs,
	"ccm":        EthJs,
	"miner":      MinerJs,
	"net":        NetJs,
	"personal":   PersonalJs,
	"rpc":        RpcJs,
	"shh":        ShhJs,
	"swarmfs":    SwarmfsJs,
	"txpool":     TxpoolJs,
	"les":        LESJs,
}

const ChequebookJs = `
web3._extend({
	property: 'chequebook',
	mccmods: [
		new web3._extend.Mccmod({
			name: 'deposit',
			call: 'chequebook_deposit',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Property({
			name: 'balance',
			getter: 'chequebook_balance',
			outputFormatter: web3._extend.utils.toDecimal
		}),
		new web3._extend.Mccmod({
			name: 'cash',
			call: 'chequebook_cash',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Mccmod({
			name: 'issue',
			call: 'chequebook_issue',
			params: 2,
			inputFormatter: [null, null]
		}),
	]
});
`

const CliqueJs = `
web3._extend({
	property: 'clique',
	mccmods: [
		new web3._extend.Mccmod({
			name: 'getSnapshot',
			call: 'clique_getSnapshot',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Mccmod({
			name: 'getSnapshotAtHash',
			call: 'clique_getSnapshotAtHash',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'getSigners',
			call: 'clique_getSigners',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Mccmod({
			name: 'getSignersAtHash',
			call: 'clique_getSignersAtHash',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'propose',
			call: 'clique_propose',
			params: 2
		}),
		new web3._extend.Mccmod({
			name: 'discard',
			call: 'clique_discard',
			params: 1
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'proposals',
			getter: 'clique_proposals'
		}),
	]
});
`

const EthashJs = `
web3._extend({
	property: 'ccmash',
	mccmods: [
		new web3._extend.Mccmod({
			name: 'getWork',
			call: 'ccmash_getWork',
			params: 0
		}),
		new web3._extend.Mccmod({
			name: 'getHashrate',
			call: 'ccmash_getHashrate',
			params: 0
		}),
		new web3._extend.Mccmod({
			name: 'submitWork',
			call: 'ccmash_submitWork',
			params: 3,
		}),
		new web3._extend.Mccmod({
			name: 'submitHashRate',
			call: 'ccmash_submitHashRate',
			params: 2,
		}),
	]
});
`

const AdminJs = `
web3._extend({
	property: 'admin',
	mccmods: [
		new web3._extend.Mccmod({
			name: 'addPeer',
			call: 'admin_addPeer',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'removePeer',
			call: 'admin_removePeer',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'addTrustedPeer',
			call: 'admin_addTrustedPeer',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'removeTrustedPeer',
			call: 'admin_removeTrustedPeer',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'exportChain',
			call: 'admin_exportChain',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Mccmod({
			name: 'importChain',
			call: 'admin_importChain',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'sleepBlocks',
			call: 'admin_sleepBlocks',
			params: 2
		}),
		new web3._extend.Mccmod({
			name: 'startRPC',
			call: 'admin_startRPC',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new web3._extend.Mccmod({
			name: 'stopRPC',
			call: 'admin_stopRPC'
		}),
		new web3._extend.Mccmod({
			name: 'startWS',
			call: 'admin_startWS',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new web3._extend.Mccmod({
			name: 'stopWS',
			call: 'admin_stopWS'
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'nodeInfo',
			getter: 'admin_nodeInfo'
		}),
		new web3._extend.Property({
			name: 'peers',
			getter: 'admin_peers'
		}),
		new web3._extend.Property({
			name: 'datadir',
			getter: 'admin_datadir'
		}),
	]
});
`

const DebugJs = `
web3._extend({
	property: 'debug',
	mccmods: [
		new web3._extend.Mccmod({
			name: 'printBlock',
			call: 'debug_printBlock',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'getBlockRlp',
			call: 'debug_getBlockRlp',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'testSignCliqueBlock',
			call: 'debug_testSignCliqueBlock',
			params: 2,
			inputFormatters: [web3._extend.formatters.inputAddressFormatter, null],
		}),
		new web3._extend.Mccmod({
			name: 'setHead',
			call: 'debug_setHead',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'seedHash',
			call: 'debug_seedHash',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'dumpBlock',
			call: 'debug_dumpBlock',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'chaindbProperty',
			call: 'debug_chaindbProperty',
			params: 1,
			outputFormatter: console.log
		}),
		new web3._extend.Mccmod({
			name: 'chaindbCompact',
			call: 'debug_chaindbCompact',
		}),
		new web3._extend.Mccmod({
			name: 'verbosity',
			call: 'debug_verbosity',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'vmodule',
			call: 'debug_vmodule',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'backtraceAt',
			call: 'debug_backtraceAt',
			params: 1,
		}),
		new web3._extend.Mccmod({
			name: 'stacks',
			call: 'debug_stacks',
			params: 0,
			outputFormatter: console.log
		}),
		new web3._extend.Mccmod({
			name: 'freeOSMemory',
			call: 'debug_freeOSMemory',
			params: 0,
		}),
		new web3._extend.Mccmod({
			name: 'setGCPercent',
			call: 'debug_setGCPercent',
			params: 1,
		}),
		new web3._extend.Mccmod({
			name: 'memStats',
			call: 'debug_memStats',
			params: 0,
		}),
		new web3._extend.Mccmod({
			name: 'gcStats',
			call: 'debug_gcStats',
			params: 0,
		}),
		new web3._extend.Mccmod({
			name: 'cpuProfile',
			call: 'debug_cpuProfile',
			params: 2
		}),
		new web3._extend.Mccmod({
			name: 'startCPUProfile',
			call: 'debug_startCPUProfile',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'stopCPUProfile',
			call: 'debug_stopCPUProfile',
			params: 0
		}),
		new web3._extend.Mccmod({
			name: 'goTrace',
			call: 'debug_goTrace',
			params: 2
		}),
		new web3._extend.Mccmod({
			name: 'startGoTrace',
			call: 'debug_startGoTrace',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'stopGoTrace',
			call: 'debug_stopGoTrace',
			params: 0
		}),
		new web3._extend.Mccmod({
			name: 'blockProfile',
			call: 'debug_blockProfile',
			params: 2
		}),
		new web3._extend.Mccmod({
			name: 'setBlockProfileRate',
			call: 'debug_setBlockProfileRate',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'writeBlockProfile',
			call: 'debug_writeBlockProfile',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'mutexProfile',
			call: 'debug_mutexProfile',
			params: 2
		}),
		new web3._extend.Mccmod({
			name: 'setMutexProfileFraction',
			call: 'debug_setMutexProfileFraction',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'writeMutexProfile',
			call: 'debug_writeMutexProfile',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'writeMemProfile',
			call: 'debug_writeMemProfile',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'traceBlock',
			call: 'debug_traceBlock',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Mccmod({
			name: 'traceBlockFromFile',
			call: 'debug_traceBlockFromFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Mccmod({
			name: 'traceBadBlock',
			call: 'debug_traceBadBlock',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Mccmod({
			name: 'standardTraceBadBlockToFile',
			call: 'debug_standardTraceBadBlockToFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Mccmod({
			name: 'standardTraceBlockToFile',
			call: 'debug_standardTraceBlockToFile',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Mccmod({
			name: 'traceBlockByNumber',
			call: 'debug_traceBlockByNumber',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Mccmod({
			name: 'traceBlockByHash',
			call: 'debug_traceBlockByHash',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Mccmod({
			name: 'traceTransaction',
			call: 'debug_traceTransaction',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Mccmod({
			name: 'preimage',
			call: 'debug_preimage',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Mccmod({
			name: 'getBadBlocks',
			call: 'debug_getBadBlocks',
			params: 0,
		}),
		new web3._extend.Mccmod({
			name: 'storageRangeAt',
			call: 'debug_storageRangeAt',
			params: 5,
		}),
		new web3._extend.Mccmod({
			name: 'getModifiedAccountsByNumber',
			call: 'debug_getModifiedAccountsByNumber',
			params: 2,
			inputFormatter: [null, null],
		}),
		new web3._extend.Mccmod({
			name: 'getModifiedAccountsByHash',
			call: 'debug_getModifiedAccountsByHash',
			params: 2,
			inputFormatter:[null, null],
		}),
	],
	properties: []
});
`

const EthJs = `
web3._extend({
	property: 'ccm',
	mccmods: [
		new web3._extend.Mccmod({
			name: 'chainId',
			call: 'ccm_chainId',
			params: 0
		}),
		new web3._extend.Mccmod({
			name: 'sign',
			call: 'ccm_sign',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter, null]
		}),
		new web3._extend.Mccmod({
			name: 'resend',
			call: 'ccm_resend',
			params: 3,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter, web3._extend.utils.fromDecimal, web3._extend.utils.fromDecimal]
		}),
		new web3._extend.Mccmod({
			name: 'signTransaction',
			call: 'ccm_signTransaction',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter]
		}),
		new web3._extend.Mccmod({
			name: 'submitTransaction',
			call: 'ccm_submitTransaction',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter]
		}),
		new web3._extend.Mccmod({
			name: 'getHeaderByNumber',
			call: 'ccm_getHeaderByNumber',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'getHeaderByHash',
			call: 'ccm_getHeaderByHash',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'getBlockByNumber',
			call: 'ccm_getBlockByNumber',
			params: 2
		}),
		new web3._extend.Mccmod({
			name: 'getBlockByHash',
			call: 'ccm_getBlockByHash',
			params: 2
		}),
		new web3._extend.Mccmod({
			name: 'getRawTransaction',
			call: 'ccm_getRawTransactionByHash',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'getRawTransactionFromBlock',
			call: function(args) {
				return (web3._extend.utils.isString(args[0]) && args[0].indexOf('0x') === 0) ? 'ccm_getRawTransactionByBlockHashAndIndex' : 'ccm_getRawTransactionByBlockNumberAndIndex';
			},
			params: 2,
			inputFormatter: [web3._extend.formatters.inputBlockNumberFormatter, web3._extend.utils.toHex]
		}),
		new web3._extend.Mccmod({
			name: 'getProof',
			call: 'ccm_getProof',
			params: 3,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter, null, web3._extend.formatters.inputBlockNumberFormatter]
		}),
	],
	properties: [
		new web3._extend.Property({
			name: 'pendingTransactions',
			getter: 'ccm_pendingTransactions',
			outputFormatter: function(txs) {
				var formatted = [];
				for (var i = 0; i < txs.length; i++) {
					formatted.push(web3._extend.formatters.outputTransactionFormatter(txs[i]));
					formatted[i].blockHash = null;
				}
				return formatted;
			}
		}),
	]
});
`

const MinerJs = `
web3._extend({
	property: 'miner',
	mccmods: [
		new web3._extend.Mccmod({
			name: 'start',
			call: 'miner_start',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Mccmod({
			name: 'stop',
			call: 'miner_stop'
		}),
		new web3._extend.Mccmod({
			name: 'setEtherbase',
			call: 'miner_setEtherbase',
			params: 1,
			inputFormatter: [web3._extend.formatters.inputAddressFormatter]
		}),
		new web3._extend.Mccmod({
			name: 'setExtra',
			call: 'miner_setExtra',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'setGasPrice',
			call: 'miner_setGasPrice',
			params: 1,
			inputFormatter: [web3._extend.utils.fromDecimal]
		}),
		new web3._extend.Mccmod({
			name: 'setRecommitInterval',
			call: 'miner_setRecommitInterval',
			params: 1,
		}),
		new web3._extend.Mccmod({
			name: 'getHashrate',
			call: 'miner_getHashrate'
		}),
	],
	properties: []
});
`

const NetJs = `
web3._extend({
	property: 'net',
	mccmods: [],
	properties: [
		new web3._extend.Property({
			name: 'version',
			getter: 'net_version'
		}),
	]
});
`

const PersonalJs = `
web3._extend({
	property: 'personal',
	mccmods: [
		new web3._extend.Mccmod({
			name: 'importRawKey',
			call: 'personal_importRawKey',
			params: 2
		}),
		new web3._extend.Mccmod({
			name: 'sign',
			call: 'personal_sign',
			params: 3,
			inputFormatter: [null, web3._extend.formatters.inputAddressFormatter, null]
		}),
		new web3._extend.Mccmod({
			name: 'ecRecover',
			call: 'personal_ecRecover',
			params: 2
		}),
		new web3._extend.Mccmod({
			name: 'openWallet',
			call: 'personal_openWallet',
			params: 2
		}),
		new web3._extend.Mccmod({
			name: 'deriveAccount',
			call: 'personal_deriveAccount',
			params: 3
		}),
		new web3._extend.Mccmod({
			name: 'signTransaction',
			call: 'personal_signTransaction',
			params: 2,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter, null]
		}),
		new web3._extend.Mccmod({
			name: 'unpair',
			call: 'personal_unpair',
			params: 2
		}),
		new web3._extend.Mccmod({
			name: 'initializeWallet',
			call: 'personal_initializeWallet',
			params: 1
		})
	],
	properties: [
		new web3._extend.Property({
			name: 'listWallets',
			getter: 'personal_listWallets'
		}),
	]
})
`

const RpcJs = `
web3._extend({
	property: 'rpc',
	mccmods: [],
	properties: [
		new web3._extend.Property({
			name: 'modules',
			getter: 'rpc_modules'
		}),
	]
});
`

const ShhJs = `
web3._extend({
	property: 'shh',
	mccmods: [
	],
	properties:
	[
		new web3._extend.Property({
			name: 'version',
			getter: 'shh_version',
			outputFormatter: web3._extend.utils.toDecimal
		}),
		new web3._extend.Property({
			name: 'info',
			getter: 'shh_info'
		}),
	]
});
`

const SwarmfsJs = `
web3._extend({
	property: 'swarmfs',
	mccmods:
	[
		new web3._extend.Mccmod({
			name: 'mount',
			call: 'swarmfs_mount',
			params: 2
		}),
		new web3._extend.Mccmod({
			name: 'unmount',
			call: 'swarmfs_unmount',
			params: 1
		}),
		new web3._extend.Mccmod({
			name: 'listmounts',
			call: 'swarmfs_listmounts',
			params: 0
		}),
	]
});
`

const TxpoolJs = `
web3._extend({
	property: 'txpool',
	mccmods: [],
	properties:
	[
		new web3._extend.Property({
			name: 'content',
			getter: 'txpool_content'
		}),
		new web3._extend.Property({
			name: 'inspect',
			getter: 'txpool_inspect'
		}),
		new web3._extend.Property({
			name: 'status',
			getter: 'txpool_status',
			outputFormatter: function(status) {
				status.pending = web3._extend.utils.toDecimal(status.pending);
				status.queued = web3._extend.utils.toDecimal(status.queued);
				return status;
			}
		}),
	]
});
`

const AccountingJs = `
web3._extend({
	property: 'accounting',
	mccmods: [
		new web3._extend.Property({
			name: 'balance',
			getter: 'account_balance'
		}),
		new web3._extend.Property({
			name: 'balanceCredit',
			getter: 'account_balanceCredit'
		}),
		new web3._extend.Property({
			name: 'balanceDebit',
			getter: 'account_balanceDebit'
		}),
		new web3._extend.Property({
			name: 'bytesCredit',
			getter: 'account_bytesCredit'
		}),
		new web3._extend.Property({
			name: 'bytesDebit',
			getter: 'account_bytesDebit'
		}),
		new web3._extend.Property({
			name: 'msgCredit',
			getter: 'account_msgCredit'
		}),
		new web3._extend.Property({
			name: 'msgDebit',
			getter: 'account_msgDebit'
		}),
		new web3._extend.Property({
			name: 'peerDrops',
			getter: 'account_peerDrops'
		}),
		new web3._extend.Property({
			name: 'selfDrops',
			getter: 'account_selfDrops'
		}),
	]
});
`

const LESJs = `
web3._extend({
	property: 'les',
	mccmods:
	[
		new web3._extend.Mccmod({
			name: 'getCheckpoint',
			call: 'les_getCheckpoint',
			params: 1
		}),
	],
	properties:
	[
		new web3._extend.Property({
			name: 'latestCheckpoint',
			getter: 'les_latestCheckpoint'
		}),
		new web3._extend.Property({
			name: 'checkpointContractAddress',
			getter: 'les_getCheckpointContractAddress'
		}),
	]
});
`
