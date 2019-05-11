package frequence

import (
	"fmt"
	"github.com/Nosp1/Is-105/is105-ica03/fileversion"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)
var filebaseFreq string = "frequence_res"
var dirFreq = "./frequence/frequenceresults/"
var dirEntriesFreq []os.FileInfo
var pathfreq string




type kv struct {
	Key   string
	Value int
}
/*
Oppgave 3c
 */
func Frequence(fileName string) {
	dirEntriesFreq, _ = ioutil.ReadDir(dirFreq)
	pathfreq = dirFreq + filebaseFreq + strconv.Itoa(len(dirEntriesFreq)) + ".txt"

	file, err := ioutil.ReadFile(fileName)
	if err !=nil{
		fmt.Println(err)
	}
	lines := strings.Split(string(file), "\n")
	//kutter ned filen
	lines = lines [:len(lines)-1]
	//printer ut alle linjer med index. kommentert ut for renere output.
	//for i, line := range lines{
		//fmt.Println(i,line)
	//}
	//printer ut total antall linjer i filen
	fmt.Println(fileName, "inneholder", len(lines), "linjer ")
	//oppretter map
	m := make(map[string]int)
	fmt.Println("dette er de fem mest brukte runene:")
	//iterer over filen og teller hver rune.
	for i := 0; i < len(file); i++ {
		m[string(file[i])] += 1
	}


	//Setter variabelen ss til å være en slice av typen k og v
	var ss []kv
	//iterer over map og legger runene inn i ss med k som string verdi for rune og v for int verdi for rune.
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	//Sorterer slicen ss etter mest brukte
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	//Oppretter ny variabel "ts" som returnerer de
	//fem mest brukte "runes" i ss
	ts := ss[:5]
	//Looper over key-values i ts-slicen.
	for _, kv := range ts {
		//Printer ut de fem mest brukte symbolene og antall ganger brukt.
		fmt.Printf("%q, %d\n", kv.Key, kv.Value)
	}
_, err1 := os.Stat(pathfreq)

if err1 == nil {
	WriteToFileFreq(fileversion.DontOverrideFileversion(pathfreq), lines,ts)
}
if err1 != nil{
	WriteToFileFreq(pathfreq,lines,ts)
}
}


/*
Hjelpe funksjon som skriver antall linjer og de fem mest brukte runene til fil.
 */
func WriteToFileFreq(filepath string, lines []string,   ts []kv ) {
	write, err := os.Create(filepath)
	if err !=nil{
		panic(err)
	}

	defer write.Close()
	fmt.Println("Writing results to file")
	fmt.Fprintln(write,"Frequence resultat: ")
	fmt.Fprintf(write, "\nAntall linjer: %v", len(lines))
	fmt.Fprintln(write, "De mest brukte runene er: ")
	for i := 0; i< 5; i ++{
		fmt.Fprintf(write, "%q, %d\n", ts[i].Key, ts[i].Value)
	}
	fmt.Println("Writing to file complete.")


}
/*
Hjelpe funksjon som isolerer linjetelling og runetelling uten å skrive til fil.
For å lettere kunne kjøre benchmarks.
 */
func FrequenceForBench(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(file), "\n")
	//kutter ned filen
	lines = lines [:len(lines)-1]
	//printer ut alle linjer med index. Kommentert ut for renere output.
	//for i, line := range lines {
	//	fmt.Println(i, line)
	//}
	//printer ut total antall linjer i filen
	fmt.Println(fileName, "inneholder", len(lines), "linjer ")
	//oppretter map
	m := make(map[string]int)
	fmt.Println("dette er de fem mest brukte runene:")
	//iterer over filen og teller hver rune.
	for i := 0; i < len(file); i++ {
		m[string(file[i])] += 1
	}

	//Setter variabelen ss til å være en slice av typen k og v
	var ss []kv
	//iterer over map og legger runene inn i ss med k som string verdi for rune og v for int verdi for rune.
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	//Sorterer slicen ss etter mest brukte
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	//Oppretter ny variabel "ts" som returnerer de
	//fem mest brukte "runes" i ss
	ts := ss[:5]
	//Looper over key-values i ts-slicen.
	for _, kv := range ts {
		//Printer ut de fem mest brukte symbolene og antall ganger brukt.
		fmt.Printf("%q, %d\n", kv.Key, kv.Value)
	}
}