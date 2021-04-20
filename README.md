# invalid-http-headers-response

This is a little toy http server that responses with an invalid http header.

You can run this server by running `go run main.go` and try it with your favorite http client of choice, like for example curl:
```
curl localhost:8090/ -v
*   Trying ::1:8090...
* Connected to localhost (::1) port 8090 (#0)
> GET / HTTP/1.1
> Host: localhost:8090
> User-Agent: curl/7.71.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< INVALID HEADAH: invalid
< Date: Tue, 20 Apr 2021 19:25:48 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
```

## But, I want another key-value-pair in my response!
The header can be modified to your liking with passing arguments to the program.
For example: `go run main.go "i n" a` returns `i n: a` in the headers instead of the nice `INVALID HEADAH: invalid`.
