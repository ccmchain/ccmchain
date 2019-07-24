// This test calls a mccmod that doesn't exist.

--> {"jsonrpc": "2.0", "id": 2, "mccmod": "invalid_mccmod", "params": [2, 3]}
<-- {"jsonrpc":"2.0","id":2,"error":{"code":-32601,"message":"the mccmod invalid_mccmod does not exist/is not available"}}
