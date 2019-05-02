package main

import (
	//Husk å endre path til egen under kjøring
	"github.com/Nosp1/Is-105/is105-ica03/frequence"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		feilmelding := "Du mangler -f som flag"
		panic(feilmelding)
	}
	if len(args) == 2 {
		feilmelding := "Du mangler filpath: prøv ./files/<filnavn>"
		panic(feilmelding)
	}

	if len(args) == 3 {
		if args[1] == "-f" {
			frequence.Frequence(args[2])
		}
	}
}
