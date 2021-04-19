---
id: e94f2a7b-8667-4e87-b1b5-0dd28bdff84c
title: 02_notre 1er Programme
desc: ''
updated: 1610816806915
created: 1610816806915
---
# Notre 1er programme

Créez un dossier "bonjour" et à l'intérieur, créez le fichier "main.go" avec votre éditeur préféré. Tapez-y les lignes suivantes :

```go
  package main
  func main() {
    println("bonjour")
  }
```

Sauvegardez le fichier et tapez en ligne de commande (toujours dans le dossier "bonjour")

```
$ go build
```

Par convention, le $ signale le prompt du terminal pour indiquer ce que vous devez taper en ligne de commande (Il ne faut évidemment pas taper le signe $ )

"go build" compile le programme en un exécutable nommé "bonjour" (du nom du dossier), que vous pouvez lancer en version Mac ou Linux :

```
$ ./bonjour
```

Et en version Windows :

```
$ bonjour.exe
```

Par défaut, Go donne à l'exécutable le nom du dossier qui le contient. Il ne doit y avoir qu'un seul fichier nommé "main.go" dans le dossier.

Il est possible de choisir un autre nom pour l'exécutable avec l'option -o comme suit :

```
$ go build -o salam
```

Le résultat est la création d'un fichier "salam" (ou "salam.exe" sous Windows).

Il est également possible, quand on n'a qu'un seul fichier, d'exécuter le programme sans le compiler avec l'option "run" comme suit :

```
$ go run main.go
```

