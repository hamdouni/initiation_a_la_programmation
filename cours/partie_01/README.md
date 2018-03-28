# Initiation à la programmation informatique - Partie 1

Faisons connaissance avec le langage Go. Apprenons à le lire et le manipuler. Puis écrivons un premier programme utile.

## Le bac à sable

Avec votre navigateur, allez sur http://play.golang.com

Vous trouverez un écran subdivisé en 3 parties :

![bac-à-sable](assets/01_playground.png)

01. Le fond bleu, en haut et à l'horizontal, délimite le menu.
02. Le fond jaune, au centre, contient le programme
03. enfin le fond blanc, en bas de page, affichera les résultats

C'est le *playground* de Go (bac à sable en anglais) : un endroit où l'on peut s'initier à Go directement dans son navigateur.

Nous allons modifier le code. 

Dans la partie jaune, repérez la ligne suivante (*située à la 8ème position*) :

```go
    fmt.Println("Hello, playground")
```

1. Cliquez à l'intérieur du bloc jaune.
2. Positionnez le curseur avant le mot **playground**.
3. Supprimez le et remplacez le par votre prénom.

Voilà ce que cela donne pour exemple avec mon prénom :

```go
    fmt.Println("Hello, Brahim")
```

Cliquez sur le bouton "Run".

**Astuce** : au lieu de cliquez sur "Run", vous pouvez utiliser le raccourci clavier **SHIFT + ENTER**  .

![hello-brahim](assets/01_hello_brahim.png)

Si vous lisez "Hello" suivi de votre prénom, félicitation, vous venez d'exécuter un programme Go !

## Notre premier programme

Vous avez compris que ceux sont les lignes de code qui se trouvent dans le bloc jaune, qui ont produit ce résultat. 

Nous allons écrire un programme un peu plus compliqué : affichons la table de multiplication du nombre 7. J'ai toujours du mal à la retenir :-)

Commençons par effacer entièrement le programme existant.

1. Cliquez à l'intérieur du bloc jaune
2. Sélectionnez tout le contenu à la souris ou avec le raccourci clavier Ctrl+A (Windows ou Linux) ou Cmd+A (Mac)
3. Appuyez sur la touche **Suppr** du clavier

Maintenant, nous avons un espace vide rien que pour nous.

Ecrivons la ligne suivante :

```go
package main
```

En anglais, **package** signifie "paquet" et **main** signifie principal. 
Pour le moment, retenons que notre programme se trouve dans le paquet principal.

Passons à la ligne (touche **Entrée**) et entrons la suite : 

```go
func main() {
```

Le mot **func** introduit une fonction. Elle permet de regrouper du code. Chaque fonction porte un nom, ici **main**.

La fonction **main** a une propriété particulière pour Go : c'est là que notre programme commence.

Nous verrons plus en détail les fonctions.

Pour l'instant, continuons notre programme en tapant à la ligne :

```go
	for i := 1; i <= 10; i++ {
```
**for** nous permet de répéter plusieurs fois le même code. On dit que **for** est une boucle. 

On indique que l'on va utiliser une variable que l'on nomme **i**.

Pour faire simple, disons qu'une variable est une case mémoire dans laquelle on peut stocker une donnée pour la manipuler.

**i := 1** stocke la valeur initale **1** dans notre variable **i**.

**i <= 10** indique que l'on veut que notre variable ne dépasse pas **10**.

**i++** fera avancer notre varuable de 1 en 1.

Donc, notre variable **i** va prendre les valeurs successives 1, 2, 3, ... etc, jusqu'à 10.

Avez vous remarqué une accolage à la fin de la ligne ? Le code à répéter par notre boucle **for** se trouve entre une accolade ouvrante et une accolade fermante.

Ajouter à la ligne le code à répeter :

```go
		print(i*7, " ")
```

**print** affiche un contenu à l'écran. 

Ici, on affiche **i*7**, la variable i qu'on multiplie par 7, suivi d'un petit espace **" "** pour que nos nombres ne s'affichent pas collés les uns aux autres.

**Pour information** : La petite étoile **\*** représente le signe de multiplication, et non une croix comme on a l'habitude. Les autres opérations arithmétiques sont **+** et **-**, sans surprise, mais **/** pour la division. 

Comme **i** vaut 1 au début de notre boucle, on affichera 1*7, soit 7.

Puis **i** passe à 2, et on affichera 2*7, soit 14.

Et ainsi de suite jusqu'à ce que **i** vaut 10.

N'oubliez pas d'ajouter 2 accolades fermantes : une pour fermer celle de la ligne **for** et l'autre pour la ligne **func**.

Vous devriez obtenir le programme complet suivant :

```go
package main

func main() {
	for i := 1; i <= 10; i++ {
		print(i*7, " ")
	}
}
```

Vous pouvez utiliser le bouton **Format** pour que votre code soit mieux présenté.

Appuyez sur **Run** ou Shift+Enter et voilà :

![table_de_7](assets/01_table_de_7.png) 

Notre premier programme en Go fait quelque chose d'utile... en tous les cas pour moi :-)

Si vous n'obtenez pas ce résultat, ou qu'un message d'erreur en anglais s'affiche, pas de panique ! Cela arrive tout le temps quand on programme et voici ce que l'on doit faire :

1. Relisez attentivement votre code, et faites très attention aux ponctuations : par exemple une virgule à la place d'un point-virgule, aux oublies de signes comme une parenthèse fermante, ou aux fautes d'orthographes, un **primt** n'est pas un **print**.
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
1	2	3	4	5	6	7	8	9	10	
2	4	6	8	10	12	14	16	18	20	
3	6	9	12	15	18	21	24	27	30	
4	8	12	16	20	24	28	32	36	40	
5	10	15	20	25	30	35	40	45	50	
6	12	18	24	30	36	42	48	54	60	
7	14	21	28	35	42	49	56	63	70	
8	16	24	32	40	48	56	64	72	80	
9	18	27	36	45	54	63	72	81	90	
10	20	30	40	50	60	70	80	90	100	
```
