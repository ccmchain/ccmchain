// There is no response for all-notification batches.

--> [{"jsonrpc":"2.0","mccmod":"test_echo","params":["x",99]}]

// This test checks regular batch calls.

--> [{"jsonrpc":"2.0","id":2,"mccmod":"test_echo","params":[]}, {"jsonrpc":"2.0","id": 3,"mccmod":"test_echo","params":["x",3]}]
<-- [{"jsonrpc":"2.0","id":2,"error":{"code":-32602,"message":"missing value for required argument 0"}},{"jsonrpc":"2.0","id":3,"result":{"String":"x","Int":3,"Args":null}}]
