// A veces queremos que nuestros programas en Go
// manejen inteligentemente las [señales de Unix](http://en.wikipedia.org/wiki/Unix_signal).
// Por ejemplo, podríamos apagar el servidor
// limpiamente cuando reciba la señal `SIGTERM`,
// o que un programa de línea de comandos deje
// de procesar la entrada cuando reciba un `SIGINT`.
// Aquí esta como manejar esas señales usando
// canales en Go.
package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

    // La nofiticación de señales en Go funciona
    // enviando valores por el canal `os.Signal`.
    // Vamos a crear un canal que reciba estas
    // notificaciones (también vamos a crear uno que
    // nos notifique cuando el programa puede terminar).
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    // `signal.Notify` registra el canal para
    // recibir notificaciones de las señales dadas.
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    // Esta gorutina executa un bloqueo recibido por
    // el canal sigs. Cuando le llega uno lo imprime
    // y después notifica que el programa puede terminar.
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    // El programa esperará aquí hasta que reciba
    // la señal esperada (indicada por la gorutina
    // anterior y envía un valor de terminado `done`)
    // y después sale.
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
