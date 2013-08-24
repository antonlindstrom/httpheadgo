# HTTP HEAD

I've started to learn Go, this is a learning repo.

What this does is:

  * Doing a HEAD request against an URL
  * Print the results

I guess this isn't useful for more than testing.

Building:

    mkdir bin
    go build -o bin/httphead httphead.go

Example:

    $ ./bin/httphead 
    Host: http://example.com
    Status: HTTP/1.1 200
    http://example.com OK
