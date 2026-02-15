
# `github.com/tep/constr`

## Install

```bash
go get github.com/tep/constr
```

## Example Usage
```go
import (
    "github.com/tep/constr"
)

const (
    ErrBadThing     = constr.Error("a bad thing happened")
    ErrOtherBadness = constr.Error("other badness")
)

...
```

## Description

It's simple really -- `constr.Error` is a drop-in replacement for `errors.New`
providing `error` values you can declare as `const` instead of `var`.  They're
still easily comparable but completely immutable.

That is all.

## Colophon

Well... other than maybe an explanation of the package name (if you're at all
interested).  It's merely a portmanteau of the words `constant` and `string`.

