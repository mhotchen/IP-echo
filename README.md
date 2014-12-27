IP echo
=======

Super simple HTTP service to echo the request's IP address as the HTTP response body.

Supports sitting behind proxies using the `X-Forwarded-For` header.

Usage
-----

Run the `ip-echo` binary file, or you can use go to build it yourself, simply run `go build -o ip-echo`, then run it.

It will reply to all HTTP requests with the origin IP (or the first IP in `X-Forwarded-For` if it's set), as well as print the ip address to stdout for logging, debugging, or whatever you fancy.

Configuring
-----------

#### Port

By default `ip-echo` will listen to port 8080. If you set up an environment variable called `PORT` then that will be used instead.

#### Logging

IP echo simply logs the IP address it found out to standard output. You can redirect standard output to any file or program you like. You can add a timestamp on most unix systems by running the command like so:

```
./ip-echo | while read line; do echo "[$(date --rfc-3339=seconds)] $line"; done
    [2014-12-27 00:01:15+00:00] 192.168.1.64
```

If you don't want to log, then you can redirect to `/dev/null`.

Performance
-----------

This isn't built to be fast, but because it's compiled and incredibly basic, it is fast. Even on a raspberry pi this is able to handle several hundred concurrent connections without issue. If by some miracle you're using this and serving more requests than Google then stick a proxy in front of it and run a few instances.
