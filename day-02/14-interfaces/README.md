# Interfaces

## Interfaces en Golang

En Golang, una interfaz es un conjunto de firmas de métodos. Si un tipo
tiene una definición para esos métodos, se dice que implementa la
interfaz. A diferencia de otros lenguajes, *la asociación de un tipo con
una interfaz es implicita*.

Un tipo puede implementar más de una interfaz.

``` go
type Permanent struct {
    empID int
    basicPay int
    pension int
}

type Contract struct {
    empID int
    basicPay int
}

func (p Permanent) calculateSalary() int {
    return p.basicPay + p.pension
}

func (p Contract) calculateSalary() int {
    return p.basicPay
}

type SalaryCalculator interface {
    calculateSalary() int
}

// In main
sc := [...]SalaryCalculator{Permanent{1, 5000, 50}, Contract{3, 7000}}
totalPayout := 0
for _, v := range sc {
    totalPayout += v.calculateSalary()
}
fmt.Println(totalPayout)
```

## Interfaces anidadas.

Las interfaces al igual que las estructuras pueden anidarse.

``` go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}

type ReadSeeker interface {
    Reader
    Seeker
}
```

Para que un tipo implemente la interfaz `ReadSeeker`{.verbatim} tiene
que implementar `Read`{.verbatim} y `Seeker`{.verbatim} a la vez.

## Interfaces y composición.

De la misma forma que una estructura \"hereda\" los métodos de otra en
la composición, implementa las interfaces de aquellas que la componen.

``` go
type Animal interface {
    Name() string
}

type Dog struct{}

func (d *Dog) Name() string {
    return "Dog"
}

func Bark(d *Dog) {
    fmt.Println("Woof!")
}

type GuideDog struct {
    *Dog
}
```

En el ejemplo anterior el tipo `GuidedDog`{.verbatim} implementa la
interfaz `Animal`{.verbatim}

## La interfaz vacía

El tipo `interface{}`{.verbatim} representa un valor comodín al estilo
del tipo `object`{.verbatim} de C# o `void *`{.verbatim} en C
(técnicamente todos los tipos implementan una interfaz sin métodos)

``` go
func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}
```

## Type Assertions. {#type-assertions}

A veces nos es necesario saber de qué tipo es el valor guardado en la
interfaz.

``` go
var i interface{} = "hello"
s := i.(string)
fmt.Println(s)

s, ok := i.(string)
fmt.Println(s, ok)

f, ok := i.(float64)
fmt.Println(f, ok)

f = i.(float64) // panic
fmt.Println(f)

```

La sintaxis `variable.(Tipo)`{.verbatim} es un **type assertion**,
concepto muy parecido al **type casting** de otros lenguajes.

### Type assertions seguras e inseguras

Los type assertions se ejecutan con alguna de las siguientes variantes

``` go
// variante segura
   v, ok := i.(Tipo)
// variante insegura
   v := i.(Tipo)
```

La variante segura retorna un par `(Tipo,bool)`{.verbatim} donde el
segundo valor representa el estado de la operación. Un estado de
`true`{.verbatim} significa que se pudo efectuar la conversión,
`false`{.verbatim} que la conversión no es posible y el primer valor de
la tupla estará con el valor nulo.

En la variante insegura si no se puede efectuar la conversión el
*runtime* de Golang lanzará un `panic`{.verbatim}.

## Type switches

Los **type switches** son una construcción especial que nos permite
determinar el tipo de una variable y actuar en consecuencia

``` go
switch v := v.(type) {
    case string:
        fmt.Printf("%v is a string\n", v)
    case int:
        fmt.Printf("%v is an int\n", v)
    default:
        fmt.Printf("The type of v is unknown\n")
}
```

En lugar de usar la sintaxis `v.(Tipo)`{.verbatim} para la conversión se
utiliza `v.(type)`{.verbatim}.
