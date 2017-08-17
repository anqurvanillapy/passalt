# passalt

Like `werkzeug.security`'s `generate_password_hash` and `check_password_hash`,
but in Go.

Format of generated hashes looks like this:

```
sha512$salt$hash
```

## Install

```bash
$ go get https://github.com/anqurvanillapy/passalt
```

## Usage

- There are only two APIs, similar to those in `werkzeug`
  + `passalt.Generate(passwd)`
  + `passalt.Check(passwd)`

```go
import "github.com/anqurvanillapy/passalt"

func main() {
	// Generate the hash that is safe to store in your database.
	hash := passalt.Generate("foo")
	passalt.Check("foo")    // => true
	passalt.Check("bar")    // => false
}
```

## License

MIT
