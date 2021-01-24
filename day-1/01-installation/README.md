# Instalacion

1. Descargar la version para tu OS desde
   [https://golang.org/dl](https://golang.org/dl).
2. Abrir el instalador y seguir las
   [instrucciones](https://golang.org/doc/install).

    Go quedara instalado por default en:

        Unix: /usr/local/go
        Windows: C:\Go

3. Para chequear que la instalacion esta correctar escribir en la
   consola

		$ go version

4. Activar variables de entorno para el trabajo con módulos

		$ go env -w GO111MODULE=on GOPROXY=https://goproxy.io,direct

4.  Para chequear que todas las variables de entorno estan correctas
    tipear en la consola

        $ go env

        GOROOT = [Directorio de instalacion]
		GOPROXY=https://goproxy.io,direct
		GO111MODULE=on

## Otras Herramientas

- [Visual Studio Code](https://code.visualstudio.com/download)
- [Visual Studio Code official Go
  extension](https://code.visualstudio.com/docs/languages/go)
- [Postman](https://www.getpostman.com)
- [MongoDB](https://www.mongodb.com)
