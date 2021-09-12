# Manejo de errores: error, defer, recover, panic

## Manejo de errores en Golang

En lugar de excepciones como Python, Java o C#, Golang sigue un enfoque
más cercano a los lenguajes funcionales donde el estado de error es
representado por un tipo de datos,

``` go
var DivisionByZero = errors.New("División por cero")

func safedivide(a int, b int) (int, error) {
    if b == 0 {
        return 0, DivisionByZero
    }
    return a / b, nil
}
```

## El tipo error {#error-type}

El tipo `error`{.verbatim} es una interfaz *built-in* del lenguaje

``` go
type error interface {
    Error() string
}
```

### Errores personalizados

Esto nos permite crear nuestros propios tipos de errores

``` go
type MyError struct {
    message: string
    status: int
}

func (m *MyError) Error() string {
    return fmt.Sprintf("Error - %s. Status %d", m.message, m.status)
}

func returnMyError() error {
    return &MyError{
     message: "None",
     status : -1,
    }
}
```

### Operando con errores

El paquete `errors`{.verbatim} contiene funciones dedicadas a manejar
tipos de error. Podemos usar la función `errors.Is`{.verbatim} para
verificar si un error es de un tipo específico.

``` go
if _, err := os.Open("non-existing"); err != nil {
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("file does not exist")
    } else {
        fmt.Println(err)
    }
}
```

## Defer

La sentencia `defer`{.verbatim} se utiliza para indicar que la función a
continuación se debe ejecutar a la salida de la función actual

Es muy usual en Golang usar `defer`{.verbatim} para destruir o liberar
cualquier tipo de recurso temporal que se obtenga en la función.

``` go
func doSomething() {
 a := getExternalResource();
 defer a.Release()
 /// Resto de la función
}
```

`defer`{.verbatim} se ejecuta sin importar las causas por las que la
función actual haya terminado, lo que garantiza en el ejemplo anterior
que `a.Release()`{.verbatim} siempre se ejecute.

## Panic y recover

Golang incluye una función especial `panic`{.verbatim} para indicar un
error que no puede ser manejado de forma correcta.

La contraparte de `panic`{.verbatim} es la función `recover`{.verbatim},
que verifica si ocurrió una llamada a `panic`{.verbatim} en el contexto
de la función actual.

Si una llamada a `panic`{.verbatim} no es seguida por un
`recover`{.verbatim}, la función termina y el contexto pasa al
invocador. Esto continúa hasta que se encuentre un recover o se llegue a
la función `main`{.verbatim} en cuyo caso el programa se detendrá con un
mensaje de error.

```{=org}
#+REVEAL: split
```
``` go
func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
```

## Panic y recover no son try y catch

Uno de los errores más comunes para los que se inician en Golang es
pensar en `panic`{.verbatim} y `recover`{.verbatim} como una alternativa
a los bloques `try-catch`{.verbatim} que aparecen en otros lenguajes.
Esto es considerado una mala práctica.

La función `panic`{.verbatim} se debe utilizar solo para indicar estados
en el flujo de una aplicación para los cuales no hay solución efectiva.
Por otro lado `recover`{.verbatim} debería utilizarse para liberar o
destruir recursos adquiridos por la aplicación antes de hacer una salida
forzosa.
