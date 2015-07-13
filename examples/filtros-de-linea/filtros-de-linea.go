// Un _filtro de línea_ es un tipo común de programa
// que lee la entrada estándar, la procesa y después
// imprime el resultado a la salida estándar. `grep`
// y `sed` son filtros de línea comunes.

// Aquí hay un ejemplo de un filtro de línea en Go
// que escribe la versión en mayúsculas de todo el
// texto de entrada. Puedes usar este patrón para
// escribir tus propios filtros de línea.
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    // Al envolver el `os.Stdin` que no usa búfer con
    // un scanner con búfer, obtenemos un scanner más
    // conveniente que mueve el scanner al siguiente
    // token; que es la siguiente línea en el scanner
    // por default.
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        // La llamada a `Text` nos regresa el token actual,
        // aquí es la siguiente línea de la entrada.
        ucl := strings.ToUpper(scanner.Text())

        // Escribirmos la línea convertida en mayúsculas.
        fmt.Println(ucl)
    }

    // Revisamos si hubieron errores durante la llamada
    // a `Scan`. El fin de arvhivo es algo esperado y no
    // es reportado por `Scan` como un error.
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
