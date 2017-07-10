http_logger [![GoDoc](http://godoc.org/github.com/mappymappy/http_logger?status.svg)](http://godoc.org/github.com/mappymappy/http_logger)
======
logging HttpRequest and Response by ltsvFormat
this library designed for myframework [ghost](https://github.com/mappymappy/ghost)

## install

```
go get github.com/mappymappy/http_logger
```

## usage

```
n := ghost.CreateEmptyGhost()
n.AddMiddleware(http_logger.Default())
```


## logging-sample

```
2017/07/10 15:56:34 request_uri:/	method:GET	remote:[::1]:52353	user_agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36	pid:54063	body:	tag:start_serve_request
2017/07/10 15:56:34 request_uri:/	method:GET	remote:[::1]:52353	user_agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36	pid:54063	latency:16.271µs	body:	status:404	tag:finish_serve_request
```
