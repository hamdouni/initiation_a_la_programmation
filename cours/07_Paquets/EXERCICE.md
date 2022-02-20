# Exemple : intéragir avec l'utilisateur via saisie

On va demander 3 phrases à l'utilisateur, qu'il va saisir au clavier, et que l'on va réafficher tout en majuscule.

Plusieurs notions déjà vues seront utilisées :

- l'import de package
- l'utilisation de fonctions
- les boucles

Voilà le code complet :

```go
package main

import (
	"bufio"
	"os"
	"strings"
)

func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func main() {
	for i := 1; i <= 3; i++ {
		print("Saisir le texte n° ", i, " : ")
		saisie := input()
		print(strings.ToUpper(saisie))
	}
}
```

Notre fonction _input_ va récupérer l'entrée standard _os.Stdin_  dans une variable _reader_. 

```
reader := bufio.NewReader(os.Stdin)
```

Les caractères saisies par l'utilisateur sont lus par la fonction _ReadString_ jusqu'à ce qu'elle tombe sur un saut de ligne _\\n_ qui correspond à la saisie de la touche Entrée du clavier :

```
text, _ := reader.ReadString('\n')
```

Remarquez que nous avons délibérément ignoré le 2ème retour de la fonction _ReadString_, qui est un code erreur (sans importance dans notre contexte).

Le résultat de la saisie est ensuite retourné comme valeur de notre fonction :

```
return text
```

Dans notre programme principal _main_, on boucle 3 fois avec la variable _i_ :

```
for i := 1; i <= 3; i++ {
```

A chaque tour de boucle, on affiche une invitation à saisir du texte :

```
	print("Saisir le texte n° ", i, " : ")
```

avec la petite subtilité qu'on indique le numéro de saisie grâce à la variable _i_.

Ensuite, on effectue l'appel à notre fonction de saisie :

```
	saisie := input()
```

donc, on récupère, dans la variable _saisie_ le texte entré par l'utilisateur.

Et enfin, on n'affiche ce même texte, mais passé par la fonction ToUpper qui aura pour conséquence de l'afficher en majuscule :

```
	print(strings.ToUpper(saisie))
```

