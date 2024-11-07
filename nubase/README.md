# ✨ NuBase

NuBase is a WebAssembly-native database, exposed via HTTP.

Importantly, this [WebAssembly Component][wasm-components] `import`s the `wasmcoud:couchbase/documents`
interface (see [related WIT](../wit/nubase.wit)), and makes that functionality available via HTTP
by exporting the [WASI][wasi] standard [`wasi:http/incoming-handler`][wasi-http].

[couchbase]: https://couchbase.com
[wit]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/WIT.md
[wasi-http]: https://github.com/WebAssembly/wasi-http
[wasm-components]: https://component-model.bytecodealliance.org/

### 📚 NuBase HTTP API

NuBase's API is simple, maximizing for developer UX, exposing simple
document-level operations -- a "baby's first [Couchbase][couchbase] API", making use of the full
power of Couchbase underneath:

| API endpoint                   | Description                                     |
|--------------------------------|-------------------------------------------------|
| `GET /api/v1/_status`          | Basic status endpoint (returns "ok") if present |
| `GET /api/v1/documents`        | Get all documents (paginated)                   |
| `GET /api/v1/documents/:id`    | Get a single document                           |
| `PUT /api/v1/documents`        | Insert a single document                        |
| `DELETE /api/v1/documents/:id` | Delete a single existing document               |

**The implementation for this component is partially finished, but you'll have to write the rest!**.

## 📦 Dependencies

| Tool     | Description                                                             |
|----------|-------------------------------------------------------------------------|
| `go`     | [Golang][golang] toolchain                                              |
| `tinygo` | [TinyGo][tinygo] toolchain                                              |
| `wash`   | [WAsmcloud][wasmcloud] SHell, a tool for running WebAssembly components |

[wasmcloud]: https://wasmcloud.com/docs
[golang]: https://go.dev
[tinygo]: https://tinygo.org
[wash]: https://wasmcloud.com/docs/ecosystem/wash/

## 👟 Quickstart

### Setup

Once you have the relevant dependencies installed, first run `go mod download`. This will install relevant
golang tooling (see `tools.go`) for working with this repository.

```console
go mod download
```

### Development

To develop this component, we can use [`wash dev`][wash-dev], a subcommand of the WAsmcloud SHell:

```bash
wash dev
```

After starting `wash dev`, you'll *automatically* have access to a [wasmCloud provider][wasmcloud-docs-provider] which
starts a local HTTP server that you can use to access this component.

Once `wash dev` has started, you should be able to simply `curl` this component:

```console
curl localhost:8080/api/v1/_status
```

> [!WARN]
> If using `wash dev`, check your output for the address of the HTTP server -- it may be `localhost:8000`

Once you're sure it's running, you can send JSON payloads over:

```console
curl localhost:8080/api/v1/documents --data '{"example": "document"}'
```

You should get back the ID of the document you inserted -- you can feed that ID into another CURL to retrieve the document:

```console
curl localhost:8080/api/v1/documents/<id>
```

[wasmcloud-docs-provider]: https://wasmcloud.com/docs/concepts/providers
[wash-dev]: https://wasmcloud.com/docs/cli/#wash-dev
