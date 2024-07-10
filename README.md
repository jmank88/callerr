# callerr

`package callerr` enables creating application specific stack traces via errors with caller info.

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
