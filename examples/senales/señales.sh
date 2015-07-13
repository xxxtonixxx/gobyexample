# Cuando corremos este programa se bloqueará 
# esperando una señal. Al escribir `ctrl-C` (que se 
# mostrará en la terminal como `^C`) podemos enviar
# la señal `SIGINT`.
$ go run señales.go
awaiting signal
^C
interrupt
exiting
