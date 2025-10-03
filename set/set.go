package set

type Set[T comparable] struct {
	set map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{
		set: make(map[T]struct{}),
	}
}

func (s *Set[T]) Set(val T) {
	s.set[val] = struct{}{}
}

func (s *Set[T]) Exists(val T) bool {
	_, exists := s.set[val]
	return exists
}

func (s *Set[T]) Erase(val T) {
	delete(s.set, val)
}
