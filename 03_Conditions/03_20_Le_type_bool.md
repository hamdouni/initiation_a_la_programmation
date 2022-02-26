## Le type bool

Une condition doit renvoyer "true" ou "false", respectivement vrai ou faux. Go appelle ce type _bool_ du nom du mathématicien Bool, qui créa l'algèbre binaire qui porte son nom (cf Wikipedia)

On peut donc écrire :

```
c := true
if c {
    print("toujours")
}
```

qui affichera systématiquement le mot "toujours". Respectivement, le code suivant n'affichera jamais rien :

```
c := false
if c {
    print("jamais")
}
```