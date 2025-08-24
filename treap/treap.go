package main

import (
	"cmp"
	"math/rand"
	"time"
)

type Treap[T cmp.Ordered] struct {
	root *node[T]
	rng  *rand.Rand
}

type node[T cmp.Ordered] struct {
	key   T
	sum   T
	prior int
	size  int
	left  *node[T]
	right *node[T]
}

func newNode[T cmp.Ordered](key T, rng *rand.Rand) *node[T] {
	return &node[T]{
		key:   key,
		prior: rng.Int(),
		size:  1,
		sum:   key,
	}
}

func newTreap[T cmp.Ordered](r *rand.Rand) *Treap[T] {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return &Treap[T]{
		root: nil,
		rng:  r,
	}
}

func size[T cmp.Ordered](n *node[T]) int {
	if n == nil {
		return 0
	}
	return n.size
}

func update[T cmp.Ordered](n *node[T]) {
	if n != nil {
		n.size = 1 + size(n.left) + size(n.right)
		// n.sum = n.key + ...
	}
}

func split[T cmp.Ordered](n *node[T], x T) (*node[T], *node[T]) {
	if n == nil {
		return nil, nil
	}
	if n.key < x {
		l, r := split(n.right, x)
		n.right = l
		update(n)
		return n, r
	} else {
		l, r := split(n.left, x)
		n.left = r
		update(n)
		return l, n
	}
}

func splitSize[T cmp.Ordered](n *node[T], k int) (*node[T], *node[T]) {
	if n == nil {
		return nil, nil
	}
	if size(n.right) >= k {
		l, r := splitSize(n.right, k)
		n.right = l
		update(n)
		return n, r
	} else {
		l, r := splitSize(n.left, k-1-size(n.right))
		n.left = r
		update(n)
		return l, n
	}
}

func merge[T cmp.Ordered](a, b *node[T]) *node[T] {
	if a == nil || b == nil {
		if a != nil {
			return a
		}
		return b
	}
	if a.prior >= b.prior {
		a.right = merge(a.right, b)
		update(a)
		return a
	} else {
		b.left = merge(a, b.left)
		update(b)
		return b
	}
}

func (t *Treap[T]) Insert(x T) {
	n := newNode(x, t.rng)
	l, r := split(t.root, x)
	t.root = merge(merge(l, n), r)
}

func (t *Treap[T]) Erase(x T) {
	t.root = erase(t.root, x)
}

func erase[T cmp.Ordered](n *node[T], x T) *node[T] {
	if n == nil {
		return nil
	}
	if n.key < x {
		n.right = erase(n.right, x)
	} else if n.key > x {
		n.left = erase(n.left, x)
	} else {
		return merge(n.left, n.right)
	}
	update(n)
	return n
}

func (t *Treap[T]) Find(x T) *T {
	cur := t.root
	for cur != nil {
		if cur.key < x {
			cur = cur.right
		} else if cur.key > x {
			cur = cur.left
		} else {
			return &cur.key
		}
	}
	return nil
}

func (t *Treap[T]) LowerBound(x T) *T {
	cur := t.root
	var res *node[T]
	for cur != nil {
		if cur.key < x {
			cur = cur.right
		} else {
			res = cur
			cur = cur.left
		}
	}
	if res != nil {
		return &res.key
	}
	return nil
}

func (t *Treap[T]) UpperBound(x T) *T {
	cur := t.root
	var res *node[T]
	for cur != nil {
		if cur.key <= x {
			cur = cur.right
		} else {
			res = cur
			cur = cur.left
		}
	}
	if res != nil {
		return &res.key
	}
	return nil
}

func (t *Treap[T]) Size() int {
	return size(t.root)
}

func (t *Treap[T]) Min() *T {
	cur := t.root
	if cur == nil {
		return nil
	}
	for cur.left != nil {
		cur = cur.left
	}
	return &cur.key
}

func (t *Treap[T]) Max() *T {
	cur := t.root
	if cur == nil {
		return nil
	}
	for cur.right != nil {
		cur = cur.right
	}
	return &cur.key
}
