# Interfaces

En Golang, una interfaz es un conjunto de firmas de metodos. Si un tipo tiene una definicion para esos metodos, se dice que implementa la interfaz. A diferencia de otros lenguajes, *la asociacion de un tipo con una interfaz es implicita*. 

La principal ventaja del uso de interfaces en un lenguaje estatico como Golang es la reutilizacion de codigo.

Un tipo puede implementar mas de una interfaz.

El siguiente ejemplo ilustra el uso de interfaces. Una compania tiene empleados en dos categorias, permanentes y contratados. Las dos categorias de empleados tienen un salario a fin de mes, pero el calculo es diferente para cada tipo de empleado.

Esta situacion se puede modelar en Golang de la siguiente manera:

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
    totalPayout += v.calculateSalary
}
fmt.Println(totalPayout)
```

En el ejemplo anterior se definen dos tipos para cada categoria de empleado, `Permanent` y `Contract`. El metodo `calculateSalary` se define para los dos tipos. El calculo es diferente para cada tipo.

Para calcular el total a pagar para todos los empleados se crea un arreglo de 3 empleados. El salario de cada empleado se calcula en un iterador `for` llamando al mismo metodo `calculateSalary` independientemente del tipo subyacente. Esto es posible por que el tipo del arreglo es la interfaz `SalaryCalculator` que tambien define un metodo `calculateSalary`. Notar que no es necesario explicitar que los tipos `Permanent` y `Contract` implementan `SalaryCalculator`.