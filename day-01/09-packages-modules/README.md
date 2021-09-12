# Paquetes y módulos

## Paquetes

En Golang, los paquetes corresponden a los archivos en un mismo
directorio. Por convención, los paquetes llevan el nombre del directorio
que los contiene. Aunque esto no es requerido es considerado una buena
práctica.

Supongamos que creamos el proyecto
`github.com/ourcompany/superproduct`{.verbatim}

```{=org}
#+REVEAL: split
```
``` example
+-- LICENSE
+-- README.md
+-- config.go  # package superproduct
+-- go.mod
+-- go.sum
+-- client
       +-- lib.go             # package client
       +-- lib_test.go        # package client
+-- cmd
       +-- super-client
           +-- main.go        # package main
       +-- super-server       # package main
           +-- main.go
+-- internal
       +-- auth
           +-- auth.go        # package auth
           +-- auth_test.go   # package auth
+-- server
       +-- lib.go             # package server
```

```{=org}
#+REVEAL: split
```
El paquete `main`{.verbatim} tiene significado especial ya que
representa el punto de entrada de la aplicación. Como se ve en el
ejemplo anterior podemos tener más de una aplicación en un mismo
proyecto Golang (`super-client`{.verbatim} y `super-server`{.verbatim}).

Al menos uno de los archivos en el paquete `main`{.verbatim} debe tener
una función `main`{.verbatim}.

## Estructura de un proyecto

1.  En la raíz se situan los archivos con funcionalidades comunes que se
    desean exportar.
2.  Las funcionalidades específicas van dentro de subpaquetes
    (`client`{.verbatim} y `server`{.verbatim} en el ejemplo). Pueden
    existir varios niveles de subpaquetes.
3.  Las funcionalidades que no son parte de nuestra API van dentro del
    paquete `internal`{.verbatim} (Golang no permite importar
    `internal`{.verbatim} o ninguno de sus subpaquetes desde módulos de
    terceros)
4.  Las aplicaciones ejecutables se ubican en subpaquetes de
    `cmd`{.verbatim}.

## Importando paquetes.

Para importar paquetes utilizamos **la ruta completa del paquete dentro
del módulo**

``` go
package main

import "fmt"
import "github.com/ourcompany/superproduct"
import "github.com/ourcompany/superproduct/client"

func main() {
  fmt.Println(superproduct.Config())
  fmt.Println(client.Hello())
}
```

### Alias con `import`{.verbatim}

Es posible usar alias para los paquetes importados

``` go
import s "github.com/ourcompany/superproduct"
```

### Importando varios paquetes

Cuando existen varios `imports`{.verbatim}, la convención es agruparlos.

``` go
import (
    "fmt"
    s "github.com/ourcompany/superproduct"
    "github.com/ourcompany/superproduct/client"
)
```

### Casos especiales

Existen dos casos especiales de alias para `import`{.verbatim}

1.  El `.`{.verbatim} incorpora todos los elementos del paquete que se
    importa a nuestro namespace.
2.  El `_`{.verbatim} no incorpora ninguno de los elementos del paquete
    importado a nuestro namespace.

## La función init

La función `init`{.verbatim} nos permite ejecutar código de
inicialización para nuestros paquetes. A diferencia de la función
`main`{.verbatim}, puede existir más de una función `init`{.verbatim}
por paquete.

``` go
// tomado de /github.com/lib/pq/conn.go

func init() {
    sql.Register("postgres", &Driver{})
}
```

## Reglas de visibilidad

Las reglas de visibilidad en Golang siguen un patrón sencillo.

1.  Todo elemento con letra inicial mayúscula es exportado.
2.  Todo elemento con letra inicial minúscula no es exportado.

En el ejemplo el tipo `square`{.verbatim} no es exportado la función
`NewSquare`{.verbatim} si.

``` go
package geometry
type square struct {
    a int
    b int
}

func NewSquare(a, b int) *square {
    return &square{a, b}
}
```

```{=org}
#+REVEAL: split
```
Es posible para un tipo no exportado tener campos o métodos que sean
exportados.

``` go
// Area of a square
func (s square) Area() int {
    return s.a * s.b
}
```

En este caso es posible acceder a los elementos exportados aunque no sea
posible declarar explicitamente que se accede al tipo.

``` go
// Inválido porque square no es exportado
//var s *geometry.square = geometry.NewSquare(length, breadth)
s := geometry.NewSquare(length, breadth)
fmt.Println("Area is", s.Area())
```

## Módulos {#paquetes-y-módulos}

A partir de la versión `1.13`{.verbatim}, Golang incluye un sistema
nativo de manejo de dependencias utilizando módulos. En versiones
anteriores el código de nuestros proyectos tenía que ubicarse en
`$GOPATH/src`{.verbatim}. Ese enfoque es ahora considerado obsoleto

Para crear un módulo ejecutamos el siguiente comando:

``` shell
go mod init <nombre del módulo>
```

```{=org}
#+REVEAL: split
```
Por convención el nombre del módulo es la URL del repositorio de control
de versiones que alberga el código.

El sistema de módulos depende de dos archivos.

1.  `go.mod`{.verbatim} que incluye la definición y las dependencias
    directas.
2.  `go.sum`{.verbatim} que incluye las dependencias directas e
    indirectas con versiones exactas y suma de verificación.

Golang incluye el subcomando `mod`{.verbatim} para ejecutar diferentes
tareas relacionadas con módulos. Para más detalles ejecutar

``` shell
go help mod
```

## Manejando dependencias

Existen varias formas de adicionar dependencias a nuestro módulo. La más
simple es importar la dependencia en el código y ejecutar el siguiente
comando.

``` shell
go mod tidy
```

Los entornos de desarrollo modernos y el Go Language Server hacen este
proceso de forma automática

## Referencias

-   [Golang standard library packages](https://golang.org/pkg/)
-   [Using Go Modules](https://blog.golang.org/using-go-modules)
-   [Tutorial: Create a Go
    module](https://golang.org/doc/tutorial/create-module)
-   [Go Modules Reference](https://golang.org/ref/mod)
