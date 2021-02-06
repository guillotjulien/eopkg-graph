package internal

import (
	"fmt"
	"sync"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

// A Dependency has various attributes which help determine what needs to
// be installed when updating or installing the package.
type Dependency struct {
	Name string `xml:",innerxml"`
}

// DependencyGraph represents a graph structure that contains relationship between packages
type DependencyGraph struct {
	nodes []*Dependency
	edges map[*Dependency][]*Dependency
	sync.RWMutex
}

// AddNode adds a node to the graph
func (g *DependencyGraph) AddNode(n *Dependency) {
	g.Lock()
	g.nodes = append(g.nodes, n)
	g.Unlock()
}

// AddEdge adds an edge to the graph
func (g *DependencyGraph) AddEdge(n1, n2 *Dependency) {
	g.Lock()
	if g.edges == nil {
		g.edges = make(map[*Dependency][]*Dependency, 0)
	}

	g.edges[n1] = append(g.edges[n1], n2)
	// g.edges[n2] = append(g.edges[n2], n1)
	g.Unlock()
}

func (g *DependencyGraph) String() {
	g.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].Name + " -> "
		near := g.edges[g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].Name + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.RUnlock()
}

// Graphviz returns a populated instance of Graphviz graph
func (g *DependencyGraph) Graphviz(showSingleNodes bool) (gviz *graphviz.Graphviz, graph *cgraph.Graph, err error) {
	g.RLock()

	nodes := make(map[string]*cgraph.Node)

	gviz = graphviz.New()
	graph, err = gviz.Graph()
	if err != nil {
		return
	}

	for _, n := range g.nodes {
		// In the case the node doesn't have subnodes, we don't display it in
		// the graph.
		if len(g.edges[n]) == 0 && !showSingleNodes {
			continue
		}

		var node *cgraph.Node
		node, err = graph.CreateNode(n.Name)
		if err != nil {
			return
		}
		node.SetShape(cgraph.RectShape)
		node.SetLabel(fmt.Sprintf("%s\n", n.Name))

		nodes[n.Name] = node
	}

	for node, e := range g.edges {
		n1 := nodes[node.Name]

		for _, n := range e {
			n1.SetLabel(fmt.Sprintf("%s\n%s", n1.Get("label"), n.Name))

			if n2, ok := nodes[n.Name]; ok {
				_, err = graph.CreateEdge("", n1, n2)
				if err != nil {
					return
				}
			}
		}
	}

	g.RUnlock()
	return
}

// HTML returns a go-charts interactive graph ready to be saved as html
func (g *DependencyGraph) HTML(rootPackage string, showSingleNodes bool) *charts.Graph {
	g.RLock()

	nodes := make(map[string]opts.GraphNode)

	graph := charts.NewGraph()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: fmt.Sprintf("Runtime Dependencies %s", rootPackage)}),
	)

	graphNodes := make([]opts.GraphNode, 0)
	graphEdges := make([]opts.GraphLink, 0)

	for _, n := range g.nodes {
		// In the case the node doesn't have subnodes, we don't display it in
		// the graph.
		if len(g.edges[n]) == 0 && !showSingleNodes {
			continue
		}

		node := opts.GraphNode{Name: n.Name}

		if node.Name == rootPackage {
			node.SymbolSize = 20
			node.Symbol = "diamond"
			node.X = 250 // FIXME: Should depends on the size of the graph
			node.Y = 250
			node.Fixed = true
			node.ItemStyle = &opts.ItemStyle{
				Color: "#50a3ba",
			}
		}

		graphNodes = append(graphNodes, node)

		nodes[n.Name] = node
	}

	for node, e := range g.edges {
		n1 := nodes[node.Name]

		for _, n := range e {
			if n2, ok := nodes[n.Name]; ok {
				edge := opts.GraphLink{Source: n1.Name, Target: n2.Name}
				graphEdges = append(graphEdges, edge)
			}
		}
	}

	graph.AddSeries("graph", graphNodes, graphEdges)
	graph.SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{
			Show:     true,
			Position: "right",
		}),
		charts.WithGraphChartOpts(opts.GraphChart{
			Layout:             "force",
			Roam:               true,
			FocusNodeAdjacency: true,
			Force:              &opts.GraphForce{Repulsion: 10000},
		}),
	)

	g.RUnlock()

	return graph
}
