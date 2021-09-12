# Instrucciones de control

Golang soporta múltiples instrucciones de control e iteradores.

## `if=/=else`{.verbatim} {#if---else-if---else}

Se pueden tener muchas instrucciones `else if`{.verbatim}. El bloque
`else`{.verbatim} es opcional:

``` go
if condition {
    // codigo a ejecutar
} else if condition {
    // codigo a ejecutar
} else {
    // codigo a ejecutar
}
```

> El bloque `else`{.verbatim} debe estar en la misma línea que la llave
> de cierre (`}`{.verbatim}).

### Inicializar la condición.

La instruccion `if`{.verbatim} soporta indicar una instrucción opcional
que se ejecuta antes de la condición:

``` go
if statement; condition {
    // código a ejecutar
}
```

## `for`{.verbatim}

Hay una sola instrucción de iteración en Golang, el iterador
`for`{.verbatim}.

``` go
for initialization; condition; post {
    // codigo a ejecutar
}
```

Los tres componentes `initialization`{.verbatim}, `condition`{.verbatim}
y `post`{.verbatim} son opcionales en Golang.

### `for`{.verbatim} en detalle

1.  La inicialización se ejecuta una sola vez.
2.  La condición se prueba antes de ejecutar el ciclo.
3.  El post se ejecuta después de cada iteración del ciclo.

### Ejemplo de `for`{.verbatim}

El siguiente ejemplo imprime números del 1 al 10 usando un iterador
`for`{.verbatim}:

``` go
for i := 1; i <= 10; i = i + 1 {
    fmt.Println(i)
}
```

## `switch`{.verbatim}

La instrucción de control `switch`{.verbatim} evalúa una expresión y la
compara contra una lista de posibles valores o expresiones que puedan
coincidir. Es una forma abreviada de escribir muchas cláusulas
`if else`{.verbatim}.

``` go
switch expression {
    case expression or value | (, expression or value)*:
        // código a ejecutar
        (break | fallthrough)
    case expression or value | (, expression or value)*:
        // código a ejecutar
    default:
        // código a ejecutar
}
```

### `break`{.verbatim}, `fallthrough`{.verbatim} y `default`{.verbatim}

1.  `break`{.verbatim} detiene el flujo y sale del `switch`{.verbatim}
2.  `fallthrough`{.verbatim} continúa al próximo `case`{.verbatim}.
3.  Si no hay coincidencias, se ejecuta el código en el bloque
    `default`{.verbatim}.
