# HTTP Client

El paquete [http](https://golang.org/pkg/net/http/) contiene una
implementacion de un cliente http que permite emular las acciones que
realiza un web browser.

Para este ejercicio se utiliza el tester HTTP https://httpbin.org
([codigo fuente](https://github.com/postmanlabs/httpbin), [imagen de
Docker para uso
local](https://hub.docker.com/r/kennethreitz/httpbin/)).

## Ejecutando un GET

```go
resp, _ := http.Get("https://httpbin.org/get")
defer resp.Body.Close()

data, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(data))
```

## Ejecutando un POST

```go
payload := "Hello world!"
resp, _ := http.Post("https://httpbin.org/post", "text/plain", strings.NewReader(payload))
defer resp.Body.Close()

data, _ := ioutil.ReadAll(resp.Body)
fmt.Println(string(data))
```

Notar que si obtenemos un error de `http.Get` o `http.Post` debemos
cerrar el body (utilizando `defer`).

## Extrayendo informacion de una respuesta JSON

Dado este JSON:

```json
{
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/1.1"
  }
}
```

Podemos extraer informacion del campo `User-Agent` de la siguiente
manera:

```go
type response struct {
	Headers struct {
		UserAgent string `json:"User-Agent"`
	} `json:"headers"`
}

var resp response
_ := json.Unmarshal([]byte(jsonData), &resp)

fmt.Println(resp.Headers.UserAgent)
```

## HTTP request timeouts

Podemos especificar un tiempo maximo de espera para que un request que
tarda mucho termine automaticamente.

Crear un nuevo objeto request:

```go
req, _ := http.NewRequest(http.MethodGet, url, nil)
```

Crear un `context` con timeout:

```go
ctx, cancel := context.WithTimeout(context.Background(), timeout)
defer cancel()
```

Ejecutar el request pasando por parametro el `context` con el timeout:

```go
resp, _ := http.DefaultClient.Do(req.WithContext(ctx))
defer resp.Body.Close()
```

## Que son los contextos

El paquete [context](https://golang.org/pkg/context/) permite crear
objectos `context` para pasar informacion entre varias funciones:

`context.Background()` crear un contexto vacio.

Los contextos se pueden encadenar:

```go
ctx := context.Background()
ctxT1, cancel := context.WithTimeout(ctx, 3*time.Second)
ctxT2, cancelNew := context.WithTimeout(ctxT1, 2*time.Second)
```
