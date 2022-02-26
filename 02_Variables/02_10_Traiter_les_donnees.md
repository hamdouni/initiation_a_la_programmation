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