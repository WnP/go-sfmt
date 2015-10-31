# sfmt

## Description

sfmt stands for string format, this package provided simple utilities to
convert any supported string format -spaced|dashed|doted|underscored|camelcased- to
any supperted string format -spaced|dashed|doted|underscored|camelcased.

All string's characters are first convert to lower case before processing,
so the result string is a lower cased formated string, `CamelCased` return
a string where the first character is lower cased, you could easely convert
the first letter to upper case using `strings.Title`.

## Install

Using `go get`:

```
go get -u github.com/WnP/sfmt
```

## Usage

for all examples don't forget to import the package:

```
import sfmt
```

### Spaced

```go
res := smft.Spaced("The.quick.brown.fox.jumps.over.the.lazy.dog")
// res value is "the quick brown fox jumps over the lazy dog"
```

### Dashed

```go
res := smft.Dashed("The quick brown fox jumps over the lazy dog")
// res value is "the-quick-brown-fox-jumps-over-the-lazy-dog"
```

### Doted

```go
res := smft.Dashed("The-quick-brown-fox-jumps-over-the-lazy-dog")
// res value is "the.quick.brown.fox.jumps.over.the.lazy.dog"
```

### Underscored

```go
res := smft.Dashed("TheQuickBrownFoxJumpsOverTheLazyDog")
// res value is "the_quick_brown_fox_jumps_over_the_lazy_dog"
```

### Camelcased

`CamelCased` return a string where the first character is lower cased, you
could easely convert the first letter to upper case using `strings.Title`.

```go
res := smft.Dashed("The_quick_brown_fox_jumps_over_the_lazy_dog")
// res value is "theQuickBrownFoxJumpsOverTheLazyDog"
```

## License (MIT)

MIT License (MIT)

Copyright (c) 2015 Steeve Chailloux

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

## Sources

All sources are hosted at: https://github.com/WnP/go-sfmt

## Contribute

Pull requests and issues are welcome.
