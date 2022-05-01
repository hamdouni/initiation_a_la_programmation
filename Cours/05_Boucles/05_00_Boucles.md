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

```go
for i := 1; i <= 10; i = i + 1 {
    println(i)
}
```

L'instruction qui nous intéresse est le _for_ qui se construit en 4 blocs :

- initialisation d'une variable : i := 1
- condition de poursuite : i &lt;= 10
- modification d'avancement : i = i + 1
- code à répéter entre accolades 
