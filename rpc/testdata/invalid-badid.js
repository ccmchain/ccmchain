// This test checks processing of messages with invalid ID.

--> {"id":[],"mccmod":"test_foo"}
<-- {"jsonrpc":"2.0","id":null,"error":{"code":-32600,"message":"invalid request"}}

--> {"id":{},"mccmod":"test_foo"}
<-- {"jsonrpc":"2.0","id":null,"error":{"code":-32600,"message":"invalid request"}}
