Seg
===

Segment matcher for paths.

Description
-----------

Seg provides two methods for consuming and capturing path segments.
A path is a string that starts with a slash and contains segments
separated by slashes, for example `/foo/bar/baz` or `/users/42`.

Usage
-----

Consider the following example:

```go
s := seg.New("/posts/42")

s.Prev()
// ""
s.Curr()
// "/posts/42"

s.Consume("posts")
// true

s.Prev()
// "/posts"

s.Curr()
// "/42"

s.Consume("42")
// true

s.Prev() 
// "/posts/42"
s.Curr()
// ""
```

The previous example shows how to walk the path by providing segments
to consume. In the following example, we'll see what happens when
we try to consume a segment with a string that doesn't match:

```go
s := seg.New("/posts/42")

s.Prev()
// ""
s.Curr()
// "/posts/42"

s.Consume("admin")
// false

s.Prev()
// ""

s.Curr()
// "/users/42"
```

As you can see, the command fails and the `Prev()` and `Curr()`
strings are not altered. Now we'll see how to capture segment values:

```ruby
s := seg.New("/users/42")

captures := make(map[string]string)

s.Prev()
// ""

s.Curr()
// "/users/42"

s.Capture("foo", captures)
// true

s.Prev()
// "/users"

s.Curr()
// "/42"

s.Capture("bar", captures)
// true

s.Prev()
// "/users/42"

s.Curr()
// ""

captures
// map[foo:users bar:42]
```

Installation
------------

Install it using the "go get" command:

    go get github.com/soveran/seg.go
