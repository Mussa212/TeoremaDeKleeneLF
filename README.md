# TeoremaDeKleeneLF

#### Instrucciones:

1. Descargar el repositorio: 

```bash
git clone https://github.com/Mussa212/TeoremaDeKleeneLF.git
```

2. SetUp del proyecto:
   - Debe tener instalado [GoLang](https://go.dev/dl/).
   - Debe tener instalado [GraphViz](https://graphviz.org/download/).

3. Ejecutar el programa:

```bash
go mod tidy 
go run main.go -re "<expresión>" [-out <archivo>.png]
```
Por ejemplo:

```bash

go run main.go -re "(a+b)*" -out automata.png

go run main.go -re "ab+ba" -out prueba.png
```
