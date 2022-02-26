## Organiser son code : définir des unités fonctionnelles

Les paquets servent également à organiser son code en regroupant les fonctionnalités cohérentes entre elles. 

Dans ce cas, chaque paquet sera logé dans un dossier pour séparer physiquement le code.

Reprenons l'exemple de notre cinéma de quartier et de son outil de vente de tickets : 

```
[+] impression/
[+] tarifs/
    main.go
```

Le paquet _impression_ regrouperait les fonctionnalités liées à l'impression des tickets.

Le paquet _tarifs_ regrouperait les fonctionnalités liées au calcul des tarifs.

Le programme principal `main.go` pourrait utiliser les fonctionnalités de ces paquets avec la syntaxe suivante :

```go
tarif := tarifs.Calcul(age)
impression.Ticket(tarif)
```

Dans le dossier `tarifs` on aurait un fichier `calcul.go` avec :

```go
package tarifs

func Calcul(age int) (tarif_cts int) {
```

## Visibilité: règle de la majuscule

Go impose une convention de nommage des fonctions pour déterminer leur visibilité. 

Une fonction dont le nom commence par une majuscule, comme dans l'exemple ci-dessus, rend la fonction utilisable par les autres paquets. Ainsi, il est possible d'appeler `tarifs.Calcul` depuis notre programme principal.

A l'inverse, une fonction, dont le nom commence par une minuscule, ne sera visible que dans le paquet où elle a été déclarée.

Si notre paquet `tarifs` contenait la fonction suivante :

```go 
func ristourne(prix_cts int) int {
```

Depuis notre programme principal `main.go`, nous n'aurions pas pu écrire :

```go
prix := tarifs.ristourne(1200)
```

La fonction `ristourne` ne peut être utilisée qu'à l'intérieur du paquet `tarifs`.