johnny5
=======

A pedagogical attempt at making a "batteries included" version of a Go http server.

Motivation
----------

The default Go http server examples all use `ListenAndServe` which doesn't allow for a graceful shutdown.  

Maybe most people don't need that, but I do.

Things to Try
-------------

`curl ":8080/" ":8080/foo"`

then Ctrl-C the server while it's processing.


