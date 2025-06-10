---
title: "Mesh Plugins Documentation"
description: "Learn how to use Mesh Plugins. Read the technical documentation."
---

# Mesh Plugins and the Plugin Catalog

Plugins are a powerful feature for extending Mesh with custom features and behaviors.
The [Plugin Catalog](https://plugins.Mesh.io/) is a software-as-a-service (SaaS) platform that provides an exhaustive list of the existing plugins.

??? note "Plugin Catalog Access"
    You can reach the [Plugin Catalog](https://plugins.Mesh.io/) from the Mesh Dashboard using the `Plugins` menu entry.

To add a new plugin to a Mesh instance, you must change that instance's static configuration.
Each plugin's **Install** section provides a static configuration example.
Many plugins have their own section in the Mesh dynamic configuration.

To learn more about Mesh plugins, consult the [documentation](https://plugins.Mesh.io/install).

!!! danger "Experimental Features"
    Plugins can change the behavior of Mesh in unforeseen ways.
    Exercise caution when adding new plugins to production Mesh instances.

## Build Your Own Plugins

Mesh users can create their own plugins and share them with the community using the Plugin Catalog.

Mesh will load plugins dynamically.
They need not be compiled, and no complex toolchain is necessary to build them. 
The experience of implementing a Mesh plugin is comparable to writing a web browser extension.

To learn more about Mesh plugin creation, please refer to the [developer documentation](https://plugins.Mesh.io/create).


