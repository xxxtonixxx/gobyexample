// A veces nuestros programas en Go necesitan create otro,
// proceso no-Go. Por ejemplo, el resaltador se sintáxis
// de este site esta [implementado](https://github.com/mmcgrana/gobyexample/blob/master/tools/generate.go)
// create un proceso de [`pygmentize`](http://pygments.org/)
// desde un programa en Go. Veamos algunos ejemplos de
// creación de procesos desde Go.
package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

    // Empezaremos con un simple comando que no recibe
    // argumentos y que solo imprime algo en la salida
    // estándar. El auxilar `exec.Command` crea un
    // objeto para representar este proces externo.
    dateCmd := exec.Command("date")

    // `.Output` es otro auxiliar que ayuda en el caso de
    // ejecutar un programa, esperar a que termine y
    // recolectar su salida. Si no hay errores, `dateOut`
    // contendra los bytes con la información solicitada.
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    // A continuación veremos un caso un poco más complejo
    // donde conectaremos los datos al proceso externo
    // usando su entrada estándar y recopilaremos los
    // resultados de su salida estándar.
    grepCmd := exec.Command("grep", "hello")

    // Aquí vamos a tomar explícitamente la salida y la
    // entrada, comenzaremos el proceso, escribiremos
    // un poco en él, leeremos el resultado de la salida
    // y finalmente esperaremos a que termine el proceso.
    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

    // Omitimos el manejo de errores en el ejemplo anterior
    // pero podrías usar el patrón usual `if err != nil`.
    // También hemos leído solamente la salida estándar
    // pero podríamos haber leído el error estándar
    // exactamente de la misma forma.
    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    // Cuando creamos un nuevo proceso necesitamos
    // proveer un comand explicitamente delineado y
    // un arreglo de arguments, en vez de pasar todo
    // en una solo string. Si necesitas crear un
    // proceso con un solo string, puedes usar
    // la opción `-c ` de `bash`
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}
