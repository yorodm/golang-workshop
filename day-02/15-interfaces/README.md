# Interfaces


## Interfaces en Golang

En Golang, una interfaz es un conjunto de firmas de metodos. Si un
tipo tiene una definicion para esos metodos, se dice que implementa la
interfaz. A diferencia de otros lenguajes, *la asociacion de un tipo
con una interfaz es implicita*.

Es convención en Golang hacer interfaces que contengan pocos métodos.

Un tipo puede implementar mas de una interfaz.

```go
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

En el ejemplo anterior se definen dos tipos para cada categoria de
empleado, `Permanent` y `Contract`. El metodo `calculateSalary` se
define para los dos tipos. El calculo es diferente para cada tipo.

El uso de la interfaz `SalaryCalculator` permite guardar en el arreglo
estructuras de tipo `Permanent` y `Contract` indistintamente. Cda
elemento del arreglo invocará al método `calculateSalary` del tipo
subyacente.

## Interfaces anidadas.

Las interfaces al igual que las estructuras pueden anidarse.

```
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

Cualquier tipo que implemente la interfaz `ReadSeeker` tiene que
implementar `Read` y `Seeker` a la vez.


## Interfaces y composición.

De la misma forma que una estructura "hereda" los métodos de otra en
la composición, una estructura implementa las interfaces de aquellas
que la componen

```go

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

En el ejemplo anterior el tipo `GuidedDog` implementa la interfaz
`Animal`

## La interfaz vacía

El tipo `interface{}` representa valore de cualquier tipo en Golang
(tecnicamente todos los tipos implementan una interfaz sin métodos)

```go
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

## Type assertions

A veces es necesario saber el tipo concreto detrás de una interfaz.

## Punteros a interfaces
