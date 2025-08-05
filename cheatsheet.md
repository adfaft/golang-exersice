
# System
- `go version` , check version
- `go env GOROOT` , check location install

# Project
- `go mod init example/hello` , initialized project
- `go mod --replace [packagename]=[path]`
- `go mod tidy`  , importing all imports
- `go run [filego]` atau `go run .` , run go file
- `go get .` untuk download semua or `go get example.com/theirmodule` untuk satuan
- `go get example.com/theirmodule@x.x.x` untuk versi tertentu (atau downgrade)
- `go list -m -u all` or `go list -m -u example.com/theirmodule`  List all modules dependencies only, along with the latest version

# Workspace untuk Multi folder project
- `go work init [path]` to initialized go workspace and add path, next will run with `go run [path]`
- `go work use [path]` untuk menambahkan folder ke go workspace
- `go run [path]` untuk menjalankan folder

# Build
- `go build .` dan `./[projectname]`
- `go build -gcflags=all="-N -l"` dan `./[projectname]` agar bisa didebug secara local 

# Library

## fmt
formatting string : https://pkg.go.dev/fmt
Default
```
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```
General
```
%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
	(floating-point infinities and NaNs print as ±Inf and NaN)
%T	a Go-syntax representation of the type of the value
%%	a literal percent sign; consumes no value
%t	the word true or false
%d	base 10
%X	base 16, with upper-case letters for A-F
%q	a single-quoted character literal safely escaped with Go syntax.
%f     default width, default precision
%9f    width 9, default precision
%.2f   default width, precision 2
%9.2f  width 9, precision 2
%9.f   width 9, precision 0
```

## govulncheck
- install `go install golang.org/x/vuln/cmd/govulncheck@latest`
- run `govulncheck ./...`
