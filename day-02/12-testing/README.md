# Pruebas unitarias

Golang incluye en el paquete [testing](https://pkg.go.dev/testing) las
herramientas necesarias para hacer pruebas unitarias. Por convención las
pruebas unitarias se incluyen en el mismo paquete de la funcionalidad a
probar, adicionando a los archivos el sufijo `_test`{.verbatim}.

Además de pruebas unitarias, Golang nos permite escribir pruebas de
rendimiento o *benchmarks* y ejemplos de prueba.

## Nuestra primera prueba unitaria.

Si tenemos una función para calcular los números de fibonacci.

``` go
func Fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}
```

### Pruebas unitarias

``` go
func TestFibonacci(t *testing.T) {
    var in int = 10
    var want int = 55

    got := Fibonacci(in)
    if got != want {
        t.Errorf("Fibonacci(%d) == %d, want %d", in, got, want)
    }
}
```

### Pruebas de ejemplo

Los ejemplos además de ejecutar pruebas, son una forma de documentar el
uso de funcionalidades.

``` go
func ExampleFibonacci() {
    fmt.Println(Fibonacci(10))
    // Output: 55
}
```

```{=org}
#+REVEAL: split
```
Tanto las pruebas unitarias como los ejemplos se ejecutan utilizando el
comando `go test`{.verbatim}.

## Benchmarks

Un *benchmark* o prueba de rendimiento nos permite examinar el desempeño
de nuestras funcionalidades.

``` go
func BenchmarkFibonacci(b *testing.B) {
        for n := 0; n < b.N; n++ {
                Fibonacci(10)
        }
}
```

## Datos de pruebas o fixtures

Por convención, Golang ignora cualquier directorio que empiece con
`.`{.verbatim}, `_`{.verbatim} o se llame `testdata`{.verbatim}. Dado
que las pruebas siempre se ejecutan en el directorio donde se encuentra
el archivo `_test`{.verbatim} podemos usar la función `Join`{.verbatim}
del paquete `path/filepath`{.verbatim} para acceder a los datos de
prueba.

```{=org}
#+REVEAL: split
```
``` go
func TestBytesInFile(t *testing.T) {
    l, err := BytesInFile(filepath.Join("testdata", "hello_world"))
    if err != nil {
        t.Errorf("Error %s", err)
        t.FailNow()
    }
    if l != 12 { // New line at end of file
        t.Errorf("Got wrong number of bytes %d", l)
    }
}
```
