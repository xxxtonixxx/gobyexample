// Los argumentos en la
// [_línea de comandos_](https://es.wikipedia.org/wiki/L%C3%ADnea_de_comandos)
// son una forma de parametrizar la execución de
// los programas. Por ejemplo, `go run hello` usa
// `run` y `hello.go` como argumentos al prgrama go.
package main

import "os"
import "fmt"

func main() {

    // `os.Args` provee acceso a los argumentos
    // en crudo. Hay que notar que el primer valor
    // en este slice es la ruta al programa y que
    // `os.Args[1:]` tiene los argumentos del
    // programa.
    argsConProg := os.Args
    argsSinProg := os.Args[1:]

    // Puedes tener los argumentos con su índice
    // normal.
    arg := os.Args[3]

    fmt.Println(argsConProg)
    fmt.Println(argsSinProg)
    fmt.Println(arg)
}
