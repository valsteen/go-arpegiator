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

func (s Set) Subtract(s2 Set) []Element {
	result := make([]Element, 0, len(s))

	for key, e := range s {
		if _, found := s2[key]; !found {
			result = append(result, e)
		}
	}

	return result
}

func (s Set) Compare(s2 Set) (added []Element, removed []Element) {
	return s2.Subtract(s), s.Subtract(s2)
}

func (s Set) String() string {
	elements := make([]Element, 0, len(s))
	for _, element := range s {
		elements = append(elements, element)
	}
	return fmt.Sprintf("%v", elements)
}
