---
id: b26f3d70-ef48-440f-b603-32acb9a68bd5
title: 04_stocker Et Manipuler Les Donnees   Variables Et Types
desc: ''
updated: 1610816806916
created: 1610816806916
---
# Stocker et manipuler les données

Pour garder une information en mémoire, nous avons besoin de variables.

Une variable est un nom que l'on donne à une case mémoire pour y ranger des données.

Par exemple :

```
var m = 123
```

donnera le nom _m_ a une case mémoire dans laquelle on va ranger le nombre 123.

## Traitement de données

Une fois que l'on sait stocker une donnée, on va pouvoir la manipuler. 

Par exemple, pour l'afficher à l'écran, on écrira le code suivant :

```
var m = 22
print(m)
```

L'instruction _print_ va d'abord chercher en mémoire ce que contient la variable _m_ (ici 22) et affichera la valeur trouvée.

On peut effectuer des calculs avec les variables comme suit  :

```
var m = 741
print(m + 147)
```

L'instruction _print(m + 147)_ commence par récupérer la valeur stockée dans la case _m_ (ici 741) puis effectue l'addition :

```
741 + 147 = 888
```

Et enfin, l'affiche à l'écran.

On peut ranger le résultat du calcul dans une autre case mémoire.

Ci-dessous, nous appelons _p_ une nouvelle case mémoire dans laquelle nous rangeons directement le résultat de notre calcul :

```
var m = 64
var p = 4 * m
print(p)
```

donnera 256, le résultat de 4 x 64.

## Type entier

Dans Go, les variables que l'on manipule doivent impérativement avoir un type qui détermine ce que l'on peut en faire. 

Par exemple, un nombre entier, qui sera de type _int_ (de l'anglais "integer" qui veut dire entier), supportera les opérations arithmétiques comme l'addition ou la multiplication.

```
var x int
var y int
x = 5
y = 14
print(2*x+y)
```

affichera 24.

Les nombres négatifs sont acceptés :

```
var n = -234
print(n)
```

Chaque type détermine la taille maximum de la variable. 

Un _int_ peut aller de -2147483648 à 2147483647 (2 puissance 31 - 1 en architecture 32 bits).

Si l'on veut manipuler des nombres entiers encore plus grand, il faudra utiliser le type _int64_, entier sur 64 bits avec une valeur max de 9223372036854775807 (2 puissance 63 - 1).

## Raccourcis pratiques

En Go, on peut indiquer le type d'une variable en même temps que sa valeur d'initialisation. C'est la notation _:=_ comme suit :

```
d := 10
```

Go déduit automatiquement le type _int_ (entier) en se basant sur la valeur 10.

On peut regrouper les déclarations d'un même type. 

```
var x int
var y int
```

est équivalent à

```
var x,y int 
```

Et Go permet également la multi-déclaration comme suit :

```
n,m,o,p := 1,2,3,4
```

qui aura pour effet d'affecter les valeurs 1, 2, 3 et 4, respectivement à n, m, o et p, et qui est beaucoup plus court à écrire que :

```
n := 1
m := 2
o := 3
p := 4
```

Une autre notation assez pratique est également possible sur les calculs qui auto-référence la variable (ie qui utilise la variable à la fois dans le calcul et pour stocker le résultat) :

```
x = x + 1
y = y - 1
```

peut s'écrire plus simplement 

```
x ++
y --
```

On trouvera aussi à la place de :

```
x = x * 27
y = y / 100
```

la forme 

```
x *= 27
y /= 100
```

## Type décimal

Si l'on veut stocker un décimal, on doit faire appel au type _float32_ ou _float64_ selon la précision que l'on souhaite. 

Flottant en français vient du nom de la méthode utilisée en informatique pour stocker les nombres à virgule.

Voici comment l'on peut déclarer un flottant _f_ dans Go :

```
var f float32
```

Par défaut, Go lui assignera la valeur zéro :

```
var f float32
print(f)
```

affichera 0.

On indique toujours la valeur par défaut sous forme d'un décimal (à l'anglaise, avec le point comme séparateur de virgule), surtout lorsqu'on veut affecter une première valeur entière. 

Dans l'exemple suivant qui plante volontairement, la variable _f_ sera bien de type _float32_ car la valeur 3.14 sera interprétée par Go comme un décimal. Par contre, la variable _j_ sera de type _int_ car 2 est un entier pour Go.

```
f := 3.14
j := 2
print(f*j)
```

génèrera l'erreur "mismatched types float64 and int", ce que l'on traduira par "mélange entre type flottant et entier". 

Pour y remédier, on met 2 sous sa forme décimal comme suit :

```
f := 3.14
j := 2.0
print(f*j)
```

affiche le résultat correct 6,28 car Go interprétera _j_ comme flottant.

## Type chaine de caractères

Pour représenter du texte, on utilisera le type _string_ (chaine en français).

```
a := "De la terre à la lune"
println(a)
```

L'addition de 2 chaines de caractères revient à les concaténer (il n'y a pas d'équivalent de la soustraction ou de la multiplication pour les chaines). Ainsi, si l'on écrit :

```
a := "De la terre "
b := "à la lune"
c := a + b
println(c)
```

on obtiendra l'affichage de "De la terre à la lune" à l'écran.

