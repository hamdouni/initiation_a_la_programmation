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

