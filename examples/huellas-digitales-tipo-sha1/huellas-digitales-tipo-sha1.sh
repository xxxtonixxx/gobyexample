# Este programa calcula un hash y lo imprime
# en formato hexadecimal legible para los humanos.
$ go run huellas-digitales-tipo-sha1.go
esta es una cadena de texto
110518500fa165c1859df82d3e32c8a127f93c1f

# Puedes calcular otros tipos de hashes siguiendo
# pasos similares a los descritos arriba. Por ejemplo,
# para calcular hashes MD5:
# `import crypto/md5` y después `md5.New()`.

# Si necesitas hashes criptográficamente seguros
# debes investigar cuidadosamente sobre
# [nivel de fortaleza de un hash] (http://es.wikipedia.org/wiki/Funci%C3%B3n_hash_criptogr%C3%A1fica)!
