---
title: "Mesh RedirectScheme Documentation"
description: "In Mesh Proxy's HTTP middleware, RedirectScheme redirects clients to different schemes/ports. Read the technical documentation."
---

# RedirectScheme

Redirecting the Client to a Different Scheme/Port


<!--
TODO: add schema
-->

The RedirectScheme middleware redirects the request if the request scheme is different from the configured scheme.

!!! warning "When behind another reverse-proxy"

    When there is at least one other reverse-proxy between the client and Mesh, 
    the other reverse-proxy (i.e. the last hop) needs to be a [trusted](../../routing/entrypoints.md#forwarded-headers) one. 
    
    Otherwise, Mesh would clean up the X-Forwarded headers coming from this last hop, 
    and as the RedirectScheme middleware relies on them to determine the scheme used,
    it would not function as intended.

## Configuration Examples

```yaml tab="Docker & Swarm"
# Redirect to https
labels:
  - "Mesh.http.middlewares.test-redirectscheme.redirectscheme.scheme=https"
  - "Mesh.http.middlewares.test-redirectscheme.redirectscheme.permanent=true"
```

```yaml tab="Kubernetes"
# Redirect to https
apiVersion: Mesh.io/v1alpha1
kind: Middleware
metadata:
  name: test-redirectscheme
spec:
  redirectScheme:
    scheme: https
    permanent: true
```

```yaml tab="Consul Catalog"
# Redirect to https
labels:
  - "Mesh.http.middlewares.test-redirectscheme.redirectscheme.scheme=https"
  - "Mesh.http.middlewares.test-redirectscheme.redirectscheme.permanent=true"
```

```yaml tab="File (YAML)"
# Redirect to https
http:
  middlewares:
    test-redirectscheme:
      redirectScheme:
        scheme: https
        permanent: true
```

```toml tab="File (TOML)"
# Redirect to https
[http.middlewares]
  [http.middlewares.test-redirectscheme.redirectScheme]
    scheme = "https"
    permanent = true
```

## Configuration Options

### `permanent`

Set the `permanent` option to `true` to apply a permanent redirection.

```yaml tab="Docker & Swarm"
# Redirect to https
labels:
  # ...
  - "Mesh.http.middlewares.test-redirectscheme.redirectscheme.permanent=true"
```

```yaml tab="Kubernetes"
# Redirect to https
apiVersion: Mesh.io/v1alpha1
kind: Middleware
metadata:
  name: test-redirectscheme
spec:
  redirectScheme:
    # ...
    permanent: true
```

```yaml tab="Consul Catalog"
# Redirect to https
labels:
  # ...
  - "Mesh.http.middlewares.test-redirectscheme.redirectscheme.permanent=true"
```

```yaml tab="File (YAML)"
# Redirect to https
http:
  middlewares:
    test-redirectscheme:
      redirectScheme:
        # ...
        permanent: true
```

```toml tab="File (TOML)"
# Redirect to https
[http.middlewares]
  [http.middlewares.test-redirectscheme.redirectScheme]
    # ...
    permanent = true
```

### `scheme`

The `scheme` option defines the scheme of the new URL.

```yaml tab="Docker & Swarm"
# Redirect to https
labels:
  - "Mesh.http.middlewares.test-redirectscheme.redirectscheme.scheme=https"
```

```yaml tab="Kubernetes"
# Redirect to https
apiVersion: Mesh.io/v1alpha1
kind: Middleware
metadata:
  name: test-redirectscheme
spec:
  redirectScheme:
    scheme: https
```

```yaml tab="Consul Catalog"
# Redirect to https
labels:
  - "Mesh.http.middlewares.test-redirectscheme.redirectscheme.scheme=https"
```

```yaml tab="File (YAML)"
# Redirect to https
http:
  middlewares:
    test-redirectscheme:
      redirectScheme:
        scheme: https
```

```toml tab="File (TOML)"
# Redirect to https
[http.middlewares]
  [http.middlewares.test-redirectscheme.redirectScheme]
    scheme = "https"
```

### `port`

The `port` option defines the port of the new URL.

```yaml tab="Docker & Swarm"
# Redirect to https
labels:
  # ...
  - "Mesh.http.middlewares.test-redirectscheme.redirectscheme.port=443"
```

```yaml tab="Kubernetes"
# Redirect to https
apiVersion: Mesh.io/v1alpha1
kind: Middleware
metadata:
  name: test-redirectscheme
spec:
  redirectScheme:
    # ...
    port: "443"
```

```yaml tab="Consul Catalog"
# Redirect to https
labels:
  # ...
  - "Mesh.http.middlewares.test-redirectscheme.redirectscheme.port=443"
```

```yaml tab="File (YAML)"
# Redirect to https
http:
  middlewares:
    test-redirectscheme:
      redirectScheme:
        # ...
        port: "443"
```

```toml tab="File (TOML)"
# Redirect to https
[http.middlewares]
  [http.middlewares.test-redirectscheme.redirectScheme]
    # ...
    port = 443
```

!!! info "Port in this configuration is a string, not a numeric value."
