# Cuando corremos nuestro programa este es 
# reemplazado por `ls`.
$ go run reemplazando-procesos-con-exec.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 
reemplazando-procesos-con-exec.go

# Go no ofrece un clásico `fork` de Unix.
# Usualmente esto no causa problems ya que 
# crear gorutinas, lanzar nuevos procesos
# y reemplazar el proceso actual con exec, 
# cubre la mayoría de los usos de `fork`.
