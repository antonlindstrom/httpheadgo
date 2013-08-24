# HTTP HEAD

I've started to learn Go, this is a learning repo.

What this does is:

  * Load a JSON file and parse
  * Doing a HEAD request against an URL
  * Print the results

I guess this isn't useful for more than testing.

To be able to load files, clone this repo to a directory like:

    $GOPATH/src/github.com/antonlindstrom/httpheadgo

Building:

    make

Example:

    $ ./bin/httphead 
    Host: http://example.com
    Status: HTTP/1.1 200
    http://example.com OK
