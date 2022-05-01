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

Avant d'aller plus loin, apprenons quelques raccourcis qui vont nous être utiles par la suite : [Raccourcis pratiques](./02_30_Raccourcis_pratiques.md).