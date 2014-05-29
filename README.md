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
$ open public/index.html
```

Para compilar constantemente en un ciclo:

```console
$ tools/build-loop
```


### Deploy Local

```bash
$ mkdir -p $GOPATH/src/github.com/dabit
$ cd $GOPATH/src/github.com/dabit
$ git clone https://github.com/dabit/gobyexample.git
$ cd gobyexample
$ go get
$ foreman start
$ foreman open
```

### Deploy en Heroku

Setup Básico:

```bash
$ export DEPLOY=$USER
$ export APP=gobyexample-$USER
$ heroku create $APP -r $DEPLOY
$ heroku config:add -a $APP
    BUILDPACK_URL=https://github.com/dabit/buildpack-go.git
    CANONICAL_HOST=$APP.herokuapp.com \
    FORCE_HTTPS=1 \
    AUTH=go:byexample
$ heroku labs:enable dot-profile-d -a $APP
$ git push $DEPLOY master
$ heroku open -a $APP
```

Agregar un dominio + SSL:

```bash
$ heroku domains:add $DOMAIN
$ heroku addons:add ssl -r $DEPLOY
# order ssl cert for domain
$ cat > /tmp/server.key
$ cat > /tmp/server.crt.orig
$ curl https://knowledge.rapidssl.com/library/VERISIGN/ALL_OTHER/RapidSSL%20Intermediate/RapidSSL_CA_bundle.pem > /tmp/rapidssl_bundle.pem
$ cat /tmp/server.crt.orig /tmp/rapidssl_bundle.pem > /tmp/server.crt
$ heroku certs:add /tmp/server.crt /tmp/server.key -r $DEPLOY
# add ALIAS record from domain to ssl endpoint dns
$ heroku config:add CANONICAL_HOST=$DOMAIN -r $DEPLOY
$ heroku open -r $DEPLOY
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
