# Initiation à la programmation Go - Partie 8

Pour cette exercice, nous nous prenons pour un cinéma de quartier qui aurait besoin d'un logiciel pour nous donner le prix de la place en fonction de l'année de naissance du client. Il existe 3 tarifs :

- Mineur (>= 18 ans) : 5 €
- Senior (>= 60 ans) : 7 €
- Pour tous les autres : 9 €

```
Réfléchissons un peu !

On a déjà écrit le code pour calculer l'âge d'une personne dans un précédent exemple. On va juste recopier cette partie. 

Pour les tarifs, comme il y en a 3, mieux vaut utiliser la construction _switch case_ plutôt que _if_. Chaque tarif sera traité par un cas, avec le tarif autre dans l'instruction _default_.

```

Voilà ce que cela pourrait donner :

```go
package main

func main() {
    naissance := 2000
    maintenant := 2017

    age := maintenant - naissance

    switch {
    case age <= 18:
        print("5 €")
    case age >= 60:
        print("7 €")
    default:
        print("9 €")
    }
}
```

