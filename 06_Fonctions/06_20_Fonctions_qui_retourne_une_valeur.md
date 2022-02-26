## Fonctions qui retourne une valeur

Une fonction peut renvoyer une valeur à celui qui l'a invoqué.

```
func addition(a,b int) (somme int) {
    return a+b
}
```

Cette fonction attend 2 valeurs en paramètres et retourne leur somme. On déclare la valeur attendue en dernier : _somme int_.

C'est l'instruction _return_ qui indique ce que la fonction va renvoyer : ici _a+b_ qui correspond bien à la somme de nos 2 paramètres.

Pour utiliser cette fonction, on écrira :

```
    resultat := addition(5,7)
```

qui mettra le résultat de la somme de 5 par 7 dans la variables _resultat_.