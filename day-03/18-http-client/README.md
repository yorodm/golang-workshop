# HTTP Client

El paquete [http](https://golang.org/pkg/net/http/) contiene una
implementacion de un cliente http que permite emular las acciones que
realiza un web browser.

Para este ejercicio se utiliza el tester HTTP <https://httpbin.org>
([codigo fuente](https://github.com/postmanlabs/httpbin), [imagen de
Docker para uso local](https://hub.docker.com/r/kennethreitz/httpbin/)).

## Ejecutando un GET

``` go
resp, _ := http.Get("https://httpbin.org/get")
defer resp.Body.Close()

data, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(data))
```

## Ejecutando un POST

``` go
payload := "Hello world!"
resp, _ := http.Post("https://httpbin.org/post", "text/plain", strings.NewReader(payload))
defer resp.Body.Close()

data, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(data))
```

Notar que si obtenemos un error de `http.Get`{.verbatim} o
`http.Post`{.verbatim} debemos cerrar el body (utilizando
`defer`{.verbatim}).

## HTTP request timeouts

Podemos especificar un tiempo maximo de espera para que un request que
tarda mucho termine automaticamente.

Crear un nuevo objeto request:

``` go
req, _ := http.NewRequest(http.MethodGet, url, nil)
```

Crear un `context`{.verbatim} con timeout:

``` go
ctx, cancel := context.WithTimeout(context.Background(), timeout)
defer cancel()
```

Ejecutar el request pasando por parametro el `context`{.verbatim} con el
timeout:

``` go
resp, _ := http.DefaultClient.Do(req.WithContext(ctx))
defer resp.Body.Close()
```
