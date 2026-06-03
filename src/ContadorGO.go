package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func contar(id int, inicio, fin uint64, ch chan<- string) {
	tiempoInicio := time.Now()

	var buf strings.Builder
	for i := inicio; i <= fin; i++ {
		buf.WriteString("Goroutine-")
		buf.WriteString(strconv.Itoa(id))
		buf.WriteString(" -> ")
		buf.WriteString(strconv.FormatUint(i, 10))
		buf.WriteByte('\n')
	}

	fmt.Fprintf(os.Stderr, "Goroutine-%d terminó en %.3f seg (rango: %d - %d)\n",
		id, time.Since(tiempoInicio).Seconds(), inicio, fin)

	ch <- buf.String()
}

func main() {
	var numHilos int
	fmt.Print("Cantidad de goroutines: ")
	fmt.Scan(&numHilos)

	var limite uint64
	fmt.Print("Límite: ")
	fmt.Scan(&limite)

	bloque := limite / uint64(numHilos)
	ch := make(chan string, numHilos)

	inicioPrograma := time.Now()

	for i := 0; i < numHilos; i++ {
		inicio := uint64(i)*bloque + 1
		fin := uint64(i+1) * bloque
		if i == numHilos-1 {
			fin = limite
		}
		go contar(i+1, inicio, fin, ch)
	}

	out := bufio.NewWriter(os.Stdout)
	for i := 0; i < numHilos; i++ {
		out.WriteString(<-ch)
	}
	out.Flush()

	fmt.Printf("\nTiempo total: %.3f seg\n", time.Since(inicioPrograma).Seconds())
}
