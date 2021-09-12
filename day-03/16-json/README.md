# Formatos de datos: JSON y XML

## El paquete `encoding`{.verbatim}. {#encoding}

El paquete `encoding`{.verbatim} define interfaces para convertir datos
desde/hacia bytes o representaciones de texto.

-   `ascii85`{.verbatim} Formato utilizado por Adobe PDF y Postcript.
-   `asn1`{.verbatim} Estructuras ASN.1 en formato DER.
-   `base32`{.verbatim} Codificación Base32, RFC 4648.
-   `base64`{.verbatim} Codificación Base64, RFC 4648.
-   `binary`{.verbatim} Traducción entre formatos numéricos y
    codificación a bytes de los mismos.
-   `csv`{.verbatim} Archivos con valores separados por coma (CSV).
-   `gob`{.verbatim} Transmisión de flujos de datos entre emisor y
    receptor ()
-   `hex`{.verbatim} Valores hexadecimales.
-   `json`{.verbatim} JSON como se define en la RFC 7159.
-   `pem`{.verbatim} Formato PEM usado actualmente para representar
    llaves TLS y certificados.
-   `xml`{.verbatim} Parser XML 1.0 con soporte para espacios de
    nombres.

## Serializando a JSON (Marshal). {#serializacion}

`json.Marshal`{.verbatim} codifica tipos Golang en JSON

-   `bool`{.verbatim} como JSON Boolean.
-   valores numéricos como JSON Number.
-   `string`{.verbatim} como JSON String, eliminando etiquetas HTML.
-   arreglos y *slices* como JSON Array.
-   `struct`{.verbatim} como JSON Object.
-   `nil`{.verbatim} como `null`{.verbatim}
-   `chan`{.verbatim}, complex y literales de funciones provocan error.

### Serializando estructuras

Las etiquetas definen como se serializan las estructuras.

``` go
type Address struct {
    // omitempty hace que los campos vacíos no
    // no se serialice.
    City  string `json:"ciudad,omitempty"`
    State string `json:"estado,omitempty"`
    // Campos marcados con - no se serializan
    Zip   strung `json:"-"`
}

type Employee struct {
    FirstName string `json:"nombre"`
    LastName  string `json:"apellido"`
    Age       int    `json:"edad"`
    // no se serializa porque no es exportado
    id        string
    Address
}
```

```{=org}
#+REVEAL: split
```
``` go
emp1 := Employee{"Peter", "Parker", 22, Address{"Manhattan", "New York"}}
d, _ := json.Marshal(&emp1)
fmt.Println(string(d))
```

## Deserializando desde JSON (Unmarshal). {#deserializacion}

`json.Unmarshal`{.verbatim} es la contraparte de
`json.Marshal`{.verbatim}.

``` go
tony := `{"nombre":"Tony","apellido":"Stark","edad":44}`
var emp3 Employee
err = json.Unmarshal([]byte(tony), &emp3)
if err != nil {
    log.Fatal("No se pudo deserializar")
}
```

### Consideraciones.

1.  `Unmarshal`{.verbatim} a un slice elimina los valores del slice.
2.  `Unmarshal`{.verbatim} a un array descarta los valores extra si el
    array destino es muy pequeño.
3.  `Unmarshal`{.verbatim} a un `map`{.verbatim} conserva las llaves
    existentes.
4.  `null`{.verbatim} se traduce a `nil`{.verbatim} o al valor que tenga
    el tipo sin inicializar
