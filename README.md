## Go con Ejemplos

Contenido, herramientas y el web server de el sitio
[Go con Ejemplos](http://goconejemplos.com).

### Información General

El sitio de Go con Ejemplos se construye extrayendo el
código y los comentarios de los archivos fuente en el
folder `examples` y mostrando esa información en el sitio
usando `templates`. Los programas que implementan
este proceso están en `tools`.

El proceso produce un directorio de archivos estáticos - `public` - perfecto
para ser servido por cualquier servidor HTTP moderno. Aparte,
se incluye un servidor ligero de Go en `server.go`.

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
$ mkdir -p $GOPATH/src/github.com/mmcgrana
$ cd $GOPATH/src/github.com/mmcgrana
$ git clone git@github.com:mmcgrana/gobyexample.git
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
    BUILDPACK_URL=https://github.com/mmcgrana/buildpack-go.git
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

Gracias a [Jeremy Ashkenas](https://github.com/jashkenas)
por [Docco](http://jashkenas.github.com/docco/), lo que
inspiró este proyecto.
