## Le type chaine de caractères

Pour représenter du texte, on utilisera le type _string_ (chaine en français).

```go
    a := "De la terre à la lune"
    println(a)
```

L'addition de 2 chaines de caractères revient à les concaténer. Ainsi, si l'on écrit :

```go
    a := "De la terre "
    b := "à la lune"
    c := a + b
    println(c)
```

on obtiendra l'affichage de "De la terre à la lune" à l'écran.

 Remarquez que les autres opérations arithmétiques ne s'appliquent pas pour les chaines de caractères (soustraction, multiplication ou division).