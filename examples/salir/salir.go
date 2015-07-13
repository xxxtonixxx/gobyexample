// Usamos `os.Exit` para terminar inmediatamente
// con un estatus determinado.
package main

import "fmt"
import "os"

func main() {

    // `defer` _no_ se ejecutará cuando se use `os.Exit`,
    // así que este `fmt.Println` nunca será invocado.
    defer fmt.Println("!")

    // El estatus de salida será 3.
    os.Exit(3)
}

// A diferencia de C por ejemplo, Go no usa un entero
// como valor de retorno de `main` para indicar el
// estatus de salida. Si se requiere salir con un
// estatus diferente de cero se requiere llamar `os.Exit`.
