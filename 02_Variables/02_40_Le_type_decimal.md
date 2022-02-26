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