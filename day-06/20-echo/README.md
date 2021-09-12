# Labstack Echo

## Introducción

**Echo** es un microframework Golang para crear servicios web.

## Rutas.

### Handlers

`echo.Context`{.verbatim} representa el acceso al estado de la solicitud
(ruta, parámetros, *handlers*, etc) y contiene los métodos para generar
las respuesta

``` go
func updateUser (c echo.Context) (err error) {
    // ...
}
```

### Rutas REST

``` go
import "github.com/labstack/echo/v4"

func main(){
    e := echo.New()
    e.POST("/users", createUser)
    e.GET("/users/:id", findUser)
    e.PUT("/users/:id", updateUser)
    e.DELETE("/users/:id", deleteUser)
    e.Any("/",home)
}
```

### Grupos

Establecer opciones similares para varias rutas.

``` go
//Todas las URLs con /v2/*
g := e.Group("/v2")
e.POST("/users", createUserV2)
e.GET("/users/:id", findUserV2)
e.PUT("/users/:id", updateUserV2)
e.DELETE("/users/:id", deleteUserV2)
```

## Middleware

Funciones que se procesan antes de un handler.

``` go
func MyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        //Hacer algo
        return next(c)
    }
}
```

### Incluyendo middleware

``` go
//Antes de ejecutarse el router
e.Pre(MyMiddleware)
//Después de ejecutar el router
e.Use(MyMiddleware)
// A nivel de grupo
admin := e.Group("/admin", MyMiddleware)
//A nivel de ruta
e.GET("/", <Handler>, <Middleware...>)
```

## Obteniendo datos

### Desde el contexto

``` go
func(c echo.Context) error {
    name := c.FormValue("name") // valor de formulario
    printFull := c.QueryParam("full") // valor de query
    return c.String(http.StatusOK, name)
}
```

### Usando `Bind`{.verbatim}

``` go
type User struct {
  Name  string `json:"name" form:"name" query:"name"`
  Email string `json:"email" form:"email" query:"email"`
}
func handle(c echo.Context) (err error) {
    u := new(User)
    if err = c.Bind(u); err != nil {
        return
    }
    // Hacer algo con el usuario
}
```

## Respuestas {#echo-response}

`echo.Context`{.verbatim} es también utilizada para generar respuestas.

``` go
// Retornar una cadena
c.String(http.StatusOK, "Hello, World!")
// Retornar HTML
c.HTML(http.StatusOK, "<p>Hello, World!</p>")
// Retorna JSON, serializa el valor de u
c.JSON(http.StatusOK, u)
// Retorna XML, serializa el valor de u
c.XML(http.StatusOK, u)
// Retorna el contenido del fichero
c.File("<PATH_TO_YOUR_FILE>")
// Retorna el contenido del fichero como flujo de datos
c.Stream(http.StatusOK, "<CONTENT_TYPE>", file)
// Redirige
c.Redirect(http.StatusMovedPermanently, "<URL>")
```
