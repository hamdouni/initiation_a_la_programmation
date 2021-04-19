---
id: cd3dfb82-3155-4ec1-b9b9-25f91caae7e4
title: 05_exemple   Calculer Son Age
desc: ''
updated: 1610816806916
created: 1610816806916
---
# Exemple : calculer son âge

Nous allons illustrer ces principes avec un petit programme qui va calculer notre âge sans tenir compte du mois et du jour pour faire simple. 

```
Réfléchissons un peu !

On nous donne l'année de naissance de la personne et on suppose que l'on connait l'année en cours.

Calculer l'âge d'une personne c'est soustraire l'année en cours avec son année de naissance : année en cours - année de naissance. 
```

Traduisons cela en code. Nous avons besoin d'une variable pour y ranger l'année en cours et d'une autre pour l'année de naissance :

```
maintenant := 2017
naissance := 1972
```

On commence par définir une variable _maintenant_ qui va contenir l'année en cours et une variable _naissance_ qui sera notre année de naissance (à ajuster selon votre cas)

Comme on l'a déjà vu, Go déduit automatiquement qu'il s'agit de type _int_ car 2017 et 1972 sont des nombres entiers.

Pour obtenir l'âge, on soustrait ces 2 variables :

```
age := maintenant - naissance
```

Le résultat de la soustraction étant aussi un entier, "age" prendra le type _int_. Le résultat est alors affiché à l'aide de l'instruction _print_ :

```
	print("J'ai ", age, " ans")
```

Voilà ce que donne le programme en entier :

```go
package main
func main() {
    maintenant := 2017
    naissance := 1972
    age := maintenant - naissance
    print("J'ai ", age, " ans")
}
```

