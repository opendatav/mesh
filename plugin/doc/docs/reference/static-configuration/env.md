---
title: "Mesh Environment Variables Documentation"
description: "Reference the environment variables for static configuration in Mesh Proxy. Read the technical documentation."
---

# Static Configuration: Environment variables

!!! warning "Environment Variable Casing"

    Mesh normalizes the environment variable key-value pairs by lowercasing them.
    This means that when you interpolate a string in an environment variable's name,
    that string will be treated as lowercase, regardless of its original casing.

    For example, assuming you have set environment variables as follows:

    ```bash
        export Mesh_ENTRYPOINTS_WEB=true
        export Mesh_ENTRYPOINTS_WEB_ADDRESS=:80

        export Mesh_CERTIFICATESRESOLVERS_myResolver=true
        export Mesh_CERTIFICATESRESOLVERS_myResolver_ACME_CASERVER=....
    ```
    
    Although the Entrypoint is named `WEB` and the Certificate Resolver is named `myResolver`, 
    they have to be referenced respectively as `web`, and `myresolver` in the configuration.

