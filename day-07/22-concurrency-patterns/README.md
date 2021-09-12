# Patrones de concurrencia

## Futures

Las *futures* representan el resultado de un cálculo que se ejecuta de
forma concurrente. Este patrón es nos permite ejecutar una operación
costosa sin que impacte la ejecución del proceso principal.

### Interfaz para Future

Definamos una interfaz que exprese el comportamiento que queremos en un
*future*

``` go
type Value interface{}

type Future interface {
    Get(c context.Context) (Value, error)
}
```

Usamos `context.Context`{.verbatim} para dar la posibilidad de cancelar
la espera.

### Implementando Futures

Una implementación simple para *future* es usar un canal que comunique
el estado de completamiento o error.

``` go
type result struct {
    value Value
    err   error
}

type futureImpl struct {
    result chan *result
}
```

Por conveniencia las estructuras son privadas.

```{=org}
#+REVEAL: split
```
El método `Get`{.verbatim} bloquea la gorutina actual en espera del
resultado del *future* o la señal `Done`{.verbatim} del contexto.

``` go
func (f *futureImpl) Get(c context.Context) (Value, error) {
    select {
        case <-ctx.Done():
            return nil, ctx.Err()
        case result := <-f.result:
            return result.value, result.err
    }
}
```

```{=org}
#+REVEAL: split
```
Finalmente una función para crear nuevas *futures*.

``` go
func NewFuture(f func() (Value, error)) Future {
    fut := &futureImpl {
        result: make(chan *result)
    }
    go func(){
        defer close(fut.result)
        value, err := f()
        f.result <- &result{value, err}
    }()
    return fut
}
```
