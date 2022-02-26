/*
Plus Grand Commun Diviseur

Rappel : le PGCD de deux nombres entiers n et m est le plus grand nombre possible, qui divise à la fois n et m.

Euclid a proposé, vers 300 avant JC, un algorithme pour résoudre ce problème comme suit :

E1. r le reste de la division entière de m par n
E2. si r = 0, le résultat est n et l'algorithme se termine
E3. m <- n, n <- r, aller E1

*/
package main

func main() {
	m := 119
	n := 544

	for {
		r := m % n
		if r == 0 {
			break
		}
		m, n = n, r
	}
	println(n)
}
