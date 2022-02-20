# Initiation à la programmation Go - Partie 6 - Les fonctions

Les fonctions en Go permettent de regrouper des opérations pour pouvoir les éxécuter plusieurs fois. Dès que du code est répété plus d'une fois, il sera nécessaire de le regrouper sous une fonction.

## Déclarons une fonction

```
func bonjour() {
    println("hello")
}
```

Nous avons déclaré une nouvelle fonction _bonjour_ qui ne fait qu'afficher "hello" à l'écran.

On indique le nom de la fonction que l'on souhaite après l'instruction _func_.

Le jeu de paranthèse vide indique que la fonction n'attend pas de paramètres (voir plus bas pour les fonctions avec paramètres).

Puis vient entre accolades le traitement de la fonction : ici un simple affichage à l'écran _println_.

Pour exécuter cette fonction, on utilise son nom comme une instruction :

```
bonjour()
bonjour()
bonjour()
```

Ici, nous avons appelé 3 fois notre nouvelle fonction _bonjour_.
On obtiendra donc à l'écran :

```
hello
hello
hello
```

Voilà le code en entier :

```go
    package main

    func bonjour() {
        println("hello")
    }

    func main() {
        bonjour()
        bonjour()
        bonjour()
    }
```

## Fonctions avec paramètres

On peut changer le comportement d'une fonction en lui passant des valeurs qui seront utilisées dans son traitement.

Prenons l'exemple du calcul d'aire d'un rectangle : si on a la longueur et la largeur, l'aire se calcul par aire = longueur \* largeur.

Soit le code suivant pour calculer l'aire de 3 terrains, le premier de 10 par 5 mètres, le second de 12 par 9 mètres et le dernier de 13 par 11 mètres :

```
var long, larg, aire int

long = 10
larg = 5
aire = long * larg
println("l'aire ", long, " * ", larg, " = ", aire)

long = 12
larg = 9
aire = long * larg
println("l'aire ", long, " * ", larg, " = ", aire)

long = 13
larg = 11
aire = long * larg
println("l'aire ", long, " * ", larg, " = ", aire)
```

On voit clairement des répétitions. Voici ce que donnerait le même code avec une fonction :

```
func affiche_aire(long, larg int) {
    aire := long * larg
    println("l'aire ", long, " * ", larg, " = ", aire)
}

func main() {
    affiche_aire(10, 5)
    affiche_aire(12, 9)
    affiche_aire(13, 11)
}
```

La partie la plus intéressante est dans le premier bloc. D'abord on déclare notre nouvelle fonction :

```
func affiche_aire(long, larg int)
```

Le nom que l'on donne à notre fonction apparaît après l'instruction _func_, ici "affiche_aire".

Entre parenthèses, on indique les paramètres : des variables qu'attend la fonction pour effectuer son traitement, et qui lui sont indiqué lors de l'appel, dans le même ordre !

Ici, on passe 2 paramètres : "long" et "larg" de type _int_.

Le contenu de la fonction est exactement le même  calcul d'aire et d'affichage à l'écran, qu'on avait précédemment.

L'appel se fait dans la fonction principale _main_ en réutilisant le nom de la fonction et en indiquant les valeurs que l'on veut passer :

```
affiche_aire(10, 5)
```

va donc appeler notre nouvelle fonction en indiquant que la longueur vaut 10 et la largeur 5 (on respecte l'ordre)

L'avantage de cette approche est de pouvoir facilement ajouter d'autres terrains à calculer. Par exemple, voilà comment on pourrait ajouter 2 autres terrains avec seulement 2 lignes de codes supplémentaires (code complet) :

```go
    package main

    func affiche_aire(long, larg int) {
        aire := long * larg
        println("l'aire ", long, " * ", larg, " = ", aire)
    }

    func main() {
        affiche_aire(10, 5)
        affiche_aire(12, 9)
        affiche_aire(13, 11)
        affiche_aire(15, 7)
        affiche_aire(9, 10)
    }
```

## Fonctions qui retourne une valeur

Une fonction peut renvoyer une valeur à celui qui l'a invoqué.

```
func addition(a,b int) (somme int) {
    return a+b
}
```

Cette fonction attend 2 valeurs en paramètres et retourne leur somme. On déclare la valeur attendue en dernier : _somme int_.

C'est l'instruction _return_ qui indique ce que la fonction va renvoyer : ici _a+b_ qui correspond bien à la somme de nos 2 paramètres.

Pour utiliser cette fonction, on écrira :

```
    resultat := addition(5,7)
```

qui mettra le résultat de la somme de 5 par 7 dans la variables _resultat_.

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

### Ignorer des valeurs de retour

Enfin, il est possible de ne pas se préoccuper d'une ou plusieurs valeurs retournées, en indiquant l'underscore comme suit :

```
a,_,_,d := multiretours()
```

indique que l'on ignore le 2ème et le 3ème retours, et qu'on ne souhaite utiliser que le premier et le dernier.

