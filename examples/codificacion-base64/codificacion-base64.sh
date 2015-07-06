# The string encodes to slightly different values with the
# standard and URL base64 encoders (trailing `+` vs `-`)
# but they both decode to the original string as desired.

# La cadaena codificada es ligeramente diferente cuando
# se usa base64 estándar y base 64 compatible con URLs
# ( los `+` al final vs. `-`)
# pero ambos pueden decodificar a la cadena original 
# como se desea.
$ go run codificacion-base64.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
