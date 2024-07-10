# callerr

`package callerr` enables creating application specific stack traces via errors with caller info.

Instead of having to choose between wrapped errors with ambiguous origins:
```
Failed to a: failed to b: failed to c: original error
```
Or stack traces which are often too dense and too long:
```
panic: original error

goroutine 1 [running]:
main.c(...)
	/tmp/sandbox185963507/prog.go:22
main.b(...)
	/tmp/sandbox185963507/prog.go:16
main.a()
	/tmp/sandbox185963507/prog.go:10 +0x25
main.main()
	/tmp/sandbox185963507/prog.go:27 +0x13
```
`package callerr` offers a third choice. Wrapped errors with caller info that format like stack traces:
```
Failed to a: 
[/home/path/to/package/file.go:11] failed to b: 
[/home/path/to/another/package/file.go:17] failed to c: 
[/home/path/to/yet/another/package/example.go:22] original error
```

# API

Two functions for creating errors are provided:
- `New(msg string) error`
- `Format(msg string, args ...any) eror`

They behave just like `errors.New` and `fmt.Errorf()`, but also include caller info.
For example:
`callerr.New("test")` will print as:
`[/home/path/to/package/file.go:42] test`

Wrapping multiple errors by using `%w` with `callerr.Format` prints as:
```
Failed to a: 
[/home/path/to/package/file.go:11] failed to b: 
[/home/path/to/another/package/file.go:17] failed to c: 
[/home/path/to/yet/another/package/example.go:22] original error
```
