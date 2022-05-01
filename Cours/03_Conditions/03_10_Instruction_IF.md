## L'instruction If

Lorsque l'on souhaite effectuer une action mais seulement lorsqu'une condition est remplie, on utilise l'instruction _if_.

```
if x > 5 {
    print("plus grand")
}
print(" et la suite")
```

Ici la condition à évaluer est x > 5.

```
if x > 5 
```

Si x est plus grand que 5, l'action entre crochets sera exécutée :

```
print("plus grand")
```

Sinon, on passe à l'instruction après le crochet fermant. Dans notre cas :

```
print(" et la suite")
```

Voilà ce que cela donnerait avec x = 12

```
x := 12
if x > 5 {
    print("plus grand")
}
print(" et la suite")
```

On obtiendrait la phrase "plus grand et la suite" car la condition x > 5 est vraie, x valant 12.

Voyons ce que cela donnerait avec un x = 3

```
x := 3
if x > 5 {
    print("plus grand")
}
print(" et la suite")
```

Cette fois, il ne s'affichera que " et la suite" car la même condition est maintenant fausse.