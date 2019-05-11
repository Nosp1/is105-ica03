package frequence

import (
	"testing"
)

/*
De tre funksjonene under kjører benchmark på 3 forskjellige i filer med forskjellige størrelse.

 */
//Benchmark test på pg100.txt med Bfrequence.
func BenchmarkBfrequencepg100(b *testing.B) {
	benchpg100(1, b)
}

//Benchmark test på bible2.txt (inneholdet i bible.txt x2)
func BenchmarkBfrequenceBible2(b *testing.B) {
	benchBible2(1, b)
}

//Benchmark test på bible.txt (hele bibelen)
func BenchmarkBfrequence1Bible(b *testing.B) {
	benchBible(1, b)
}

//Benchmark test på pg200.txt
func BenchmarkBfrequencepg200(b *testing.B) {
	benchpg200(1, b)
}

func BenchmarkBfrequenceFile(b *testing.B) {
	benchFile(1, b)
}

// funksjonene under er hjelpe funksjoner for benchmarking. Blant annet, valg av fil og hvor mange ganger benchmark skal kjøres
func benchpg100(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		BfrequenceForBench("../files/pg100.txt")

	}

}
func benchpg200(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		BfrequenceForBench("../files/pg200.txt")
	}
}
func benchBible2(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		BfrequenceForBench("../files/bible2.txt")
	}
}
func benchBible(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		BfrequenceForBench("../files/bible.txt")
	}
}

func benchFile(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		BfrequenceForBench("../files/file.txt")
	}
}
