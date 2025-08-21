package main

type dsu struct {
	parent, rank []int
}

func (d *dsu) MakeSet(v int) {
	d.parent[v] = v
	d.rank[v] = 0
}

func (d *dsu) Init(n int) {
	d.parent = make([]int, n+1)
	d.rank = make([]int, n+1)
	for v := range n + 1 {
		d.MakeSet(v)
	}
}

func (d *dsu) FindSet(v int) int {
	if v == d.parent[v] {
		return v
	}
	d.parent[v] = d.FindSet(d.parent[v])
	return d.parent[v]
}

func (d *dsu) UnionSets(a, b int) {
	a, b = d.FindSet(a), d.FindSet(b)
	if a != b {
		if d.rank[a] < d.rank[b] {
			a, b = b, a
		}
		d.parent[b] = a
		if d.rank[a] == d.rank[b] {
			d.rank[a]++
		}
	}
}
