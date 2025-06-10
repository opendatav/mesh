---
title: "Mesh Docker Routing Documentation"
description: "This guide will teach you how to attach labels to your containers, to route traffic and load balance with Mesh and Docker."
---

# Mesh & Docker

A Story of Labels & Containers

![Docker](../../assets/img/providers/docker.png)

Attach labels to your containers and let Mesh do the rest!

One of the best feature of Mesh is to delegate the routing configuration to the application level.
With Docker, Mesh can leverage labels attached to a container to generate routing rules.

!!! warning "Labels & sensitive data"

    We recommend to *not* use labels to store sensitive data (certificates, credentials, etc).
    Instead, we recommend to store sensitive data in a safer storage (secrets, file, etc).

## Configuration Examples

??? example "Configuring Docker & Deploying / Exposing one Service"

    Enabling the docker provider

    ```yaml tab="File (YAML)"
    providers:
      docker: {}
    ```

    ```toml tab="File (TOML)"
    [providers.docker]
    ```

    ```bash tab="CLI"
    --providers.docker=true
    ```

    Attaching labels to containers (in your docker compose file)

    ```yaml
    version: "3"
    services:
      my-container:
        # ...
        labels:
          - Mesh.http.routers.my-container.rule=Host(`example.com`)
    ```

??? example "Specify a Custom Port for the Container"

    Forward requests for `http://example.com` to `http://<private IP of container>:12345`:

    ```yaml
    version: "3"
    services:
      my-container:
        # ...
        labels:
          - Mesh.http.routers.my-container.rule=Host(`example.com`)
          # Tell Mesh to use the port 12345 to connect to `my-container`
          - Mesh.http.services.my-service.loadbalancer.server.port=12345
    ```

    !!! important "Mesh Connecting to the Wrong Port: `HTTP/502 Gateway Error`"
        By default, Mesh uses the first exposed port of a container.

        Setting the label `Mesh.http.services.xxx.loadbalancer.server.port`
        overrides that behavior.

??? example "Specifying more than one router and service per container"

    Forwarding requests to more than one port on a container requires referencing the service loadbalancer port definition using the service parameter on the router.

    In this example, requests are forwarded for `http://example-a.com` to `http://<private IP of container>:8000` in addition to `http://example-b.com` forwarding to `http://<private IP of container>:9000`:

    ```yaml
    version: "3"
    services:
      my-container:
        # ...
        labels:
          - Mesh.http.routers.www-router.rule=Host(`example-a.com`)
          - Mesh.http.routers.www-router.service=www-service
          - Mesh.http.services.www-service.loadbalancer.server.port=8000
          - Mesh.http.routers.admin-router.rule=Host(`example-b.com`)
          - Mesh.http.routers.admin-router.service=admin-service
          - Mesh.http.services.admin-service.loadbalancer.server.port=9000
    ```

## Routing Configuration

!!! info "Labels"

    - Labels are case-insensitive.
    - The complete list of labels can be found in [the reference page](../../reference/dynamic-configuration/docker.md).

### General

Mesh creates, for each container, a corresponding [service](../services/index.md) and [router](../routers/index.md).

The Service automatically gets a server per instance of the container,
and the router automatically gets a rule defined by `defaultRule` (if no rule for it was defined in labels).

#### Service definition


??? example "Automatic assignment with one Service"

    With labels in a compose file

    ```yaml
    labels:
      - "Mesh.http.routers.myproxy.rule=Host(`example.net`)"
      # service myservice gets automatically assigned to router myproxy
      - "Mesh.http.services.myservice.loadbalancer.server.port=80"
    ```

??? example "Automatic service creation with one Router"

    With labels in a compose file

    ```yaml
    labels:
      # no service specified or defined and yet one gets automatically created
      # and assigned to router myproxy.
      - "Mesh.http.routers.myproxy.rule=Host(`example.net`)"
    ```

??? example "Explicit definition with one Service"

    With labels in a compose file

    ```yaml
    labels:
      - Mesh.http.routers.www-router.rule=Host(`example-a.com`)
      # Explicit link between the router and the service
      - Mesh.http.routers.www-router.service=www-service
      - Mesh.http.services.www-service.loadbalancer.server.port=8000
    ```

### Routers

To update the configuration of the Router automatically attached to the container,
add labels starting with `Mesh.http.routers.<name-of-your-choice>.` and followed by the option you want to change.

For example, to change the rule, you could add the label ```Mesh.http.routers.my-container.rule=Host(`example.com`)```.

!!! warning "The character `@` is not authorized in the router name `<router_name>`."

??? info "`Mesh.http.routers.<router_name>.rule`"

    See [rule](../routers/index.md#rule) for more information.

    ```yaml
    - "Mesh.http.routers.myrouter.rule=Host(`example.com`)"
    ```

??? info "`Mesh.http.routers.<router_name>.entrypoints`"

    See [entry points](../routers/index.md#entrypoints) for more information.

    ```yaml
    - "Mesh.http.routers.myrouter.entrypoints=ep1,ep2"
    ```

??? info "`Mesh.http.routers.<router_name>.middlewares`"

    See [middlewares](../routers/index.md#middlewares) and [middlewares overview](../../middlewares/overview.md) for more information.

    ```yaml
    - "Mesh.http.routers.myrouter.middlewares=auth,prefix,cb"
    ```

??? info "`Mesh.http.routers.<router_name>.service`"

    See [service](../routers/index.md#service) for more information.

    ```yaml
    - "Mesh.http.routers.myrouter.service=myservice"
    ```

??? info "`Mesh.http.routers.<router_name>.tls`"

    See [tls](../routers/index.md#tls) for more information.

    ```yaml
    - "Mesh.http.routers.myrouter.tls=true"
    ```

??? info "`Mesh.http.routers.<router_name>.tls.certresolver`"

    See [certResolver](../routers/index.md#certresolver) for more information.

    ```yaml
    - "Mesh.http.routers.myrouter.tls.certresolver=myresolver"
    ```

??? info "`Mesh.http.routers.<router_name>.tls.domains[n].main`"

    See [domains](../routers/index.md#domains) for more information.

    ```yaml
    - "Mesh.http.routers.myrouter.tls.domains[0].main=example.org"
    ```

??? info "`Mesh.http.routers.<router_name>.tls.domains[n].sans`"

    See [domains](../routers/index.md#domains) for more information.

    ```yaml
    - "Mesh.http.routers.myrouter.tls.domains[0].sans=test.example.org,dev.example.org"
    ```

??? info "`Mesh.http.routers.<router_name>.tls.options`"

    See [options](../routers/index.md#options) for more information.

    ```yaml
    - "Mesh.http.routers.myrouter.tls.options=foobar"
    ```

??? info "`Mesh.http.routers.<router_name>.observability.accesslogs`"

    See accesslogs [option](../routers/index.md#accesslogs) for more information.
    
    ```yaml
    - "Mesh.http.routers.myrouter.observability.accesslogs=true"
    ```

??? info "`Mesh.http.routers.<router_name>.observability.metrics`"

    See metrics [option](../routers/index.md#metrics) for more information.
    
    ```yaml
    - "Mesh.http.routers.myrouter.observability.metrics=true"
    ```

??? info "`Mesh.http.routers.<router_name>.observability.tracing`"

    See tracing [option](../routers/index.md#tracing) for more information.
    
    ```yaml
    - "Mesh.http.routers.myrouter.observability.tracing=true"
    ```

??? info "`Mesh.http.routers.<router_name>.priority`"

    See [priority](../routers/index.md#priority) for more information.

    ```yaml
    - "Mesh.http.routers.myrouter.priority=42"
    ```

### Services

To update the configuration of the Service automatically attached to the container,
add labels starting with `Mesh.http.services.<name-of-your-choice>.`, followed by the option you want to change.

For example, to change the `passHostHeader` behavior,
you'd add the label `Mesh.http.services.<name-of-your-choice>.loadbalancer.passhostheader=false`.

!!! warning "The character `@` is not authorized in the service name `<service_name>`."

??? info "`Mesh.http.services.<service_name>.loadbalancer.server.port`"

    Registers a port.
    Useful when the container exposes multiples ports.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.server.port=8080"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.server.scheme`"

    Overrides the default scheme.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.server.scheme=http"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.server.url`"

    Defines the service URL.
    This option cannot be used in combination with `port` or `scheme` definition.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.server.url=http://foobar:8080"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.serverstransport`"

    Allows to reference a ServersTransport resource that is defined either with the File provider or the Kubernetes CRD one.
    See [serverstransport](../services/index.md#serverstransport) for more information.

    ```yaml
    - "Mesh.http.services.<service_name>.loadbalancer.serverstransport=foobar@file"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.passhostheader`"

    See [pass Host header](../services/index.md#pass-host-header) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.passhostheader=true"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.headers.<header_name>`"

    See [health check](../services/index.md#health-check) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.healthcheck.headers.X-Foo=foobar"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.hostname`"

    See [health check](../services/index.md#health-check) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.healthcheck.hostname=example.org"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.interval`"

    See [health check](../services/index.md#health-check) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.healthcheck.interval=10s"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.unhealthyinterval`"

    See [health check](../services/index.md#health-check) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.healthcheck.unhealthyinterval=10s"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.path`"

    See [health check](../services/index.md#health-check) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.healthcheck.path=/foo"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.method`"

    See [health check](../services/index.md#health-check) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.healthcheck.method=foobar"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.status`"

    See [health check](../services/index.md#health-check) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.healthcheck.status=42"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.port`"

    See [health check](../services/index.md#health-check) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.healthcheck.port=42"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.scheme`"

    See [health check](../services/index.md#health-check) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.healthcheck.scheme=http"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.timeout`"

    See [health check](../services/index.md#health-check) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.healthcheck.timeout=10s"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.followredirects`"

    See [health check](../services/index.md#health-check) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.healthcheck.followredirects=true"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.sticky.cookie=true"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie.httponly`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.sticky.cookie.httponly=true"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie.name`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.sticky.cookie.name=foobar"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie.path`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.sticky.cookie.path=/foobar"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie.secure`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.sticky.cookie.secure=true"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie.samesite`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.sticky.cookie.samesite=none"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie.domain`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.sticky.cookie.domain=foo.com"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie.maxage`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.sticky.cookie.maxage=42"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.responseforwarding.flushinterval`"

    See [response forwarding](../services/index.md#response-forwarding) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.responseforwarding.flushinterval=10"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.strategy`"

    See [load balancing strategy](../services/index.md#load-balancing-strategy) for more information.

    ```yaml
    - "Mesh.http.services.myservice.loadbalancer.strategy=p2c"
    ```

### Middleware

You can declare pieces of middleware using labels starting with `Mesh.http.middlewares.<name-of-your-choice>.`,
followed by the middleware type/options.

For example, to declare a middleware [`redirectscheme`](../../middlewares/http/redirectscheme.md) named `my-redirect`,
you'd write `Mesh.http.middlewares.my-redirect.redirectscheme.scheme=https`.

More information about available middlewares in the dedicated [middlewares section](../../middlewares/overview.md).

!!! warning "The character `@` is not authorized in the middleware name."

??? example "Declaring and Referencing a Middleware"

    ```yaml
       services:
         my-container:
           # ...
           labels:
             # Declaring a middleware
             - Mesh.http.middlewares.my-redirect.redirectscheme.scheme=https
             # Referencing a middleware
             - Mesh.http.routers.my-container.middlewares=my-redirect
    ```

!!! warning "Conflicts in Declaration"

    If you declare multiple middleware with the same name but with different parameters, the middleware fails to be declared.

### TCP

You can declare TCP Routers and/or Services using labels.

??? example "Declaring TCP Routers with one Service"

    ```yaml
       services:
         my-container:
           # ...
           labels:
             - "Mesh.tcp.routers.my-router.rule=HostSNI(`example.com`)"
             - "Mesh.tcp.routers.my-router.tls=true"
             - "Mesh.tcp.services.my-service.loadbalancer.server.port=4123"
    ```

!!! warning "TCP and HTTP"

    If you declare a TCP Router/Service, it will prevent Mesh from automatically creating an HTTP Router/Service (like it does by default if no TCP Router/Service is defined).
    You can declare both a TCP Router/Service and an HTTP Router/Service for the same container (but you have to do so manually).

#### TCP Routers

??? info "`Mesh.tcp.routers.<router_name>.entrypoints`"

    See [entry points](../routers/index.md#entrypoints_1) for more information.

    ```yaml
    - "Mesh.tcp.routers.mytcprouter.entrypoints=ep1,ep2"
    ```

??? info "`Mesh.tcp.routers.<router_name>.rule`"

    See [rule](../routers/index.md#rule_1) for more information.

    ```yaml
    - "Mesh.tcp.routers.mytcprouter.rule=HostSNI(`example.com`)"
    ```

??? info "`Mesh.tcp.routers.<router_name>.service`"

    See [service](../routers/index.md#services) for more information.

    ```yaml
    - "Mesh.tcp.routers.mytcprouter.service=myservice"
    ```

??? info "`Mesh.tcp.routers.<router_name>.tls`"

    See [TLS](../routers/index.md#tls_1) for more information.

    ```yaml
    - "Mesh.tcp.routers.mytcprouter.tls=true"
    ```

??? info "`Mesh.tcp.routers.<router_name>.tls.certresolver`"

    See [certResolver](../routers/index.md#certresolver_1) for more information.

    ```yaml
    - "Mesh.tcp.routers.mytcprouter.tls.certresolver=myresolver"
    ```

??? info "`Mesh.tcp.routers.<router_name>.tls.domains[n].main`"

    See [domains](../routers/index.md#domains_1) for more information.

    ```yaml
    - "Mesh.tcp.routers.mytcprouter.tls.domains[0].main=example.org"
    ```

??? info "`Mesh.tcp.routers.<router_name>.tls.domains[n].sans`"

    See [domains](../routers/index.md#domains_1) for more information.

    ```yaml
    - "Mesh.tcp.routers.mytcprouter.tls.domains[0].sans=test.example.org,dev.example.org"
    ```

??? info "`Mesh.tcp.routers.<router_name>.tls.options`"

    See [options](../routers/index.md#options_1) for more information.

    ```yaml
    - "Mesh.tcp.routers.mytcprouter.tls.options=mysoptions"
    ```

??? info "`Mesh.tcp.routers.<router_name>.tls.passthrough`"

    See [TLS](../routers/index.md#tls_1) for more information.

    ```yaml
    - "Mesh.tcp.routers.mytcprouter.tls.passthrough=true"
    ```

??? info "`Mesh.tcp.routers.<router_name>.priority`"

    See [priority](../routers/index.md#priority_1) for more information.

    ```yaml
    - "Mesh.tcp.routers.myrouter.priority=42"
    ```

#### TCP Services

??? info "`Mesh.tcp.services.<service_name>.loadbalancer.server.port`"

    Registers a port of the application.

    ```yaml
    - "Mesh.tcp.services.mytcpservice.loadbalancer.server.port=423"
    ```

??? info "`Mesh.tcp.services.<service_name>.loadbalancer.server.tls`"

    Determines whether to use TLS when dialing with the backend.

    ```yaml
    - "Mesh.tcp.services.mytcpservice.loadbalancer.server.tls=true"
    ```

??? info "`Mesh.tcp.services.<service_name>.loadbalancer.proxyprotocol.version`"

    See [PROXY protocol](../services/index.md#proxy-protocol) for more information.

    ```yaml
    - "Mesh.tcp.services.mytcpservice.loadbalancer.proxyprotocol.version=1"
    ```

??? info "`Mesh.tcp.services.<service_name>.loadbalancer.serverstransport`"

    Allows to reference a ServersTransport resource that is defined either with the File provider or the Kubernetes CRD one.
    See [serverstransport](../services/index.md#serverstransport_2) for more information.

    ```yaml
    - "Mesh.tcp.services.<service_name>.loadbalancer.serverstransport=foobar@file"
    ```

### UDP

You can declare UDP Routers and/or Services using labels.

??? example "Declaring UDP Routers with one Service"

    ```yaml
       services:
         my-container:
           # ...
           labels:
             - "Mesh.udp.routers.my-router.entrypoints=udp"
             - "Mesh.udp.services.my-service.loadbalancer.server.port=4123"
    ```

!!! warning "UDP and HTTP"

    If you declare a UDP Router/Service, it will prevent Mesh from automatically creating an HTTP Router/Service (like it does by default if no UDP Router/Service is defined).
    You can declare both a UDP Router/Service and an HTTP Router/Service for the same container (but you have to do so manually).

#### UDP Routers

??? info "`Mesh.udp.routers.<router_name>.entrypoints`"

    See [entry points](../routers/index.md#entrypoints_2) for more information.

    ```yaml
    - "Mesh.udp.routers.myudprouter.entrypoints=ep1,ep2"
    ```

??? info "`Mesh.udp.routers.<router_name>.service`"

    See [service](../routers/index.md#services_1) for more information.

    ```yaml
    - "Mesh.udp.routers.myudprouter.service=myservice"
    ```

#### UDP Services

??? info "`Mesh.udp.services.<service_name>.loadbalancer.server.port`"

    Registers a port of the application.

    ```yaml
    - "Mesh.udp.services.myudpservice.loadbalancer.server.port=423"
    ```

### Specific Provider Options

#### `Mesh.enable`

```yaml
- "Mesh.enable=true"
```

You can tell Mesh to consider (or not) the container by setting `Mesh.enable` to true or false.

This option overrides the value of `exposedByDefault`.

#### `Mesh.docker.network`

```yaml
- "Mesh.docker.network=mynetwork"
```

Overrides the default docker network to use for connections to the container.

If a container is linked to several networks, be sure to set the proper network name (you can check this with `docker inspect <container_id>`),
otherwise it will randomly pick one (depending on how docker is returning them).

!!! warning
    When deploying a stack from a compose file `stack`, the networks defined are prefixed with `stack`.
