package ds

type Set[E comparable]map[E] struct{}

func NewSet[E comparable](items  ...E) Set[E] {
	set := make(Set[E])
	for _, item := range items {
		set.Add(item)
	}
	return set
}
func (s Set[E]) Add(item E) {
	s[item] = struct{}{}
}
func (s Set[E]) Remove(item E) {
	delete(s, item)
}
func (s Set[E]) Contains(item E) bool {
	_, contains := s[item]
	return contains
}