# Initiation à la programmation Go - Partie 7 - Les paquets

Les paquets interviennent à 2 niveaux : afin d'utiliser des fonctionnalités écrites par d'autres, et pour organiser son propre code.

## Utiliser les packages : import

On peut aisément ajouter des fonctionnalités supplémentaires au langage Go grâce à l'ajout de packages. Un package est simplement un ensemble de fonctions, cohérentes entre elles, qui permettent d'étendre les capacités de son programme. 

Il en existe de très nombreux, dans des domaines très variés, comme le calcul mathématiques ou les manipulations de fichiers. Cf la liste sur le [site de Golang](https://pkg.go.dev/std) à l'adresse :

```
https://pkg.go.dev/std
```

Pour ajouter un package à notre programme, on utilise l'instruction _import_ suivie du nom du package souhaité entre guillemets :

```
import "strings"
```

Si on a besoin de plusieurs packages, plutôt que d'écrire :

```
import "strings"
import "os"
```

On peut utiliser la forme suivante :

```
import (
    "strings"
    "os"
)
```

Pour voir ce que propose le package _strings_, allons sur <https://pkg.go.dev/strings>. 

On y trouve, par exemple, la fonction _ToUpper_ qui permet de mettre  une chaîne de caractère en lettres majuscules.

```
println(strings.ToUpper("je me vois en grand"))
```

affichera à l'écran

```
JE ME VOIS EN GRAND
```

Pour utiliser cette fonction, on a précisé le package devant : _strings.ToUpper_. 

Ainsi, on évite les "collisions" de nom de fonction comme dans l'exemple ci-dessous, où nous avons déclaré notre propre fonction _ToUpper_ :

```go
package main

import "strings"

func ToUpper() string {
	return "je suis ToUpper"
}
func main() {
	println(strings.ToUpper("je me vois en grand"))
	println(ToUpper())
}
```

affichera à l'écran :

```
JE ME VOIS EN GRAND
je suis ToUpper
```

## Organiser son code : définir des unités fonctionnelles

Les paquets servent également à organiser son code en regroupant les fonctionnalités cohérentes entre elles. 

Dans ce cas, chaque paquet sera logé dans un dossier pour séparer physiquement le code.

Reprenons l'exemple de notre cinéma de quartier et de son outil de vente de tickets : 

```
[+] impression/
[+] tarifs/
    main.go
```

Le paquet _impression_ regrouperait les fonctionnalités liées à l'impression des tickets.

Le paquet _tarifs_ regrouperait les fonctionnalités liées au calcul des tarifs.

Le programme principal `main.go` pourrait utiliser les fonctionnalités de ces paquets avec la syntaxe suivante :

```go
tarif := tarifs.Calcul(age)
impression.Ticket(tarif)
```

Dans le dossier `tarifs` on aurait un fichier `calcul.go` avec :

```go
package tarifs

func Calcul(age int) (tarif_cts int) {
```

## Visibilité: règle de la majuscule

Go impose une convention de nommage des fonctions pour déterminer leur visibilité. 

Une fonction dont le nom commence par une majuscule, comme dans l'exemple ci-dessus, rend la fonction utilisable par les autres paquets. Ainsi, il est possible d'appeler `tarifs.Calcul` depuis notre programme principal.

A l'inverse, une fonction, dont le nom commence par une minuscule, ne sera visible que dans le paquet où elle a été déclarée.

Si notre paquet `tarifs` contenait la fonction suivante :

```go 
func ristourne(prix_cts int) int {
```

Depuis notre programme principal `main.go`, nous n'aurions pas pu écrire :

```go
prix := tarifs.ristourne(1200)
```

La fonction `ristourne` ne peut être utilisée qu'à l'intérieur du paquet `tarifs`.