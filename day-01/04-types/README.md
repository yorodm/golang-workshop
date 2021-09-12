# Tipos

## Tipos simples

Los siguientes tipos de datos están disponibles en Golang:

-   bool
-   Tipos numéricos:
    -   int8, int16, int32, int64, int
    -   uint8, uint16, uint32, uint64, uint
    -   float32, float64
    -   complex64, complex128
    -   byte
    -   rune
-   string
-   error
-   chan

### Booleanos

`bool`{.verbatim} representa valores boleanos `true`{.verbatim} o
`false`{.verbatim}.

### Tipos numéricos

Para tipos numéricos, el número que sigue al tipo indica la cantidad de
bits que se usa para representar el valor en memoria. El número de bits
determina qué tan grande el número puede ser.

`int8`{.verbatim}: 8 bit enteros con signo

`int16`{.verbatim}: 16 bit entero con signo

`int32`{.verbatim}: 32 bit entero con signo

`int`{.verbatim}: 32 or 64 bit entero con signo (dependiende de la
plataforma).

`uint`{.verbatim} : 32 or 64 bit entero sin signo (dependiende de la
plataforma).

`float32`{.verbatim}: 32 bit número flotante.

`float64`{.verbatim}: 64 bit número flotante.

`byte`{.verbatim} alias de `uint8`{.verbatim}

`rune`{.verbatim} alias de `int32`{.verbatim}

### Números complejos

`complex64`{.verbatim}: números complejos con parte real e imaginaria
`float32`{.verbatim}

`complex128`{.verbatim}: números complejos con parte real e imaginaria
`float64`{.verbatim}

La función `complex`{.verbatim} se utiliza para construir un número
complejo con partes reales e imaginarias:

``` go
func complex(r, i FloatType) ComplexType
```

### El tipo string

`string`{.verbatim} es una colección de caracteres(runas) encerrados
entre comillas.

``` go
first := "Allen"
last := "Varghese"
name := first +" "+ last
fmt.Println("My name is",name)
```

### Errores y canales.

`error`{.verbatim} es un tipo especial para indicar condiciones de
error.

`chan`{.verbatim} es un tipo especial que representa un canal de
comunicación.

## Conversion de tipos

No hay conversión automática de tipos en Golang. Se utilizan funciones
de tipo para hacer una conversión.

``` go
int(<float value>)
float64(<integer value>)
```

## Constantes

Las constantes no cambian sus valores una vez asignadas. Se usa la
palabra reservada `const`{.verbatim} en vez de `var`{.verbatim} para
declarar una constante.

``` go
const a bool = true
const b int32 = 32890
const c string = "Something"
```
