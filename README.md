# Knitja

A utility for converting Ninja build files to Knitfiles. It uses a modified [samurai](https://github.com/michaelforney/samurai) under the hood to parse and convert the ninja file. This is simply a Go wrapper around samurai.

Build:

```
go build
```

To convert, just run `knitja` instead of `ninja`, and it will print out a corresponding Knitfile.

This tool is experimental and in-progress!
