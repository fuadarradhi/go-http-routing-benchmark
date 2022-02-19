Go HTTP Router Benchmark
========================

# Original Source
https://github.com/pkieltyka/go-http-routing-benchmark
https://github.com/julienschmidt/go-http-routing-benchmark


# Usage

```bash
go test -bench=.
```

You can bench specific frameworks only by using a regular expression as the value of the `bench` parameter:
```bash
go test -bench="Martini|Gin|HttpMux"
```
