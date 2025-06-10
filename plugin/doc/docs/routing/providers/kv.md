---
title: "Mesh Routing Configuration with KV stores"
description: "Read the technical documentation to learn the Mesh Routing Configuration with KV stores."
---

# Mesh & KV Stores

A Story of key & values

## Routing Configuration

!!! info "Keys"

    - Keys are case-insensitive.
    - The complete list of keys can be found in [the reference page](../../reference/dynamic-configuration/kv.md).

### Routers

!!! warning "The character `@` is not authorized in the router name `<router_name>`."

??? info "`Mesh/http/routers/<router_name>/rule`"

    See [rule](../routers/index.md#rule) for more information.
    
    | Key (Path)                           | Value                      |
    |--------------------------------------|----------------------------|
    | `Mesh/http/routers/myrouter/rule` | ```Host(`example.com`)```  |

??? info "`Mesh/http/routers/<router_name>/entrypoints`"

    See [entry points](../routers/index.md#entrypoints) for more information.

    | Key (Path)                                    | Value       |
    |-----------------------------------------------|-------------|
    | `Mesh/http/routers/myrouter/entrypoints/0` | `web`       |
    | `Mesh/http/routers/myrouter/entrypoints/1` | `websecure` |

??? info "`Mesh/http/routers/<router_name>/middlewares`"

    See [middlewares](../routers/index.md#middlewares) and [middlewares overview](../../middlewares/overview.md) for more information.

    | Key (Path)                                    | Value       |
    |-----------------------------------------------|-------------|
    | `Mesh/http/routers/myrouter/middlewares/0` | `auth`      |
    | `Mesh/http/routers/myrouter/middlewares/1` | `prefix`    |
    | `Mesh/http/routers/myrouter/middlewares/2` | `cb`        |

??? info "`Mesh/http/routers/<router_name>/service`"

    See [rule](../routers/index.md#service) for more information.

    | Key (Path)                              | Value       |
    |-----------------------------------------|-------------|
    | `Mesh/http/routers/myrouter/service` | `myservice` |

??? info "`Mesh/http/routers/<router_name>/tls`"

    See [tls](../routers/index.md#tls) for more information.

    | Key (Path)                          | Value  |
    |-------------------------------------|--------|
    | `Mesh/http/routers/myrouter/tls` | `true` |
    
??? info "`Mesh/http/routers/<router_name>/tls/certresolver`"

    See [certResolver](../routers/index.md#certresolver) for more information.

    | Key (Path)                                       | Value        |
    |--------------------------------------------------|--------------|
    | `Mesh/http/routers/myrouter/tls/certresolver` | `myresolver` |    

??? info "`Mesh/http/routers/<router_name>/tls/domains/<n>/main`"

    See [domains](../routers/index.md#domains) for more information.

    | Key (Path)                                         | Value         |
    |----------------------------------------------------|---------------|
    | `Mesh/http/routers/myrouter/tls/domains/0/main` | `example.org` |
    
??? info "`Mesh/http/routers/<router_name>/tls/domains/<n>/sans/<n>`"

    See [domains](../routers/index.md#domains) for more information.

    | Key (Path)                                           | Value              |
    |------------------------------------------------------|--------------------|
    | `Mesh/http/routers/myrouter/tls/domains/0/sans/0` | `test.example.org` |
    | `Mesh/http/routers/myrouter/tls/domains/0/sans/1` | `dev.example.org`  |
    
??? info "`Mesh/http/routers/<router_name>/tls/options`"

    See [options](../routers/index.md#options) for more information.

    | Key (Path)                                  | Value    |
    |---------------------------------------------|----------|
    | `Mesh/http/routers/myrouter/tls/options` | `foobar` |

??? info "`Mesh/http/routers/<router_name>/observability/accesslogs`"

    See accesslogs [option](../routers/index.md#accesslogs) for more information.

    | Key (Path)                                               | Value  |
    |----------------------------------------------------------|--------|
    | `Mesh/http/routers/myrouter/observability/accesslogs` | `true` |

??? info "`Mesh/http/routers/<router_name>/observability/metrics`"

    See metrics [option](../routers/index.md#metrics) for more information.

    | Key (Path)                                            | Value  |
    |-------------------------------------------------------|--------|
    | `Mesh/http/routers/myrouter/observability/metrics` | `true` |

??? info "`Mesh/http/routers/<router_name>/observability/tracing`"

    See tracing [option](../routers/index.md#tracing) for more information.

    | Key (Path)                                            | Value  |
    |-------------------------------------------------------|--------|
    | `Mesh/http/routers/myrouter/observability/tracing` | `true` |

??? info "`Mesh/http/routers/<router_name>/priority`"

    See [priority](../routers/index.md#priority) for more information.

    | Key (Path)                               | Value |
    |------------------------------------------|-------|
    | `Mesh/http/routers/myrouter/priority` | `42`  |

### Services

!!! warning "The character `@` is not authorized in the service name `<service_name>`."

??? info "`Mesh/http/services/<service_name>/loadbalancer/servers/<n>/url`"

    See [servers](../services/index.md#servers) for more information.

    | Key (Path)                                                      | Value                                   |
    |-----------------------------------------------------------------|-----------------------------------------|
    | `Mesh/http/services/myservice/loadbalancer/servers/0/url`    | `http://<ip-server-1>:<port-server-1>/` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/serverstransport`"

    Allows to reference a ServersTransport resource that is defined either with the File provider or the Kubernetes CRD one.
    See [serverstransport](../services/index.md#serverstransport) for more information.

    | Key (Path)                                                      | Value         |
    |-----------------------------------------------------------------|---------------|
    | `Mesh/http/services/myservice/loadbalancer/serverstransport` | `foobar@file` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/passhostheader`"

    See [pass Host header](../services/index.md#pass-host-header) for more information.

    | Key (Path)                                                      | Value  |
    |-----------------------------------------------------------------|--------|
    | `Mesh/http/services/myservice/loadbalancer/passhostheader`   | `true` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/healthcheck/headers/<header_name>`"

    See [health check](../services/index.md#health-check) for more information.

    | Key (Path)                                                               | Value    |
    |--------------------------------------------------------------------------|----------|
    | `Mesh/http/services/myservice/loadbalancer/healthcheck/headers/X-Foo` | `foobar` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/healthcheck/hostname`"

    See [health check](../services/index.md#health-check) for more information.

    | Key (Path)                                                          | Value         |
    |---------------------------------------------------------------------|---------------|
    | `Mesh/http/services/myservice/loadbalancer/healthcheck/hostname` | `example.org` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/healthcheck/interval`"

    See [health check](../services/index.md#health-check) for more information.

    | Key (Path)                                                          | Value |
    |---------------------------------------------------------------------|-------|
    | `Mesh/http/services/myservice/loadbalancer/healthcheck/interval` | `10`  |

??? info "`Mesh/http/services/<service_name>/loadbalancer/healthcheck/unhealthyinterval`"

    See [health check](../services/index.md#health-check) for more information.

    | Key (Path)                                                                   | Value |
    |------------------------------------------------------------------------------|-------|
    | `Mesh/http/services/myservice/loadbalancer/healthcheck/unhealthyinterval` | `10`  |

??? info "`Mesh/http/services/<service_name>/loadbalancer/healthcheck/path`"

    See [health check](../services/index.md#health-check) for more information.

    | Key (Path)                                                      | Value  |
    |-----------------------------------------------------------------|--------|
    | `Mesh/http/services/myservice/loadbalancer/healthcheck/path` | `/foo` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/healthcheck/method`"

    See [health check](../services/index.md#health-check) for more information.

    | Key (Path)                                                        | Value    |
    |-------------------------------------------------------------------|----------|
    | `Mesh/http/services/myservice/loadbalancer/healthcheck/method` | `foobar` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/healthcheck/status`"

    See [health check](../services/index.md#health-check) for more information.

    | Key (Path)                                                        | Value |
    |-------------------------------------------------------------------|-------|
    | `Mesh/http/services/myservice/loadbalancer/healthcheck/status` | `42`  |

??? info "`Mesh/http/services/<service_name>/loadbalancer/healthcheck/port`"

    See [health check](../services/index.md#health-check) for more information.

    | Key (Path)                                                      | Value |
    |-----------------------------------------------------------------|-------|
    | `Mesh/http/services/myservice/loadbalancer/healthcheck/port` | `42`  |

??? info "`Mesh/http/services/<service_name>/loadbalancer/healthcheck/scheme`"

    See [health check](../services/index.md#health-check) for more information.

    | Key (Path)                                                        | Value  |
    |-------------------------------------------------------------------|--------|
    | `Mesh/http/services/myservice/loadbalancer/healthcheck/scheme` | `http` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/healthcheck/timeout`"

    See [health check](../services/index.md#health-check) for more information.

    | Key (Path)                                                         | Value |
    |--------------------------------------------------------------------|-------|
    | `Mesh/http/services/myservice/loadbalancer/healthcheck/timeout` | `10`  |

??? info "`Mesh/http/services/<service_name>/loadbalancer/sticky`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    | Key (Path)                                            | Value  |
    |-------------------------------------------------------|--------|
    | `Mesh/http/services/myservice/loadbalancer/sticky` | `true` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/sticky/cookie/httponly`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    | Key (Path)                                                            | Value  |
    |-----------------------------------------------------------------------|--------|
    | `Mesh/http/services/myservice/loadbalancer/sticky/cookie/httponly` | `true` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/sticky/cookie/name`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    | Key (Path)                                                        | Value    |
    |-------------------------------------------------------------------|----------|
    | `Mesh/http/services/myservice/loadbalancer/sticky/cookie/name` | `foobar` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/sticky/cookie/path`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    | Key (Path)                                                        | Value     |
    |-------------------------------------------------------------------|-----------|
    | `Mesh/http/services/myservice/loadbalancer/sticky/cookie/path` | `/foobar` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/sticky/cookie/secure`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    | Key (Path)                                                          | Value  |
    |---------------------------------------------------------------------|--------|
    | `Mesh/http/services/myservice/loadbalancer/sticky/cookie/secure` | `true` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/sticky/cookie/samesite`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    | Key (Path)                                                            | Value  |
    |-----------------------------------------------------------------------|--------|
    | `Mesh/http/services/myservice/loadbalancer/sticky/cookie/samesite` | `none` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/sticky/cookie/domain`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    | Key (Path)                                                            | Value     |
    |-----------------------------------------------------------------------|-----------|
    | `Mesh/http/services/myservice/loadbalancer/sticky/cookie/domain`   | `foo.com` |

??? info "`Mesh/http/services/<service_name>/loadbalancer/sticky/cookie/maxage`"

    See [sticky sessions](../services/index.md#sticky-sessions) for more information.

    | Key (Path)                                                          | Value |
    |---------------------------------------------------------------------|-------|
    | `Mesh/http/services/myservice/loadbalancer/sticky/cookie/maxage` | `42`  |

??? info "`Mesh/http/services/<service_name>/loadbalancer/responseforwarding/flushinterval`"

    See [response forwarding](../services/index.md#response-forwarding) for more information.

    | Key (Path)                                                                      | Value |
    |---------------------------------------------------------------------------------|-------|
    | `Mesh/http/services/myservice/loadbalancer/responseforwarding/flushinterval` | `10`  |

??? info "`Mesh/http/services/<service_name>/loadbalancer/strategy`"

    See [load balancing strategy](../services/index.md#load-balancing-strategy) for more information.

    | Key (Path)                                              | Value |
    |---------------------------------------------------------|-------|
    | `Mesh/http/services/myservice/loadbalancer/strategy` | `p2c` |

??? info "`Mesh/http/services/<service_name>/mirroring/service`"

    | Key (Path)                                               | Value    |
    |----------------------------------------------------------|----------|
    | `Mesh/http/services/<service_name>/mirroring/service` | `foobar` |

??? info "`Mesh/http/services/<service_name>/mirroring/mirrors/<n>/name`"

    | Key (Path)                                                        | Value    |
    |-------------------------------------------------------------------|----------|
    | `Mesh/http/services/<service_name>/mirroring/mirrors/<n>/name` | `foobar` |

??? info "`Mesh/http/services/<service_name>/mirroring/mirrors/<n>/percent`"

    | Key (Path)                                                           | Value |
    |----------------------------------------------------------------------|-------|
    | `Mesh/http/services/<service_name>/mirroring/mirrors/<n>/percent` | `42`  |

??? info "`Mesh/http/services/<service_name>/weighted/services/<n>/name`"

    | Key (Path)                                                        | Value    |
    |-------------------------------------------------------------------|----------|
    | `Mesh/http/services/<service_name>/weighted/services/<n>/name` | `foobar` |

??? info "`Mesh/http/services/<service_name>/weighted/services/<n>/weight`"

    | Key (Path)                                                          | Value |
    |---------------------------------------------------------------------|-------|
    | `Mesh/http/services/<service_name>/weighted/services/<n>/weight` | `42`  |

??? info "`Mesh/http/services/<service_name>/weighted/sticky/cookie/name`"

    | Key (Path)                                                         | Value    |
    |--------------------------------------------------------------------|----------|
    | `Mesh/http/services/<service_name>/weighted/sticky/cookie/name` | `foobar` |

??? info "`Mesh/http/services/<service_name>/weighted/sticky/cookie/secure`"

    | Key (Path)                                                           | Value  |
    |----------------------------------------------------------------------|--------|
    | `Mesh/http/services/<service_name>/weighted/sticky/cookie/secure` | `true` |

??? info "`Mesh/http/services/<service_name>/weighted/sticky/cookie/samesite`"

    | Key (Path)                                                             | Value  |
    |------------------------------------------------------------------------|--------|
    | `Mesh/http/services/<service_name>/weighted/sticky/cookie/samesite` | `none` |

??? info "`Mesh/http/services/<service_name>/weighted/sticky/cookie/domain`"

    | Key (Path)                                                             | Value     |
    |------------------------------------------------------------------------|-----------|
    | `Mesh/http/services/<service_name>/weighted/sticky/cookie/domain`   | `foo.com` |

??? info "`Mesh/http/services/<service_name>/weighted/sticky/cookie/httpOnly`"

    | Key (Path)                                                             | Value  |
    |------------------------------------------------------------------------|--------|
    | `Mesh/http/services/<service_name>/weighted/sticky/cookie/httpOnly` | `true` |

??? info "`Mesh/http/services/<service_name>/weighted/sticky/cookie/maxage`"

    | Key (Path)                                                           | Value |
    |----------------------------------------------------------------------|-------|
    | `Mesh/http/services/<service_name>/weighted/sticky/cookie/maxage` | `42`  |

??? info "`Mesh/http/services/<service_name>/weighted/sticky/cookie/path`"

    | Key (Path)                                                           | Value     |
    |----------------------------------------------------------------------|-----------|
    | `Mesh/http/services/<service_name>/weighted/sticky/cookie/path`   | `/foobar` |

### Middleware

More information about available middlewares in the dedicated [middlewares section](../../middlewares/overview.md).

!!! warning "The character `@` is not authorized in the middleware name."

!!! warning "Conflicts in Declaration"

    If you declare multiple middleware with the same name but with different parameters, the middleware fails to be declared.

### TCP

You can declare TCP Routers and/or Services using KV.

#### TCP Routers

??? info "`Mesh/tcp/routers/<router_name>/entrypoints`"

    See [entry points](../routers/index.md#entrypoints_1) for more information.

    | Key (Path)                                      | Value |
    |-------------------------------------------------|-------|
    | `Mesh/tcp/routers/mytcprouter/entrypoints/0` | `ep1` |
    | `Mesh/tcp/routers/mytcprouter/entrypoints/1` | `ep2` |
    
??? info "`Mesh/tcp/routers/<router_name>/rule`"

    See [rule](../routers/index.md#rule_1) for more information.

    | Key (Path)                           | Value                        |
    |--------------------------------------|------------------------------|
    | `Mesh/tcp/routers/my-router/rule` | ```HostSNI(`example.com`)``` |  

??? info "`Mesh/tcp/routers/<router_name>/service`"

    See [service](../routers/index.md#services) for more information.
    
    | Key (Path)                                | Value       |
    |-------------------------------------------|-------------|
    | `Mesh/tcp/routers/mytcprouter/service` | `myservice` |

??? info "`Mesh/tcp/routers/<router_name>/tls`"

    See [TLS](../routers/index.md#tls_1) for more information.

    | Key (Path)                            | Value  |
    |---------------------------------------|--------|
    | `Mesh/tcp/routers/mytcprouter/tls` | `true` |

??? info "`Mesh/tcp/routers/<router_name>/tls/certresolver`"

    See [certResolver](../routers/index.md#certresolver_1) for more information.

    | Key (Path)                                         | Value        |
    |----------------------------------------------------|--------------|
    | `Mesh/tcp/routers/mytcprouter/tls/certresolver` | `myresolver` |

??? info "`Mesh/tcp/routers/<router_name>/tls/domains/<n>/main`"

    See [domains](../routers/index.md#domains_1) for more information.

    | Key (Path)                                           | Value         |
    |------------------------------------------------------|---------------|
    | `Mesh/tcp/routers/mytcprouter/tls/domains/0/main` | `example.org` |
        
??? info "`Mesh/tcp/routers/<router_name>/tls/domains/<n>/sans`"

    See [domains](../routers/index.md#domains_1) for more information.

    | Key (Path)                                             | Value              |
    |--------------------------------------------------------|--------------------|
    | `Mesh/tcp/routers/mytcprouter/tls/domains/0/sans/0` | `test.example.org` |
    | `Mesh/tcp/routers/mytcprouter/tls/domains/0/sans/1` | `dev.example.org`  |
    
??? info "`Mesh/tcp/routers/<router_name>/tls/options`"

    See [options](../routers/index.md#options_1) for more information.

    | Key (Path)                                    | Value    |
    |-----------------------------------------------|----------|
    | `Mesh/tcp/routers/mytcprouter/tls/options` | `foobar` |

??? info "`Mesh/tcp/routers/<router_name>/tls/passthrough`"

    See [TLS](../routers/index.md#tls_1) for more information.

    | Key (Path)                                        | Value  |
    |---------------------------------------------------|--------|
    | `Mesh/tcp/routers/mytcprouter/tls/passthrough` | `true` |

??? info "`Mesh/tcp/routers/<router_name>/priority`"

    See [priority](../routers/index.md#priority_1) for more information.

    | Key (Path)                               | Value |
    |------------------------------------------|-------|
    | `Mesh/tcp/routers/myrouter/priority`  | `42`  |

#### TCP Services

??? info "`Mesh/tcp/services/<service_name>/loadbalancer/servers/<n>/address`"

    See [servers](../services/index.md#servers) for more information.

    | Key (Path)                                                         | Value            |
    |--------------------------------------------------------------------|------------------|
    | `Mesh/tcp/services/mytcpservice/loadbalancer/servers/0/address` | `xx.xx.xx.xx:xx` |
    
??? info "`Mesh/tcp/services/<service_name>/loadbalancer/proxyprotocol/version`"

    See [PROXY protocol](../services/index.md#proxy-protocol) for more information.

    | Key (Path)                                                             | Value |
    |------------------------------------------------------------------------|-------|
    | `Mesh/tcp/services/mytcpservice/loadbalancer/proxyprotocol/version` | `1`   |

??? info "`Mesh/tcp/services/<service_name>/loadbalancer/serverstransport`"

    Allows to reference a ServersTransport resource that is defined either with the File provider or the Kubernetes CRD one.
    See [serverstransport](../services/index.md#serverstransport_2) for more information.

    | Key (Path)                                                      | Value         |
    |-----------------------------------------------------------------|---------------|
    | `Mesh/tcp/services/myservice/loadbalancer/serverstransport` | `foobar@file` |

??? info "`Mesh/tcp/services/<service_name>/weighted/services/<n>/name`"

    | Key (Path)                                                          | Value    |
    |---------------------------------------------------------------------|----------|
    | `Mesh/tcp/services/<service_name>/weighted/services/0/name`      | `foobar` |

??? info "`Mesh/tcp/services/<service_name>/weighted/services/<n>/weight`"

    | Key (Path)                                                       | Value |
    |------------------------------------------------------------------|-------|
    | `Mesh/tcp/services/<service_name>/weighted/services/0/weight` | `42`  |
