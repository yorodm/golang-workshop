# Structs y punteros

Un struct es un tipo de dato definido por el usuario que representa
una coleccion de campos (es similar al concepto al struct en C). Se
usa en lugares donde tiene sentido agrupar diferentes tipos de datos.

```go
type Employee struct {
    firstName string
    lastName string
    age int
}
```

En el ejemplo anterior se define un `struct` con el nombre `Employee`
que contiene atributos de un empleado como `first name`, `last name` y
`age` con los respectivos tipos de datos de cada uno.

Los campos individuales de un `struct` se acceden con `.`.

```go
var emp := Employee {
    firstName: "Something"
}

fmt.Println(emp.firstName)

> Something
```

## Campos struct

Golang permite que los campos de un `struct` puedan ser de
tipo`struct`.

```go
type Address struct {
    city, state string
}

type Employee struct {
    firstName, lastName string
    age int
    address Address
}
```

```go
var emp Employee

emp.firstName = "Peter"
emp.address = Address {
    city  : "New York",
    state : "New York",
}

fmt.Println(emp)

> {Peter  {New York New York}}
```

Los campos anidados se acceden con multiples niveles de notacion de
punto.

```go
fmt.Println(emp.address.city)

> AA
```

## Composición

Un `struct` puede tambien "heredar" todos los campos de otro usando
composición. Para esto existe una notación especial

```
type Address struct {
    city, state string
}

type Employee struct {
    firstName, lastName string
    age int
    Address
}
```

Nótese que aquí no se incluye un campo para `Address`, con lo que
indicamos que nuestro tipo `Employee` está compuesto por `Address`

Al crear una estructura compuesta tenemos que especificar todos los
campos

```go
 emp := Employee{
	firstName: "Peter",
	lastName: "Parker"
	age: 22,
	city: "New York",
	state: "New York",
 }
```

Para acceder a los campos de la estructura compuesta utilizamos la
notación usual

```go
```go
fmt.Println(emp.city)

> New York
```

## Visibilidad

Un struct y sus campos son visibles por fuera del paquete que los
contiene dependiendo de la primera letra de su nombre. Si el struct
employee es llamado `employee` en vez de `Employee`, no sera visible
por fuera del paquete donde se declaro. Eso se llama un tipo
**exportable** . La misma condicion se aplica a los campos de un
struct.

## Igualdad

Un struct es un tipo por valor. Dos variables struct se consideran
iguales si todos los valores de sus campos son iguales. Esto quiere
decir que si un struct tiene campos que no se pueden comparar, como un
`map`, la operacion de comparacion (`==`) va a fallar.

## Methods/Receivers

Los metodos son funciones asociadas a un tipo especifico. Son
similares al concepto de metodos de clase en el mundo OOP. La sintaxis
es la siguiente:

```go
func (t Type) MethodName(parameter list) {
    // codigo del metodo
}
```

Usualmente se define el codigo del metodo en el mismo archivo que el
tipo que lo contiene a fin de tener todo el codigo relacionado junto.

En el siguiente ejemplo se agrega el metodo `print` al tipo
`Employee`, definido en los ejemplo anteriores, para imprimir el
contenido del registro mas claramente:

```go
func (e Employee) Print() {
    fmt.Println("Employee Record:")
    fmt.Println("Name:", e.firstName, e.lastName)
    fmt.Println("Address:", e.address)
}

var emp Employee

emp.Print()

> Employee Record:
  Name: Allen Varghese
  Address: {AA CO}
```



## Punteros

Golang soporta punteros para actualizar valores pero no admite
aritmetica de punteros como en C. `*` se usa como prefijo para definir
un puntero para de un tipo dado.

El valor por defecto de los punteros en Golang es `nil`, este valor
también se utiliza para indicar que un puntero es nulo.


```go
func (t *Type) MethodName(parameter list) {
    // body
}
```

Tener en cuenta que un puntero solo permite recibir punteros de su
tipo y no otros. Los punteros se usan normalmente cuando es muy
costoso copiar datos, y es mas facil modificar los valores de las
variables y estructuras originales. Esta mecanica deja bien explicito
cuando se usa una referencia, o un valor.

```go
func (e *Employee) updateAge(newAge int) {
    e.age = newAge
}
emp := Employee{
    age: 33,
}

fmt.Println("Before:", emp.age)
emp.updateAge(34)
fmt.Println("After:", emp.age)

> Before: 33
  After: 34
```

En le ejemplo anterior el campo `age` se modifica en la variable
original. Para indicar este comportamiento notar el `*` antes del tipo
`Employee` en la definicion del metodo `updateAge`.
