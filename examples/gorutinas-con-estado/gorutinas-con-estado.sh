# Al correr nuestro programa vemos que la versión
# manejando el estado por la gorutina llega a cerca
# de 800,000 operaciones por segundo. 
$ go run gorutinas-con-estado.go
ops: 807434

# En este caso particular usando gorutinas fue un 
# poco más complicado que el anterior usando mutexes.
# Podría ser útil bajo ciertas circunstancias, por 
# ejemplo cuando otros canales están involucrados 
# o cuando el tener múltimple mutexes puede ser 
# suceptible a fallos. Se debe usar el que sea 
# más natural en cada escenario, especialmente 
# teniendo en cuenta la facilidad de entendimiento
# lo correcto de tu programa.
