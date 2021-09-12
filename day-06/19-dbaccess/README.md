# Acceso a bases de datos

## Bases de datos SQL {#database-sql}

El paquete `database/sql`{.verbatim} contiene interfaces genéricas para
el acceso a bases de datos SQL. El concepto es similar a **ADO.NET** en
.NET Framework o **JDBC** en Java.

### Definiendo la base de datos.

Para el resto de los ejemplo asumamos que tenemos una base de datos con
una tabla única.

``` sql
CREATE TABLE `userinfo` (
       `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
       `username` VARCHAR(64) NULL,
       `departname` VARCHAR(64) NULL,
       `created` DATE NULL
);
```

### Estableciendo conexión

Para acceder a una base de datos usamos la función
`sql.Open`{.verbatim}. La función recibe el nombre del driver y el DSN
(ver documentación del driver)

``` go
import _ "github.com/mattn/go-sqlite3"

// En una función
db, err := sql.Open("sqlite3", "./data.db")
if err != nil {
    panic("Error accediendo a la base de datos")
}
defer db.Close() // Cerrar la base de datos siempre
```

### Consultas

La función `db.Query`{.verbatim} nos permite consultar la base de datos.
Las llamadas a `db.Query`{.verbatim} retornan una estructura
`db.Rows`{.verbatim} que nos permite iterar sobre las filas recibidas e
inspeccionar su valor.

``` go
// Consultar
rows, err := db.Query("SELECT * FROM userinfo")
var uid int
var username string
var department string
var created time.Time

// Verdadero si existen más filas
for rows.Next() {
    // Tomar los valores de la fila
    err = rows.Scan(&uid, &username, &department, &created)
    checkErr(err)
    fmt.Println(uid)
    fmt.Println(username)
    fmt.Println(department)
    fmt.Println(created)
}
rows.Close() // Close libera recursos del iterador
```

### Manipulación de datos

Para ejecutar consultas que manipulen datos debemos crear una estructura
`sql.Stmt`{.verbatim}.

``` go
// Insertar
stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
res, err := stmt.Exec("pparker", "Avengers", "2021-01-01")
// Actualizar
stmt, err = db.Prepare("update userinfo set username=? where uid=?")
res, err = stmt.Exec("spiderman", id)
// Eliminar
stmt, err = db.Prepare("delete from userinfo where uid=?")
res, err = stmt.Exec(id)
// Llamar Close para liberar recursos
err := stmt.Close()
```

## Bases de datos NOSQL: MongoDB {#database-nosql}

La mayoriá de los *drivers* para bases de dato NOSQL de Golang no
implementan las interfaces en `database/sql`{.verbatim} por lo que cada
biblioteca maneja sus propios tipos e interfaces.

### Obteniendo los paqutes necesarios

``` go
// Necesitamos estos imports
import (
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

```

### Definiendo los datos

1.  Las etiquetas `bson`{.verbatim} para serializar y deserializar hacia
    la base de datos.
2.  Las etiquetas `json`{.verbatim} para comunición entre servicios.

``` go
type Item struct {
    ID     int    `json:"id,omitempty" bson:"id,omitempty"`
    Title  string `json:"title,omitempty" bson:"title,omitempty"`
    IsDone bool   `json:"isdone,omitempty" bson:"isdone,omitempty"`
}
```

### Conectar y desconectar.

``` go
type MongoDB struct {
    *mongo.Client
}

func NewMongoDB(ctx context.Context) (*MongoDB, error) {
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        return nil, err
    }
    return &MongoDB{client}, nil
}

func (m *MongoDB) Disconnect(ctx context.Context) {
    defer m.Client.Disconnect(ctx)
}
```

### Crear y actualizar

``` go
func (m *MongoDB) CreateItem(ctx context.Context, newItem Item) string {

    collection := m.Database("todo").Collection("items")
    result, _ := collection.InsertOne(ctx, newItem)

    return result.InsertedID.(primitive.ObjectID).Hex()
}

func (m *MongoDB) UpdateItem(ctx context.Context, item Item) {
    update := bson.M{"$set": bson.M{"title": item.Title, "isdone": item.IsDone}}

    collection := m.Database("todo").Collection("items")
    collection.UpdateOne(ctx, Item{ID: item.ID}, update)
}
```

### Obtener elementos

``` go
func (m *MongoDB) GetItems(ctx context.Context) (items []Item) {
    collection := m.Database("todo").Collection("items")
    cursor, _ := collection.Find(ctx, bson.M{})

    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var oneItem Item
        cursor.Decode(&oneItem)
        items = append(items, oneItem)
    }

    return
}
```

```{=org}
#+REVEAL: split
```
``` go
func (m *MongoDB) GetItem(ctx context.Context, id int) (item Item) {

    collection := m.Database("todo").Collection("items")
    collection.FindOne(ctx, Item{ID: id}).Decode(&item)
    return
}
```

### Eliminar.

``` go
func (m *MongoDB) DeleteItem(ctx context.Context, id int) {
    collection := m.Database("todo").Collection("items")
    collection.DeleteMany(ctx, Item{ID: id})
    return
}
```

## Referencias

1.  [Documentación de
    `database/sql`{.verbatim}](https://pkg.go.dev/database/sql)
2.  [Documentación del driver de
    MongoDB](https://pkg.go.dev/go.mongodb.org/mongo-driver)
