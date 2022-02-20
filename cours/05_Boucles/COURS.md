# Initiation à la programmation Go - Partie 5 - Les boucles

En programmation, nous faisons appel régulièrement à la notion de répétition : nous devons répéter le même traitement sur des données différentes. Nous avons à notre disposition un outil puissant : les boucles.

## Répéter des opérations : les boucles

Pour faire compter l'ordinateur de 1 à 10, et compte tenu de ce que l'on connait déjà, on pourrait écrire :

```
println(1)
println(2)
println(3)
println(4)
println(5)
println(6)
println(7)
println(8)
println(9)
println(10)
```

On peut faire plus court en utilisant une structure de boucle. 

En informatique, une boucle est un ensemble d'opérations que l'on peut répéter. Voici notre exemple réécrit :

```
for i := 1; i <= 10; i++ {
    println(i)
}
```

L'instruction qui nous intéresse est le _for_ qui se construit en 4 blocs :

- initialisation d'une variable : i := 1
- condition de poursuite : i &lt;= 10
- modification d'avancement : i = i + 1
- code à répéter entre accolades 

## Initialisation de la boucle

L'initialisation consiste à définir la valeur initiale d'une variable.

Cette variable peut servir à compter le nombre de répétitions.

Elle peut également servir au sein de notre boucle, comme dans notre cas avec  _println(i)_

## Condition de poursuite

La condition de poursuite indique si l'on doit ou non arrêter la boucle. Tant que la condition est remplie (ici i &lt; ou = à 10), on continue. Dès que ce n'est plus le cas, on sort du _for_ et on poursuit le programme après l'accolade fermante.

Attention à toujours vous assurer que la condition sera remplie un jour, au risque d'avoir une boucle qui ne s'arrête jamais. 

L'exemple suivant montre un tel cas :

```
for a := 1; a >= 1; a++ {
    println("hey!")
}
```

La condition a >= 1 sera toujours vraie et donc la boucle sera infinie !

## Avancement de la boucle

En 3eme position du _for_, la modification d'avancement permet de spécifier les changements à opérer à chaque passage dans la boucle. Ici, avec i = i + 1, on indique que l'on souhaite augmenter de 1 la variable i à chaque tour (i prendra successivement les valeurs de 1 à 10)

## Code à répéter

Enfin le code à répéter peut, mais ce n'est pas une obligation, utiliser la variable de boucle : ici i. Dans notre cas, on l'utilise car i prend les valeurs de 1 à 10 et c'est exactement ce que l'on veut afficher. 

# Imbrication de boucles

On peut imbriquer les boucles si on prend soin de ne pas mélanger les variables. Par exemple, voici comment on pourrait afficher les tables de multiplication :

```
for x:=1; x<=10; x++ {
    for y:=1; y<=10; y++{
        println(x,"x",y,"=",x*y)
    }
}
```

# Autre forme du _for_

Il est également possible d'omettre certaines parties du _for_.

On peut, par exemple, ne garder que la condition de poursuite comme dans l'exemple ci-dessous qui affiche les premiers nombres de la suite de [Fibonacci](https://fr.wikipedia.org/wiki/Suite_de_Fibonacci) :

```
n,p,f := 1,1,0
for f < 1000 {
    f = n + p
    n = p
    p = f
    println(f)
}
```

La boucle se répète tant que le nombre de Fibonacci, rangé dans la variable _f_, reste inférieur à 1000. Dès qu'il dépasse, on s'arrête.

