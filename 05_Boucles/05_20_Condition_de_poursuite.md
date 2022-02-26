## Condition de poursuite

```go
... i <= 10; ...
```

La condition de poursuite indique si l'on doit ou non arrêter la boucle. Tant que la condition est remplie, on continue. 
Dans notre exemple, tant que la variable i est inférieure ou égale à 10, on répète notre boucle.

Dès que ce n'est plus le cas, on sort de la boucle, c'est-à-dire que le programme se poursuit après l'accolade fermante, qui symbolise la fin de notre boucle.

Attention à toujours vous assurer que la condition sera remplie un jour, au risque d'avoir une boucle qui ne s'arrête jamais (boucle infinie).

L'exemple suivant montre un tel cas :

```go
for a := 1; a >= 1; a++ {
    println("hey!")
}
```

La condition a >= 1 sera toujours vraie et donc la boucle ne se terminera pas !
