---
id: 32e7a2d7-5244-492d-a791-b086e7e74b01
title: 12_utiliser Les Packages   Import
desc: ''
updated: 1610816806919
created: 1610816806919
---
# Utiliser les packages : import

On peut aisément ajouter des fonctionnalités supplémentaires au langage Go grâce à l'ajout de packages. Un package est simplement un ensemble de fonctions, cohérentes entre elles, qui permettent d'étendre les capacités de son programme. 

Il en existe de très nombreux, dans des domaines très variés, comme le calcul mathématiques ou les manipulations de fichiers. Cf la liste sur le [site de Golang](https://golang.org/pkg/) à l'adresse :

```
https://golang.org/pkg/
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

Pour voir ce que propose le package _strings_, allons sur <https://golang.org/pkg/strings/>. 

On y trouve, par exemple, la fonction _ToUpper_ qui permet de mettre en capital une chaîne de caractère.

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

