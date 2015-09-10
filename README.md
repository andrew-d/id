# id

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/andrew-d/id) [![Build Status](https://travis-ci.org/andrew-d/id.svg?branch=master)](https://travis-ci.org/andrew-d/id)

This package implements a simple `ID` type which can be generated from some
input bytes.  The ID has a couple of nice properties:

- When encoded as a string, it uses the [Luhn algorithm][luhn] to add a simple
  checksum to the output string.
- When decoding an ID from a string, the Luhn checksum will be validated to
  ensure validity.
- The encoding doesn't use characters `0`, `1`, or `8`, which minimizes the
  potential for typos.  Decoding from a string will automatically replace those
  characters with the 'correct' values (`O`, `I`, and `B`).
- Comparing two IDs using the `Equals` function will use a constant-time
  comparison algorithm, to prevent timing attacks.


## License

MIT.

Note that this code is mostly a copy of the code for the IDs that are used in
the [syncthing][sync] project.  Portions of code are taken from the
[protocol][proto] project, and the license for that code is included in this
repository.


[luhn]: https://en.wikipedia.org/wiki/Luhn_algorithm
[sync]: https://syncthing.net/
[proto]: https://github.com/syncthing/protocol
