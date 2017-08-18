# passalt

Like `werkzeug.security`'s `generate_password_hash` and `check_password_hash`,
but in Go.  The format of generated hashes is:

```
sha512$salt$hash
```

## Install

```bash
$ go get github.com/anqurvanillapy/passalt
```

## Usage

- There are only two APIs, similar to those in `werkzeug.security`
  + `passalt.New(passwd)`
  + `passalt.Check(hash, passwd)`

```go
import "github.com/anqurvanillapy/passalt"

func main() {
	// Generate the hash that is safe to store in your database.
	hash := passalt.New("foo")
	passalt.Check(hash, "foo")  // => true
	passalt.Check(hash, "bar")  // => false
}
```

## License

MIT
