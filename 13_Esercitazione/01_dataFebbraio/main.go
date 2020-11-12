package main

import (
		"fmt"
		"strconv"
)

type data struct {
	g, m, a int
}

func isBisestile(anno int) bool {
	return anno % 4 == 0 && (anno & 100 != 0 || anno % 400 == 0)
}

func ultimoGiornoDiFebbraio(anno int) data {
	if isBisestile(anno) {
		return data{29, 2, anno}
	}
	return data{28, 2, anno}
}

func stringa(d data) string {
	return strconv.Itoa(d.g) + "/" + strconv.Itoa(d.m) + "/" + strconv.Itoa(d.a)
}

func main()  {
	var n int
	fmt.Print("Inserisci un anno: ")
	fmt.Scan(&n)
	
	s := stringa(ultimoGiornoDiFebbraio(n))
	fmt.Println(s)


}

// Scrivere una funzione che prenda tre 
// - argomenti un puntatore a una data
// - una stringa s
// - un intero x
//
// la funzione deve modificare uno dei tre campi della data (il giorno se s = g, il mese se s = "m") ponendolo uguale a x.