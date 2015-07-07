// Go provee soporte built-in para la
// codificación/decodificación en [base64](https://es.wikipedia.org/wiki/Base64).

package main

// Esta sintáxis importa el paquete `encoding/base64`
// bajo el nombre `b64` en vez del default `base64`.
// Nos ahorrará algo de espacio más adelante.
import b64 "encoding/base64"
import "fmt"

func main() {

    // Aquí esta la cadena que codificaremos/decodificaremos
    data := "abc123!?$*&()'-=@~"

    // Go soporta ambas formas: base64 estándar y compatible
    // con URLs. El codificador necesita un `[]byte` así que
    // hacemos cast a nuestra cadena a ese tipo.
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    // La decodificación puede regresar un error, que puede
    // ser revisada si es que no se sabe si la entrada está
    // bien formada.
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    // Esto codifica/decodifica usando un formato base64
    // compatible con URL.
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
