# Creando APIs REST

## Servidor HTTP.

El paquete `net/http`{.verbatim} contiene implementaciones de servidor y
cliente para este protocolo. Crear un programa que responda a peticiones
HTTP es una tarea sencilla.

``` go
func main() {
    http.HandleFunc("/hello", func(w http.ResponseWWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world")
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Las estructuras `Server`{.verbatim} y `ServeMux`{.verbatim} son las
encargadas de operar el servidor y el ciclo de petición y respuesta. La
variable `DefaultServeMux`{.verbatim} es utilizada por el servidor por
defecto

## Creando nuestro propio ServeMux

Usar `DefaultServeMux`{.verbatim} tiene sentido para proyectos pequeños,
pero lo usual es que utilicemos nuestra propio `ServeMux`{.verbatim}

``` go
mux := http.NewServeMux()
mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Hello world")
})
log.Fatal(http.ListenAndServe(":8080", mux))
```

Alternativamente podemos crear nuestro propio `Server`{.verbatim}.

``` go
// Create a server listening on port 8000
s := &http.Server{
    Addr:    ":8000",
    Handler: mux,
}

// Continue to process new requests until an error occurs
log.Fatal(s.ListenAndServe())
```

## APIS REST con ServeMux.

Tanto `Server`{.verbatim} como `ServeMux`{.verbatim} proveen el nivel
mínimo de abstracción necesario para manejar peticiones HTTP.

Tomemos por ejemplo el código para hacer un API REST sencilla

``` go
type H = func(w http.ResponseWriter, r *http.Request)

func dispatchResource(get, post, put, delete H) H {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        switch r.Method {
        case "GET":
            get(w, r)
        case "POST":
            post(w, r)
        case "PUT":
            put(w, r)
        case "DELETE":
            delete(w, r)
        default:
            w.WriteHeader(http.StatusNotFound)
            w.Write([]byte(`{"message": "not found"}`))
        }
    }
}

func main() {
    // initialize mux
    getFunc := func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "get called"}`))
    }
    postFunc := func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte(`{"message": "post called"}`))
    }
    m.HandleFunc("/resource", dispatchResource(getFunc, postFunc, nil, nil))
    // use mux
}
```

## Referencias

-   [Documentación de net/http](https://pkg.go.dev/net/http)
-   [Gorilla Mux](https://github.com/gorilla/mux)
-   [Gin](https://github.com/gin-gonic/gin)
