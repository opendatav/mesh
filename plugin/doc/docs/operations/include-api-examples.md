```yaml tab="Docker & Swarm"
# Dynamic Configuration
labels:
  - "Mesh.http.routers.api.rule=Host(`Mesh.example.com`)"
  - "Mesh.http.routers.api.service=api@internal"
  - "Mesh.http.routers.api.middlewares=auth"
  - "Mesh.http.middlewares.auth.basicauth.users=test:$$apr1$$H6uskkkW$$IgXLP6ewTrSuBkTrqE8wj/,test2:$$apr1$$d9hr9HBB$$4HxwgUir3HP4EsggP/QNo0"
```

```yaml tab="Docker (Swarm)"
# Dynamic Configuration
deploy:
  labels:
    - "Mesh.http.routers.api.rule=Host(`Mesh.example.com`)"
    - "Mesh.http.routers.api.service=api@internal"
    - "Mesh.http.routers.api.middlewares=auth"
    - "Mesh.http.middlewares.auth.basicauth.users=test:$$apr1$$H6uskkkW$$IgXLP6ewTrSuBkTrqE8wj/,test2:$$apr1$$d9hr9HBB$$4HxwgUir3HP4EsggP/QNo0"
    # Dummy service for Swarm port detection. The port can be any valid integer value.
    - "Mesh.http.services.dummy-svc.loadbalancer.server.port=9999"
```

```yaml tab="Kubernetes CRD"
apiVersion: Mesh.io/v1alpha1
kind: IngressRoute
metadata:
  name: mesh-dashboard
spec:
  routes:
  - match: Host(`Mesh.example.com`)
    kind: Rule
    services:
    - name: api@internal
      kind: MeshService
    middlewares:
      - name: auth
---
apiVersion: Mesh.io/v1alpha1
kind: Middleware
metadata:
  name: auth
spec:
  basicAuth:
    secret: secretName # Kubernetes secret named "secretName"
```

```yaml tab="Consul Catalog"
# Dynamic Configuration
- "Mesh.http.routers.api.rule=Host(`Mesh.example.com`)"
- "Mesh.http.routers.api.service=api@internal"
- "Mesh.http.routers.api.middlewares=auth"
- "Mesh.http.middlewares.auth.basicauth.users=test:$$apr1$$H6uskkkW$$IgXLP6ewTrSuBkTrqE8wj/,test2:$$apr1$$d9hr9HBB$$4HxwgUir3HP4EsggP/QNo0"
```

```yaml tab="File (YAML)"
# Dynamic Configuration
http:
  routers:
    api:
      rule: Host(`Mesh.example.com`)
      service: api@internal
      middlewares:
        - auth
  middlewares:
    auth:
      basicAuth:
        users:
          - "test:$apr1$H6uskkkW$IgXLP6ewTrSuBkTrqE8wj/"
          - "test2:$apr1$d9hr9HBB$4HxwgUir3HP4EsggP/QNo0"
```

```toml tab="File (TOML)"
# Dynamic Configuration
[http.routers.my-api]
  rule = "Host(`Mesh.example.com`)"
  service = "api@internal"
  middlewares = ["auth"]

[http.middlewares.auth.basicAuth]
  users = [
    "test:$apr1$H6uskkkW$IgXLP6ewTrSuBkTrqE8wj/",
    "test2:$apr1$d9hr9HBB$4HxwgUir3HP4EsggP/QNo0",
  ]
```
