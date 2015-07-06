# Para experimentar con las banderas es mejor si 
# primero compilamos el programa y ejecutamos el 
# binario resultante directamente.
$ go build command-line-flags.go

# Probemos primero el programa pasandole valores para
# todas las banderas.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Si se omiten banderas tomarán automáticamente sus 
# valores por default.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# Los argumentos restantes pueden ser proporcionados
# después de cualquier bandera.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# El paquete `flag` necesita que todas las banderas
# aparezcan antes del resto de los argumentos (de otra
# manera las banderas serán interpretadas como argumentos
# posicionales)
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
trailing: [a1 a2 a3 -numb=7]

# Usa las banderas `-h` o `-help`  para obtener 
# automáticamente un texto de ayuda para el programa.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Si se usa una bandera que no fue especificada en con el 
# paquete `flag` el programa terminara con un mensaje
# de error y mostrará el texto de ayuda de nuevo.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

# A continuación veremos las variables de entorno. Otra 
# forma común de parametrizar programas.
