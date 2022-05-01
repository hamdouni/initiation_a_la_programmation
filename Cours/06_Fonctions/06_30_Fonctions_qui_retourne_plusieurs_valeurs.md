## Fonctions qui retourne plusieurs valeurs

On peut également retourner plusieurs valeurs en même temps. On indique la liste des valeurs à renvoyer dans la dernière partie de la déclaration de la fonction :

```
func multiretours() (ret1, ret2, ret3 int, ret4 string)
```

Ici, on retourne 3 entiers et une chaîne de caractères.

Le traitement devra comporter un _return_ en conséquence. Par exemple :

```
return 12, 14, 3, "hello"
```

Et la fonction appelante devra également en tenir compte :

```
a,b,c,d := multiretours()
```

avec succéssivement a, b et c entiers, et d une chaine.

Notons également qu'il est possible d'ommetre le nom des variables de retours. Go ne s'intéresse qu'à leur type. Ainsi on pourrait écrire :

```
func multiretours() (int, int, int, string)
```

Voilà une illustration complète : 

```go
package main

func multiretours() (int, int, int, string) {
	return 1, 2, 3, "hello"
}
func main() {
	a,b,c,d := multiretours()
	
	println(a,b,c,d)
}
```