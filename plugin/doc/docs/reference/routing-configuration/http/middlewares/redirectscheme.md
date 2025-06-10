---
title: "Mesh RedirectScheme Documentation"
description: "In Mesh Proxy's HTTP middleware, RedirectScheme redirects clients to different schemes/ports. Read the technical documentation."
---

The `RedirectScheme` middleware redirects the request if the request scheme is different from the configured scheme.

!!! warning "When behind another reverse-proxy"

    When there is at least one other reverse-proxy between the client and Mesh, 
    the other reverse-proxy (i.e. the last hop) needs to be a [trusted](../../../install-configuration/entrypoints.md#configuration-options) one. 
    
    Otherwise, Mesh would clean up the X-Forwarded headers coming from this last hop, 
    and as the RedirectScheme middleware relies on them to determine the scheme used,
    it would not function as intended.

## Configuration Examples

```yaml tab="Structured (YAML)"
# Redirect to https
http:
  middlewares:
    test-redirectscheme:
      redirectScheme:
        scheme: https
        permanent: true
```

```toml tab="Structured (TOML)"
# Redirect to https
[http.middlewares]
  [http.middlewares.test-redirectscheme.redirectScheme]
    scheme = "https"
    permanent = true
```

```yaml tab="Labels"
# Redirect to https
labels:
  - "Mesh.http.middlewares.test-redirectscheme.redirectscheme.scheme=https"
  - "Mesh.http.middlewares.test-redirectscheme.redirectscheme.permanent=true"
```

```json tab="Tags"
// Redirect to https
{
  // ...
  "Tags": [
    "Mesh.http.middlewares.test-redirectscheme.redirectscheme.scheme=https"
    "Mesh.http.middlewares.test-redirectscheme.redirectscheme.permanent=true"
  ]
}

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

## Configuration Options

| Field                        | Description                                             | Default | Required |
|:-----------------------------|----------------------------------------------------------|:--------|:---------|
| `scheme` | Scheme of the new URL. | "" | Yes |
| `permanent` | Enable a permanent redirection. | false | No |
| `port` | Port of the new URL.<br />Set a string, **not** a numeric value. | "" | No |
