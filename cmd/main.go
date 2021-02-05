package main

import (
	"log"

	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/goccy/go-graphviz"
	"github.com/guillotjulien/eopkg-deps/internal"
)

type graphArgs struct {
	Package string `desc:"Name of the package for which graph is generated"`
	Path    string `desc:"Saving path of the generated graph"`
}

func main() {
	root := &cmd.Root{
		Name:    "eopkg-deps",
		Short:   "",
		Version: "0.0.1",
	}

	graphCmd := cmd.Sub{
		Name:  "graph",
		Short: "Generate graph of a package dependencies from eopkg metadata",
		Args:  &graphArgs{},
		Run:   graphRun,
	}

	cmd.Register(&graphCmd)

	root.Run()
}

func graphRun(r *cmd.Root, s *cmd.Sub) {
	args := s.Args.(*graphArgs)

	p, err := internal.NewPackage(args.Package)
	if err != nil {
		log.Fatal(err)
	}

	d, err := p.DependencyGraph()
	if err != nil {
		log.Fatal(err)
	}

	gviz, g, err := d.Graphviz()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := g.Close(); err != nil {
			log.Fatal(err)
		}
		gviz.Close()
	}()

	if err := gviz.RenderFilename(g, graphviz.PNG, args.Path); err != nil {
		log.Fatal(err)
	}
}
