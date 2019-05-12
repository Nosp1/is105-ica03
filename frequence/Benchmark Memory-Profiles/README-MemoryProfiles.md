I denne mappen finner du grafiske fremstillinger av Bfrequence.go og
 Frequence.go.


BfreqFileMedSprintF.PDF viser minne allokering under kjøring av bfrequence.go
på filen File.txt. Merk i denne grafiske rapporten har ikke koden blitt endret


BfreFIleUtenSprintF.pdf viser minne allokering uder kjøring av bfrequence.go
på filen File.txt med kode endring i bfrequence. Merk at minne alokeringen
er drastisk redusert.

FreqFileresults Memory Map.pdf viser kjøring av frequence.go på filen File.txt
Vi har kun lagt med en kjøring av frequence da resultatene var tilnærmet like
hver gang.

De andre filene i mappen er binære memory profiler som go generer under
benchmark med -Memprofile. For å lese disse filene brukes Pprof, et ekstra
verktøy i GO. Skriv Pprof <fil> for å laste den inn pprof. Skriv deretter
pdf for å lage pdf av den binære filen med grafisk representasjon. 
