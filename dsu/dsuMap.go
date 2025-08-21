package main

type dsuMap[K comparable] struct {
	parent map[K]K
	rank   map[K]int
}

func (d *dsuMap[K]) Init() {
	d.parent = make(map[K]K)
	d.rank = make(map[K]int)
}

func (d *dsuMap[K]) MakeSet(v K) {
	d.parent[v] = v
	d.rank[v] = 0
}

func (d *dsuMap[K]) FindSet(v K) K {
	if v == d.parent[v] {
		return v
	}
	d.parent[v] = d.FindSet(d.parent[v])
	return d.parent[v]
}

func (d *dsuMap[K]) UnionSets(a, b K) {
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
