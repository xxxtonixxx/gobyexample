// Extraer números contenidos en cadenas de texto es una tarea básica
// pero común en muchos programas; he aquí cómo hacerlo en Go.

package main

// El paquete integrado `strconv` ([documentación aquí](http://golang.org/pkg/strconv/)) provee,
// mediante 4 funciones, lo necesario para la extracción de números:
// [`ParseFloat`](http://golang.org/pkg/strconv/#ParseFloat),
// [`ParseInt`](http://golang.org/pkg/strconv/#ParseInt),
// [`ParseUnit`](http://golang.org/pkg/strconv/#ParseUint) y
// [`Atoi`](http://golang.org/pkg/strconv/#Atoi) (que viene de Ascii to Integer, Ascii a Entero).

import "strconv"
import "fmt"

func main() {

    // `ParseFloat`: el segundo argumento, el número `64`, especifica
    // cuántos bits de precision se deben extraer.
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    // `ParseInt`: en este caso, el `0` especifica que se debe
    // inferir la base dependiendo del valor de la cadena.
    //`64` requiere que el resultado quepa en 64 bits.
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    // `ParseInt` reconoce números formateados en hexadecimal.
    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    // `ParseUint` (para enteros sin signo)
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    // `Atoi` es una función para conversión básica
    // de base 10 a `int` (entero);
    // es equivalente a `ParseInt("135", 10, 0)`
    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    // Este tipo de funciones regresan un error
    // si el argumento no es del tipo correcto,
    // pero regresan `0` como primer valor.
    z, e := strconv.Atoi("wat")
    fmt.Println(e)
    fmt.Println(z)
}
