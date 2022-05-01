# Les variables

Dans cette partie un peu plus théorique, nous allons voir de plus près la
notion de variable et de type. Les variables nous servent tout le temps en
programmation, et pour commencer, elles servent à stocker de l'information que
l'on pourra utiliser ultérieurement.

## Stocker les données

Une variable est un nom que l'on donne à une case mémoire pour y ranger des
données.

Par exemple :

```go
var m int
```

donnera le nom _m_ a une case mémoire dans laquelle on va ranger un nombre
entier: **int** est le diminutif du mot anglais **integer** qui signifie
entier.

Donner un nom à une case mémoire s'appelle déclarer la variable : nous
déclarons que la variable **m** est un nombre entier.

**Remarque:** par défaut, _m_ prendra la valeur 0.

Le choix du nom d'une variable est libre mais il faut respecter quelques règles
:

* un nom de variable doit commencer par une lettre ou le signe *_*

* seuls les caractères alpha-numériques (a-z, A-Z, 0-9) et *_* sont autorisés :
  pas d'espace ou de signes de ponctuations ou de caractères spéciaux

* la casse est signifiante : age, Age et AGE sont 3 variables différentes

* vous ne pouvez pas utiliser l'un des mots clés du langage Go, comme *print*
  ou *func*.

Il est possible de spécifier une valeur lors de la déclaration. On appelle cela
une _initialisation_.

Par exemple, si l'on veut _initialiser_ la variable _m_ avec la valeur 123,
nous écririons :

```go
var m int = 123
```

Ainsi, nous déclarons que la variable m est un entier que nous initialisons à
la valeur 123.

Lorsqu'on initialise une variable, on peut omettre de préciser le type car il
peut être induit par la valeur que l'on affecte. L'exemple ci-dessus peut
s'écrire plus simplement :

```go
var m = 123
```

Comme 123 est un entier, la variable _m_ sera par défaut de type _int_.

Nous allons voir maintenant comment utiliser ces variables : [Traiter les
données](./02_10_Traiter_les_donnees.md).
