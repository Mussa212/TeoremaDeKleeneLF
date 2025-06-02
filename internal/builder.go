package internal

type State struct {
	ID        int
	Lambda    []*State
	Trans     map[rune][]*State
	Accepting bool
}
type AFN struct{ Start, Accept *State }

func NewState(id int) *State {
	return &State{ID: id, Trans: make(map[rune][]*State)}
}

func BuildAFN(postfix []Token) *AFN {
	var st []*AFN
	id := 0
	pop := func() *AFN { n := st[len(st)-1]; st = st[:len(st)-1]; return n }
	for _, t := range postfix {
		switch t.Type {
		case Char:
			s1, s2 := NewState(id), NewState(id+1)
			id += 2
			s1.Trans[t.Val] = []*State{s2}
			s2.Accepting = true
			st = append(st, &AFN{s1, s2})
		case Concat:
			b, a := pop(), pop()
			a.Accept.Lambda = append(a.Accept.Lambda, b.Start)
			a.Accept.Accepting = false
			st = append(st, &AFN{a.Start, b.Accept})
		case Union:
			b, a := pop(), pop()
			s := NewState(id)
			id++
			f := NewState(id)
			id++
			s.Lambda = append(s.Lambda, a.Start, b.Start)
			a.Accept.Lambda = append(a.Accept.Lambda, f)
			b.Accept.Lambda = append(b.Accept.Lambda, f)
			f.Accepting = true
			st = append(st, &AFN{s, f})
		case Star:
			a := pop()
			s := NewState(id)
			id++
			f := NewState(id)
			id++
			s.Lambda = append(s.Lambda, a.Start, f)
			a.Accept.Lambda = append(a.Accept.Lambda, a.Start, f)
			f.Accepting = true
			st = append(st, &AFN{s, f})
		}
	}
	return st[0]
}
