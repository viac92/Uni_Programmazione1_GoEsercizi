// Esercizio: 2
// Dato n calcola il numero di numeri primi minori o uguali a n. Usarlo per verificare sperimentalmente che questo P(n) è asintoti.

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	var numCorrente int = n
	var trovaPrimi int = 2
	var contaPrimi int

	for cicliNum := 0; cicliNum < n; cicliNum++ {
		
		for trovaPrimi < numCorrente {
			if numCorrente % trovaPrimi == 0 {
				break 
			}
			trovaPrimi++
		}
		if  trovaPrimi == numCorrente {
			contaPrimi++
		}
		trovaPrimi = 2
		numCorrente--

	}

	fmt.Println("I numeri primi <=", n, "sono:", contaPrimi)
	var confronto float64 = float64(n)/math.Log(float64(n))
	fmt.Println("Funzione da confrontare: ", confronto)
	fmt.Println("Rapporto:", float64(contaPrimi)/confronto)

}
