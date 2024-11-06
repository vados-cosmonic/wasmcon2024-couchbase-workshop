# NuDB

This component implements a HTTP-API driven WebAssembly-native Database powered by [Couchbase][couchbase].

The interface for this database is written in the [WebAssembly Interface Types ("WIT")][wit] (see [`nudb.wit`](../wit/nudb.wit)).

Importantly, the interface `import`s `wasmcloud:couchbase` for access to Couchbase, and exports [`wasi:http/incoming-handler`][wasi-http],
which enables providing an interface driven by HTTP.

**This implementation of this component is *incomplete***. Your mission, should you choose to accept it, is to complete the implementation for the database interface.

[couchbase]: https://couchbase.com
[wit]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/WIT.md
[wasi-http]: https://github.com/WebAssembly/wasi-http

## Prerequisites

| Tool     | Description                                                             |
|----------|-------------------------------------------------------------------------|
| `go`     | [Golang][golang] toolchain                                                        |
| `tinygo` | [TinyGo][tinygo] toolchain                                              |
| `wash`   | [WAsmcloud][wasmcloud] SHell, a tool for running WebAssembly components |

[wasmcloud]: https://wasmcloud.com/docs
[golang]: https://go.dev
[tinygo]: https://tinygo.org
[wash]: https://wasmcloud.com/docs/ecosystem/wash/

## Quickstart

### Development

To develop this component, we can use [`wash dev`][wash-dev], a subcommand of the WAsmcloud SHell:

```bash
wash dev
```

> [!NOTE]
> Want to do things the hard way? Read [`./docs/build-the-hard-way.md`](./docs/build-the-hard-way.md)

After starting `wash dev`, you'll *automatically* have access to a [wasmCloud provider][wasmcloud-docs-provider] which
starts a local HTTP server that you can use to access your component.

Once `wash dev` has started, you should be able to simply `curl` your component:

```console
curl localhost:8080/api/v1/_status
```

[wasmcloud-docs-provider]: https://wasmcloud.com/docs/concepts/providers
[wash-dev]: https://wasmcloud.com/docs/cli/#wash-dev

### Build Build Build!

Once you're up and running with `wash dev` and can access the serveer, you can start implementing the
missing API functionality (marked by `// TODO`s) in `main.go`!

As a reminder, the APIs we'll be implementing are the following:

| API endpoint                   | Description                             |
|--------------------------------|-----------------------------------------|
| `GET /api/v1/documents`        | Get all documents (paginated)           |
| `GET /api/v1/documents/:id`    | Get a single document                   |
| `PUT /api/v1/documents`        | Insert a single document                |
| `DELETE /api/v1/documents/:id` | Delete a single existing document       |
| `GET /api/v1/documents/latest` | Get the most recently inserted document |
