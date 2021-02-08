package internal

import "testing"

func BenchmarkDependencyGraph(b *testing.B) {
	p, _ := NewPackage("pango")
	for i := 0; i < b.N; i++ {
		p.DependencyGraph()
	}
}

func BenchmarkDependencyGraphDFS(b *testing.B) {
	p, _ := NewPackage("pango")
	for i := 0; i < b.N; i++ {
		p.DependencyGraphDFS()
	}
}

func BenchmarkDependencyGraphBFS(b *testing.B) {
	p, _ := NewPackage("pango")
	for i := 0; i < b.N; i++ {
		p.DependencyGraphBFS()
	}
}

func BenchmarkDependencyGraphThread(b *testing.B) {
	p, _ := NewPackage("pango")
	for i := 0; i < b.N; i++ {
		p.DependencyGraphThread()
	}
}

func BenchmarkDependencyGraphDFSThread(b *testing.B) {
	p, _ := NewPackage("pango")
	for i := 0; i < b.N; i++ {
		p.DependencyGraphDFSThread()
	}
}

func BenchmarkDependencyGraphBFSThread(b *testing.B) {
	p, _ := NewPackage("pango")
	for i := 0; i < b.N; i++ {
		p.DependencyGraphBFSThread()
	}
}
