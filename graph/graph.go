package main

import (
	"fmt"

	"example.com/m/io"
)

type IO struct {
	*io.IO
}

type graph struct {
	used  []bool
	edges [][]int
	dist  []int
}

func (g *graph) Init(n int) {
	g.used = make([]bool, n+1)
	g.edges = make([][]int, n+1)
	g.dist = make([]int, n+1)
}

// Indexes starts from 1
func (g *graph) ReadTree(io *IO) {
	n := len(g.used) - 1
	for i := 2; i <= n; i++ {
		var v int
		io.Read(&v)
		g.edges[i] = append(g.edges[i], v)
		g.edges[v] = append(g.edges[v], i)
	}
}

func (g *graph) Read(io *IO, m int) {
	for range m {
		var v, i int
		io.Read(&v, &i)
		g.edges[i] = append(g.edges[i], v)
		g.edges[v] = append(g.edges[v], i)
	}
}

func (g *graph) setZeroUsed() {
	g.used = make([]bool, len(g.used))
}

func (g *graph) dfs(n int) {
	g.used[n] = true
	for _, v := range g.edges[n] {
		if !g.used[v] {
			g.dfs(v)
		}
	}
}

func (g *graph) bfs(start int) {
	q := []int{start}
	g.dist[start] = 0
	for len(q) > 0 {
		v := q[0]
		g.used[v] = true
		q = q[1:]
		for _, u := range g.edges[v] {
			if !g.used[u] {
				q = append(q, u)
				g.dist[u] = g.dist[v] + 1
			}
		}
	}
}

func main() {
	io := &IO{io.NewIO()}
	g := graph{}
	g.Init(4)
	g.ReadTree(io)
	g.bfs(1)
	for i, d := range g.dist {
		fmt.Printf("%d: %d\n", i, d)
	}
}
