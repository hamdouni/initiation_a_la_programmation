## Le bac à sable

Avec votre navigateur, allez sur [http://play.golang.com](http://play.golang.com)

Vous trouverez un écran subdivisé en 3 parties :

![bac-à-sable](../assets/01_playground.png)

01. Le fond bleu, en haut et à l'horizontal, délimite le menu.

02. Le fond jaune, au centre, contient le programme.

03. Enfin le fond blanc, en bas de page, affichera les résultats.

C'est le **bac à sable** de Go (*playground* en anglais) : un endroit où l'on peut s'initier à la programmation, sans rien installer sur son ordinateur, avec un simple navigateur.

On constate qu'il y a déjà du code présent dans la partie jaune. Appuyons sur le bouton **Run**.

Vous devriez constater le message suivant dans la zone blanche :

``` 
Hello, playground
```

Nous venons d'exécuter un programme écrit en Go. Tentons de le modifier. 

Dans la partie jaune, repérez la ligne suivante (*située à la 8ème position*) :

```go
    fmt.Println("Hello, playground")
```

1. Cliquez à l'intérieur du bloc jaune.

2. Positionnez le curseur avant le mot **playground**.

3. Supprimez ce dernier et remplacez le par votre prénom.

Voilà ce que cela donne pour exemple avec mon prénom :

```go
    fmt.Println("Hello, Brahim")
```

Cliquez sur le bouton **Run**.

**Astuce** : au lieu de cliquez sur "Run", vous pouvez utiliser le raccourci clavier **SHIFT + ENTER**  .

![hello-brahim](../assets/01_hello_brahim.png)

Si vous lisez "Hello" suivi de votre prénom, félicitation ! Vous avez réussis notre premier exercice : modifier du code !

Nous allons maintenant écrire notre propre code dans notre [premier programme](./01_20_Premier_programme.md).