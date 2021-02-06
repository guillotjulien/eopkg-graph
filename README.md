# eopkg-graph

A CLI utility to generate the runtime dependency graph of a package from its eopkg metadata.

![GNOME-shell dependency graph](https://github.com/guillotjulien/eopkg-graph/blob/main/examples/graph-gnome-shell.png "GNOME-shell dependency graph")
<figcaption style="text-align: center; font-weight: bold;">GNOME-shell dependency graph</figcaption>

## Installation

`make && make install`

## Usage

The command `eopkg-graph generate` will generate the dependency graph of a package.
To reduce visual complexity, we do not create nodes for packages that do not pull anything else, they will be listed under the package that pulled them.

`eopkg-graph generate <package> <graph_path>`

Example: `eopkg-graph generate budgie-desktop budgie-desktop-graph.png`