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