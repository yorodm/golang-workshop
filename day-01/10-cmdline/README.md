# Línea de comandos {#linea-de-comandos}

## Argumentos

Interactuar con la línea de comandos es muy útil y muy común al ejecutar
scripts y programas que no tienen GUI. El paquete `os`{.verbatim} se usa
para tener acceso a los argumentos de la línea de comandos al ejecutar
un programa de Golang.

```{=org}
#+REVEAL: split
```
Los argumentos de la línea de comandos están disponibles en el arreglo
`os.Args`{.verbatim}:

``` go
import "os"

os.Args // Lista completa incluyendo el nombre del programa
os.Args[1:] // Solo los argumentos
```

## Flags

Especificar los argumentos como valores separados por espacios en la
terminal es muy básico y difícil de extender. El uso de
`flags`{.verbatim} provee mas flexibilidad en cómo los argumentos se
especifican, sobre todo los opcionales. Golang cuenta con un paquete
llamado `flag`{.verbatim} para soportarlos.

``` go
import "flag"

flag.<type>(nombre del flag, valor default, descripcion)
```

```{=org}
#+REVEAL: split
```
Un flag se define con la sintaxis anterior. El `<type>`{.verbatim} puede
ser `string`{.verbatim}, `int`{.verbatim} o `bool`{.verbatim}.

Un flag opcional se puede indicar controlando si el valor es el mismo
que el valor especificado por defecto.

Un flag se puede leer independientemente de la posición en que el
usuario los hubiera especificado.

## Environment Variables

Las variables de entorno se usan para configurar aspectos de sistema en
la mayoría de los \*NIX y también en Windows. Son pares de clave-valor
usualmente disponibles para todos los programas que corren en la
terminal/shell. Se pueden definir variables de entorno custom solo
disponibles durante la ejecución de un script de shell.

```{=org}
#+REVEAL: split
```
El paquete `os`{.verbatim} proporciona dos funciones,
`Setenv`{.verbatim} y `Getenv`{.verbatim}, para escribir y leer
variables de entorno. Un string vacío se retorna si la clave no existe.

``` go
  import "os"

  os.Setenv("FOO", "1")
  fmt.Println("FOO:", os.Getenv("FOO"))
  fmt.Println("BAR:", os.Getenv("BAR"))

//Output: FOO: 1
//        BAR:
```

```{=org}
#+REVEAL: split
```
La función `Environ`{.verbatim} retorna la lista completa de todas las
variables de entorno existentes.

``` go
import "strings"

for _, e := range os.Environ() {
    pair := strings.Split(e, "=")
    fmt.Println(pair[0])
}
```
