# Funciones

Una funcion es un grupo de instrucciones que se ejecutan todas juntas
como un bloque. Una funcion puede o no tener argumentos de entrada, y
retornar valores.

En Golang una funcion se define con el keyword `func` con el siguiente
formato:

```go
func <name>(<arg1> <type of arg1>, ...) <return type> {
    // cuerpo de la funcion (una o mas instrucciones)
}
```
El siguiente es un ejemplo de una funcion que suma dos enteros:

```go
func AddIntegers(a int, b int) int {
    return a + b
}
```
El keyword `return` es usado para indicar que valor va a retornar la funcion.

Para llamar una funcion:

```go
AddIntegers(10, 20) => 30
```
Golang soporta que una funcion devuelva multiples valores:

```go
return a, b
```

```go
func SumDifference(a int, b int) (int, int) {
    return a + b, a - b
}
```

En el ejemplo anterior se calcula los valores en la misma linea junto
con la instruccion `return`.

## Blank Identifier

Se utiliza el __blank identifier__ en el lugar de un valor que se
quiere descartar al llamar a una funcion:

```go
var _, diff = SumDifference(10, 20)

fmt.Println("Difference is ", diff)
```


## Named Return Values

Cuando se define una funcion se le puede asignar un nombre al tipo de
dato de retorno para luego referenciarlo en el codigo de la funcion.


```go
func Product(a int, b int) (prod int) {
    prod = a * b
    return
}
```

En el ejemplo anterior se asigna el nombre `prod` al valor de retorno
de la funcion. Luego `prod` se usa como una variable y se le asigna un
valor. Como `prod` ya tiene un valor asignado no hace falta indicar un
valor en la instruccion `return`.

## Funciones anónimas y clausuras
