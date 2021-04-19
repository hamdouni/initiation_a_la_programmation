---
id: e0e18d7f-8805-4a30-a0c1-cfb2a55bae77
title: 06_géRer Des Conditions   If
desc: ''
updated: 1610816806917
created: 1610816806917
---
# Gérer des conditions

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

## Le type bool

Une condition doit renvoyer "true" ou "false", respectivement vrai ou faux. Go appelle ce type _bool_ du nom du mathématicien Bool, qui créa l'algèbre binaire qui porte son nom (cf Wikipedia)

On peut donc écrire :

```
c := true
if c {
    print("toujours")
}
```

qui affichera systématiquement le mot "toujours". Respectivement, le code suivant n'affichera jamais rien :

```
c := false
if c {
    print("jamais")
}
```

## L'instruction else

Go permet également d'indiquer une alternative dans le cas où la condition est fausse, avec l'instruction _else_ qu'on traduirait par "sinon". 

```
if x > 5 {
     print("plus")
}
else {
    print("moins")
}
print(" et la suite")
```

Le code entre crochets qui suit le _else_ ne sera exécuté que si la condition est fausse, c'est-à-dire, dans notre cas, seulement si x plus petit ou égal à 5.

Littéralement, on pourrait traduire ce code en français par :

```
si x est strictement plus grand que 5
alors afficher "plus"
sinon afficher "moins"
afficher " et la suite"
```

