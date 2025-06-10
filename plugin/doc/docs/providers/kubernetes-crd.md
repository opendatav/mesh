---
title: "Kubernetes IngressRoute & Mesh CRD"
description: "The Mesh team developed a Custom Resource Definition (CRD) for an IngressRoute type, to provide a better way to configure access to a Kubernetes cluster."
---

# Mesh & Kubernetes

The Kubernetes Ingress Controller, The Custom Resource Way.

In early versions, Mesh supported Kubernetes only through the [Kubernetes Ingress provider](./kubernetes-ingress.md), which is a Kubernetes Ingress controller in the strict sense of the term.

However, as the community expressed the need to benefit from Mesh features without resorting to (lots of) annotations,
the Mesh engineering team developed a [Custom Resource Definition](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)
(CRD) for an IngressRoute type, defined below, in order to provide a better way to configure access to a Kubernetes cluster.

## Requirements

!!! tip "All Steps for a Successful Deployment"

    * Add/update **all** the Mesh resources [definitions](../reference/dynamic-configuration/kubernetes-crd.md#definitions)
    * Add/update the [RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) for the Mesh custom resources
    * Use [Helm Chart](../getting-started/install-Mesh.md#use-the-helm-chart) or use a custom Mesh Deployment
        * Enable the kubernetesCRD provider
        * Apply the needed kubernetesCRD provider [configuration](#provider-configuration)
    * Add all necessary Mesh custom [resources](../reference/dynamic-configuration/kubernetes-crd.md#resources)

!!! example "Installing Resource Definition and RBAC"

    ```bash
    # Install Mesh Resource Definitions:
    kubectl apply -f https://raw.githubusercontent.com/Mesh/Mesh/v3.4/docs/content/reference/dynamic-configuration/kubernetes-crd-definition-v1.yml
    
    # Install RBAC for Mesh:
    kubectl apply -f https://raw.githubusercontent.com/Mesh/Mesh/v3.4/docs/content/reference/dynamic-configuration/kubernetes-crd-rbac.yml
    ```

## Resource Configuration

When using KubernetesCRD as a provider,
Mesh uses [Custom Resource Definition](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) to retrieve its routing configuration.
Mesh Custom Resource Definitions are a Kubernetes implementation of the Mesh concepts. The main particularities are:

* The usage of `name` **and** `namespace` to refer to another Kubernetes resource.
* The usage of [secret](https://kubernetes.io/docs/concepts/configuration/secret/) for sensitive data (TLS certificates and credentials).
* The structure of the configuration.
* The requirement to declare all the [definitions](../reference/dynamic-configuration/kubernetes-crd.md#definitions).

The Mesh CRDs are building blocks that you can assemble according to your needs.
See the list of CRDs in the dedicated [routing section](../routing/providers/kubernetes-crd.md).

## LetsEncrypt Support with the Custom Resource Definition Provider

By design, Mesh is a stateless application, meaning that it only derives its configuration from the environment it runs in, without additional configuration.
For this reason, users can run multiple instances of Mesh at the same time to achieve HA, as is a common pattern in the kubernetes ecosystem.

When using a single instance of Mesh with Let's Encrypt, you should encounter no issues. However, this could be a single point of failure.
Unfortunately, it is not possible to run multiple instances of Mesh Proxy 2.0 with Let's Encrypt enabled, because there is no way to ensure that the correct instance of Mesh will receive the challenge request and subsequent responses.
Early versions (v1.x) of Mesh used a [KV store](https://doc.Mesh.io/Mesh/v1.7/configuration/acme/#storage) to attempt to achieve this, but due to sub-optimal performance that feature was dropped in 2.0.

If you need Let's Encrypt with HA in a Kubernetes environment, we recommend using [Mesh Enterprise](https://Mesh.io/mesh-enterprise/), which includes distributed Let's Encrypt as a supported feature.

If you want to keep using Mesh Proxy, high availability for Let's Encrypt can be achieved by using a Certificate Controller such as [Cert-Manager](https://cert-manager.io/docs/).
When using Cert-Manager to manage certificates, it creates secrets in your namespaces that can be referenced as TLS secrets in your [ingress objects](https://kubernetes.io/docs/concepts/services-networking/ingress/#tls).
When using the Mesh Kubernetes CRD Provider, unfortunately Cert-Manager cannot yet interface directly with the CRDs.
A workaround is to enable the [Kubernetes Ingress provider](./kubernetes-ingress.md) to allow Cert-Manager to create ingress objects to complete the challenges.
Please note that this still requires manual intervention to create the certificates through Cert-Manager, but once the certificates are created, Cert-Manager keeps them renewed.

## Provider Configuration

### `endpoint`

_Optional, Default=""_

The Kubernetes server endpoint URL.

When deployed into Kubernetes, Mesh reads the environment variables `KUBERNETES_SERVICE_HOST` and `KUBERNETES_SERVICE_PORT` or `KUBECONFIG` to construct the endpoint.

The access token is looked up in `/var/run/secrets/kubernetes.io/serviceaccount/token` and the SSL CA certificate in `/var/run/secrets/kubernetes.io/serviceaccount/ca.crt`.
Both are mounted automatically when deployed inside Kubernetes.

The endpoint may be specified to override the environment variable values inside a cluster.

When the environment variables are not found, Mesh tries to connect to the Kubernetes API server with an external-cluster client.
In this case, the endpoint is required.
Specifically, it may be set to the URL used by `kubectl proxy` to connect to a Kubernetes cluster using the granted authentication and authorization of the associated kubeconfig.

```yaml tab="File (YAML)"
providers:
  kubernetesCRD:
    endpoint: "http://localhost:8080"
    # ...
```

```toml tab="File (TOML)"
[providers.kubernetesCRD]
  endpoint = "http://localhost:8080"
  # ...
```

```bash tab="CLI"
--providers.kubernetescrd.endpoint=http://localhost:8080
```

### `token`

_Optional, Default=""_

Bearer token used for the Kubernetes client configuration.

```yaml tab="File (YAML)"
providers:
  kubernetesCRD:
    token: "mytoken"
    # ...
```

```toml tab="File (TOML)"
[providers.kubernetesCRD]
  token = "mytoken"
  # ...
```

```bash tab="CLI"
--providers.kubernetescrd.token=mytoken
```

### `certAuthFilePath`

_Optional, Default=""_

Path to the certificate authority file.
Used for the Kubernetes client configuration.

```yaml tab="File (YAML)"
providers:
  kubernetesCRD:
    certAuthFilePath: "/my/ca.crt"
    # ...
```

```toml tab="File (TOML)"
[providers.kubernetesCRD]
  certAuthFilePath = "/my/ca.crt"
  # ...
```

```bash tab="CLI"
--providers.kubernetescrd.certauthfilepath=/my/ca.crt
```

### `namespaces`

_Optional, Default: []_

Array of namespaces to watch.
If left empty, Mesh watches all namespaces.

```yaml tab="File (YAML)"
providers:
  kubernetesCRD:
    namespaces:
    - "default"
    - "production"
    # ...
```

```toml tab="File (TOML)"
[providers.kubernetesCRD]
  namespaces = ["default", "production"]
  # ...
```

```bash tab="CLI"
--providers.kubernetescrd.namespaces=default,production
```

### `labelselector`

_Optional, Default: ""_

A label selector can be defined to filter on specific resource objects only,
this applies only to Mesh [Custom Resources](../routing/providers/kubernetes-crd.md#custom-resource-definition-crd)
and has no effect on Kubernetes `Secrets`, `EndpointSlices` and `Services`.
If left empty, Mesh processes all resource objects in the configured namespaces.

See [label-selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors) for details.

!!! warning

    Because the label selector is applied to all Mesh Custom Resources, they all must match the filter.

```yaml tab="File (YAML)"
providers:
  kubernetesCRD:
    labelSelector: "app=Mesh"
    # ...
```

```toml tab="File (TOML)"
[providers.kubernetesCRD]
  labelSelector = "app=Mesh"
  # ...
```

```bash tab="CLI"
--providers.kubernetescrd.labelselector="app=Mesh"
```

### `ingressClass`

_Optional, Default: ""_

Value of `kubernetes.io/ingress.class` annotation that identifies resource objects to be processed.

If the parameter is set, only resources containing an annotation with the same value are processed.
Otherwise, resources missing the annotation, having an empty value, or the value `Mesh` are processed.

```yaml tab="File (YAML)"
providers:
  kubernetesCRD:
    ingressClass: "mesh-internal"
    # ...
```

```toml tab="File (TOML)"
[providers.kubernetesCRD]
  ingressClass = "mesh-internal"
  # ...
```

```bash tab="CLI"
--providers.kubernetescrd.ingressclass=mesh-internal
```

### `throttleDuration`

_Optional, Default: 0_

The `throttleDuration` option defines how often the provider is allowed to handle events from Kubernetes. This prevents
a Kubernetes cluster that updates many times per second from continuously changing your Mesh configuration.

If left empty, the provider does not apply any throttling and does not drop any Kubernetes events.

The value of `throttleDuration` should be provided in seconds or as a valid duration format,
see [time.ParseDuration](https://golang.org/pkg/time/#ParseDuration).

```yaml tab="File (YAML)"
providers:
  kubernetesCRD:
    throttleDuration: "10s"
    # ...
```

```toml tab="File (TOML)"
[providers.kubernetesCRD]
  throttleDuration = "10s"
  # ...
```

```bash tab="CLI"
--providers.kubernetescrd.throttleDuration=10s
```

### `allowEmptyServices`

_Optional, Default: false_

If the parameter is set to `true`,
it allows the creation of an empty [servers load balancer](../routing/services/index.md#servers-load-balancer) if the targeted Kubernetes service has no endpoints available.
With IngressRoute resources,
this results in `503` HTTP responses instead of `404` ones.

```yaml tab="File (YAML)"
providers:
  kubernetesCRD:
    allowEmptyServices: true
    # ...
```

```toml tab="File (TOML)"
[providers.kubernetesCRD]
  allowEmptyServices = true
  # ...
```

```bash tab="CLI"
--providers.kubernetesCRD.allowEmptyServices=true
```

### `allowCrossNamespace`

_Optional, Default: false_

If the parameter is set to `true`,
IngressRoute are able to reference resources in namespaces other than theirs.

```yaml tab="File (YAML)"
providers:
  kubernetesCRD:
    allowCrossNamespace: true
    # ...
```

```toml tab="File (TOML)"
[providers.kubernetesCRD]
  allowCrossNamespace = true
  # ...
```

```bash tab="CLI"
--providers.kubernetescrd.allowCrossNamespace=true
```

### `allowExternalNameServices`

_Optional, Default: false_

If the parameter is set to `true`, IngressRoutes are able to reference ExternalName services.

```yaml tab="File (YAML)"
providers:
  kubernetesCRD:
    allowExternalNameServices: true
    # ...
```

```toml tab="File (TOML)"
[providers.kubernetesCRD]
  allowExternalNameServices = true
  # ...
```

```bash tab="CLI"
--providers.kubernetescrd.allowexternalnameservices=true
```

### `nativeLBByDefault`

_Optional, Default: false_

Defines whether to use Native Kubernetes load-balancing mode by default.
For more information, please check out the IngressRoute `nativeLB` option [documentation](../routing/providers/kubernetes-crd.md#load-balancing).

```yaml tab="File (YAML)"
providers:
  kubernetesCRD:
    nativeLBByDefault: true
    # ...
```

```toml tab="File (TOML)"
[providers.kubernetesCRD]
  nativeLBByDefault = true
  # ...
```

```bash tab="CLI"
--providers.kubernetescrd.nativeLBByDefault=true
```

## Full Example

For additional information, refer to the [full example](../user-guides/crd-acme/index.md) with Let's Encrypt.


