## Go con Ejemplos

Contenido, herramientas y servidor web de [Go con Ejemplos][1].

### Información General

El sitio de [Go con Ejemplos][1] se construye analizando el código y los
comentarios de los archivos fuente en el folder `examples` y mostrando esta
información en el sitio usando `templates` (plantillas). Los programas que
realizan este proceso de publicación se encuentran en el directorio `tools`.

El proceso de publicación produce un directorio de archivos estáticos
(`public`) perfecto para ser servido por cualquier servidor HTTP moderno.
Además, se incluye un servidor web Go ligero en `server.go`.

### Compilar

Para compilar el sitio:

```console
$ go get github.com/russross/blackfriday
$ tools/build
```

Para compilar constantemente en un ciclo:

```console
$ tools/build-loop
```

### Deploy Local

```console
$ cd tools
$ go run server.go
```

### License

This work is copyright Mark McGranaghan and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

The Go Gopher is copyright [Renée French](http://reneefrench.blogspot.com/) and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).


### Traducciones

Algunas traducciones hechas por contribuidores:

* [Chino](http://everyx.github.io/gobyexample/) by [everyx](https://github.com/everyx)

### Gracias

Gracias a [Jeremy Ashkenas](https://github.com/jashkenas) por
[Docco](http://jashkenas.github.com/docco/), lo que inspiró este proyecto.

[1]: http://goconejemplos.com
