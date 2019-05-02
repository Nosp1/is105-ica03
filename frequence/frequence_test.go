package frequence

import (
	"testing"
)

/*
Benchmark av FrequenceForBench metoden
Tre benchmark funksjoner for å teste tre forskjellige filer av varierende størrelse.
*/
//Tester tekst fil med alle Shakespears verker som innhold
func BenchmarkFrequencepg100(b *testing.B) {
	benchFrequencePG100(1,b)

}
//Tester tekstfil med bibelen som innhold
func BenchmarkFrequenceBible(b *testing.B) {
	benchFrequenceBible(1,b)
}
//tester tekstfil med 2x bibelen som innhold.
func BenchmarkFrequenceBible2(b *testing.B) {
	benchmarkFrequenceBible2(1,b)
}
/*
Funksjonene under setter absolutt path til filen for benchtesting av hver fil.
Funksjonene blir kalt inn i testene over.
 */
func benchFrequencePG100(i int, b *testing.B) {
	for n := 0; n < b.N; n++{
		b.StopTimer()
		b.StartTimer()
		FrequenceForBench("../files/pg100.txt")
	}
}
func benchFrequenceBible(i int, b*testing.B){
	for n := 0; n < b.N; n++{
		b.StopTimer()
		b.StartTimer()
		FrequenceForBench("../files/bible.txt")
	}
}

func benchmarkFrequenceBible2(i int,b *testing.B) {
	for n := 0; n < b.N; n++{
		b.StopTimer()
		b.StartTimer()
		FrequenceForBench("../files/bible2.txt")
	}

}