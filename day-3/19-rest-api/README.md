# Creando APIs REST

## Servidor HTTP.

El paquete `net/http` contiene implementaciones de servidor y cliente
para este protocolo. Crear un programa que responda a peticiones HTTP
es una tarea sencilla.

```go
func main() {
	http.HandleFunc("/hello", func(w http.ResponseWWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Las estructuras `Server` y `ServeMux` son las encargadas de operar el
servidor y el ciclo de petición y respuesta. La variable
`DefaultServeMux` es utilizada por el servidor por defecto

## Creando nuestro propio ServeMux

Usar `DefaultServeMux` tiene sentido para proyectos pequeños, pero lo
usual es que utilicemos nuestra propia instancia de `ServeMux`

```go
mux := http.NewServeMux()
mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world")
})
log.Fatal(http.ListenAndServe(":8080", mux))
```


## APIS REST con ServeMux.
