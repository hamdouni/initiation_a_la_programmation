# Initiation à la programmation informatique - Partie 1

Faisons connaissance avec le langage Go. Apprenons à le lire et le manipuler. Puis écrivons un premier programme utile.

## Le bac à sable

Avec votre navigateur, allez sur http://play.golang.com

Vous trouverez un écran subdivisé en 3 parties :

![bac-à-sable](assets/01_playground.png)

01. Le fond bleu, en haut et à l'horizontal, délimite le menu.
02. Le fond jaune, au centre, contient le programme
03. enfin le fond blanc, en bas de page, affichera les résultats

C'est le **bac à sable** de Go (*playground* en anglais) : un endroit où l'on peut s'initier à Go, sans rien installer sur son ordinateur, avec un simple navigateur.

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

![hello-brahim](assets/01_hello_brahim.png)

Si vous lisez "Hello" suivi de votre prénom, félicitation ! Vous avez réussis notre premier exercice.

## Notre premier programme

Vous avez compris que les lignes de code, qui se trouvent dans le bloc jaune, sont responsables de l'affichage. 

Voyons comment nous pouvons écrire un programme un peu plus compliqué : affichons la table de multiplication du nombre 7. J'ai toujours du mal à la retenir :-)

Commençons par effacer entièrement le programme existant.

1. Cliquez à l'intérieur du bloc jaune
2. Sélectionnez tout le contenu à la souris ou avec le raccourci clavier Ctrl+A (Windows ou Linux) ou Cmd+A (Mac)
3. Appuyez sur la touche **Suppr** du clavier

Maintenant, nous avons un espace vide rien que pour nous.

Ecrivons la ligne suivante :

```go
package main
```

En anglais, **package** signifie "paquet" et **main** signifie principal : notre programme fait partie du paquet principal.

Passons à la ligne (touche **Entrée**) et entrons la suite : 

```go
func main() {
```

Le mot **func** introduit une fonction. 

Chaque fonction porte un nom, ici **main**.

Une fonction permet de regrouper du code. Tout ce qui se trouve entre l'accolade ouvrante et fermante fait partie de la fonction.

Nous verrons plus en détail à quoi servent les fonctions.

Pour l'instant, retenons que la fonction **main** a une propriété particulière pour Go car c'est là que débutera notre programme.

Continuons notre programme en tapant à la ligne :

```go
	for i := 1; i <= 10; i++ {
```

**for** nous permet de définir une boucle. 

Une boucle répéte plusieurs fois le code qui se trouve entre accolades.

Pour cela, **for** a besoin de 3 parties séparées par un point-virgule :

![for expliqué](assets/01_for.png)

1. **i := 1** : une variable que l'on nomme **i** et à laquelle on affecte la valeur 1. Une variable est une case mémoire dans laquelle on peut stocker une donnée pour la manipuler.
2. **i <= 10** indique que l'on tourne dans la boucle tant que notre variable ne dépasse pas **10**.
3. **i++** fera avancer notre variable **i** de 1 en 1, à chaque tour de boucle.

Notre variable **i** va donc prendre les valeurs 1 puis 2, 3, ... etc, jusqu'à 10.

Allons à la ligne pour écrire le code qui sera répeté :

```go
		print(i*7, " ")
```

**print** affiche un contenu à l'écran. 

Le contenu a afficher se trouve entre parenthèses :

* **i*7**, la variable i qu'on multiplie par 7
* **" "**, un petit espace pour que nos nombres ne s'affichent pas collés les uns aux autres.

**Deux remarques** : 
1. La multiplication est représentée par une petite étoile **\***, et non une croix comme on a l'habitude. Les autres opérations arithmétiques sont **+** et **-**, sans surprise, mais **/** pour la division. 
2. Notre petit espace est entouré par des guillemets. C'est ainsi qu'on délimite une chaine de caractères, notion qu'on aura l'occasion de recroiser.

Comme **i** vaut 1 au début de notre boucle, on affichera 1*7, soit 7.

Puis **i** passe à 2, et on affichera 2*7, soit 14.

Et ainsi de suite jusqu'à ce que **i** vaut 10, pour afficher 70.

N'oubliez pas d'ajouter les 2 accolades fermantes : une pour fermer l'accolade ouvrante de la ligne **for** et l'autre pour la ligne **func**.

Appuyez sur le bouton **Format** pour que votre code soit mieux présenté.

Vous devriez obtenir le programme complet suivant :

```go
package main

func main() {
	for i := 1; i <= 10; i++ {
		print(i*7, " ")
	}
}
```

Appuyez sur **Run** ou Shift+Enter et voilà :

![table_de_7](assets/01_table_de_7.png) 

Notre premier programme en Go fait quelque chose d'utile... en tous les cas pour moi :-)

Si vous n'obtenez pas ce résultat, ou qu'un message d'erreur en anglais s'affiche, pas de panique ! Cela arrive tout le temps quand on programme et voici ce que l'on doit faire :

1. Lisez le message d'erreur pour repérer le numéro de la ligne incriminée et le type d'erreur.
2. Relisez attentivement votre code, et faites très attention aux ponctuations : par exemple une virgule à la place d'un point-virgule, aux oublies de signes comme une parenthèse fermante, ou aux fautes d'orthographes, **primt** n'est pas **print**.
2. A chaque fois que vous corrigez, relancez le programme pour vérifier que cela marche (**Run** ou Shift+Enter).
3. Si cela ne fonctionne pas, n'hésitez pas à copier/coller le programme complet ci-dessus.

## Conclusion

Nous avons appris à nous servir du bac-à-sable du langage Go, pour écrire un programme, l'exécuter et consulter le résultat.

Nous avons aussi fait connaissance avec quelques mots de ce langage pour afficher des données à l'écran (**print**) ou pour répéter plusieurs fois les mêmes actions (**for**).

Nous avons également effleurer plusieurs notions que nous approfondirons plus tard, comme les paquets (**package**), les fonctions (**func**) ou encore les variables.

Pour vous exercer, vous pouvez changer le programme pour afficher la table d'autres nombres, par exmple celle de 9.

Vous pouvez également expliciter la multiplication en affichant l'opération, par exemple : **7 x 2 = 14**

Et pour les plus téméraires, vous pouvez réfléchir à afficher toutes les tables sous la forme d'un tableau comme suit :

```
1  2  3  4  5  6  7  8  9  10
2  4  6  8  10 12 14 16 18 20
3  6  9  12 15 18 21 24 27 30
4  8  12 16 20 24 28 32 36 40
5  10 15 20 25 30 35 40 45 50
6  12 18 24 30 36 42 48 54 60
7  14 21 28 35 42 49 56 63 70
8  16 24 32 40 48 56 64 72 80
9  18 27 36 45 54 63 72 81 90
10 20 30 40 50 60 70 80 90 100
```

## Prochaine partie

Nous verrons la corrections dans les
[Premiers exercices](../partie_02/README.md).