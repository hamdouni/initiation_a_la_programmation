# Imbrication de boucles

On peut imbriquer les boucles si on prend soin de ne pas mélanger les variables. Par exemple, voici comment on pourrait afficher les tables de multiplication :

```go
for x:=1; x<=10; x++ {
    for y:=1; y<=10; y++{
        println(x,"x",y,"=",x*y)
    }
}
```
