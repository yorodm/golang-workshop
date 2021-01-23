# Packages

La buena organizacion del codigo es fundamental en cualquier proyecto. Organizar codigo en multiples directorios y multiples archivos permite mantener y probar el codigo mas facilmente. Los directorios son llamados `packages` en Golang. Puede haber uno o mas archivos en un `package`. Para importar paquetes se debe indicar el path al directorio que lo contiene.

La siguiente estructura de proyecto es la indicada para utilizar paquetes:

```bash
Project Root directory
- bin
- pkg
- src
```

El directorio `bin` contiene los binarios del proyecto que se crean con el comando `go install`.

El directorio `pkg` contiene las versiones compiladas de los paquetes con extension `.a` (package archives).

El directorio `src` contiene el codigo fuente del proyecto organizado en multiples archivos y directorios.

---
Para ejecutar el ejercicio:

1. Copiar los directorio `app` y `geometry` al workspace.
1. Correr el comando `go install` dentro del directorio `geometry`.
1. Correr el comando `go build` dentro del directorio `app`.
1. Chequear que se genero el ejecutable.
---

## Referencias

* [Golang standard library packages](https://golang.org/pkg/)
* [Everything you need to know about packages](https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc)
