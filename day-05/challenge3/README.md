# Challenge 3

-   Convertir las siguientes funciones para que utilicen *gorutines* y
    *channels*.
-   Crear dos *benchmarks* para comparar la velocidad de ejecucion entre
    las versiones con y sin uso de *gorutines*.

``` go
func Squares(number int) (sum int) {
    sum = 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit
        number /= 10
    }
    return
}

func Cubes(number int) (sum int) {
    sum = 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit * digit
        number /= 10
    }
    return
}

func SquaresPlusCubes(number int) int {
    return Squares(number) + Cubes(number)
}
```
