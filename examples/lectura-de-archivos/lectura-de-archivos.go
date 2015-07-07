// Leer y escribir archivos son tareas básicas para
// muchos programas en Go. Primero vamos a ver algunos
// ejemplos de como leer archivos.

package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

// Al leer archivos revisamos si hubo error en la llamada.
// Este función auxiliar nos ayudará en el código que sigue.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // Quizá la tarea más básica de lectura de archivos
    // es consumir todo el contenido en memoria.
    dat, err := ioutil.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

    // Muchas veces se necesita tener más control sobre
    // cómo y qué partes del archivo se leen. Para ello
    // hay que empezar abriendo el archivo con la llamada
    // a `Open` y obteniendo un archivo `os.File`.
    f, err := os.Open("/tmp/dat")
    check(err)

    // Leemos algunos bytes del inicio del archivo.
    // Dejamos que se leean hasta 5, pero también
    // revisamos cuantos fueron leídos.
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))

    // También se puede buscar con `Seek` para saber
    // la ubicación en el archivo y leer con `Read`
    // a partir de ahí.
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

    // El paquete `io` tiene funciones que pueden ser
    // utiles para leer archivos. Por ejemplo, una
    // lectura como la anterior puede ser implementada
    // más robustamente con `ReadAtLeast`
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // No hay un rebobinado built-in, pero se puede
    // lograr con `Seek(0,0)`
    _, err = f.Seek(0, 0)
    check(err)

    // El paquete `bufio` implementa un lector con buffer
    // que puede ser útil tanto por su eficiencia con
    // muchas lecturas pequeñas como por los métodos
    // adicionales de lectura que ofrece.
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    // Hay que cerrar el archivo cuando hayamos terminado
    // (usualmente esto se indica justo después de abrir
    // el archivo usando `defer`)
    f.Close()
}
