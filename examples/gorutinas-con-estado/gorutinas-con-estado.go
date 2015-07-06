// En el ejemplo anterior usamos un lock explícito
// con mutexes para sincronizar el acceso al estado
// compartido por muchas gorutinas. Otra opción sería
// usar las caractarísticas de sincronización buit-in
// de las gorutinas y de los canales para obtener el mismo
// resultado. Esta forma se alinea con la idea de Go
// de compartir memoria a través de comunicación y que cada
// pieza de dato sea propiedad de exactamente una
// gorutina.

package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

// En este ejemplo nuestro estado le pertenecerá
// a una sola gorutina. Esto garantiza que los
// datos jamás se corromperán por el acceso
// concurrente. Para poder leer o escribir a ese
// estado, otras gorutinas tienen que enviar
// mensajes a la gorutina dueña y recibir las
// respuestas correspondientes. Las estructuras
// `opLeer` y `opEscribir` encapsulan esas
// peticiones y son una forma de que la gorutina
// dueña pueda responder.
type opLeer struct {
    llave int
    resp  chan int
}
type opEscribir struct {
    llave int
    val   int
    resp  chan bool
}

func main() {

    // Como antes, vamos a contar cuantas operaciones
    // se realizan.
    var ops int64 = 0

    // Los canales `lecturas` y `escrituras` serán usados
    // por otras gorutinas para hacer peticiones de
    // lectura y escritura.
    lecturas := make(chan *opLeer)
    escrituras := make(chan *opEscribir)

    // Aquí está la gorutina que es dueña del `estado` que
    // era un mapa en el ejemplo anterior pero ahora es de
    // acceso privado para esta gorutina con estado. Esta
    // gorutina seleccionea repetidamente de los canales
    // `lecturas` y `escrituras`, respondiendo a las peticiones
    // según llegan. Una respuesta es ejecutada primero
    // respondiendo a la operación solicitada y después
    // enviando un valor en el canal de respuesta `resp`
    // para indicar éxito (y el valor deseado en el caso del
    // canal `lecturas`).
    go func() {
        var estado = make(map[int]int)
        for {
            select {
            case leer := <-lecturas:
                leer.resp <- estado[leer.llave]
            case escribir := <-escrituras:
                estado[escribir.llave] = escribir.val
                escribir.resp <- true
            }
        }
    }()

    // Esto inicia 100 gorutinas para solicitar lecturas
    // la gorutina dueña del estado via el canal `lecturas`.
    // Cada lectura require construir una estructura `opLeer`,
    // envairlo por el canal `lecturas` y recibir el
    // resultado en el cadal `resp` provisto.
    for r := 0; r < 100; r++ {
        go func() {
            for {
                leer := &opLeer{
                    llave: rand.Intn(5),
                    resp:  make(chan int)}
                lecturas <- leer
                <-leer.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    // Iniciamos también 10 escrituras usando un
    // mecanismo similar.
    for w := 0; w < 10; w++ {
        go func() {
            for {
                escribir := &opEscribir{
                    llave: rand.Intn(5),
                    val:   rand.Intn(100),
                    resp:  make(chan bool)}
                escrituras <- escribir
                <-escribir.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    // Dejamos que las gorutinas trabajen
    // por un segundo.
    time.Sleep(time.Second)

    // Finalmente capturamos y reportamos el conteo
    // de `ops`
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)
}
