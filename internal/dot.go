package internal

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
)

// ToDOT genera el grafo DOT pero usando nombres q0, q1, …
func (afn *AFN) ToDOT() string {
	// helper que mapea un entero a "q<entero>"
	name := func(id int) string {
		return fmt.Sprintf("q%d", id)
	}

	g := gographviz.NewEscape()
	g.SetName("AFN")
	g.SetDir(true)
	g.AddAttr("AFN", "rankdir", "LR")

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
		g.AddNode("AFN", name(s.ID), attrs)

		// transiciones con letra
		for ch, dsts := range s.Trans {
			for _, d := range dsts {
				g.AddEdge(name(s.ID), name(d.ID), true,
					map[string]string{"label": string(ch)})
				visit(d)
			}
		}
		// transiciones ε
		for _, d := range s.Lambda {
			g.AddEdge(name(s.ID), name(d.ID), true,
				map[string]string{"label": "λ"})
			visit(d)
		}
	}
	// arranca el recorrido desde el estado inicial
	visit(afn.Start)

	// nodo "start" apuntando a q<start.ID>
	g.AddNode("AFN", "start", map[string]string{"shape": "point"})
	g.AddEdge("start", name(afn.Start.ID), true, nil)

	return g.String()
}

func (afn *AFN) Renumber() {
	// BFS para recolectar todos los estados alcanzables
	seen := make(map[*State]bool)
	order := []*State{afn.Start}
	seen[afn.Start] = true

	for i := 0; i < len(order); i++ {
		s := order[i]
		// recorré transiciones ε
		for _, d := range s.Lambda {
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
