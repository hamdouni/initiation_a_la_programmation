# Les conditions

La gestion des conditions est une notion importante en programmation. 

Elle permet d'effectuer des actions seulement si des critères sont remplis.

Et pour cela, on utilise l'instruction _if_.

## L'instruction if

Voici un exemple où l'on affiche le texte "x est plus grand" seulement si la
variable x a une valeur strictement plus grande que 5 :

```
if x > 5 {
    println("x est plus grand")
}
print("fin du programme")
```

Ici la condition à évaluer est x > 5.

```
if x > 5 
```

Si x est plus grand que 5, l'action entre accolades sera exécutée :

```
println("x est plus grand")
```

Sinon, on passe à l'instruction après l'accolade fermante. Dans notre cas :

```
print("fin du programme")
```

Voilà ce que cela donnerait avec x = 12

```
x := 12
if x > 5 {
    println("x est plus grand")
}
print("fin du programme")
```

On obtiendrait la phrase "x est plus grand" suivie du message "fin du
programme" car la condition x > 5 est vraie, x valant 12.

Voyons ce que cela donnerait avec un x = 3

```
x := 3
if x > 5 {
    println("x est plus grand")
}
print("fin du programme")
```

Cette fois, il ne s'affichera que "fin du programme" car la condition est
maintenant fausse.


## Le type bool

Une condition doit renvoyer "true" ou "false", respectivement vrai ou faux. Go
appelle ce type _bool_ du nom du mathématicien Bool, qui créa l'algèbre binaire
qui porte son nom (cf Wikipedia)

On peut donc écrire :

```
c := true
if c {
    print("toujours")
}
```

qui affichera systématiquement le mot "toujours". Respectivement, le code
suivant n'affichera jamais rien :

```
c := false
if c {
    print("jamais")
}
```

## L'instruction else

Go permet également d'indiquer une alternative dans le cas où la condition est
fausse, avec l'instruction _else_ qu'on traduirait par "sinon". 

```
if x > 5 {
     println("plus")
}
else {
    println("moins")
}
print("fin programme")
```

Le code entre crochets qui suit le _else_ ne sera exécuté que si la condition
est fausse, c'est-à-dire, dans notre cas, seulement si x plus petit ou égal à
5.

Littéralement, on pourrait traduire ce code en français par :

```
si x est strictement plus grand que 5
alors afficher "plus"
sinon afficher "moins"
afficher "fin programme"
```

## Exercice

Nous allons illustrer l'utilisation des conditions pour faire un petit
programme qui nous indique si on a le droit de vote selon l'année de notre
naissance. 

Pour simplifier, on ne se préoccupera pas du mois et du jour.

```
Réfléchissons un peu !

On nous donne l'année de naissance de la personne et on suppose que l'on
connait l'année en cours.

Savoir si on a le droit de vote, c'est déterminer si on a au moins 18 ans. 

Donc, on doit d'abord calculer l'âge en effectuant la soustraction : année en
cours - année de naissance.

Ensuite, on compare l'âge avec 18.  Si l'âge est égale ou plus grand que 18,
c'est bon. Sinon, c'est qu'on est trop jeune pour voter.
```

Traduisons cela en code. Nous avons besoin d'une variable pour indiquer une
année de naissance et une autre variable pour indiquer l'année en cours :

```
naissance := 1998
maintenant := 2017
```

On calcule l'âge de la personne :

```
age := maintenant - naissance
```

Si l'âge de la personne est supérieur ou égal à 18, c'est qu'elle est en droit
de voter. 

```
if age >= 18 {
    print("vous avez le droit de vote")
}
```

Sinon, on lui indique de patienter :

```
else {
    print("patientez encore un peu")
}
```

Voici le programme en entier :

```go
package main

func main() {
    naissance := 2000
    maintenant := 2017

    age := maintenant - naissance

    if age >= 18 {
        print("Vous avez ", age, " ans, vous avez le droit de vote")
    } else {
        print("Vous n'avez que ", age, " ans, patientez encore un peu")
    }
}
```

