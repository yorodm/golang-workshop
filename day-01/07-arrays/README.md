# Arreglos

## Arreglos en Golang

Una colección de elementos de un mismo tipo es un arreglo. Como Golang
es un lenguaje fuertemente tipado no es posible mezclar valores de
diferentes tipos en un arreglo.

Un arreglo se define como `[size]type`{.verbatim}:

``` go
  var a [3]int
  fmt.Println(a)

// Output: [0 0 0]
```

### Comenzando en 0.

El primer índice de un arreglo es `0`{.verbatim} en vez de
`1`{.verbatim}. Se accede a los valores de un arreglo usando el número
del índice, y se asignan valores con el operador `=`{.verbatim}.

``` go
  var a [3]int
  a[0] = 1
  a[1] = 2
  a[2] = 3
  fmt.Println(a)
// Output: [1 2 3]
```

### Inicializando arreglos.

Un arreglo se puede inicializar en la declaración:

``` go
a := [3]int{1, 2, 3}
```

Se puede omitir el número de elementos si se inicializa con valores en
la declaración:

``` go
a := [...]int{1, 2, 3}
```

### Tipos valor

Un detalle importante a tener en cuenta es que un arreglo es un *tipo
valor*, esto quiere decir que cada vez que un arreglo es asignado a una
nueva variable, se hace una copia del arreglo original. Si se cambian
valores en la nueva variable, esto no se ve reflejado en la variable
original.

``` go
a := [...]string{"IRL", "IND", "US", "CAN"}
b := a

b[1] = "CHN"

fmt.Println("Original:", a)
fmt.Println("Copy    :", b)
```

Esto afecta el modo en que los arreglos son pasados por parámetros.

## Funciones para operar con arreglos.

### `len`{.verbatim}

La función `len`{.verbatim} se usa para conocer el tamaño de un arreglo:

``` go
  a := [...]int{1, 2, 3}
  fmt.Println(len(a))

// Output: 3
```

### `range`{.verbatim}

Una forma de interactuar con un arreglo es utilizar el keyword
`range`{.verbatim}. Este retorna el índice del arreglo y el valor.

``` go
  a := [...]int{1, 2, 3, 4, 5}
  sum := 0
  for i, v := range a {
      fmt.Println("Index:", i, " Value is:", v)
      sum += v
  }
  fmt.Println("Sum:", sum)

// Output: 15
```

### `range`{.verbatim} y el blank identifier

Se puede usar el **blank identifier** (`_`{.verbatim}) si no nos
interesa alguno de los valores que retorna el keyword
`range`{.verbatim}.

``` go
_, v := range a
```

## Arreglos multidimensionales

Se pueden crear arreglos de más de una dimensión de la siguiente manera:

``` go
a := [3][2]int
```

### Inicializando arreglos multidimensionales

``` go
  a := [3][2]string{
      {"lion", "tiger"},
      {"cat", "dog"},
      {"pigeon", "peacock"},
  }
  fmt.Println(a)
// Output: [[lion tiger] [cat dog] [pigeon peacock]]
```

## Slices

Los arreglos tienen tamaño fijo, reservan memoria, y son tipos valor. Un
slice es una forma flexible de acceder a una arreglo. Un slice no tiene
datos, solo apunta a un arreglo.

### Creando slices

Un slice vacío se crea así:

``` go
var sa []int
```

El valor de un slice vacío es `nil`{.verbatim}

Un slice también se puede crear apuntando a un subconjunto de valores de
un arreglo:

``` go
  a := [...]int{1, 2, 3, 4, 5}
  sa := a[1: 4]
  fmt.Println(sa)
// Output: [2 3 4]
```

### Tipos referencia

Como un slice es un *tipo referencia*, modificar un valor en un elemento
del slice modifica el arreglo original.

``` go
  a := [...]int{1, 2, 3, 4, 5}
  sa := a[1: 4]

  fmt.Println("Before:", a)
  sa[0] = 22

  fmt.Println("After:", a)

// Output: Before: [1 2 3 4 5]
// Output: After: [1 22 3 4 5]
```

### La función `make`{.verbatim}

Un slice también se puede crear utilizando la funcion `make`{.verbatim},
especificando el tipo, y el tamaño, y opcionalmente la capacidad (que
indica el máximo tamaño que el slice puede crecer):

``` go
#+begin_src go
  i := make([]int, 5, 5)
  fmt.Println(i)
// Output: [0 0 0 0 0]
```

Crear un slice con `make`{.verbatim} inicializa todos sus valores con
los valores por defecto del tipo del slice.

### Cambiando el tamaño de un slice

El tamaño de un slice se puede incrementar utilizando la función
`append`{.verbatim}.

``` go
  sa := []int{1, 2, 3}
  newSa := append([]int{}, sa...)
  fmt.Println(newSa)
// Output: [1, 2, 3]
```

En vez de valores se puede indicar otro slice. El operador
`...`{.verbatim} se usa para expandir el slice en sus valores.
