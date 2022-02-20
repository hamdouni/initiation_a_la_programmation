# Initiation à la programmation Go - Partie 3 - Les variables

Aujourd'hui, dans cette partie un peu plus théorique, nous allons voir de plus près la notion de variables et de type. 

## Stocker les données

Pour garder une information en mémoire, nous avons besoin de variables.

Une variable est un nom que l'on donne à une case mémoire pour y ranger des données.

Par exemple :

```go
    var m int
```

donnera le nom _m_ a une case mémoire dans laquelle on va ranger un nombre entier: **int** est le diminutif du mot anglais **integer** qui signifie entier.

Donner un nom à une case mémoire s'appelle déclarer la variable : nous déclarons que la variable **m** est un nombre entier.

**Remarque:** par défaut, _m_ prendra la valeur 0.

Le choix du nom d'une variable est libre mais il faut respecter quelques règles :

* un nom de variable doit commencer par une lettre ou le signe *_*
* seuls les caractères alpha-numériques (a-z, A-Z, 0-9) et *_* sont autorisés : pas d'espace ou de signes de ponctuations ou de caractères spéciaux
* la casse est signifiante : age, Age et AGE sont 3 variables différentes
* vous ne pouvez pas utiliser l'un des mots clés du langage Go, comme *print* ou *func*.

Il est possible de spécifier une valeur lors de la déclaration. On appelle cela une _initialisation_.

Par exemple, si l'on veut _initialiser_ la variable _m_ avec la valeur 123, nous écririons :

```go
    var m int = 123
```

Ainsi, nous déclarons que la variable m est un entier que nous initialisons à la valeur 123.

Lorsqu'on initialise une variable, on peut omettre de préciser le type car il peut être induit par la valeur que l'on affecte. L'exemple ci-dessus peut s'écrire plus simplement :

```go
    var m = 123
```

Comme 123 est un entier, la variable _m_ sera par défaut de type _int_.

## Traiter les données

Une fois que l'on sait stocker une donnée, on va pouvoir la manipuler. 

Par exemple, pour l'afficher à l'écran, on écrira le code suivant :

```go
    var m = 22
    print(m)
```

L'instruction _print_ va d'abord chercher en mémoire ce que contient la variable _m_ (ici 22) et affichera la valeur trouvée.

On peut effectuer des calculs avec les variables comme suit  :

```go
    var m = 741
    print(m + 147)
```

L'instruction _print(m + 147)_ commence par récupérer la valeur stockée dans la case _m_ (ici 741) puis effectue l'addition :

    741 + 147 = 888

Et enfin, l'instruction affiche le résultat à l'écran.

On peut ranger le résultat du calcul dans une autre case mémoire.

Ci-dessous, nous appelons _p_ une nouvelle case mémoire dans laquelle nous rangeons directement le résultat de notre calcul :

```go
    var m = 64
    var p = 4 * m
    print(p)
```

donnera 256, le résultat de 4 x 64.

## Le type entier

Dans Go, les variables que l'on manipule doivent impérativement avoir un type qui détermine ce que l'on peut en faire. Par exemple, un nombre entier, qui sera de type _int_ (de l'anglais "integer" qui veut dire entier), supportera les opérations arithmétiques comme l'addition ou la multiplication.

```go
    var x int
    var y int
    x = 5
    y = 14
    print(2*x+y)
```

affichera 24.

Les nombres négatifs sont acceptés :

```go
    var n = -234
    print(n)
```

Chaque type détermine la taille maximum de la variable. 

Un _int_ peut aller de -2147483648 à 2147483647 (2 puissance 31 - 1 en architecture 32 bits).

Si l'on veut manipuler des nombres entiers encore plus grand, il faudra utiliser le type _int64_, entier sur 64 bits avec une valeur max de 9223372036854775807 (2 puissance 63 - 1).

## Raccourcis pratiques

En Go, on peut indiquer le type d'une variable en même temps que sa valeur d'initialisation. C'est la notation _:=_ comme suit :

```go
    d := 10
```

Go déduit automatiquement le type _int_ (entier) en se basant sur la valeur 10.

On peut regrouper les déclarations d'un même type. 

```go
    var x int
    var y int
```

est équivalent à

```go
    var x,y int 
```

Et Go permet également la multi-déclaration comme suit :

```go
    n,m,o,p := 1,2,3,4
```

qui aura pour effet d'affecter les valeurs 1, 2, 3 et 4, respectivement à n, m, o et p, et qui est beaucoup plus court à écrire que :

```go
    n := 1
    m := 2
    o := 3
    p := 4
```

Une autre notation assez pratique est également possible sur les calculs qui auto-référence la variable (ie qui utilise la variable à la fois dans le calcul et pour stocker le résultat), comme dans les 2 exemples ci-dessous :

```go
    x = x + 1
    y = y - 1
```

peuvent s'écrire plus simplement 

```go
    x ++
    y --
```

On trouvera aussi l'équivalent pour la multiplication :

    x = x * 27     peut s'écrire     x *= 27

et pour la division : 

    y = y / 100   peut s'écrire    y /= 100

## Le type décimal

Si l'on veut stocker un nombre décimal, on doit faire appel au type _float32_ ou _float64_ selon la précision que l'on souhaite. 

**Remarque** : _float_ ou flottant en français vient du nom de la méthode utilisée en informatique pour stocker les nombres à virgule, et c'est toute une histoire :-).

Voici comment l'on peut déclarer un flottant _f_ dans Go :

```go
    var f float32
```

Par défaut, Go lui assignera la valeur zéro

```go
    var f float32
    print(f)
```

affichera +0.000000e+000 soit l'équivalent de zéro.

Pour enlever toute ambiguité, nous devons écrire :

```go
var f = 1.0
```

pour que Go comprenne que notre variable f est de type flotant. Si nous avions écris

```go
var f = 1
```

il aurait deviner que f était un int (entier).

Dans l'exemple suivant qui plante volontairement, la variable _f_ sera bien de type _float32_ car la valeur 3.14 sera interprétée par Go comme un décimal. Par contre, la variable _j_ sera de type _int_ car 2 est un entier pour Go.

```go
    f := 3.14
    j := 2
    print(f*j)
```

génèrera l'erreur "mismatched types float64 and int", ce que l'on traduira par "mélange entre type flottant et entier". Go n'accepte pas d'opération sur des variables de type différent.

Pour y remédier, on met 2 sous sa forme décimal comme suit :

```go
    f := 3.14
    j := 2.0
    print(f+j)
```

Cette fois, Go interprétera _j_ comme flottant et 
affichera le résultat correct 6,28.

## Le type chaine de caractères

Pour représenter du texte, on utilisera le type _string_ (chaine en français).

```go
    a := "De la terre à la lune"
    println(a)
```

L'addition de 2 chaines de caractères revient à les concaténer. Ainsi, si l'on écrit :

```go
    a := "De la terre "
    b := "à la lune"
    c := a + b
    println(c)
```

on obtiendra l'affichage de "De la terre à la lune" à l'écran.

 Remarquez que les autres opérations arithmétiques ne s'appliquent pas pour les chaines de caractères (soustraction, multiplication ou division).

## Conclusion

Nous avons vu, dans le détail, les variables et les types ; 2 notions que nous allons utiliser fréquemment. 

La plupart des programmes informatiques vont mobiliser soit des nombres entiers, soit des nombres décimaux, soit des chaînes de caractères, mais la plupart du temps, un mix de tout cela.

Nous verrons qu'il existe d'autres types, plus complexes : par exemple, des regroupements de plusieurs valeurs de même type (**slice**) ou de types différents (**struct**). 

## Exercices

En attendant, pour vous exercer, vous devrez écrire un programme qui calcule et affiche l'âge d'une personne dont on connait l'année de naissance. On supposera qu'on ne s'occupe pas du jour et du mois, et qu'on connait l'année en cours. 

Et pour les plus téméraires, écrire un programme qui affiche l'âge des personnes nées entre 1972 et 1992, sous la forme :

```
    Les personnes nées en 1972 ont 50 ans.
    Les personnes nées en 1973 ont 51 ans.
    ...
```

## Prochaine partie

Pour la prochaine fois, nous utiliserons ce que nous avons appris des variables pour mettre en pratique quelques cas dans la [Pratique des variables](../partie_04/README.md)