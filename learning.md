# Definition

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

Key point
- concurrency for multicore / network machines
- modular
- compiles quickly
- garbage collection
- run-time reflection
- statically typed - compiled language, feels like dynamically types for interpreted language

# Installing

## Linux
```sh
# remove previous version, and extract current versionn (dont untar directly in /usr/local/go, it will be proken)
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.5.linux-amd64.tar.gz

# export environment variable
export PATH=$PATH:/usr/local/go/bin

# confirm, via check version
go version
```

## Windows
- download msi installer : https://go.dev/doc/install
- check go version `go version`

## Multiple Version
```sh
$ go install golang.org/dl/go1.10.7@latest
$ go1.10.7 download
$ go1.10.7 version
go version go1.10.7 linux/amd64

# check location install
$ go1.10.7 env GOROOT
```


# Tutorial
## Getting Started

- `go mod init example/hello` untuk initialize go module
it will create a file `go mod` dengan isi
```go
module example/hello
go 1.24.5
```
- file `hello.go`
```go
package main

// library string formatting dari go tools
import "fmt"
// library quote
import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```
- `go mod tidy` , import semua go yang terdefinisikan
- `go run .` run

## Create Module

### Part 1 : create a new module

- `mkdir greetings && cd greetings`create directory "greetings"
- `go mod init example.com/greetings`
- `vi Hello.go` dan isi
```go
package greetings
// declare greetings package, to collect related functions
import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

### Part 2 : use module

- `mkdir hello && cd hello`, create directory "hello"
- `go mod init example.com/hello`
- `vi hellog.go`
```go
package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
```
- `go mod edit --replace example.com/greetings=../greetings`
- `go mod tidy`
- `go run .`

### Part 3 : add error
- in module "greetings", add package error and return multivalue as string,error
```go
package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
```
- in package "hello", receive error
```go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}
```

### Part 4 : add array / map
greetings/greetings.go
```go
package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := make(map[string]string)
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with
        // the name.
        messages[name] = message
    }
    return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return one of the message formats selected at random.
    return formats[rand.Intn(len(formats))]
}

```

file `hello\hello.go`
```go
package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}
```

### part 5 : add testing
- Go will run test with filaname `[file]_test.go`
- Go will run test function with prefix `func Test[name](t *testing.T) {}` , where params `t *testing.T` is a package for test reporting and logging
- `go test` or `go test -v`
example `greeting_test.go`
```go
package greetings

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}
```

### Part 6 : build and install

Build
- `go build` in 'hello' directory, it will generate executable `hello`
- run `./hello`

Install
- `go list -f '{{.Target}}'` will return where the installation path
- `go install` in hello directory and run `hello`
- export installation path, if not yet to environment path
```sh
$ export PATH=$PATH:/path/to/your/install/directory
```

## GO Workspace

Digunakan untuk mengelola banyak module dalam multiple directory
- `go work init`
- `go work use ./folder`
- `go run ./folder`
Jika menggunakan `go work`,  maka `go mod --replace` tidak lagi diperlukan karena sudah ditemukan oleh go work. Sehingga line `replace` di `go.mod` bisa dihapus

## Accessing DB (low level)

- example data `albums.sql`
```sql
DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO album
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
```
- create database `recordings` and import sql diatas
- export environment variable
```sh
export DB_USERNAME=root
export DB_PASSWORD=(your password)
```
- create go project `go mod init example/accessdb`
- get mysql library driver `go get github.com/go-sql-driver/mysql`
- create `main.go` dengan package `main` dan func `main()`
  NOTE: hanya package `main` dan func `main` yang akan otomatis dipanggil
```
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USERNAME")
	cfg.Passwd = os.Getenv("DB_PASSWORD")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected.")

	// by artist
	artist := "John Coltrane"
	albums, err := albumByArtist(db, artist)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album by %v : %v \n", artist, albums)

	// by id
	album, err := albumById(db, 4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album by %v : %v \n", artist, album)

	// add a new item
	albId, err := addAlbum(db, Album{
		Title:  "lorem ipsum dolor sit amit",
		Artist: "john doe",
		Price:  49.99,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Add Album, insert with ID %d", albId)

}

func albumByArtist(db *sql.DB, artist string) ([]Album, error) {

	var albums []Album

	rows, err := db.Query("SELECT * from album where artist = ? ", artist)
	if err != nil {
		return nil, fmt.Errorf("albumByArtist : %q: %v ", artist, err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumByArtist : %q: %v ", artist, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumByArtist : %q: %v ", artist, err)
	}

	return albums, nil

}

func albumById(db *sql.DB, id int64) (Album, error) {

	var alb Album

	rows := db.QueryRow("SELECT * from album where id = ?", id)

	if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumById : data not found")
		}
		return alb, fmt.Errorf("albumById : %v: %v", id, err)
	}

	return alb, nil

}

func addAlbum(db *sql.DB, alb Album) (int64, error) {

	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?,  ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil

}

```

## Go API with GIN

- Initialize project `go mod init example/api-via-gin`
- Create a sample data
```go
// filename main.go
package main

// json:"id" adalah return type as jsoin dan key-nya berupa lowercase
type Album struct {
	ID     string  `json:"id"`
	Artist string  `json:"artist"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main(){

}
```
- import gin `go get github.com/gin-gonic/gin` agar bisa autocomplete
- api list data `GET /albums`
```go
func getAlbums(c *gin.Context) {
	c.IndentedJSON(200, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8000")
}
```

Test
```sh
# di terminal
go run .

# di another terminal
curl http://localhost:8000/albums
```
- api create data `POST /albums`
```go
func postAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(200, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8000")
}
```

Test
```sh
# di terminal
go run .

# di another terminal
curl http://localhost:8000/albums \
--request "POST" \
--header "Content-Type: application/json"
--data '{"id": "4", "title": "All is love", "artist": "john doe", "price": 6.99}'
```
- api get by id `GET /albums/:id`
```go
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(200, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)

	router.Run("localhost:8000")
}
```

Test
```sh
# di terminal
go run .

# di another terminal
curl http://localhost:8000/albums/2
```

## Testing  and Fuzz input
- file : `main.go`
```go
package main

import (
    "errors"
    "fmt"
    "unicode/utf8"
)

func main() {
    input := "The quick brown fox jumped over the lazy dog"
    rev, revErr := Reverse(input)
    doubleRev, doubleRevErr := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
    fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}
```

file : `reverse_test.go`
```go
package main

import (
    "testing"
    "unicode/utf8"
)

// untuk fuzz, function diawali Fuzz*
// dan parameter-nya *testing.F
func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
	    // tambahkan ke daftar test
        f.Add(tc) // Use f.Add to provide a seed corpus
    }

	// akan mengirimkan input fuzz default 10s, bisa diganti `--fuzztime {x}s`
    f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := Reverse(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
            return
        }
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
```

- run with `go test -fuzz=Fuzz` atau `go test -fuzz=Fuzz -fuzztime 30s`

## Go Vulnerability Check
- install `go install golang.org/x/vuln/cmd/govulncheck@latest`
- run `govulncheck ./...`
- cara testing, bisa downgrade salah satu library, ex: `go get golang.org/x/text@v0.3.5` , jika versinya kurang dari versi sekarang, akan otomatis di downgrade (bisa di check di `go.mod`)


# IDE via VSCODE

##  Plugin
- Install extension VSCODE GO :
- klik `{} Go` di kanan bawah status bar, untuk install tools lainnya seperti:
	- gopls : go language server
	- gotests : go unit test
	- impl : go generate stubs forimplementing interfaces
	- goplay : go playground
	- dlv : Delve, go debugger
	- staticcheck : go linter

## Debugging

- Debug via local process yang menggunakan port, hanya bisa dilakukan jika dibuild via `go build -gcflags=all="-N -l"` 
- Jika tidak bisa attach debugging ke local process
  `echo 0 | sudo tee /proc/sys/kernel/yama/ptrace_scope` untuk temporarily disable  ptrace_scope
  atau, ganti default value agar permanent `sudo vi /etc/sysctl.d/10-ptrace.conf 


# GO Tour

## Definition
- `package name`, khusus nama `main` adalah default package untuk self-contained program
- `import`, gunakan lastname sebagai package name to be used
- Export, only capitalized function name is automatically exported
- `func name(p int) int {}`  , function format
```go
func func_name(p int, s string) (string, string) {
	return s, ""+p
}

// jika tipe-nya sama, bisa pasang yang terakhir saja
func sum(x , y int) int {
	return x+y
}

// ex: naked return dengan name return values
// dimana return dikasih nama, seakan-akan di-assign variable
// naked return otomatis mengembalikan variable tersebut
// NOT RECOMENDED, karena susah dibaca
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum-x
	return
}
```
- `var x, y, z int` , variable bisa diset di package atau function. Sama seperti argument, jika sama tipe-nya, bisa ditulis versi terakhir saja
	- value dari variable dapat di definisikan dengan `=`, ex: 
	  `var c, python, java =true, false "no!"`
	  `var i, j int = 1, 2`
	- short declaration dengan `:=`, sama seperti `a := 3` sama dengan `var a = 5` tapi dengan implicit type (NOTE tidak boleh digunakan diluar function)
	- zero values, secara default jika implicit, isinya zero value seperti 0 = numeric, false = boolean, "" untuk string 
- `const Pi = 3.14` , constant cannot be declared using short syntax `:=` dan isinya bisa string, boolean atau numeric
	- other syntax 
```go
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 9
)
```

## Basic Type

Basic Type
```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

// note: int, uint, dan uintptr adalah 32bit/64bit tergantung os
// recommend to use int kecuali ubtuh sized / unsigned integer

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

Type Conversion
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// equals
i := 42
f := float64(i)
u := uint(f)

// other example
var x, y int = 3, 4
var f float64 = math.Sqrt(float64(x*x + y*y))
var z uint = uint(f)
fmt.Println(x, y, z)

```

Type inference, jika tipe-nya implicit, maka akan bergantung ke value terdekat
```go
var i int
j := i // j is an in

i := 42           // int
f := 3.142        // float64
f := 5e10         // float64
g := 0.867 + 0.5i // complex128
```

## Flow Control
- `for`, tanpa `()` dan wajib ada `{}`
```go
sum := 0
for i= 0; i<10; i++ {
	sum += i
}
// "init" part dan "post" part are optional
for ; sum<100; {
	sum += sum
}
// for without init/post part, adalah pengganti while (; bisa dihapus juga)
for sum < 100 {
	sum += sum
}
```
- `if` sama seperti  for loop, tanpa `()` dan wajib ada `{}`
```go
if x < 0 {
	return sqrt(-x) + "i"
} else if x > 0.5 {
	return sqrt(-x)
} else {
	return x;
}

// bisa diawali oleh "init" part seperti for
if v := math.Pow(x, n); v < lim {
	return v
}else{
	// walau v di-initialized di "init" part, dapat dibaca di else
	fmt.Printf("%g >= %g\n", v, lim)
}
```
- `switch`, setiap case ada otomatis break
```go
switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
}

switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
}
// switch without conditions, adalah alternative long deep if/else
t := time.Now()
switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
}
```
- `defer` will only return after function surrounding it finish returning
```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
/
counting
done
9
8
7
6
5
4
3
2
1
0
/
```

## Advanced Types
- pointers, `var p *int` hold memory of a variable. Zero value is `nil`
	- `var k *int` create a pointer to memory with value of int
	- `&i` generate pointer from other variable
	- `*p` read the value of p pointer
```go
i, j := 42, 2701

var k *int
p := &i         // point to i
fmt.Println(*p) // read i through the pointer
*p = 21         // set i through the pointer
fmt.Println(i)  // see the new value of i

p = &j         // point to j
*p = *p / 37   // divide j through the pointer
fmt.Println(j) // see the new value of j
```
- `stuct`,  collection of fields
```go
type Vertex struct {
	X int
	Y int
}
var j Vertex = Vertex{1,2}
k := Vertex{1, 2}
fmt.Println(k.Y)

// via pointer
p := &k
(*p).X = 1e9
p.X = 1e9 // boleh seperti ini khusus struct

// other example
type Vertex struct {
	X, Y int
}
// assign literally
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)
```
- `array / slice` via `[n]T`
```go
var a [2]string
a[0] = "Hello"
a[1] = "World"

primes := [6]int{2, 3, 5, 7, 11, 13}

// slices, will be references (change will be applied ot others)
primes := [6]int{2, 3, 5, 7, 11, 13}
var s []int = primes[1:4]  // [low key:high key]

// slice literal without length (first create and slice it without length)
q := []int{2, 3, 5, 7, 11, 13}
s := []struct {
	i int
	b bool
}{
	{2, true},
	{3, false},
	{5, true},
	{7, true},
	{11, false},
	{13, true},
}

// this is equivalent
var a [10]int // 
a[0:10]
a[:10]
a[0:]
a[:]

len(a) // length of a
cap(a) // capacity of a, from the first element
s := []int{2, 3, 5, 7, 11, 13} // len 6, cap 6
s = s[:2] // len 2, cap 6
s = s[1:4] // len 3, cap 5 (ambil 3 item dari key1-key4, total capacity berkurang 1 karena low key dimulai dari 1)

var s []int // default value is `nil` for literal slice, cap0, len0

// create a dynamic capacityy slice with `make`
a := make([]int, 5) // length 5, cap 5 => will default [0,0,0,0,0]
a := make([]int, 5, 10) // length 5, cap 10 => will default [0,0,0,0,0]
a := make([]int, 0, 5) // length 0, cap 5 => will default []

// slice of slices
board := [][]string{
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
}

// The players take turns.
board[0][0] = "X"
board[2][2] = "O"
board[1][2] = "X"
board[1][0] = "O"
board[0][2] = "X"

// function
append(s, 0) // append value 0 to slice s
append(s, 2,3) // append value 2 and 3 to slice s
// use range to "for loop"
for i, v := range pow {
	fmt.Printf("2%d = %d\n", i, v)
}
for i := range pow // index only
for i, _ := range pow // skip value, index only
for _, value := range pow // skip injdex, value only

```
- `map` adalah key + value, default value adalah `nil`
```go
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
m = make(map[string]Vertex)
m["Bell Labs"] = Vertex{
	40.68433, -74.39967,
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

delete(m, key)
// test if key exist, ok will be `false` if not exist
elem, ok := m[key]

```
- `func` untuk function. Dan function juga bisa berupa value dan dikirimkan ke function lain
	- secara default, *parameters* adalah copy by value, bukan reference
	- gunakan pointer jika ingin sebagai reference agar bisa merubah datanya
```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// function sebagai closure, agar variable diluar function bisa dimiliki sendiri
func adder() func(int) int {
	sum := 0 // ini beridiri sendiri baik oleh `pos` maupun `neg`
	return func(x int) int {
		sum += x
		return sum
	}
}
pos, neg := adder(), adder()
for i := 0; i < 10; i++ {
	fmt.Println(
		pos(i),
		neg(-2*i),
	)
}

// sebagai pointer agar bisa mengubah data
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
v := Vertex{3, 4}
Scale(&v, 10)
```

## Methods and Interface
- methods karena tidak ada class, method dibuat lewat `func` yang memiliki special *receiver* argument, berada diantara `func` dan `func_name`. Example below, *receiver* adalah `(v Vertex)`
	- `receiver` tidak harus `struct` , bisa tipe lainnya
	- `receiver` hanya bisa di-declare di package yang sama dengan tipe-nya
	- `receiver` biasanya berupa pointer, agar bisa megubah value receiver tersebut
	- `receiver` method dari sebuah tipe, jangan di-mix. BIasanya semuanya pointer, atau semuanya value. Pointer direkomendasi karena tidak meng-copy value dari large object
```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// contoh bukan struct
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// jika receiver berupa pointer, maka sebagai inference dan bisa mengubah value secara langsung
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// contoh sebagai pointer
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10) // shortcut dari (&v).Scale(10)
	fmt.Println(v.Abs())
}
```
- `interface` sebagai definisi method
	- empty interface dapat digunakan untuk menerima *any* type
```go
type Abser interface {
	Abs() float64
}
type T struct {
	S string
}
// disini type T implements Abser, walau secara implicit
func (t *T) Abs() {
	fmt.Println(t.S)
}

i = &T{"hello"}
i.Abs()

// contoh empty interfaced
var i interface{}
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
i = 42
describe(i)
i = "hello"
describe(i)
s, ok := i.(string) // testing value as string
fmt.Println(s, ok)

// switch between type
switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
}

// common used interface, adalah Stringer (tidak perlu didefinisikan lagi karena built-in)
type Stringer interface {
    String() string
}

// contoh
type Person struct {
	Name string
	Age  int
}
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
a := Person{"Arthur Dent", 42}
z := Person{"Zaphod Beeblebrox", 9001}
fmt.Println(a, z)
```
- `Error`  interface, seperti `Stringer`, `Error` adalah built ini interface yang bisa duigunkaan
```go
// NOTE: tidak perlu di definisikan lagi
type error interface {
    Error() string
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// example biasanya
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}

```
- `Reader` interface, builtin banyak diimplementasikan seperti file, network, compressor,h, ciphers, etc
```go
func (T) Read(b []byte) (n int, err error)

// example
r := strings.NewReader("Hello, Reader!")

b := make([]byte, 8)
for {
	n, err := r.Read(b)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	fmt.Printf("b[:n] = %q\n", b[:n])
	if err == io.EOF {
		break
	}
}
```
- `image` package, built ini untuk image
```go
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}

// example
package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
```

## Generics
- *type parameter* sebagai generic `func Index[T comparable](s []T, x T) int` didalam kurung kotak
```go
// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

si := []int{10, 20, 15, -10}
fmt.Println(Index(si, 15))

// Index also works on a slice of strings
ss := []string{"foo", "bar", "baz"}
fmt.Println(Index(ss, "hello"))

```
- *generic type*, selain sebagai parameter function juga bisa punya type sendiri
```go
// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

```

## Concurrency
- goroutine via `go f(x,y,z)` , dapat menjalankan fungsi di thread goroutine sendiri. Note, masih berada di *same address space* sehingga *access to shared memory must be synchronized*
```go
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

```
- channel via `<-` , adalah teknik untuk *send and receive* dari go routine tanpa blocking secara automatically setelah selesai. 
	- create channel  `ch: make(chan int)` atau buffered `ch: make(chan int, 5)`
	- send value to channel `ch <- v`
	- receive channel `v := <-ch`
	- check apakah channel sudah close `v, ok := <-ch`
	- close channel `close(ch)` can only be set by the sender function
```go
// format
ch := make(chan int) // create channel
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.

// example
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

s := []int{7, 2, 8, -9, 4, 0}

c := make(chan int)
go sum(s[:len(s)/2], c)
go sum(s[len(s)/2:], c)
x, y := <-c, <-c // receive from c

// example via make untuk buffered channel
ch := make(chan int, 2)
ch <- 1
ch <- 2
fmt.Println(<-ch)
fmt.Println(<-ch)

// example closing channel
```
- `select` akan menjalankan multiple goroutine dan menunggu berbaagai operasi hingga selesai
```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

// select dengan default
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return
		default:
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}


```
- `sync.Mutex` , memastikan hanya 1 goroutine yang boleh jalan via `Lock` dan `Unlock`
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

```
# Library

## fmt
String formatting
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
```
Boolean
```
%t	the word true or false
```
Integer
```
%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"
```
Floating Points
```
%b	decimalless scientific notation with exponent a power of two,
	in the manner of strconv.FormatFloat with the 'b' format,
	e.g. -123456p-78
%e	scientific notation, e.g. -1.234456e+78
%E	scientific notation, e.g. -1.234456E+78
%f	decimal point but no exponent, e.g. 123.456
%F	synonym for %f
%g	%e for large exponents, %f otherwise. Precision is discussed below.
%G	%E for large exponents, %F otherwise
%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

The exponent is always a decimal integer.
For formats other than %b the exponent is at least two digits.

%f     default width, default precision
%9f    width 9, default precision
%.2f   default width, precision 2
%9.2f  width 9, precision 2
%9.f   width 9, precision 0
```
String or slice of bytes
```
%s	the uninterpreted bytes of the string or slice
%q	a double-quoted string safely escaped with Go syntax
%x	base 16, lower-case, two characters per byte
%X	base 16, upper-case, two characters per byte
```

# Reading  
- Go Tutorial : https://go.dev/doc/tutorial/
- Go by Example : https://gobyexample.com/
- Golang by Project : https://zerotomastery.io/blog/golang-practice-projects/
- Learn Go 1000+ currated Go Example : https://github.com/inancgumus/learngo
- Native Environment variable / ENV file : https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66/