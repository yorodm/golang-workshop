# Maps

## Tipos asociativos, `map`{.verbatim} {#maps-inicial}

Un `map`{.verbatim} es un tipo de dato incluido en Golang que asocia una
clave con un valor. El valor puede ser recuperado a partir de la clave.

Un `map`{.verbatim} vacío se crea con la función `make`{.verbatim}:

``` go
make(map[type of key]type of value)

eg: make(map[string]int)
```

## Agregando valores

Los valores en un `map`{.verbatim} se referencian igual que en un
arreglo. Un valor se agrega a un `map`{.verbatim} asignándole una clave,
si la clave ya existe el valor para esa clave se sobrescribe.

``` go
  m := make(map[string]int)
  m["a"] = 1
  m["b"] = 2

  fmt.Println(m)

// Output: map[a:1 b:2]
```

### Inicializando un `map`{.verbatim}

Un `map`{.verbatim} se puede inicializar con valores en su declaración
igual que un arreglo:

``` go
m := map[string]int {
    "a": 1,
    "b": 2,
}
```

## Obteniendo valores {#retrieving-values}

Una manera de chequear si una clave existe es la siguiente:

``` go
data, ok := m["some key"]
```

Si `ok`{.verbatim} es `true`{.verbatim}, la clave existe y la variable
`data`{.verbatim} contendrá la información recuperada, si
`ok`{.verbatim} es `false`{.verbatim} la clave no existe.

Acceder a una clave inexistente en un `map`{.verbatim} causa un
`panic`{.verbatim}.

## Iteraciones

El iterador `for`{.verbatim} se utiliza para recorrer las claves y
valores de un `map`{.verbatim}

``` go
  m := map[string]int {
      "a": 1,
      "b": 2,
      "c": 3,
  }
  for key, value := range m {
      fmt.Println("Key:", key, " Value:", value)
  }
// Output: Key: a  Value 1
//         Key: b  Value 2
//         Key: c  Value 3
```

```{=org}
#+REVEAL: split
```
Se puede utilizar el identificador en blanco (`_`{.verbatim}) para
descartar cualquiera de los elementos del par.

> El orden en que se recorre el `map`{.verbatim} no esta garantizado, y
> cada ejecución puede tener un order diferente.

## Eliminacion

Los valores se eliminan de un `map`{.verbatim} con la función
`delete`{.verbatim}. La función no retorna nada. La sintaxis es la
siguiente:

``` go
delete(m, key)
```

## Tamaño

El tamaño de un `map`{.verbatim} es retornado por la funcion
`len`{.verbatim}. Esta funcion retorna la cantidad de claves que hay en
un map. La sintaxis es la siguiente:

``` go
len(m)
```

## Tipo

Un `map`{.verbatim} es un *tipo referencia*. Por lo tanto si un
`map`{.verbatim} se asigna a otra variable, ambas variable apuntan a los
mismos pares (clave, valor). Lo mismo sucede si se pasa como argumento
en una funcion.
