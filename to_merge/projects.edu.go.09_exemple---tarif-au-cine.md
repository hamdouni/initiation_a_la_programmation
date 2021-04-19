---
id: 1a1a8e59-8234-45f4-9050-1c94ef2a9d35
title: 09_exemple   Tarif Au Cine
desc: ''
updated: 1610816806918
created: 1610816806918
---
# Exemple : tarif au ciné

Nous tenons un petit cinéma de quartier et nous souhaitons mettre en place un petit logiciel qui nous donne le prix de la place en fonction de l'année de naissance du client. Il existe 3 tarifs :

- Mineur (&lt; 18 ans) : 5 €
- Senior (> 60 ans) : 7 €
- Pour tous les autres : 9 €

```
Réfléchissons un peu !

On a déjà écrit le code pour calculer l'âge d'une personne dans un précédent exemple. On va juste recopier cette partie. 

Pour les tarifs, comme il y en a 3, mieux vaut utiliser la construction _switch case_ plutôt que _if_. Chaque tarif sera traité par un cas, avec le tarif autre dans l'instruction _default_.

```

Voilà ce que cela pourrait donner :

```
package main

func main() {
    naissance := 2000
    maintenant := 2017

    age := maintenant - naissance

    switch {
    case age < 18:
        print("5 €")
    case age > 60:
        print("7 €")
    default:
        print("9 €")
    }
}
```

