---
title: "Mesh Installation Documentation"
description: "There are several flavors to choose from when installing Mesh Proxy. Get started with Mesh Proxy, and read the technical documentation."
---

# Install Mesh

You can install Mesh with the following flavors:

* [Use the official Docker image](./#use-the-official-docker-image)
* [Use the Helm Chart](./#use-the-helm-chart)
* [Use the binary distribution](./#use-the-binary-distribution)
* [Compile your binary from the sources](./#compile-your-binary-from-the-sources)

## Use the Official Docker Image

Choose one of the [official Docker images](https://hub.docker.com/_/Mesh) and run it with one sample configuration file:

* [YAML](https://raw.githubusercontent.com/Mesh/Mesh/v3.4/Mesh.sample.yml)
* [TOML](https://raw.githubusercontent.com/Mesh/Mesh/v3.4/Mesh.sample.toml)

```shell
docker run -d -p 8080:8080 -p 80:80 \
    -v $PWD/Mesh.yml:/etc/Mesh/Mesh.yml Mesh:v3.4
```

For more details, go to the [Docker provider documentation](../providers/docker.md)

!!! tip

    * Prefer a fixed version than the latest that could be an unexpected version.
    ex: `Mesh:v3.4`
    * Docker images are based from the [Alpine Linux Official image](https://hub.docker.com/_/alpine).
    * Any orchestrator using docker images can fetch the official Mesh docker image.

## Use the Helm Chart

Mesh can be installed in Kubernetes using the Helm chart from [](https://github.com/Mesh/mesh-helm-chart).

Ensure that the following requirements are met:

* Kubernetes 1.22+
* Helm version 3.9+ is [installed](https://helm.sh/docs/intro/install/)

Add Mesh Labs chart repository to Helm:

```bash
helm repo add Mesh https://Mesh.github.io/charts
```

You can update the chart repository by running:

```bash
helm repo update
```

And install it with the Helm command line:

```bash
helm install Mesh Mesh/Mesh
```

!!! tip "Helm Features"

    All [Helm features](https://helm.sh/docs/intro/using_helm/) are supported.

    Examples are provided [here](https://github.com/Mesh/mesh-helm-chart/blob/master/EXAMPLES.md).

    For instance, installing the chart in a dedicated namespace:

    ```bash tab="Install in a Dedicated Namespace"
    kubectl create ns mesh-v2
    # Install in the namespace "mesh-v2"
    helm install --namespace=mesh-v2 \
        Mesh Mesh/Mesh
    ```

??? example "Installing with Custom Values"

    You can customize the installation by specifying custom values,
    as with [any helm chart](https://helm.sh/docs/intro/using_helm/#customizing-the-chart-before-installing).

    All parameters are documented in the default [`values.yaml`](https://github.com/Mesh/mesh-helm-chart/blob/master/Mesh/values.yaml).

    You can also set Mesh command line flags using `additionalArguments`.
    Example of installation with logging set to `DEBUG`:

    ```bash tab="Using Helm CLI"
    helm install --namespace=mesh-v2 \
        --set="additionalArguments={--log.level=DEBUG}" \
        Mesh Mesh/Mesh
    ```

    ```yml tab="With a custom values file"
    # File custom-values.yml
    ## Install with "helm install --values=./custom-values.yml Mesh Mesh/Mesh
    additionalArguments:
      - "--log.level=DEBUG"
    ```

## Use the Binary Distribution

Grab the latest binary from the [releases](https://github.com/Mesh/Mesh/releases) page.

??? info "Check the integrity of the downloaded file"

    ```bash tab="Linux"
    # Compare this value to the one found in mesh-${Mesh_version}_checksums.txt
    sha256sum ./Mesh_${Mesh_version}_linux_${arch}.tar.gz
    ```

    ```bash tab="macOS"
    # Compare this value to the one found in mesh-${Mesh_version}_checksums.txt
    shasum -a256 ./Mesh_${Mesh_version}_darwin_amd64.tar.gz
    ```

    ```powershell tab="Windows PowerShell"
    # Compare this value to the one found in mesh-${Mesh_version}_checksums.txt
    Get-FileHash ./Mesh_${Mesh_version}_windows_${arch}.zip -Algorithm SHA256
    ```

??? info "Extract the downloaded archive"

    ```bash tab="Linux"
    tar -zxvf Mesh_${Mesh_version}_linux_${arch}.tar.gz
    ```

    ```bash tab="macOS"
    tar -zxvf ./Mesh_${Mesh_version}_darwin_amd64.tar.gz
    ```

    ```powershell tab="Windows PowerShell"
    Expand-Archive Mesh_${Mesh_version}_windows_${arch}.zip
    ```

And run it:

```bash
./Mesh --help
```

## Compile your Binary from the Sources

All the details are available in the [Contributing Guide](../contributing/building-testing.md)

