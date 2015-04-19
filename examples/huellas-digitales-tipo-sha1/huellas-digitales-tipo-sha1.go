// [_Las huellas digitales tipo SHA-1_](http://es.wikipedia.org/wiki/Secure_Hash_Algorithm#SHA-1)
// (en lo sucesivo, hashes SHA-1),
// son usadas frecuentemente para calcular pequeños identificadores
// para objetos binarios o de texto. Por ejemplo, el [sistema de
// control de versiones Git]((http://git-scm.com/) usa SHA-1s
// para identificar archivos y directorios versionados.
// He aquí cómo calcular hashes SHA-1 en Go.

package main

// Go implementa diferentes funciones hash en varios
// paquetes ubicados en `crypto/*`; para ver la lista completa
// de funciones hash disponibles haga [click aquí](http://golang.org/src/crypto/crypto.go#L23).
import "crypto/sha1"
import "fmt"

func main() {
    c := "esta es una cadena de texto"

    // Los pasos para generar un hash son `sha1.New()`,
    // `sha1.Write(bytes)` y, por último, `sha1.Sum([]byte{})`.
    // Iniciamos creando un nuevo hash.
    h := sha1.New()

    // `Write` espera bytes como argumento. Si tienes una cadena de texto `c`,
    // usa `[]byte(c)` para forzar la conversión a bytes.
    h.Write([]byte(c))

    // `Sum` regresa el hash generado como un [slice](slices) de bytes
    // El argumento de `Sum` puede ser usado para agregar
    // el resultado a un slice de bytes existente; normalmente no es utilizado.
    // El nombre de la función `Sum` viene de _checksum_ (suma de verificación)
    // también conocido como _hash sum_.
    sb := h.Sum(nil)

    // Los valores SHA1 frecuentemente se imprimen en sistema hexadecimal,
    // por ejemplo, en commits de git. Usa `%x` para convertir un hash
    // a una cadena en hexadecimal.
    fmt.Println(c)
    fmt.Printf("%x\n", sb)
}
