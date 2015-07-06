// En el ejemplo previo vimos como
// [crear un nuevo proceso](craando-procesos). Hacemos
// esto cuando necesitamos que el proceso sea accesible
// al proceso en Go que se está ejecutando. A veces
// necesitamos reemplazar completamente el proceso
// actual en Go con otro proceso (quizá no-Go). Para
// ello usaremo la implementación en Go de la clásica
// función <a href="http://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>.
package main

import "syscall"
import "os"
import "os/exec"

func main() {

    // Para nuestro ejemplo usaremos `ls`. Go necesita
    // la ruta absoluta al binario que queremos ejecutar,
    // así que usaremos `exec.LookPath` para encontrarla
    // (probablemente en `/bin/ls`).
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    // `Exec` necesita los argumentos en forma de slice
    // (en vez de un gran string). Le daremos a `ls`
    // algunos argumentos comunes. El primer argumento
    // debe de ser el nombre del programa.
    args := []string{"ls", "-a", "-l", "-h"}

    // `Exec` también necesita un conjunto de [variables de entorno](variables-de-entorno).
    // Aquí solamente le pasaremos nuestro ambiente
    // actual.
    env := os.Environ()

    // Aquí está la llamada a `os.Exec`. Si la llamada
    // es exitósa, la ejecución de nuestro proceso
    // terminará y será reemplazada por el proceso
    // `/bin/ls -a -l -h`. Si existe algún error,
    // recibiremos un valor de retorn.
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
