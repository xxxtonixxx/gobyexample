# La cadaena codificada es ligeramente diferente cuando
# se usa base64 estándar y base 64 compatible con URLs
# ( los `+` al final en vez de `-`)
# pero ambos pueden decodificar a la cadena original 
# como se desea.
$ go run codificacion-base64.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
