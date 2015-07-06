# Para probar nuestro filtro de línea, primero creamos
# un archivo con algunas líneas en minúscula.
$ echo 'hola'   > /tmp/lines
$ echo 'filtro' >> /tmp/lines

# Después usamos el filtro de línea para obtener 
# lineas convertidas a mayúsculas.
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
