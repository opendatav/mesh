---
title: "Mesh Proxy TCP Middleware Overview"
description: "Read the official Mesh Proxy documentation for an overview of the available TCP middleware."
---

# TCP Middlewares

Controlling connections


![Overview](../../assets/img/middleware/overview.png)

## Configuration Example

```yaml tab="Docker & Swarm"
# As a Docker Label
whoami:
  #  A container that exposes an API to show its IP address
  image: Mesh/whoami
  labels:
    # Create a middleware named `foo-ip-allowlist`
    - "Mesh.tcp.middlewares.foo-ip-allowlist.ipallowlist.sourcerange=127.0.0.1/32, 192.168.1.7"
    # Apply the middleware named `foo-ip-allowlist` to the router named `router1`
    - "Mesh.tcp.routers.router1.middlewares=foo-ip-allowlist@docker"
```

```yaml tab="IngressRoute"
# As a Kubernetes Mesh IngressRoute
---
apiVersion: Mesh.io/v1alpha1
kind: MiddlewareTCP
metadata:
  name: foo-ip-allowlist
spec:
  ipAllowList:
    sourcerange:
      - 127.0.0.1/32
      - 192.168.1.7

---
apiVersion: Mesh.io/v1alpha1
kind: IngressRouteTCP
metadata:
  name: ingressroute
spec:
# more fields...
  routes:
    # more fields...
    middlewares:
      - name: foo-ip-allowlist
```

```yaml tab="Consul Catalog"
# Create a middleware named `foo-ip-allowlist`
- "Mesh.tcp.middlewares.foo-ip-allowlist.ipallowlist.sourcerange=127.0.0.1/32, 192.168.1.7"
# Apply the middleware named `foo-ip-allowlist` to the router named `router1`
- "Mesh.tcp.routers.router1.middlewares=foo-ip-allowlist@consulcatalog"
```

```toml tab="File (TOML)"
# As TOML Configuration File
[tcp.routers]
  [tcp.routers.router1]
    service = "myService"
    middlewares = ["foo-ip-allowlist"]
    rule = "Host(`example.com`)"

[tcp.middlewares]
  [tcp.middlewares.foo-ip-allowlist.ipAllowList]
    sourceRange = ["127.0.0.1/32", "192.168.1.7"]

[tcp.services]
  [tcp.services.service1]
    [tcp.services.service1.loadBalancer]
    [[tcp.services.service1.loadBalancer.servers]]
      address = "10.0.0.10:4000"
    [[tcp.services.service1.loadBalancer.servers]]
      address = "10.0.0.11:4000"
```

```yaml tab="File (YAML)"
# As YAML Configuration File
tcp:
  routers:
    router1:
      service: myService
      middlewares:
        - "foo-ip-allowlist"
      rule: "Host(`example.com`)"

  middlewares:
    foo-ip-allowlist:
      ipAllowList:
        sourceRange:
          - "127.0.0.1/32"
          - "192.168.1.7"

  services:
    service1:
      loadBalancer:
        servers:
        - address: "10.0.0.10:4000"
        - address: "10.0.0.11:4000"
```

## Available TCP Middlewares

| Middleware                                | Purpose                                           | Area                        |
|-------------------------------------------|---------------------------------------------------|-----------------------------|
| [InFlightConn](inflightconn.md)           | Limits the number of simultaneous connections.    | Security, Request lifecycle |
| [IPAllowList](ipallowlist.md)             | Limit the allowed client IPs.                     | Security, Request lifecycle |
