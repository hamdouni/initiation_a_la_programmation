---
id: 7a58a8f9-cd04-4068-ae89-65755922880d
title: 07_exemple   Droit De Vote
desc: ''
updated: 1610816806917
created: 1610816806917
---
# Exemple : droit de vote

Nous allons illustrer l'utilisation des conditions pour faire un petit programme qui nous indique si on a le droit de vote selon l'année de notre naissance. Pour simplifier, on ne se préoccupera pas du mois et du jour.

```
Réfléchissons un peu !

On nous donne l'année de naissance de la personne et on suppose que l'on connait l'année en cours.

Savoir si on a le droit de vote, c'est déterminer si on a au moins 18 ans. 

Donc, on doit d'abord calculer l'âge en effectuant la soustraction : année en cours - année de naissance.

Ensuite, on compare l'âge avec 18. 
Si l'âge est égale ou plus grand que 18, c'est bon. Sinon, c'est qu'on est trop jeune pour voter.
```

Traduisons cela en code. Nous avons besoin d'une variable pour indiquer une année de naissance et une autre variable pour indiquer l'année en cours :

```
naissance := 1998
maintenant := 2017
```

On calcule l'âge de la personne :

```
age := maintenant - naissance
```

Si l'âge de la personne est supérieur ou égal à 18, c'est qu'elle est en droit de voter. 

```
if age >= 18 {
    print("vous avez le droit de vote")
}
```

Sinon, on lui indique de patienter :

```
else {
    print("patiencez encore un peu")
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

