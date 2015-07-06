// Muchas veces necesitamos que nuestros programas
// ejecuten operaciones sobre colecciones de datos,
// como seleccionar todos los elementos que satisfagan
// una predicado o mapear todos los elementos
// a una nueva colección con una función custom.

// En algunos lenguajes es idiomático usar estucturas
// de datos y algorítmos [genéricos](https://es.wikipedia.org/wiki/Programaci%C3%B3n_gen%C3%A9rica).
// Go no soporta genéricos; en Go es común tener
// funciones que actuan sobre colecciones de datos si
// tu programa y tipos de datos lo necesitan.

// Aquí hay algunos ejemplo de funciones sobre colecciones
// para slices de `strings`. Puedes usar estos ejemplos para
// construir tus propias funciones. Hay que notar que
// en ocasiones puede ser más claro simplemente usar
// el código que manipula la colección directamente
// en vez de crear y llamar a una función auxiliar.

package main

import "strings"
import "fmt"

// Regresa el primer índice del string `t`, o -1 si no
// se encuentra.
func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

// Regresa `true` si el string `t`esta en el slice.
func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

// Regresa `true` si uno de los strings del slice
// satisface el predicado `f`.
func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

// Regresa `true` si todos los strings dentro del
// slice satisfacen el predicado `f`
func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

// Regresa un nuevo slice que contiene todos los strings
// en el slice que satisfacen el predicado `f`
func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

// Regresa un nuevo slice que contiene el resultado de
// aplicar la funcion `f` a cada cadena en el slice original.
func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func main() {

    // Aquí podemos probar nuestras funciones
    var strs = []string{"durazno", "manzana", "pera",
        "ciruela"}

    fmt.Println(Index(strs, "pera"))

    fmt.Println(Include(strs, "uva"))

    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))

    // Los ejemplos anteriores utilizan todos, funciones
    // anónimas pero también puedes usar funciones nombradas
    // econ el tipo correcto.
    fmt.Println(Map(strs, strings.ToUpper))
}
