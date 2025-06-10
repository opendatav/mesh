---
title: "Mesh ECS Documentation"
description: "AWS ECS is a provider for routing and load balancing in Mesh Proxy. Read the technical documentation to get started."
---

# Mesh & ECS

One of the best feature of Mesh is to delegate the routing configuration to the application level.
With ECS, Mesh can leverage labels attached to a container to generate routing rules.

!!! warning "Labels & sensitive data"

    We recommend to *not* use labels to store sensitive data (certificates, credentials, etc).
    Instead, we recommend to store sensitive data in a safer storage (secrets, file, etc).

## Routing Configurationred

!!! info "labels"
    
    Labels are case-insensitive.

!!! tip "TLS Default Generated Certificates"

    To learn how to configure Mesh default generated certificate, refer to the [TLS Certificates](../http/tls/tls-certificates.md#acme-default-certificate) page.

### General

Mesh creates, for each elastic service, a corresponding [service](../http/load-balancing/service.md) and [router](../http/router/rules-and-priority.md).

The Service automatically gets a server per elastic container, and the router gets a default rule attached to it, based on the service name.

### Routers

To update the configuration of the Router automatically attached to the service, add labels starting with `Mesh.routers.{name-of-your-choice}.` and followed by the option you want to change.

For example, to change the rule, you could add the label ```Mesh.http.routers.my-service.rule=Host(`example.com`)```.

!!! warning "The character `@` is not authorized in the router name `<router_name>`."

??? info "`Mesh.http.routers.<router_name>.rule`"
    
    See [rule](../http/router/rules-and-priority.md#rules) for more information.
    
    ```yaml
    Mesh.http.routers.myrouter.rule=Host(`example.com`)
    ```

??? info "`Mesh.http.routers.<router_name>.ruleSyntax`"

    !!! warning

        RuleSyntax option is deprecated and will be removed in the next major version.
        Please do not use this field and rewrite the router rules to use the v3 syntax.

    See [ruleSyntax](../http/router/rules-and-priority.md#rulesyntax) for more information.
    
    ```yaml
    Mesh.http.routers.myrouter.ruleSyntax=v3
    ```

??? info "`Mesh.http.routers.<router_name>.entrypoints`"
    
    See [entry points](../../install-configuration/entrypoints.md) for more information.
    
    ```yaml
    Mesh.http.routers.myrouter.entrypoints=web,websecure
    ```

??? info "`Mesh.http.routers.<router_name>.middlewares`"
    
    See [middlewares overview](../http/middlewares/overview.md) for more information.
    
    ```yaml
    Mesh.http.routers.myrouter.middlewares=auth,prefix,cb
    ```

??? info "`Mesh.http.routers.<router_name>.service`"
    
    See [service](../http/load-balancing/service.md) for more information.
    
    ```yaml
    Mesh.http.routers.myrouter.service=myservice
    ```

??? info "`Mesh.http.routers.<router_name>.tls`"
    
    See [tls](../http/tls/overview.md) for more information.
    
    ```yaml
    Mesh.http.routers.myrouter.tls=true
    ```

??? info "`Mesh.http.routers.<router_name>.tls.certresolver`"
    
    See [certResolver](../../install-configuration/tls/certificate-resolvers/overview.md) for more information.
    
    ```yaml
    Mesh.http.routers.myrouter.tls.certresolver=myresolver
    ```

??? info "`Mesh.http.routers.<router_name>.tls.domains[n].main`"
    
    See [domains](../../install-configuration/tls/certificate-resolvers/acme.md#domain-definition) for more information.
    
    ```yaml
    Mesh.http.routers.myrouter.tls.domains[0].main=example.org
    ```

??? info "`Mesh.http.routers.<router_name>.tls.domains[n].sans`"
    
    See [domains](../../install-configuration/tls/certificate-resolvers/acme.md#domain-definition) for more information.
    
    ```yaml
    Mesh.http.routers.myrouter.tls.domains[0].sans=test.example.org,dev.example.org
    ```

??? info "`Mesh.http.routers.<router_name>.tls.options`"
    
    ```yaml
    Mesh.http.routers.myrouter.tls.options=foobar
    ```

??? info "`Mesh.http.routers.<router_name>.observability.accesslogs`"
    
    The accessLogs option controls whether the router will produce access-logs.
    
    ```yaml
     "Mesh.http.routers.myrouter.observability.accesslogs=true"
    ```

??? info "`Mesh.http.routers.<router_name>.observability.metrics`"
    
    The metrics option controls whether the router will produce metrics.

    ```yaml
     "Mesh.http.routers.myrouter.observability.metrics=true"
    ```

??? info "`Mesh.http.routers.<router_name>.observability.tracing`"
    
    The tracing option controls whether the router will produce traces.

    ```yaml
     "Mesh.http.routers.myrouter.observability.tracing=true"
    ```

??? info "`Mesh.http.routers.<router_name>.priority`"

    See [priority](../http/router/rules-and-priority.md#priority-calculation) for more information.

    ```yaml
    Mesh.http.routers.myrouter.priority=42
    ```

### Services

To update the configuration of the Service automatically attached to the service,
add labels starting with `Mesh.http.services.{name-of-your-choice}.`, followed by the option you want to change.

For example, to change the `passHostHeader` behavior,
you'd add the label `Mesh.http.services.{name-of-your-choice}.loadbalancer.passhostheader=false`.

!!! warning "The character `@` is not authorized in the service name `<service_name>`."

??? info "`Mesh.http.services.<service_name>.loadbalancer.server.port`"
    
    Registers a port.
    Useful when the service exposes multiples ports.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.server.port=8080
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.server.scheme`"
    
    Overrides the default scheme.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.server.scheme=http
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.serverstransport`"
    
    Allows to reference a ServersTransport resource that is defined either with the File provider or the Kubernetes CRD one.
    See [serverstransport](../http/load-balancing/serverstransport.md) for more information.
    
    ```yaml
    Mesh.http.services.<service_name>.loadbalancer.serverstransport=foobar@file
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.passhostheader`"

    ```yaml
    Mesh.http.services.myservice.loadbalancer.passhostheader=true
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.headers.<header_name>`"
    
    See [health check](../http/load-balancing/service.md#health-check) for more information.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.healthcheck.headers.X-Foo=foobar
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.hostname`"
    
    See [health check](../http/load-balancing/service.md#health-check) for more information.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.healthcheck.hostname=example.org
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.interval`"
    
    See [health check](../http/load-balancing/service.md#health-check) for more information.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.healthcheck.interval=10
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.unhealthyinterval`"
    
    See [health check](../http/load-balancing/service.md#health-check) for more information.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.healthcheck.unhealthyinterval=10
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.path`"
    
    See [health check](../http/load-balancing/service.md#health-check) for more information.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.healthcheck.path=/foo
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.method`"
    
    See [health check](../http/load-balancing/service.md#health-check) for more information.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.healthcheck.method=foobar
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.status`"
    
    See [health check](../http/load-balancing/service.md#health-check) for more information.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.healthcheck.status=42
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.port`"
    
    See [health check](../http/load-balancing/service.md#health-check) for more information.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.healthcheck.port=42
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.scheme`"
    
    See [health check](../http/load-balancing/service.md#health-check) for more information.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.healthcheck.scheme=http
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.timeout`"
    
    See [health check](../http/load-balancing/service.md#health-check) for more information.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.healthcheck.timeout=10
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.healthcheck.followredirects`"
    
    See [health check](../http/load-balancing/service.md#health-check) for more information.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.healthcheck.followredirects=true
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie`"
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.sticky.cookie=true
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie.httponly`"
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.sticky.cookie.httponly=true
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie.name`"
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.sticky.cookie.name=foobar
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie.path`"

    ```yaml
     "Mesh.http.services.myservice.loadbalancer.sticky.cookie.path=/foobar"
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie.secure`"

    ```yaml
    Mesh.http.services.myservice.loadbalancer.sticky.cookie.secure=true
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie.samesite`"

    ```yaml
    Mesh.http.services.myservice.loadbalancer.sticky.cookie.samesite=none
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.sticky.cookie.maxage`"

    ```yaml
    Mesh.http.services.myservice.loadbalancer.sticky.cookie.maxage=42
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.responseforwarding.flushinterval`"
        
    `FlushInterval` specifies the flush interval to flush to the client while copying the response body.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.responseforwarding.flushinterval=10
    ```

### Middleware

You can declare pieces of middleware using labels starting with `Mesh.http.middlewares.{name-of-your-choice}.`, followed by the middleware type/options.

For example, to declare a middleware [`redirectscheme`](../http/middlewares/redirectscheme.md) named `my-redirect`, you'd write `Mesh.http.middlewares.my-redirect.redirectscheme.scheme: https`.

More information about available middlewares in the dedicated [middlewares section](../http/middlewares/overview.md).

!!! warning "The character `@` is not authorized in the middleware name."

??? example "Declaring and Referencing a Middleware"
    
    ```yaml
    # ...
    # Declaring a middleware
    Mesh.http.middlewares.my-redirect.redirectscheme.scheme=https
    # Referencing a middleware
    Mesh.http.routers.my-service.middlewares=my-redirect
    ```

!!! warning "Conflicts in Declaration"

    If you declare multiple middleware with the same name but with different parameters, the middleware fails to be declared.

### TCP

You can declare TCP Routers and/or Services using labels.

??? example "Declaring TCP Routers and Services"

    ```yaml
    Mesh.tcp.routers.my-router.rule=HostSNI(`example.com`)
    Mesh.tcp.routers.my-router.tls=true
    Mesh.tcp.services.my-service.loadbalancer.server.port=4123
    ```

!!! warning "TCP and HTTP"

    If you declare a TCP Router/Service, it will prevent Mesh from automatically creating an HTTP Router/Service (like it does by default if no TCP Router/Service is defined).
    You can declare both a TCP Router/Service and an HTTP Router/Service for the same elastic service (but you have to do so manually).

#### TCP Routers

??? info "`Mesh.tcp.routers.<router_name>.entrypoints`"
    
    See [entry points](../../install-configuration/entrypoints.md) for more information.
    
    ```yaml
    Mesh.tcp.routers.mytcprouter.entrypoints=ep1,ep2
    ```

??? info "`Mesh.tcp.routers.<router_name>.rule`"
    
    See [entry points](../../install-configuration/entrypoints.md) for more information.
    
    ```yaml
    Mesh.tcp.routers.mytcprouter.rule=HostSNI(`example.com`)
    ```

??? info "`Mesh.tcp.routers.<router_name>.ruleSyntax`"

    !!! warning

        RuleSyntax option is deprecated and will be removed in the next major version.
        Please do not use this field and rewrite the router rules to use the v3 syntax.

    configure the rule syntax to be used for parsing the rule on a per-router basis.
    
    ```yaml
    Mesh.tcp.routers.mytcprouter.ruleSyntax=v3
    ```
    
??? info "`Mesh.tcp.routers.<router_name>.service`"
    
    See [service](../tcp/service.md) for more information.
    
    ```yaml
    Mesh.tcp.routers.mytcprouter.service=myservice
    ```

??? info "`Mesh.tcp.routers.<router_name>.tls`"
    
    See [TLS](../tcp/tls.md) for more information.
    
    ```yaml
    Mesh.tcp.routers.mytcprouter.tls=true
    ```

??? info "`Mesh.tcp.routers.<router_name>.tls.certresolver`"
    
    See [certResolver](../tcp/tls.md#configuration-options) for more information.
    
    ```yaml
    Mesh.tcp.routers.mytcprouter.tls.certresolver=myresolver
    ```

??? info "`Mesh.tcp.routers.<router_name>.tls.domains[n].main`"
    
    See [TLS](../tcp/tls.md) for more information.
    
    ```yaml
    Mesh.tcp.routers.mytcprouter.tls.domains[0].main=example.org
    ```

??? info "`Mesh.tcp.routers.<router_name>.tls.domains[n].sans`"
    
    See [TLS](../tcp/tls.md) for more information.
    
    ```yaml
    Mesh.tcp.routers.mytcprouter.tls.domains[0].sans=test.example.org,dev.example.org
    ```

??? info "`Mesh.tcp.routers.<router_name>.tls.options`"
    
    See [TLS](../tcp/tls.md) for more information.
    
    ```yaml
    Mesh.tcp.routers.mytcprouter.tls.options=mysoptions
    ```

??? info "`Mesh.tcp.routers.<router_name>.tls.passthrough`"
    
    See [Passthrough](../tcp/tls.md#passthrough) for more information.
    
    ```yaml
    Mesh.tcp.routers.mytcprouter.tls.passthrough=true
    ```

??? info "`Mesh.tcp.routers.<router_name>.priority`"

    See [priority](../tcp/router/rules-and-priority.md#priority) for more information.

    ```yaml
    Mesh.tcp.routers.mytcprouter.priority=42
    ```

#### TCP Services

??? info "`Mesh.tcp.services.<service_name>.loadbalancer.server.port`"
    
    Registers a port of the application.
    
    ```yaml
    Mesh.tcp.services.mytcpservice.loadbalancer.server.port=423
    ```

??? info "`Mesh.tcp.services.<service_name>.loadbalancer.server.tls`"
    
    Determines whether to use TLS when dialing with the backend.
    
    ```yaml
    Mesh.tcp.services.mytcpservice.loadbalancer.server.tls=true
    ```

??? info "`Mesh.http.services.<service_name>.loadbalancer.server.weight`"

    Overrides the default weight.
    
    ```yaml
    Mesh.http.services.myservice.loadbalancer.server.weight=42
    ```

??? info "`Mesh.tcp.services.<service_name>.loadbalancer.proxyprotocol.version`"
        
    See [PROXY protocol](../tcp/service.md#proxy-protocol) for more information.
    
    ```yaml
    Mesh.tcp.services.mytcpservice.loadbalancer.proxyprotocol.version=1
    ```

??? info "`Mesh.tcp.services.<service_name>.loadbalancer.serverstransport`"

    Allows to reference a ServersTransport resource that is defined either with the File provider or the Kubernetes CRD one.
    See [serverstransport](../tcp/serverstransport.md) for more information.
    
    ```yaml
    Mesh.tcp.services.<service_name>.loadbalancer.serverstransport=foobar@file
    ```

### UDP

You can declare UDP Routers and/or Services using tags.

??? example "Declaring UDP Routers and Services"

    ```yaml
    Mesh.udp.routers.my-router.entrypoints=udp
    Mesh.udp.services.my-service.loadbalancer.server.port=4123
    ```

!!! warning "UDP and HTTP"

    If you declare a UDP Router/Service, it will prevent Mesh from automatically creating an HTTP Router/Service (like it does by default if no UDP Router/Service is defined).
    You can declare both a UDP Router/Service and an HTTP Router/Service for the same elastic service (but you have to do so manually).

#### TCP Middleware

You can declare pieces of middleware using tags starting with `Mesh.tcp.middlewares.{name-of-your-choice}.`, followed by the middleware type/options.

For example, to declare a middleware [`InFlightConn`](../tcp/middlewares/inflightconn.md) named `test-inflightconn`, you'd write `Mesh.tcp.middlewares.test-inflightconn.inflightconn.amount=10`.

More information about available middlewares in the dedicated [middlewares section](../tcp/middlewares/overview.md).

??? example "Declaring and Referencing a Middleware"
    
    ```yaml
    # ...
    # Declaring a middleware
    Mesh.tcp.middlewares.test-inflightconn.amount=10
    # Referencing a middleware
    Mesh.tcp.routers.my-service.middlewares=test-inflightconn
    ```

!!! warning "Conflicts in Declaration"

    If you declare multiple middleware with the same name but with different parameters, the middleware fails to be declared.

#### UDP Routers

??? info "`Mesh.udp.routers.<router_name>.entrypoints`"
    
    See [entry points](../../install-configuration/entrypoints.md) for more information.
    
    ```yaml
    Mesh.udp.routers.myudprouter.entrypoints=ep1,ep2
    ```

??? info "`Mesh.udp.routers.<router_name>.service`"
    
    See [service](../udp/service.md) for more information.
    
    ```yaml
    Mesh.udp.routers.myudprouter.service=myservice
    ```

#### UDP Services

??? info "`Mesh.udp.services.<service_name>.loadbalancer.server.port`"
    
    Registers a port of the application.
    
    ```yaml
    Mesh.udp.services.myudpservice.loadbalancer.server.port=423
    ```

### Specific Provider Options

#### `Mesh.enable`

```yaml
Mesh.enable=true
```

You can tell Mesh to consider (or not) the ECS service by setting `Mesh.enable` to true or false.

This option overrides the value of `exposedByDefault`.
