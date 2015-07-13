# Si se ejecuta `salir.go` usando `go run`, 
# entonces la salida será tomada por `go`
# e impresa.
$ go run salir.go
exit status 3

# Al construir y ejecutar el binario se puede
# ver el estatus en la terminal.
$ go build salir.go
$ ./salir
$ echo $?
3

# Nótese que el `!` de nuestro programa jamás
# fue impreso.
