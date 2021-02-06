package main

import (
	"io"
	"log"
	"os"

	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/goccy/go-graphviz"
	"github.com/guillotjulien/eopkg-graph/internal"
)

type generateArgs struct {
	Package string `desc:"Name of the package for which graph is generated"`
	Path    string `desc:"Saving path of the generated graph"`
}

type generateFlags struct {
	SingleNodes bool   `long:"single-nodes" desc:"Display nodes without dependencies as a node"`
	Format      string `long:"format" short:"f" desc:"Graph output format, can be png (default) or html"`
}

func main() {
	root := &cmd.Root{
		Name:    "eopkg-graph",
		Short:   "",
		Version: "0.0.1",
	}

	generateCmd := cmd.Sub{
		Name:  "generate",
		Short: "Generate graph of a package dependencies from eopkg metadata",
		Args:  &generateArgs{},
		Flags: &generateFlags{},
		Run:   generateRun,
	}

	cmd.Register(&generateCmd)

	root.Run()
}

func generateRun(r *cmd.Root, s *cmd.Sub) {
	args := s.Args.(*generateArgs)
	flags := s.Flags.(*generateFlags)

	p, err := internal.NewPackage(args.Package)
	if err != nil {
		log.Fatal(err)
	}

	d, err := p.DependencyGraph()
	if err != nil {
		log.Fatal(err)
	}

	switch flags.Format {
	case "html":
		page := components.NewPage()
		page.AddCharts(d.HTML(args.Package, flags.SingleNodes))

		f, err := os.Create(args.Path)
		if err != nil {
			panic(err)

		}
		page.Render(io.MultiWriter(f))
	default:
		gviz, g, err := d.Graphviz(flags.SingleNodes)
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
}
