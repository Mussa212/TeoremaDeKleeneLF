package internal

type State struct {
	ID        int
	Epsilon   []*State
	Trans     map[rune][]*State
	Accepting bool
}
type NFA struct{ Start, Accept *State }

func NewState(id int) *State {
	return &State{ID: id, Trans: make(map[rune][]*State)}
}

func BuildNFA(postfix []Token) *NFA {
	var st []*NFA
	id := 0
	pop := func() *NFA { n := st[len(st)-1]; st = st[:len(st)-1]; return n }
	for _, t := range postfix {
		switch t.Type {
		case Char:
			s1, s2 := NewState(id), NewState(id+1)
			id += 2
			s1.Trans[t.Val] = []*State{s2}
			s2.Accepting = true
			st = append(st, &NFA{s1, s2})
		case Concat:
			b, a := pop(), pop()
			a.Accept.Epsilon = append(a.Accept.Epsilon, b.Start)
			a.Accept.Accepting = false
			st = append(st, &NFA{a.Start, b.Accept})
		case Union:
			b, a := pop(), pop()
			s := NewState(id)
			id++
			f := NewState(id)
			id++
			s.Epsilon = append(s.Epsilon, a.Start, b.Start)
			a.Accept.Epsilon = append(a.Accept.Epsilon, f)
			b.Accept.Epsilon = append(b.Accept.Epsilon, f)
			f.Accepting = true
			st = append(st, &NFA{s, f})
		case Star:
			a := pop()
			s := NewState(id)
			id++
			f := NewState(id)
			id++
			s.Epsilon = append(s.Epsilon, a.Start, f)
			a.Accept.Epsilon = append(a.Accept.Epsilon, a.Start, f)
			f.Accepting = true
			st = append(st, &NFA{s, f})
		}
	}
	return st[0]
}
