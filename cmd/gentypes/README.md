The schema for DAP messages is defined in JSON at
https://github.com/microsoft/vscode-debugadapter-node/blob/master/debugProtocol.json

In this directory we have a copy of the schema, which is licensed by Microsoft
with a [MIT
License](https://github.com/microsoft/vscode-debugadapter-node/blob/master/License.txt)

To generate Go types from the schema, run:

```
$ go run cmd/gentypes/gentypes.go cmd/gentypes/debugProtocol.json > types.go
```

The generated ``types.go`` is also checked in, so there is no need to regenerate
it unless the schema changes.
