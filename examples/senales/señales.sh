# When we run this program it will block waiting for a
# signal. By typing `ctrl-C` (which the
# terminal shows as `^C`) we can send a `SIGINT` signal,
# causing the program to print `interrupt` and then exit.
# Cuando corremos este programa se bloqueara 
# esperando una señal. Al escribir `ctrl-C` (que se 
# mostrará en la terminal como `^C`) podemos enviar
# la señal `SIGINT`.
$ go run señales.go
awaiting signal
^C
interrupt
exiting
