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
