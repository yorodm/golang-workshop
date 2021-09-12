# Variables

## Definiendo variables

Una variable es un nombre que se le asigna a una porción de memoria para
almacenar un valor de un determinado tipo.

El keyword `var`{.verbatim} se usa para declarar variables. La sintaxis
es `var <name> <type>`{.verbatim} seguido (opcionalmente) por un valor
inicial.

### Ejemplos de variables

``` go
var a int // Inicializado por defecto

var b = 10 // Se infiere el tipo

var c, d = 20, 30 // Asignando valores múltiples

e, f := 40, 50 // Crear y asignar
```

### Valor inicial

Si no se especifica un valor inicial Golang asigna uno por defecto
dependiento del tipo de dato de la variable.

Para imprimir el valor de una variable se puede usar
`Println`{.verbatim}.

`fmt.Println("a = ", a)`{.verbatim}
