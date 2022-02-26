# Initiation à la programmation Go - Partie 6 - Les fonctions

Les fonctions en Go permettent de regrouper des opérations pour pouvoir les éxécuter plusieurs fois. Dès que du code est répété plus d'une fois, il sera nécessaire de le regrouper sous une fonction.

## Déclarons une fonction

```
func bonjour() {
    println("hello")
}
```

Nous avons déclaré une nouvelle fonction _bonjour_ qui ne fait qu'afficher "hello" à l'écran.

On indique le nom de la fonction que l'on souhaite après l'instruction _func_.

Le jeu de paranthèse vide indique que la fonction n'attend pas de paramètres (voir plus bas pour les fonctions avec paramètres).

Puis vient entre accolades le traitement de la fonction : ici un simple affichage à l'écran _println_.

Pour exécuter cette fonction, on utilise son nom comme une instruction :

```
bonjour()
bonjour()
bonjour()
```

Ici, nous avons appelé 3 fois notre nouvelle fonction _bonjour_.
On obtiendra donc à l'écran :

```
hello
hello
hello
```

Voilà le code en entier :

```go
    package main

    func bonjour() {
        println("hello")
    }

    func main() {
        bonjour()
        bonjour()
        bonjour()
    }
```
