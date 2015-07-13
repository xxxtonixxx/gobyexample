# Para experimentar con los argumentos en la línea de 
# comandos es mejor construir un binario primero con 
# `go build`.
$ go build argumentos-en-la-linea-de-comandos.go
$ ./argumentos-en-la-linea-de-comandos a b c d
[./argumentos-en-la-linea-de-comandos a b c d]       
[a b c d]
c

# Después veremos procesmiento más avanzado de la línea 
# de comandos con banderas usando el paquete flag
