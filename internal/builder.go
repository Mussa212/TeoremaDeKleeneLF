package internal

type State struct {
	ID        int
	Lambda    []*State          // transiciones espontáneas (ε / λ)
	Trans     map[rune][]*State // transiciones etiquetadas
	Accepting bool
}
type AFN struct {
	Start, Accept *State
}

func NewState(id int) *State {
	return &State{ID: id, Trans: make(map[rune][]*State)}
}

func BuildAFN(postfix []Token) *AFN {
	var st []*AFN
	id := 0
	pop := func() *AFN {
		n := st[len(st)-1]
		st = st[:len(st)-1]
		return n
	}

	for _, t := range postfix {
		switch t.Type {
		case Char:
			// NFA elemental para un símbolo (ej. 'a'):
			//   s1 -a-> s2 (s2 es estado de aceptación)
			s1, s2 := NewState(id), NewState(id+1)
			id += 2
			s1.Trans[t.Val] = []*State{s2}
			s2.Accepting = true
			st = append(st, &AFN{s1, s2})

		case Lambda:
			// NFA elemental para 'λ':
			//   s1 -ε-> s2 (s2 es estado de aceptación)
			s1, s2 := NewState(id), NewState(id+1)
			id += 2
			s1.Lambda = append(s1.Lambda, s2)
			s2.Accepting = true
			st = append(st, &AFN{s1, s2})

		case Concat:
			// Concatenación de dos NFAs A·B:
			//   conecta A.Accept -ε-> B.Start, y A.Accept deja de ser acepting
			b, a := pop(), pop()
			a.Accept.Lambda = append(a.Accept.Lambda, b.Start)
			a.Accept.Accepting = false
			st = append(st, &AFN{a.Start, b.Accept})

		case Union:
			// Unión de dos NFAs A + B:
			//   crea s, f; s -ε-> A.Start y s -ε-> B.Start;
			//   A.Accept -ε-> f y B.Accept -ε-> f
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
			// Cerradura de Kleene A*:
			//   crea s, f; s -ε-> A.Start y s -ε-> f;
			//   A.Accept -ε-> A.Start y A.Accept -ε-> f
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

	// Al final la pila debe tener un único AFN completo
	return st[0]
}
