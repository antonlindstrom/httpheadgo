# HTTP HEAD Checker

I've started to learn Go, this is what I've come up with.

What this does is:

  * Load a JSON file and parse
  * Doing a HEAD request against an URL
  * Print the results
  * If it fails, it can boot an image in Docker

I guess this isn't useful until a way to register the app in a proxy is
implemented.

To be able to load files, clone this repo to a directory like:

    $GOPATH/src/github.com/antonlindstrom/httpheadgo

Building:

    make

Example:

    $ ./bin/httphead 
    Checking Example (http://example.com)
    Status: HTTP/1.1 200
    OK - Example
