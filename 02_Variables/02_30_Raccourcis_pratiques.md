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