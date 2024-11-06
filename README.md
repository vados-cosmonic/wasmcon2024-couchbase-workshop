# 🛋 WASMCON 2024: Implementing a Wasm Native Database API with Couchbase

Presented by [Couchbase][couchbase] and [Cosmonic][cosmonic] maintainers.

This repository contains everything you need to get started with the workshop:

[on Sched](https://wasmcon24.sched.com/event/ceac3a35d773d3c7498156218db7e22e)

[cosmonic]: https://cosmonic.com/
[couchbase]: https://couchbase.com/

## 👷 Workshop

In this workshop, we'll be building **NuDB**, an experimental new database with an HTTP API,
powered by Couchbase.

By the end of the workshop, we will:

- Examine the HTTP & Couchbase native WebAssembly Interface Type ("WIT") interfaces
- Build a HTTP API [WebAssembly-native database component][wasmcloud-docs-component], powered by [Couchbase][couchbase]
- Run your database implementation locally with [wasmCloud][wasmcloud]
- Run your database in the browser, against a pre-existing test suite

[wasmcloud-docs-component]: https://wasmcloud.com/docs/concepts/components
[wasmCloud]: https://wasmcloud.com

### ✨ NuDB API

NuDB's API is simple, maximizing for developer UX:

| API endpoint                   | Description                             |
|--------------------------------|-----------------------------------------|
| `GET /api/v1/documents`        | Get all documents (paginated)           |
| `GET /api/v1/documents/:id`    | Get a single document                   |
| `PUT /api/v1/documents`        | Insert a single document                |
| `DELETE /api/v1/documents/:id` | Delete a single existing document       |
| `GET /api/v1/documents/latest` | Get the most recently inserted document |

While some of these endpoints are implemented, not all of them are -- that's where you come in!

## 📂 Organization

Here's a run-down of the files in this repository:

| Folder | Description                                                           |
|--------|-----------------------------------------------------------------------|
| `wit`  | WIT definitions we'll be using during the talk                        |
| `nudb` | Code for the Database component we'll finish as part of this workshop |

## 🌎 Environment setup

### GitPod

If using [GitPod][gitpod], you can launch `.gitpod.yml` file in this repository, and make an account at GitPod.io if necessary to get into an environment that "just works"!

[gitpod]: https://gitpod.io

### Devcontainers

To use [GitHub devcontainers][devcontainers] to run this project, you can run

```console
just start-devcontainer
```

> [!NOTE]
> If you prefer to *not* use the `just` task runner, run `devcontainer up --workspace-folder .`
>
> See `Justfile` for more recipes and the commands they run.

[devcontainers]: https://github.com/devcontainers/cli

### Manual/Local

To run manually, ensure that you have the following tools installed:

| Dependency                                      | Description                                               |
|-------------------------------------------------|-----------------------------------------------------------|
| [`just`][just]                                  | Task runner (similar to GNU `make`)                       |
| [`wash`][wash]                                  | WAsmcloud SHell - a tool for managing wasmCloud instances |
| [`wit-deps`][wit-deps]                          | Manual downloading of WIT interfaces                      |
| [`tinygo`][tinygo]                              | [TinyGo][tinygo] binary                                   |

[just]: https://github.com/casey/just
[wash]: https://wasmcloud.com/docs/installation
[tinygo]: https://tinygo.org/
[wit-deps]: https://github.com/bytecodealliance/wit-deps

Once you have the required tooling installed, run the following:

```console
just setup-manual build
```

This step will make sure that required/expected tooling is installed
