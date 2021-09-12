# Structs y punteros

## Tipos compuestos, `struct`{.verbatim}

Un `struct`{.verbatim} es un tipo de dato definido por el usuario que
representa una colección de campos.

``` go
type Employee struct {
    firstName string
    lastName string
    age int
}
```

```{=org}
#+REVEAL: split
```
Los campos individuales de un `struct`{.verbatim} se acceden con
`.`{.verbatim}.

``` go
  var emp := Employee {
      firstName: "Something"
  }

  fmt.Println(emp.firstName)

//Output: Something
```

## Campos struct

Golang permite que los campos de un `struct`{.verbatim} puedan ser de
tipo `struct`{.verbatim}.

``` go
type Address struct {
    city, state string
}

type Employee struct {
    firstName, lastName string
    age int
    address Address
}
```

### Inicializando un `struct`{.verbatim}

Para inicializar un campo `struct`{.verbatim} simplemente nos referimos
al mismo en el bloque de inicialización

``` go
  var emp Employee

  emp := Employee{
      firstName: "Peter",
      lastName:  "Parker",
      age:       22,
      address: Address{
          city:  "New York",
          state: "New York",
      },
  }
  fmt.Println(emp)
//Output: {Peter  {New York New York}}
```

```{=org}
#+REVEAL: split
```
Los campos anidados se acceden con múltiples niveles de notación de
punto.

``` go
  fmt.Println(emp.address.city)

//Output: New York
```

## Etiquetas {#tags}

Los campos en las estructuras pueden opcionalmente tener etiquetas. Las
etiquetas funcionan como metadatos y no afectan la forma en que se
representan los datos de la estructura.

Las etiquetas pueden opcionalmente ser pares `llave:"valor"`{.verbatim}.

``` go
type Employee struct {
    firstName string `json:"nombre"`
    lastName  string `json:"apellido"`
    age       int    `json:"edad"`
}
```

Por convención la llave corresponde al nombre del paquete encargado de
procesar la anotación.

## Composición

Un `struct`{.verbatim} puede también \"heredar\" todos los campos de
otro usando composición. Para esto existe una notación especial

``` go
type Address struct {
    city, state string
}

type Employee struct {
    firstName, lastName string
    age int
    Address
}
```

### Inicializar `struct`{.verbatim} compuesto

Al crear una estructura compuesta tenemos que especificar todos los
campos

``` go
  emp := Employee{
      firstName: "Peter",
      lastName:  "Parker",
      age:       22,
      Address: Address{
          city:  "New York",
          state: "New York",
      },
  }
  fmt.Println(emp.city)
//Output: New York
```

## Igualdad

Un `struct`{.verbatim} es un tipo por valor. Dos variables
`struct`{.verbatim} se consideran iguales si todos los valores de sus
campos son iguales. Esto quiere decir que si un `struct`{.verbatim}
tiene campos que no se pueden comparar, como un `map`{.verbatim}, la
operación (`==`{.verbatim}) va a fallar.

## Métodos/Receptores {#methodsreceivers}

Los métodos (también llamados receptores) son funciones asociadas a un
tipo especifico. Son similares al concepto de métodos de clase en el
mundo OOP. La sintaxis es la siguiente:

``` go
func (t Type) MethodName(parameter list) {
    // codigo del metodo
}
```

### Definiendo métodos

Usualmente se define el código del método en el mismo archivo que el
tipo que lo contiene.

``` go
  // Receptor por valor
  func (e Employee) Print() {
      fmt.Println("Employee Record:")
      fmt.Println("Name:", e.firstName, e.lastName)
      fmt.Println("Address:", e.address)
  }
  // en main
  var emp Employee
  emp.Print()
// Outpput  Employee Record:
//          Name: Allen Varghese
//          Address: {AA CO}
```

## Punteros

Golang soporta punteros para actualizar valores pero no admite
aritmética de punteros como en C. `*`{.verbatim} se usa como prefijo
para definir un puntero para un tipo dado. El operador `&`{.verbatim} se
usa para crear punteros a tipos.

### Inicializando punteros

El valor por defecto de los punteros en Golang es `nil`{.verbatim}, este
valor también se utiliza para indicar que un puntero es nulo.

Tener en cuenta que un puntero solo permite recibir punteros de su tipo
y no otros.

``` go
var emp *Employee // puntero nil

emp = &Employee{...} // puntero a Employee
```

### Tipos referencias

Un puntero es una referencia a un tipo, por lo que podemos utilizarlo
para modificar el valor original. Los **receptores por puntero o por
referencia** son una aplicación directa de este concepto.

``` go
  // Receptor por puntero
  func (e *Employee) updateAge(newAge int) {
      e.age = newAge
  }
  // En main
  emp := Employee{
      age: 33,
  }
  fmt.Println("Before:", emp.age)
  emp.updateAge(34)
  fmt.Println("After:", emp.age)

//Output: Before: 33
//        After: 34
```

## Receptores: puntero o valor {#receptores-puntero-valor}

-   Elección del tipo de receptor.
    1.  Usar solo una variante de receptor para un mismo tipo.
    2.  Ante la duda, usar receptores por puntero.
-   Usar receptores por puntero.
    1.  Cuando el método es inmutable.
    2.  Para estructuras que contienen campos que no se deben copiar
        (ej. `sync.Mutex`{.verbatim}).
    3.  Para arreglos o estructuras de gran tamaño.
-   Usar receptores por valor.
    1.  Para tipos `map`{.verbatim}, `func`{.verbatim} o
        `chan`{.verbatim}
    2.  Para tipos básicos como `int`{.verbatim} o `string`{.verbatim}
    3.  Cuando el tipo del receptor no contiene valores mutables
