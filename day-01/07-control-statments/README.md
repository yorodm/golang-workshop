# Instrucciones de control

Golang soporta multiples instrucciones de control e iteradores.

## If - else if - else

La siguiente es una instrucction de control condicional:

```go
if condition {
    // codigo a ejecutar
}
```

Se pueden tener muchas instrucciones `else if`. El bloque `else` es
opcional:

```go
if condition {
    // codigo a ejecutar
} else if condition {
    // codigo a ejecutar
} else {
    // codigo a ejecutar
}
```

**Nota:** El bloque `else` debe estar en la misma linea que la llave
de cierre (`}`) por que de esta manera Golang automaticamente inserta
el `;`. Mas detalles [aqui](https://golang.org/ref/spec#Semicolons).

La instruccion `if` soporta indicar una instruccion opcional que se
ejecuta antes de la condicion:

```go
if statement; condition {
    // codigo a ejecutar
}
```

Esto es util, por ejemplo. cuando se requiere una inicializacion. El
`;` se usa para separar la instruccion de la condicion.

## for

Hay una sola instruccion de iteracion en Golang, el iterador `for`. No
existen los iteradores `while` o `do .. while` que estan presentes en
otros lenguajes.

La sintaxis es:

```go
for initialization; condition; post {
    // codigo a ejecutar
}
```

La inicializacion se ejecuta solo una vez. Luego de que el iterador es
inicializado, se chequea la condicion. Si la condicion se evalua
verdadera, el codigo dentro de `{ }` se ejecuta seguido de la
instruccion indicada en el `post`. La instruccion `post` se ejecuta
luego de cada iteracion. Luego de que la instruccion `post` se
ejecuta, la condicion se vuele a chequear. Si es verdadera, el
iterador se vuelve a ejecutar, si es falsa el iterador termina.

Los tres componentes `initialization`, `condition` y `post` son
opcionales en Golang.

El siguiente ejemplo imprime numeros del 1 al 10 usando un iterador
`for`:

```go
for i := 1; i <= 10; i = i + 1 {
    fmt.Println(i)
}
```

## switch

La instruccion de control `switch` evalua una expresion y la compara
contra una lista de posibles valores o expresiones que puedan
coincidir. Es una forma abreviada de escribir muchas clausulas `if
else`.

```go
switch expression {
    case expression or value:
        // codigo a ejecutar
    case expression or value:
        // codigo a ejecutar
    default:
        // codigo a ejecutar
}
```

En el codigo anterior la expresion es evaluada y el valor resultante
es comparado contra el valor de cada instruccion `case`. La
comparacion termina cuando se encuentra la primer coincidencia. Si no
hay coincidencias, se ejecuta el codigo en el bloque `default`. Esto
es opcional, pero es una buena forma de soportar casos de borde.
