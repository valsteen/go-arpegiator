package set

import "fmt"

type Hash string

type Element interface {
	Hash() Hash
}

type Set map[Hash]Element

func (s Set) Delete(e Element) {
	delete(s, e.Hash())
}

func (s Set) Add(e Element) {
	s[e.Hash()] = e
}

func (s Set) Diff(s2 Set) (added []Element, removed []Element) {
	max := len(s)
	if len(s2) > max {
		max = len(s2)
	}

	added = make([]Element, 0, max)
	removed = make([]Element, 0, max)

	for key, e := range s {
		if _, found := s2[key]; !found {
			removed = append(removed, e)
		}
	}

	for key, e := range s2 {
		if _, found := s[key]; !found {
			added = append(added, e)
		}
	}

	return
}


func (s Set) String() string {
	elements := make([]Element, 0, len(s))
	for _, element := range s {
		elements = append(elements, element)
	}
	return fmt.Sprintf("%v", elements)
}
