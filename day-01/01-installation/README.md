# Primeros pasos {#first-steps}

## Instalación {#install}

1.  Descargar la versión para tu OS desde <https://golang.org/dl>
    (algunas distribuciones de Linux incluyen el paquete en sus
    repositorios).

2.  Abrir el instalador y seguir las
    [instrucciones](https://golang.org/doc/install).

    ``` example
    (Unix| Linux): /usr/local/go (varía según el sistema)
    Windows: C:\Go
    ```

3.  Para chequear que la instalación está correcta escribir en la
    consola

    ``` example
    $ go version
    ```

## Configuración inicial {#config-initial}

1.  Activar variables de entorno para el trabajo con módulos

    ``` example
    $ go env -w GO111MODULE=on GOPROXY=https://goproxy.io,direct
    ```

2.  Para chequear que todas las variables de entorno están correctas

escribir en la consola

``` example
$ go env
GOROOT = [Directorio de instalación]
GOPROXY=https://goproxy.io,direct
GO111MODULE=on
```

## Completamiento en editores {#gpls}

1.  Instalar la herramienta `gopls`{.verbatim}.

``` example
$ go get golang.org/x/tools/gopls@latest
```

## Otras Herramientas

-   [Visual Studio Code](https://code.visualstudio.com/download)
-   [Visual Studio Code Go
    extension](https://code.visualstudio.com/docs/languages/go)
-   [Postman](https://www.getpostman.com)
-   [MongoDB](https://www.mongodb.com)
