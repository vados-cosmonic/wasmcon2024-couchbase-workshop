# 🛋 WASMCON 2024: Implementing a Wasm Native Database API with Couchbase

Presented by [Couchbase][couchbase] and [Cosmonic][cosmonic] maintainers.

This repository contains everything you need to get started with the workshop:

[on Sched](https://wasmcon24.sched.com/event/ceac3a35d773d3c7498156218db7e22e)

[cosmonic]: https://cosmonic.com/
[couchbase]: https://couchbase.com/

## 👷 Workshop

In this workshop, we'll be building **NuBase**, an experimental new database with an HTTP API,
powered by Couchbase.

By the end of the workshop, we will:

- Examine the HTTP & Couchbase native WebAssembly Interface Type ("WIT") interfaces
- Build a HTTP API [WebAssembly-native database component][wasmcloud-docs-component], powered by [Couchbase][couchbase]
- Run your database implementation locally with [wasmCloud][wasmcloud]

[wasmcloud-docs-component]: https://wasmcloud.com/docs/concepts/components
[wasmCloud]: https://wasmcloud.com

## 📂 Organization

Here's a run-down of the files in this repository:

| Folder            | Description                                                           |
|-------------------|-----------------------------------------------------------------------|
| `wit`             | WIT definitions we'll be using during the talk                        |
| `nubase`          | Code for the Database component we'll finish as part of this workshop |

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
> If you prefer to *not* use [`just`][just], run `devcontainer up --workspace-folder .`
>
> See [`Justfile`](./Justfile) for more recipes and the commands they run.

[devcontainers]: https://github.com/devcontainers/cli

### Manual/Local

To run manually, ensure that you have the following tools installed:

| Dependency             | Description                                               |
|------------------------|-----------------------------------------------------------|
| [`just`][just]         | Task runner (similar to GNU `make`)                       |
| [`wash`][wash]         | WAsmcloud SHell - a tool for managing wasmCloud instances |
| [`wit-deps`][wit-deps] | Manual downloading of WIT interfaces                      |
| [`tinygo`][tinygo]     | [TinyGo][tinygo] toolchain                                |
| [`go`][go]             | [Go][go] toolchain                                        |

[just]: https://github.com/casey/just
[wash]: https://wasmcloud.com/docs/installation
[tinygo]: https://tinygo.org/
[go]: https://go.dev/
[wit-deps]: https://github.com/bytecodealliance/wit-deps

You can easily check which tools are not installed by running:

```console
just setup-manual
```

Once you have all required tools installed, you can build the project by running:

```console
just build
```

Then, start Couchbase, using `docker compose`:

```console
just start-couchbase
```

From here, you can start developing the component:

```console
just dev
```
