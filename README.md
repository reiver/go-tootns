# go-tootns

Package **tootns** provides an implementation of the (Mastodon) **toot** JSON-LD namespace for the ActivityPub and ActivityStreams protocols, as they are used on the Fediverse, for the Go programming language.

The **toot** namespace is the one that uses the URL `http://joinmastodon.org/ns#`

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-tootns

[![GoDoc](https://godoc.org/github.com/reiver/go-tootns?status.svg)](https://godoc.org/github.com/reiver/go-tootns)

## Example

```golang
import (
	"github.com/reiver/go-tootns"
	"github.com/reiver/go-jsonld"
)

// ...

var toot = tootns.Toot{
	Discoverable: opt.Something(true),
	Indexable:    opt.Something(true),
}

// ...

bytes, err := jsonld.Marshal(toot)
```

`jsonld.Marshal()` can be used to mix multiple namespaces.
For example:

```golang
import (
	"github.com/reiver/go-act"
	"github.com/reiver/go-tootns"
	"github.com/reiver/go-jsonld"
)

// ...

var person act.Person{
	// ...
}

var toot = tootns.Toot{
	Discoverable: opt.Something(true),
	Indexable:    opt.Something(true),
}

// ...

bytes, err := jsonld.Marshal(person, toot)
```

## Import

To import package **tootns** use `import` code like the follownig:
```
import "github.com/reiver/go-tootns"
```

## Installation

To install package **tootns** do the following:
```
GOPROXY=direct go get github.com/reiver/go-tootns
```

## Author

Package **tootns** was written by [Charles Iliya Krempeaux](http://reiver.link)
