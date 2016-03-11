## About

Minimal Golang url shortener service using [MurmurHash](https://en.wikipedia.org/wiki/MurmurHash)

## Why

Hacked together while attending Couchbase Day Toronto. There's a executable that uses Couchbase for persistence.

## Running In Memory Store example

### Start the server

```
$ go run cmd/shortener/main.go

```

## Running Couchbase based example

make sure to change couchbase settings in the `main.go`.

### Start the server

```
$ go run cmd/shortener-couchbase/main.go

```

### Get Short URL

```
$ curl -i -XPOST http://localhost:4040/?longurl=http://gmarik.info
HTTP/1.1 200 OK
Date: Fri, 11 Mar 2016 20:20:07 GMT
Content-Length: 66
Content-Type: text/plain; charset=utf-8

{"shorturl":"http://u.ly/5595bafb","longurl":"http://gmarik.info"}
```

### Get Long URL

```
$ curl -i  http://localhost:4040/?shorturl=http://u.ly/5595bafb
HTTP/1.1 200 OK
Date: Fri, 11 Mar 2016 20:18:12 GMT
Content-Length: 66
Content-Type: text/plain; charset=utf-8

{"shorturl":"http://u.ly/5595bafb","longurl":"http://gmarik.info"}
```
