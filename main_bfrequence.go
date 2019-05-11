package main

import (
	"github.com/Nosp1/Is-105/is105-ica03/frequence"
	"os"
)

func main() {
	args := os.Args
	// om kommandolinjearugment er for kort kommer feilmelding.
	if len(args) == 1{
		feilmelding := "Du mangler -f som flagg. Prøv dette: Go run main_bfrequence.go -f <filepath> "
		panic(feilmelding)
	}
	if len(args) == 2 {
		feilmelding := "Du mangler en gyldig filepath. For eksempel: ./files/pg100.txt"
		panic(feilmelding)
	}
	// om lengden på argumentene i kommandolinje er  3 og posisjon 1 har "-f" så kjører Befrequence.
	if len(args) == 3 {
		if args[1] == "-f" {
			frequence.Bfrequence(args[2])
		}
	}
}
