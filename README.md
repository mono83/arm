# arm

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mono83/arm)
![GitHub Release](https://img.shields.io/github/v/release/mono83/arm)

Small, generic utility library for Go 1.18+. The core package has no external dependencies.

## Installation

```shell
go get -u github.com/mono83/arm
```

## Design principles

1. The core package `github.com/mono83/arm` has **no external dependencies**.
2. The core package keeps standard-library imports to a minimum.
3. Sub-packages are prefixed with `arm` (`armhash`, `armsql`, …) so they are visually distinct from standard packages without import aliasing.
4. Sub-packages may pull in external dependencies where appropriate.

---

## Core package — `arm`

```go
import "github.com/mono83/arm"
```

### Control flow

| Function | Signature | Description |
|---|---|---|
| `If` | `If[T any](cond bool, a, b T) T` | Ternary operator |
| `Must` | `Must[T any](t T, err error) T` | Returns `t`; panics if `err != nil` |
| `Try` | `Try(f func() error) error` | Calls `f` in a protected mode, converting panics into errors (with stack trace) |

```go
val := arm.If(x > 0, "positive", "non-positive")

data := arm.Must(os.ReadFile("config.json"))

err := arm.Try(func() error {
    riskyOperation()
    return nil
})
```

### Defaults and fallbacks

| Function | Signature | Description |
|---|---|---|
| `Or` | `Or[T comparable](candidates ...T) T` | Returns the first non-zero value |
| `OrUnref` | `OrUnref[T any](v *T, def T) T` | Dereferences `v`, or returns `def` if `v` is nil |
| `OrProvide` | `OrProvide[T comparable](t T, provide func() T) T` | Returns `t` if non-zero, otherwise calls `provide()` |
| `Ref` | `Ref[T any](t T) *T` | Returns a pointer to the given value |

```go
port := arm.Or(os.Getenv("PORT"), "8080")

timeout := arm.OrUnref(cfg.Timeout, 30*time.Second)

t := arm.OrProvide(req.CreatedAt, time.Now)

ptr := arm.Ref("hello") // *string
```

### Provider combinators — `AllOfProvided`

Call multiple typed provider functions in order. Returns early on the first `nil` provider or the first error. All return values are populated on success.

| Function | Providers |
|---|---|
| `AllOfProvided2[T1, T2]` | 2 |
| `AllOfProvided3[T1, T2, T3]` | 3 |
| `AllOfProvided4[T1, T2, T3, T4]` | 4 |
| `AllOfProvided5[T1, T2, T3, T4, T5]` | 5 |

```go
user, account, err := arm.AllOfProvided2(
    func() (*User, error)    { return db.FindUser(id) },
    func() (*Account, error) { return db.FindAccount(id) },
)
```

### Development stubs — `Todo` / `Todoe`

Placeholder functions for code paths that are not yet implemented. The code compiles, but any execution panics (or returns an error) with a descriptive message that includes the expected type.

| Function | Behaviour |
|---|---|
| `Todo[T any]() T` | Panics with an `ErrTodo` describing the type |
| `Todoe[T any]() (T, error)` | Returns zero value + `ErrTodo` |

```go
func newHandler() http.Handler {
    return arm.Todo[http.Handler]()
    // panics: accessing not implemented value of type "http.Handler"
}
```

---

## Uber Fx integration — `armfx`

```go
import "github.com/mono83/arm/armfx"
```

| Function | Description |
|---|---|
| `ProvideStruct(x any, anno ...fx.Annotation) fx.Option` | Builds an `fx.Provide` constructor for the given struct, injecting each field from the DI container in declaration order |

```go
app := fx.New(
    fx.Provide(newDB, newLogger),
    armfx.ProvideStruct(Service{}),
)
```

---

## Hashing — `armhash`

```go
import "github.com/mono83/arm/armhash"
```

### `Hasher[T]` type

```go
type Hasher[T any] func(io.Reader) (T, error)
```

A function type that reads from an `io.Reader` and produces a hash value. All built-in hash functions satisfy this type, enabling a uniform API across different algorithms and input forms.

### Convenience wrappers

| Function | Description |
|---|---|
| `Bytes[T](hash Hasher[T], b []byte) (T, error)` | Hash a byte slice |
| `String[T](hash Hasher[T], s string) (T, error)` | Hash a string |

### Built-in hashers

| Function | Output | Notes |
|---|---|---|
| `CRC32(r io.Reader) (uint32, error)` | `uint32` | IEEE polynomial; compatible with MySQL, PHP, gzip, zip, PNG |
| `SHA256(r io.Reader) ([]byte, error)` | `[]byte` | Standard SHA-256 |

Both stream data through the hash without buffering the full input into memory.

```go
checksum, err := armhash.String(armhash.CRC32, "hello")

digest, err := armhash.Bytes(armhash.SHA256, fileBytes)
```

### BCrypt — `armhash/armbcrypt`

```go
import "github.com/mono83/arm/armhash/armbcrypt"
```

| Function | Description |
|---|---|
| `NewHasher(cost int) armhash.Hasher[string]` | BCrypt hasher with the given cost (≤ 0 falls back to default cost 10) |
| `NewDefaultHasher() armhash.Hasher[string]` | BCrypt hasher with default cost |
| `Verify(hash string, password []byte) error` | Compares hash against plaintext; returns an error on mismatch |
| `IsValid(hash string, password []byte) bool` | Like `Verify` but returns a bool |

```go
hasher := armbcrypt.NewDefaultHasher()
hash, err := armhash.String(hasher, "s3cr3t")

ok := armbcrypt.IsValid(hash, []byte("s3cr3t")) // true
```

---

## SQL helpers — `armsql`

```go
import "github.com/mono83/arm/armsql"
```

| Function | Description |
|---|---|
| `One[T any](slice []T, err error) (*T, error)` | Asserts exactly one row: propagates an incoming error, returns `sql.ErrNoRows` for an empty slice, and a descriptive error when more than one row is present |

```go
row, err := armsql.One(repo.QueryByFilter(ctx, filter))
// err == sql.ErrNoRows → no match
// err != nil           → too many rows or query failed
// row != nil           → exactly one result
```

---

## String utilities — `armstr`

```go
import "github.com/mono83/arm/armstr"
```

| Function | Description |
|---|---|
| `Len(s string) int` | UTF-8 rune count — more correct than `len(s)`, but does not handle multi-codepoint grapheme clusters such as composite emoji |

For full grapheme-cluster accuracy use [rivo/uniseg](https://github.com/rivo/uniseg).

```go
armstr.Len("héllo") // 5, not 6
```
