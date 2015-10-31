# sfmt

## Description

**sfmt** stands for string format, this package provided simple utilities to
convert any supported string format -spaced|dashed|doted|underscored|camelcased- to
any supperted string format -spaced|dashed|doted|underscored|camelcased.

All string's characters are first convert to lower case before processing,
so the result string is a lower cased formated string.

## Install

Using `go get`:

```
go get -u github.com/WnP/sfmt
```

## Usage

for all examples don't forget to import the package:

```go
import "github.com/WnP/sfmt"
```

### CamelCased

`CamelCased` return a string where the first character is lower cased, you
could easily convert the first letter to upper case using
[strings.Title](https://golang.org/pkg/strings/#Title).

```go
fmt.Println(sfmt.CamelCased("The_quick_brown_fox_jumps_over_the_lazy_dog"))
// Output: theQuickBrownFoxJumpsOverTheLazyDog
```

### Dashed

```go
fmt.Println(sfmt.Dashed("The quick brown fox jumps over the lazy dog"))
// Output: the-quick-brown-fox-jumps-over-the-lazy-dog
```

### Doted

```go
fmt.Println(sfmt.Doted("The-quick-brown-fox-jumps-over-the-lazy-dog"))
// Output: the.quick.brown.fox.jumps.over.the.lazy.dog
```

### Spaced

```go
fmt.Println(sfmt.Spaced("The.quick.brown.fox.jumps.over.the.lazy.dog"))
// Output: the quick brown fox jumps over the lazy dog
```

### Underscored

```go
fmt.Println(sfmt.Underscored("TheQuickBrownFoxJumpsOverTheLazyDog"))
// Output: the_quick_brown_fox_jumps_over_the_lazy_dog
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
