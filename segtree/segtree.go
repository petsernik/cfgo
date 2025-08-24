package segtree

type node[T comparable] struct {
	sum      T
	deferred T
	needSet  bool
}

type SegTree[T comparable] struct {
	tree    []node[T]
	n       int
	addOp   func(*T, T, int, int)
	setOp   func(*T, T, int, int)
	merge   func(T, T) T
	neutral T
}

func NewSegTree[T comparable](data []T,
	addOp func(*T, T, int, int),
	setOp func(*T, T, int, int),
	mergeOp func(T, T) T,
	neutral T) *SegTree[T] {

	n := len(data)
	st := &SegTree[T]{
		tree:    make([]node[T], 4*n+1),
		n:       n,
		addOp:   addOp,
		setOp:   setOp,
		merge:   mergeOp,
		neutral: neutral,
	}

	for i := range st.tree {
		st.tree[i].sum = neutral
		st.tree[i].deferred = neutral
	}

	st.build(1, 0, n, data)
	return st
}

func (st *SegTree[T]) build(v int, l int, r int, data []T) {
	if l == r-1 {
		st.tree[v].sum = data[l]
		return
	}
	m := (l + r) / 2
	st.build(2*v, l, m, data)
	st.build(2*v+1, m, r, data)
	st.tree[v].sum = st.merge(st.tree[2*v].sum, st.tree[2*v+1].sum)
}

func (st *SegTree[T]) push(v int, l int, r int) {
	m := (l + r) / 2
	node := &st.tree[v]
	if node.needSet {
		st.deferredSet(2*v, l, m, node.deferred)
		st.deferredSet(2*v+1, m, r, node.deferred)
	} else if node.deferred != st.neutral {
		st.deferredAdd(2*v, l, m, node.deferred)
		st.deferredAdd(2*v+1, m, r, node.deferred)
	}
	node.needSet = false
	node.deferred = st.neutral
}

func (st *SegTree[T]) deferredAdd(v int, l int, r int, x T) {
	node := &st.tree[v]
	node.deferred = st.merge(node.deferred, x)
	if node.needSet {
		st.setOp(&node.sum, node.deferred, l, r)
	} else {
		st.addOp(&node.sum, x, l, r)
	}
}

func (st *SegTree[T]) deferredSet(v int, l int, r int, x T) {
	node := &st.tree[v]
	node.needSet = true
	node.deferred = x
	st.setOp(&node.sum, x, l, r)
}

func (st *SegTree[T]) Add(l int, r int, x T) {
	st.operation(1, 0, st.n, l, r+1, x, st.deferredAdd)
}

func (st *SegTree[T]) Set(l int, r int, x T) {
	st.operation(1, 0, st.n, l, r+1, x, st.deferredSet)
}

func (st *SegTree[T]) operation(v int, l int, r int, ql int, qr int, x T, deferredOp func(int, int, int, T)) {
	if qr <= l || r <= ql {
		return
	}
	if ql <= l && r <= qr {
		deferredOp(v, l, r, x)
		return
	}
	m := (l + r) / 2
	st.push(v, l, r)
	st.operation(2*v, l, m, ql, qr, x, deferredOp)
	st.operation(2*v+1, m, r, ql, qr, x, deferredOp)
	st.tree[v].sum = st.merge(st.tree[2*v].sum, st.tree[2*v+1].sum)
}

func (st *SegTree[T]) Query(l int, r int) T {
	return st.query(1, 0, st.n, l, r+1)
}

func (st *SegTree[T]) query(v int, l int, r int, ql int, qr int) T {
	if qr <= l || r <= ql {
		return st.neutral
	}
	if ql <= l && r <= qr {
		return st.tree[v].sum
	}
	m := (l + r) / 2
	st.push(v, l, r)
	left := st.query(2*v, l, m, ql, qr)
	right := st.query(2*v+1, m, r, ql, qr)
	return st.merge(left, right)
}
