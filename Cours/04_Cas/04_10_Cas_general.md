## Cas général

Dans sa forme complète, l'utilisation du couple d'instruction _switch-case_ se fait comme suit :

```
nombre := 2
switch nombre {
case 1:
    print("un")
case 2:
    print("deux")
case 3:
    print("trois")
default:
    print("autre")
}
```

On indique la variable à évaluer derrière l'instruction _switch_ :

```
switch nombre {
```

Puis on décrit chaque cas avec la valeur que devrait avoir notre variable. Par exemple, pour le cas où notre variable vaudrait 3, notre _case_ ressemblerait à :

```
case 3:
    print("trois")
```

