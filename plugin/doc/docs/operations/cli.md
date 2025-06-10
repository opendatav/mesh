---
title: "Mesh CLI Documentation"
description: "Learn the basics of the Mesh Proxy command line interface (CLI). Read the technical documentation."
---

# CLI

The Mesh Command Line

## General

```bash
Mesh [command] [flags] [arguments]
```

Use `Mesh [command] --help` for help on any command.

Commands:

- `healthcheck` Calls Mesh `/ping` to check the health of Mesh (the API must be enabled).
- `version` Shows the current Mesh version.

Flag's usage:

```bash
# set flag_argument to flag(s)
Mesh [--flag=flag_argument] [-f [flag_argument]]

# set true/false to boolean flag(s)
Mesh [--flag[=true|false| ]] [-f [true|false| ]]
```

All flags are documented in the [(static configuration) CLI reference](../reference/static-configuration/cli.md).

!!! info "Flags are case-insensitive."

### `healthcheck`

Calls Mesh `/ping` to check the health of Mesh.
Its exit status is `0` if Mesh is healthy and `1` otherwise.

This can be used with Docker [HEALTHCHECK](https://docs.docker.com/engine/reference/builder/#healthcheck) instruction
or any other health check orchestration mechanism.

!!! info
    The [`ping` endpoint](../operations/ping.md) must be enabled to allow the `healthcheck` command to call `/ping`.

Usage:

```bash
Mesh healthcheck [command] [flags] [arguments]
```

Example:

```bash
$ Mesh healthcheck
OK: http://:8082/ping
```

### `version`

Shows the current Mesh version.

Usage:

```bash
Mesh version
```
