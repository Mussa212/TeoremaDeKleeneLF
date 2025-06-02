package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"LenguajesFormales/internal"
)

func main() {
	re := flag.String("re", "", "Expresión regular")
	out := flag.String("out", "nfa.png", "Imagen de salida")
	flag.Parse()
	if *re == "" {
		log.Fatal("faltó -re")
	}

	toks := internal.Scan(*re)
	postfix := internal.ToPostfix(toks)
	nfa := internal.BuildAFN(postfix)
	nfa.Renumber()
	dot := nfa.ToDOT()
	// escribe dot a archivo temporal
	f, _ := os.CreateTemp("", "nfa-*.dot")
	defer os.Remove(f.Name())
	f.WriteString(dot)
	f.Close()

	// requiere graphviz instalado (dot)
	if err := exec.Command("dot", "-Tpng", f.Name(), "-o", *out).Run(); err != nil {
		log.Fatalf("graphviz: %v", err)
	}
	fmt.Println("Generado", *out)
}
