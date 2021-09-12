# Hello World

Vamos a programar el tradicional `Hello World`{.verbatim}.

## Pasos {#steps}

1.  Crear un directorio llamado `hello-world`{.verbatim}.
2.  Crear un archivo de texto llamado `main.go`{.verbatim} dentro del
    directorio creado en el paso 1.
3.  Agregar el siguiente código al archivo `main.go`{.verbatim}

``` go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

-   Ejecutar el programa con el comando `go run main.go`{.verbatim}

## ¿Qué sucede exactamente? {#deep-dive}

En esta sección vamos a revisar cada línea del programa
`Hello World`{.verbatim}.

### Packages e imports

Linea 1 `> =package main`{.verbatim}

Los **packages** son la forma de agrupar funcionalidad común en Golang.
El package default es `main`{.verbatim} y usualmente significa que el
archivo que lo contiene es el punto de entrada.

Linea 3 `> =import "fmt"`{.verbatim}

En esta línea especificamos que el package `fmt`{.verbatim} (que se
incluye como parte de Golang), es requerido para ejecutar este programa.

### Funciones

Linea 5 `> =func main() {`{.verbatim}

Las funciones en Golang se definen con el keyword `func`{.verbatim}. En
esta línea estamos definiendo una función llamada `main`{.verbatim} sin
argumentos. El código de la función se escribe entre `{ }`{.verbatim}.

Para crear un ejecutable en Golang tiene que existir una funcion llamada
`main`{.verbatim} dentro de un package también llamado
`main`{.verbatim}. Solo puede haber una funcion `main`{.verbatim},
incluso si tenemos muchos archivos `.go`{.verbatim} (solo puede existir
un entrypoint por ejecutable)

### Bibliotecas

Línea 6 `> =fmt.Println("Hello World")`{.verbatim}

Esta línea imprime `Hello World`{.verbatim} en la consola.
`Println`{.verbatim} es una función dentro del package `fmt`{.verbatim}
que recibe como argumento el string `"Hello World"`{.verbatim}.

La función `Println`{.verbatim} hace dos cosas::

-   Imprime la entrada que el usuario específica en el argumento.
-   Agrega una nueva línea después de imprimir la entrada especificada.
