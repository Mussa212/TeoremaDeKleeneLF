# TeoremaDeKleeneLF

#### Instrucciones:

1. Descargar el repositorio: 

```bash
git clone https://github.com/Mussa212/TeoremaDeKleeneLF.git

go mod tidy 
```

2. SetUp del proyecto:
   - Debe tener instalado GraphViz.
3. Ejecutar el programa:

```bash
go mod tidy 
go run main.go -re "<expresión>" [-out <archivo>.png]
```
Por ejemplo:

```bash

go run ./cmd/kleene -re "(a+b)*" -out automata.png

go run ./cmd/kleene -re "ab+ba" -out prueba.png
```