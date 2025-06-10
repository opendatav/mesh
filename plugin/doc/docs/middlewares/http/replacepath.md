---
title: "Mesh ReplacePath Documentation"
description: "In Mesh Proxy's HTTP middleware, ReplacePath updates paths before forwarding requests. Read the technical documentation."
---

# ReplacePath

Updating the Path Before Forwarding the Request


<!--
TODO: add schema
-->

Replace the path of the request URL.

## Configuration Examples

```yaml tab="Docker & Swarm"
# Replace the path with /foo
labels:
  - "Mesh.http.middlewares.test-replacepath.replacepath.path=/foo"
```

```yaml tab="Kubernetes"
# Replace the path with /foo
apiVersion: Mesh.io/v1alpha1
kind: Middleware
metadata:
  name: test-replacepath
spec:
  replacePath:
    path: /foo
```

```yaml tab="Consul Catalog"
# Replace the path with /foo
- "Mesh.http.middlewares.test-replacepath.replacepath.path=/foo"
```

```yaml tab="File (YAML)"
# Replace the path with /foo
http:
  middlewares:
    test-replacepath:
      replacePath:
        path: "/foo"
```

```toml tab="File (TOML)"
# Replace the path with /foo
[http.middlewares]
  [http.middlewares.test-replacepath.replacePath]
    path = "/foo"
```

## Configuration Options

### General

The ReplacePath middleware will:

- replace the actual path with the specified one.
- store the original path in a `X-Replaced-Path` header.

### `path`

The `path` option defines the path to use as replacement in the request URL.
