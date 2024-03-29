---
id: 3063946a-25e9-43ce-abac-e2e7a18e3734
title: 08_gerer Des Cas   Switch Case
desc: ''
updated: 1610816806917
created: 1610816806917
---
# Gérer des cas

## Remplacement du _if_

Lorsqu'on doit effectuer plus de 2 comparaisons, le _if_ n'est pas toujours le plus pratique. 

Imaginons qu'on ait besoin de tester les 3 possibilités pour un nombre par rapport à 10 : soit il est égal à 10, soit il est inférieur, soit il est plus grand.

On pourrait écrire :

```
nombre := 7
if nombre == 10 {
    print("égal")
} else if nombre < 10 {
    print("inférieur")
} else if nombre > 10 {
    print("supérieur")
}
```

Remarquez que lorsque les 2 premiers tests échouent, le dernier test est forcément vrai. On pourrait simplifier le code en ommetant le dernier test comme suit :

```
nombre := 7
if nombre == 10 {
    print("égal")
} else if nombre < 10 {
    print("inférieur")
} else {
    print("supérieur")
}
```

Une forme plus élégante serait d'utiliser les instructions _switch_ et _case_ comme suit :

```
nombre := 7
switch {
case nombre == 10:
    print("égal")
case nombre < 10:
    print("inférieur")
default:
    print("supérieur")
}
```

Chaque cas à tester se décrit avec l'instruction _case_ suivie du test à réaliser et de deux points.

```
case nombre == 10:
```

On ajoute à la ligne les différentes instructions à exécuter si la condition est vraie.

```
case nombre == 10:
    print("égal")
```

Si aucun cas n'est vérifié, c'est l'instruction _default_ qui est exécutée :

```
default:
    print("supérieur")
```

Le _switch_ délimite tous les cas possibles en les entourant d'une accolades ouvrantes et fermantes.

```
switch {
...
}
```

## Cas général

Dans sa forme complète, l'utilisation du couple d'instruction _switch-case_ se fait comme suit :

```
nombre := 2
switch nombre {
case 1:
    print("un")
case 2:
    print("deux")
case 3:
    print("trois")
default:
    print("autre")
}
```

On indique la variable à évaluer derrière l'instruction _switch_ :

```
switch nombre {
```

Puis on décrit chaque cas avec la valeur que devrait avoir notre variable. Par exemple, pour le cas où notre variable vaudrait 3, notre _case_ ressemblerait à :

```
case 3:
    print("trois")
```

