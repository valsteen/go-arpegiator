package set

type Element interface {
	Less(Element) bool
}

type Set []Element

func equals(e1, e2 Element) bool {
	return !e1.Less(e2) && !e2.Less(e1)
}

func (s Set) Delete(e Element) (out Set) {
	out = make(Set, 0, len(s))
	for _, e2 := range s {
		if !equals(e, e2) {
			out = append(out, e2)
		}
	}
	return out
}

func (s Set) Add(e Element) Set {
	out := make(Set, 0, len(s)+1)
	var i int
	for ; i < len(s); i++ {
		if s[i].Less(e) {
			out = append(out, s[i])
		} else if e.Less(s[i]) {
			out = append(out, e)
			break
		} else {
			// equals, replace with newest, skip to next
			out = append(out, e)
			i++
			break
		}
	}
	for ; i < len(s); i++ {
		out = append(out, s[i])
	}
	if len(out) < len(s)+1 {
		out = append(out, e)
	}
	return out
}

func (s Set) Subtract(s2 Set) Set {
	out := make(Set, len(s))
	copy(out, s)

	for _, e := range s2 {
		out = out.Delete(e)
	}

	return out
}

func (s Set) Compare(s2 Set) (added Set, removed Set) {
	return s2.Subtract(s), s.Subtract(s2)
}

func (s Set) Length() int {
	return len(s)
}

func (s Set) At(i int) Element {
	return s[i]
}
