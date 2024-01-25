## Stress Test - Go

This is a simple stress test written in Go. It's a simple tester for load testing a web server.

### Usage - local (without docker)

```bash
# run the app
go run cmd/main.go test --url=http://google.com --requests=1000 --concurrency=10
```

### Usage - local (with docker)

```bash
# build the image
docker build -t stress-test-go .
# run the image
docker run stress-test --url=http://google.com --requests=1000 --concurrency=10
```

### Usage - docker hub

```bash
# run the image
docker run gabrielsc1998/stress-test --url=http://google.com --requests=1000 --concurrency=10
```

### Result

The result will be something like this:

```bash

---------- Results ----------

- Total time: 1s
- Total requests: 100
- Status codes:
 - 200: 50
 - 400: 40
 - 500: 10
- Errors: 0
 - Get "http://google.com": dial tcp

-----------------------------

```