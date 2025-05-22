package internal

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
)

// ToDOT genera el grafo DOT pero usando nombres q0, q1, …
func (nfa *NFA) ToDOT() string {
	// helper que mapea un entero a "q<entero>"
	name := func(id int) string {
		return fmt.Sprintf("q%d", id)
	}

	g := gographviz.NewEscape()
	g.SetName("NFA")
	g.SetDir(true)
	g.AddAttr("NFA", "rankdir", "LR")

	seen := map[int]*State{}
	var visit func(s *State)
	visit = func(s *State) {
		if _, ok := seen[s.ID]; ok {
			return
		}
		seen[s.ID] = s

		// etiqueta del nodo: q<ID>
		attrs := map[string]string{"label": name(s.ID)}
		if s.Accepting {
			attrs["shape"] = "doublecircle"
		}
		// crea el nodo con el nombre "q<ID>"
		g.AddNode("NFA", name(s.ID), attrs)

		// transiciones con letra
		for ch, dsts := range s.Trans {
			for _, d := range dsts {
				g.AddEdge(name(s.ID), name(d.ID), true,
					map[string]string{"label": string(ch)})
				visit(d)
			}
		}
		// transiciones ε
		for _, d := range s.Epsilon {
			g.AddEdge(name(s.ID), name(d.ID), true,
				map[string]string{"label": "λ"})
			visit(d)
		}
	}
	// arranca el recorrido desde el estado inicial
	visit(nfa.Start)

	// nodo "start" apuntando a q<start.ID>
	g.AddNode("NFA", "start", map[string]string{"shape": "point"})
	g.AddEdge("start", name(nfa.Start.ID), true, nil)

	return g.String()
}

func (nfa *NFA) Renumber() {
	// BFS para recolectar todos los estados alcanzables
	seen := make(map[*State]bool)
	order := []*State{nfa.Start}
	seen[nfa.Start] = true

	for i := 0; i < len(order); i++ {
		s := order[i]
		// recorré transiciones ε
		for _, d := range s.Epsilon {
			if !seen[d] {
				seen[d] = true
				order = append(order, d)
			}
		}
		// recorré transiciones por símbolo
		for _, dsts := range s.Trans {
			for _, d := range dsts {
				if !seen[d] {
					seen[d] = true
					order = append(order, d)
				}
			}
		}
	}

	// reasigná IDs secuenciales según el orden BFS
	for idx, s := range order {
		s.ID = idx
	}
}
