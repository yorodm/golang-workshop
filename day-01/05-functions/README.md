# Funciones

## Funciones en Golang {#golang-funcs}

Una función es un grupo de instrucciones que se ejecutan todas juntas
como un bloque. Una función puede o no tener argumentos de entrada, y
retornar valores.

En Golang una función se define con el keyword `func`{.verbatim}. El
siguiente es un ejemplo de una función que suma dos enteros:

``` go
func AddIntegers(a int, b int) int {
    return a + b
}
```

## Retornando valores {#golang-returns}

la palabra reservada `return`{.verbatim} es usado para indicar qué valor
va a retornar la función.

Golang soporta que una función devuelva múltiples valores:

``` go
func SumDifference(a int, b int) (int, int) {
    return a + b, a - b
}
```

## Identificador en blanco {#blank-identifier}

Se utiliza el **identificador en blanco** en el lugar de un valor que se
quiere descartar al llamar a una función:

``` go
var _, diff = SumDifference(10, 20)

fmt.Println("Difference is ", diff)
```

## Valores de retorno con nombres {#named-return-values}

Cuando se define una función se le puede asignar un nombre al tipo de
dato de retorno para luego referenciarlo en el código de la función.

``` go
func Product(a int, b int) (prod int) {
    prod = a * b
    return
}
```

Al asignar un valor de retorno con nombre no hace falta incluirlo en la
sentencia `return`{.verbatim}.

## Funciones anónimas y clausuras

Golang soporta funciones anónimas y clausuras. Tomemos por ejemplo la
función `sort.Slice`{.verbatim}

``` go
func Slice(slice interface{}, less func(i, j int) bool)
```

El parámetro `less`{.verbatim} describe una función que toma dos enteros
y retorna un valor `bool`{.verbatim}

### Funciones anónimas.

Podemos asignar una función a una variable o definirla en el momento de
su uso. Las funciones anónimas forman **clausuras**

``` go
c := []int{20, 11, 12, 1, 5}
less := func(i int, j int) bool {
    return c[i] < c[j]
}
sort.Slice(c, less)
fmt.Println("Despues del sort", c)
//Output: [1 5 11 12 20]
c := []int{20, 11, 12, 1, 5}
sort.Slice(c, func(i int, j int) bool {
    return c[i] < c[j]
})
fmt.Println("Despues del sort", c)
//Output: [1 5 11 12 20]
```
