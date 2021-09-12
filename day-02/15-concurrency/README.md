# Concurrencia en Go.

## Gorutinas

Las gorutinas son funciones o métodos que se ejecutan de modo
concurrente. Para los efectos se pueden considerar como hilos o threads
ligeros.

Las gorutinas se crean usando la palabre clave `go`{.verbatim} antes de
la llamada a una función.

``` go
// llamar name como una gorutina
go name()

// gorutina  anónima como
go func() {
// código
}()
```

### Sincronizando gorutinas

Las gorutinas se ejecutan de modo **concurrente**. Para sincronizar
gorutinas es común utilizar elementos de sincronización como
`sync.WaitGroup`{.verbatim}.

``` go
func main() {
    wg := &sync.WaitGroup{}
    wg.Add(4)
    for i:=0; i < 4; i++ {
        go worker(wg)
    }
    // Sin esto es posible que main termine
    // antes de las gorutinas.
    wg.Wait()
}

func worker(w *sync.WaitGroup) {
    defer wg.Done()
    DoWork()
}
```

## Canales

Los canales son un mecanismo para sincronización y comunicación entre
gorutinas. Los usuarios de sistemas basados en Unix pueden pensar en
canales como *pipes*.

Para crear un canal usamos la función `make`{.verbatim}.

``` go
c := make(chan string) // canal sin buffer
cb := make(chan string, 200) // canal con buffer de 200
```

### Tipos de canales

La comunicación vía canales en síncrona. Cuando se envían datos a un
canal la gorutina se bloquea hasta que estos datos sean recibidos. Lo
mismo ocurre si se intenta recibir sin haber datos disponibles.

Los canales con buffer permiten acumular datos enviados y solo bloquean
si

1.  Se han enviado `n`{.verbatim} datos, donde `n`{.verbatim} es el
    tamaño del buffer.
2.  Se intenta recibir desde el canal y el buffer no contiene datos.

### Terminando la comunicación.

Para cerrar un canal se utiliza `close`{.verbatim}. Un canal cerrado no
puede usarse para enviar datos. Recibir datos de un canal cerrado
siempre retornará el *valor cero* del tipo de datos del canal.

``` go
// Enviar
c <- "Hola"
// Recibir
v := <- c
// Recibir mientras el canal esté abierto
for v := range c {
    processValue(v)
}
```

### Multiplexing

La sentencia `select`{.verbatim} nos permite esperar por el resultado de
más de un canal.

``` go
// Esperamos por mensaje o timeout
for {
    select {
    case v := <-input:
        doOperation(v)
    // Canal de tipo <- chan Time
    case <- time.After(time.Duration(80) * time.Millisecond):
        fmt.Println("Timeout!")
        return results
    }
}
```

## Referencias

1.  [Concurrency is Not Parallelism](https://blog.golang.org/waza-talk)
2.  [Gorutines - A Tour of Go](https://tour.golang.org/concurrency/1)
3.  [Channels - A Tour of Go](https://tour.golang.org/concurrency/2)
4.  [Buffered Channels - A Tour of
    Go](https://tour.golang.org/concurrency/3)
