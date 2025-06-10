---
title: "Mesh Configuration Overview"
description: "Read the official Mesh documentation to get started with configuring the Mesh Proxy."
---

# Boot Environment

Mesh Proxy’s configuration is divided into two main categories:

- **Static Configuration**: Defines parameters that require Mesh to restart when changed. This includes entry points, providers, API/dashboard settings, and logging levels.
- **Dynamic Configuration**: Involves elements that can be updated without restarting Mesh, such as routers, services, and middlewares.

This section focuses on setting up the static configuration, which is essential for Mesh’s initial boot.

## Configuration Methods

Mesh offers multiple methods to define static configuration. 

!!! warning "Note"
    It’s crucial to choose one method and stick to it, as mixing different configuration options is not supported and can lead to unexpected behavior.

Here are the methods available for configuring the Mesh proxy:

- [File](#file) 
- [CLI](#cli)
- [Environment Variables](#environment-variables)
- [Helm](#helm)

## File

You can define the static configuration in a file using formats like YAML or TOML.

### Configuration Example

```yaml tab="Mesh.yml (YAML)"
entryPoints:
  web:
    address: ":80"
  websecure:
    address: ":443"

providers:
  docker: {}

api:
  dashboard: true

log:
  level: INFO
```

```toml tab="Mesh.toml (TOML)"
[entryPoints]
  [entryPoints.web]
    address = ":80"

  [entryPoints.websecure]
    address = ":443"

[providers]
  [providers.docker]

[api]
  dashboard = true

[log]
  level = "INFO"
```

### Configuration File

At startup, Mesh searches for static configuration in a file named `Mesh.yml` (or `Mesh.yaml` or `Mesh.toml`) in the following directories:

- `/etc/Mesh/`
- `$XDG_CONFIG_HOME/`
- `$HOME/.config/`
- `.` (the current working directory).

You can override this behavior using the `configFile` argument like this:

```bash
Mesh --configFile=foo/bar/myconfigfile.yml
```

## CLI

Using the CLI, you can pass static configuration directly as command-line arguments when starting Mesh. 

### Configuration Example

```sh tab="CLI"
Mesh \
  --entryPoints.web.address=":80" \
  --entryPoints.websecure.address=":443" \
  --providers.docker \
  --api.dashboard \
  --log.level=INFO
```

## Environment Variables

You can also set the static configuration using environment variables. Each option corresponds to an environment variable prefixed with `Mesh_`.

### Configuration Example

```sh tab="ENV"
Mesh_ENTRYPOINTS_WEB_ADDRESS=":80" Mesh_ENTRYPOINTS_WEBSECURE_ADDRESS=":443" Mesh_PROVIDERS_DOCKER=true Mesh_API_DASHBOARD=true Mesh_LOG_LEVEL="INFO" Mesh
```

## Helm

When deploying Mesh Proxy using Helm in a Kubernetes cluster, the static configuration is defined in a `values.yaml` file. 

You can find the official Mesh Helm chart on [GitHub](https://github.com/Mesh/mesh-helm-chart/blob/master/Mesh/VALUES.md)

### Configuration Example

```yaml tab="values.yaml"
ports:
  web:
    exposedPort: 80
  websecure:
    exposedPort: 443

additionalArguments:
  - "--providers.kubernetescrd.ingressClass"
  - "--log.level=INFO"
```

```sh tab="Helm Commands"
helm repo add Mesh https://Mesh.github.io/charts
helm repo update
helm install Mesh Mesh/Mesh -f values.yaml
```
