package frequence

import (
	"bufio"
	"fmt"
	"github.com/Nosp1/Is-105/is105-ica03/fileversion"
	"io/ioutil"
	"os"
	"sort"
	"strconv"

	//"path/filepath"
)
//prøver å endre variabler til stor bokstav. Husk å endre tilbake
var path string
var dirEntries []os.FileInfo
var fileBase string = "​bfrequence_res"
var dir = "./frequence/bfrequenceresults/"


/*
 Buffered filereader.
Tar in filenavn, teller linjer med hjelpe funksjonen LinesinFilebuffered
Skriver så til fil med hjelpemetoden Writetofile


 */

func Bfrequence(fileName string) {
	dirEntries, _ = ioutil.ReadDir(dir)
	path = dir + fileBase + strconv.Itoa(len(dirEntries)) + ".txt"

	f, _ := os.Open(fileName)
	//for index, line := range LinesInFileBuffered(fileName) {

		//fmt.Sprintf("Index = %v, line = %v\n", index, line)
	//}
	defer f.Close()
	// Get count of lines.
	lines := LinesInFileBuffered(fileName)
	fmt.Println("Antall linjer: ", len(lines))
	//oppretter scanner
	input := bufio.NewScanner(f)
	//splitter scanna runer.
	input.Split(bufio.ScanRunes)
	//oppretter map av lest fil
	m := make(map[string]int)
	//scanner filen.
	for input.Scan() {
		ru := input.Text()
		//inkrementerer runer
		m[ru]++
	}
	//gjør variablel list til lflist og lengden på m.
	list := make(lflist, 0, len(m))
	// iterer over map
	for l, f := range m {
		//legger til freq og string til listen.
		list = append(list, &letterFreq{l, f})
	}
	// sorterer listen
	sort.Sort(list)
	fmt.Println("dette er de fem mest brukte runene: ")
	counter := 0
	for _, lf := range list {
		if counter >= 5 {
			break
		}
		fmt.Printf("%+q %7d\n", lf.string, lf.freq)
		counter++
	}
	// Filen finnes alt
	_, err := os.Stat(path)

	if err == nil {
		writeToFile(fileversion.DontOverrideFileversion(path), list, lines)

	}
	// Filen finnes ikke
	if err != nil {
		// Lager ny
		writeToFile(path, list, lines)

	}
}
/*
Funksjonen writeToFile håndterer skriving av lest data til fil.

 */
func writeToFile(filepath string, list lflist, lines []string) string {
	write, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer write.Close()

	w := bufio.NewWriter(write)
	fmt.Println("Writing results to file")
	fmt.Fprintln(w, "Bfrequence resultat: ")
	fmt.Fprint(w, "De mest brukte runene er: ", )
	for i := 0; i < 5; i ++ {
		_, err = fmt.Fprintf(w, "\n%+q %7d\n", list[i].string, list[i].freq)
	}
	_, err = fmt.Fprintf(w, "\nAntall linjer: %v", len(lines))
	w.Flush()
	fmt.Println("Writing to file is complete. ")

	return "sucess!"
}

type letterFreq struct {
	string
	freq int
}
type lflist []*letterFreq

func (list lflist) Len() int {
	return len(list)
}
func (list lflist) Less(i, j int) bool {
	switch t := list[i].freq - list[j].freq; {
	case t < 0:
		return false
	case t > 0:
		return true
	}
	return list[i].string < list[j].string
}
func (list lflist) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
/*
Hjelpe funksjon for å telelinjer i fil
 */

func LinesInFileBuffered(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result

}
/*
Funksjon som isolerer ut linjeteller og runeteller uten å skrive til fil
for å lettere kunne kjøre benchmarks.
 */
func BfrequenceForBench(fileName string) {
	f, _ := os.Open(fileName)
	//for index, line := range LinesInFileBuffered(fileName) {

	//	fmt.Sprintf("Index = %v, line = %v\n", index, line)
	//}
	defer f.Close()
	// Get count of lines.
	lines := LinesInFileBuffered(fileName)
	fmt.Println("Antall linjer: ", len(lines))
	//oppretter scanner
	input := bufio.NewScanner(f)
	//splitter scanna runer.
	input.Split(bufio.ScanRunes)
	//oppretter map av lest fil
	m := make(map[string]int)
	//scanner filen.
	for input.Scan() {
		ru := input.Text()
		//inkrementerer runer
		m[ru]++
	}
	//gjør variablel list til lflist og lengden på m.
	list := make(lflist, 0, len(m))
	// iterer over map
	for l, f := range m {
		//legger til freq og string til listen.
		list = append(list, &letterFreq{l, f})
	}
	// sorterer listen
	sort.Sort(list)
	fmt.Println("dette er de fem mest brukte runene: ")
	counter := 0
	for _, lf := range list {
		if counter >= 5 {
			break
		}
		fmt.Printf("%+q %7d\n", lf.string, lf.freq)
		counter++
	}

}


















