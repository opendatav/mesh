---
title: "Certificates Resolver"
description: "Automatic Certificate Management using Let's Encrypt/Vault and Tailscale."
---


In Mesh, TLS Certificates can be generated using Certificates Resolvers.

In Mesh, two certificate resolvers exist:

- [`acme`](./acme.md): It allows generating ACME certificates stored in a file (not distributed).
- [`tailscale`](./tailscale.md): It allows provisioning TLS certificates for internal Tailscale services.

The Certificates resolvers are defined in the static configuration.

!!! note Referencing a certificate resolver
    Defining a certificate resolver does not imply that routers are going to use it automatically.
    Each router or entrypoint that is meant to use the resolver must explicitly reference it.


