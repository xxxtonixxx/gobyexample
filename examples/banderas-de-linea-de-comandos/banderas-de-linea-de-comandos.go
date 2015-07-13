// Las banderas en la línea de comandos son una forma
// común de especificar opciones para los programas.
// Por ejemplo en `wc -l` el `-l` es una bandera
// de línea de comando.

package main

// Go provee un paquete para banderas llamado `flag`
// que provee análisis básico de banderas de línea de
// comandos. Usaremos este paquete para implementar
// nuestro ejemplo.
import "flag"
import "fmt"

func main() {

    // Existen declaraciones básicas de banderas para los
    // tipos string, integer y boolean. Aquí declaramos
    // un string con la bandera `word` y un valor default
    // `"foo"` y una descripción corta. Esta función
    // `flag.String` regresa un apuntador a string (no
    // un valor de string); veremos como usar este apuntador
    // más adelante.
    wordPtr := flag.String("word", "foo", "a string")

    // Esto declara las banderas `numb` y `fork` usando
    // un mecanísmo similar al de la bandera `word`
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    // También es posible declarar una opción que utilice
    // una variable existente declarada en algún otro lado
    // en el programa.
    // Hay que notar que necesitamos pasarle a la función
    // el apuntador a la bandera.
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // Una vez que todas las banderas han sido declaradas,
    // llamamos a `flag.Parse()` para ejeuctar el análisis.
    flag.Parse()

    // Aquí simplemente imprimimos las opciones analizadas
    // y cualquier argumento posicional restante. Hay
    // que notar que debemos desreferenciar los apuntadores
    // ejem: `*wordPtr` para obtener los valores.
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
