# ✨ NuBase

NuBase is a WebAssembly-native database, exposed via HTTP.

Importantly, this [WebAssembly Component][wasm-components] `import`s the `wasmcoud:couchbase/documents`
interface (see [related WIT](../wit/nubase.wit)), and makes that functionality available via HTTP
by exporting the [WASI][wasi] standard [`wasi:http/incoming-handler`][wasi-http].

[couchbase]: https://couchbase.com
[wit]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/WIT.md
[wasi-http]: https://github.com/WebAssembly/wasi-http
[wasm-components]: https://component-model.bytecodealliance.org/
[wasi]: https://wasi.dev

### 📚 NuBase HTTP API

NuBase's API is simple, maximizing for developer UX, exposing simple
document-level operations -- a "baby's first [Couchbase][couchbase] API", making use of the full
power of Couchbase underneath:

| API endpoint                        | Description                                     |
|-------------------------------------|-------------------------------------------------|
| `GET /api/v1/_status`               | Basic status endpoint (returns "ok") if present |
| `GET /api/v1/documents`             | Get all documents (paginated)                   |
| `GET /api/v1/documents/:id`         | Get a single document                           |
| `PUT /api/v1/documents`             | Insert a single document                        |
| `POST /api/v1/documents`            | Insert a single document                        |
| `POST /api/v1/batch/insert`         | Insert multiple documents                       |
| `DELETE /api/v1/documents/:id`      | Delete a single existing document               |
| `POST /api/v1/documents/:id/upsert` | Upsert a single document                        |
| `POST /api/v1/batch/upsert`         | Perform multiple upserts                        |

**The implementation for this component is partially finished, but you'll have to write the rest!**.

## 📦 Dependencies

| Tool         | Description                                                             |
|--------------|-------------------------------------------------------------------------|
| `go`         | [Golang][golang] toolchain                                              |
| `tinygo`     | [TinyGo][tinygo] toolchain                                              |
| `wash`       | [WAsmcloud][wasmcloud] SHell, a tool for running WebAssembly components |
| `wasm-tools` | [WebAssembly toolkit][wasm-tools] (used during code generation)         |

[wasmcloud]: https://wasmcloud.com/docs
[golang]: https://go.dev
[tinygo]: https://tinygo.org
[wash]: https://wasmcloud.com/docs/ecosystem/wash/
[wasm-tools]: https://github.com/bytecodealliance/wasm-tools

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

## Local Usage guide

Once `wash dev` is running, you should be able to simply `curl` this component:

```console
curl localhost:8080/api/v1/_status
```

Let's dive into some more usage patterns.

### Add a document

Once you're sure it's running, you can send JSON payloads over:

```console
curl localhost:8080/api/v1/documents/1 --data '{"example": "document"}'
```

### Retrieve a document

You should get back the ID of the document you inserted -- you can feed that ID into another CURL to retrieve the document we just wrote (by ID):

```console
curl localhost:8080/api/v1/documents/1
```

### Upsert a document

We can insert or update-with-[JSON-patch][json-patch] a document:

```console
curl localhost:8080/api/v1/documents/1/upsert --data-binary @- <<EOF
{
  "docId": "1",
  "patches": [
    { "op": "replace", "path": "/example", "value": "replaced" }
  ],
  "insert": {
    "example": "document"
  }
}
EOF
```

[json-patch]: https://jsonpatch.com/

### Batch insert a document

We can insert multiple documents at the same time:

```console
curl localhost:8080/api/v1/batch/insert --data-binary @- <<EOF
{
  "docs": [
    {
      "docId": "1",
      "doc": {"example": "document"}
    },
    {
      "docId": "2",
      "doc": {"example": "document-2"}
    }
  ]
}
EOF
```

### Delete an existing document

You can delete the document you just created as well:

```console
curl -X DELETE localhost:8080/api/v1/documents/1
```

### Batch upsert a document

**TODO: Fill this out!**

```console
TODO
```

[wasmcloud-docs-provider]: https://wasmcloud.com/docs/concepts/providers
[wash-dev]: https://wasmcloud.com/docs/cli/#wash-dev
