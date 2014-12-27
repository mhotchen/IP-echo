IP echo
=======

Super simple HTTP service to echo the request's IP address as the HTTP response body.

Supports sitting behind proxies using the `X-Forwarded-For` header.

It will reply to all HTTP requests with the origin IP (or the first IP in `X-Forwarded-For` if it's set), as well as print the ip address to stdout for logging, debugging, or whatever you fancy. Eg.

```
$ curl localhost:8080
    127.0.0.1
```

It also has IPv6 support:

```
$ curl localhost:8080
    [::1]
```

Usage
-----

Install golang then compile:

```
$ sudo apt-get install golang
$ go build -o ip-echo
```

This creates a binary at `./ip-echo`. To run the service:

```
$ ./ip-echo
```

Configuring
-----------

#### Port

By default `ip-echo` will listen to port 8080. If you set up an environment variable called `PORT` then that will be used instead.

```
$ export PORT=80; ./ip-echo
```

#### Logging

IP echo simply logs the IP address it found out to standard output. You can redirect standard output to any file or program you like. You can add a timestamp on most unix systems by running the command like so:

```
$ ./ip-echo | while read line; do echo "[$(date --rfc-3339=seconds)] $line"; done
    [2014-12-27 00:01:15+00:00] 192.168.1.64
```

If you don't want to log, then you can redirect to `/dev/null`.

#### Daemonizing

If you're using Ubuntu or Debian you can copy `example-init.d-script/ip-echo` in to `/etc/init.d/`.

This will allow you to start, stop, and restart `ip-echo` using service:

```
$ sudo service ip-echo status
    [ ok ] ip-echo service is running.

$ sudo service ip-echo stop
    [ ok ] Stopping ip-echo server: ip-echo.

$ sudo service ip-echo status
    [FAIL] ip-echo service is not running ... failed!

$ sudo service ip-echo start
    [ ok ] Starting ip-echo service: ip-echo.
````

This runs `ip-echo` with the following default parameters:

```
DAEMON=/usr/sbin/ip-echo
USER=ip-echo
PORT=8080
LOG_FILE=/var/log/ip-echo.log
```

You can override these parameters by creating a file located at `/etc/default/ip-echo` and overriding any of the parameters listed in the same syntax and casing.

To configure Ubuntu or Debian to work with these defaults all you need to do is copy the `ip-echo` executable to `/usr/sbin/`:

```
$ sudo cp ip-echo /usr/sbin/
```

And add a user like so:

```
$ sudo useradd ip-echo
```

It's also easy to use `ip-echo` with chroot but that requires modifications to the init.d script.

---

To start `ip-echo` automatically when the server turns on run:

```
sudo update-rc.d ip-echo defaults
```
