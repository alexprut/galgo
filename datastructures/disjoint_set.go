package datastructures

type DisjointSet struct {
	sets map[interface{}]*Element
}

func (dj *DisjointSet) MakeSet(value interface{}) *Element {
	var e = NewElement(value)
	dj.sets[value] = &e
	return &e
}

func (dj *DisjointSet) Union(a *Element, b *Element) {
	var parentA = dj.FindSet(a)
	var parentB = dj.FindSet(b)

	if parentA.value == parentB.value {
		return
	}

	if parentA.rank > parentB.rank {
		parentB.parent = parentA
	} else {
		parentA.parent = parentB
		if parentA.rank == parentB.rank {
			parentB.rank = parentB.rank + 1
		}
	}
}

func (dj *DisjointSet) FindSetElement(s *Element) *Element {
	if s.parent.value == s.value {
		return s
	}

	s.parent = dj.FindSetElement(s.parent)
	return s.parent
}

func (dj *DisjointSet) FindSet(value interface{}) *Element {
	return dj.FindSet(dj.GetElement(value))
}

func (dj *DisjointSet) GetElement(value interface{}) *Element {
	var e, ok = dj.sets[value]
	if ok {
		return e
	} else {
		return nil
	}
}

type Element struct {
	parent *Element
	rank   int
	value  interface{}
}

func NewElement(value interface{}) Element {
	var e = Element{nil, 0, value}
	e.parent = &e
	return e
}
