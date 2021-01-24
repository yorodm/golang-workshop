# Maps

Un map un tipo de dato incluido en Golang que asocia una clave con un
valor. El valor puede ser recuperado a partir de la clave.

Un map se crea con la funcion `make`:

```go
make(map[type of key]type of value)

eg: make(map[string]int)
```

En el ejemplo anterior se crea un map con una clave de tipo `string`
que puede almacenar y recuperar valores de tipo `integer`.

## Agregando valores

Los valores en un map se referencian igual que en un arreglo. Un valor
se agrega a un map asignandole una clave, si la clave ya existe el
valor para esa clave se sobreescribe.

```go
m := make(map[string]int)
m["a"] = 1
m["b"] = 2

fmt.Println(m)

> map[a:1 b:2]
```

En el ejemplo anterior un nuevo map se crea utilizando la funcion
`make`. Dos claves, `a` y `b` son asignadas con los valores `1` y `2`
respectivamente. Notar que los tipos de datos de las claves y valores
deben coincidir con los especificados en la declaracion del map.

Un map se puede inicializar con valores en su declaracion igual que un
arreglo:

```go
m := map[string]int {
    "a": 1,
    "b": 2,
}
```

## Retrieving Values

Los valores se recuperan de manera similar a como son asignados, pero
al reves. Un error se produce si se intenta recuperar un valor de una
clave inexistente.

Una manera de chequear si una clave existe es la siguiente:

```go
data, ok := m["some key"]
```

Si `ok` es `true`, la clave existe y la variable `data` contendra la
informacion recuperada, si `ok` es `false` la clave no existe.

## Iteraciones

El iterador `for` se puede usar para iterar a traves, de un map. En
vez de de retornar el indice y el valor como en un arreglo, se retorna
la clave y el valor.

```go
m := map[string]int {
    "a": 1,
    "b": 2,
    "c": 3,
}
for key, value := range m {
    fmt.Println("Key:", key, " Value:", value)
}

> Key: a  Value 1
  Key: b  Value 2
  Key: c  Value 3
```

El iterador `for` itera a traves del map hasta que todos los pares
(clave, valor) se terminan.

```text
El orden en que se recorre el map no esta garantizado, y cada ejecucion puede tener un order diferente.
```

## Eliminacion

Los valores se eliminan de un map con la funcion `delete`. La funcion
no retorna nada. La sintaxis es la siguiente:

```go
delete(m, key)
```

## Tamaño

El tamaño de un map es retornado por la funcion `len`. Esta funcion
retorna la cantidad de claves que hay en un map. La sintaxis es la
siguiente:

```go
len(m)
```

## Tipo

Un map es un tipo de dato por referencia, igual que un slice. Por lo
tanto si un map se asigna a una otra variable, ambas variable apuntan
a los mismos pares (clave, valor). Lo mismo sucede si se pasa como
argumento en una funcion.
