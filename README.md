ARM
===

Arm is simple, small and zero-dependency collection of handy tools.


Installation
------------

```shell
go get -u github.com/mono83/arm
```


Design principles
-----------------

1. Core package `github.com/mono83/arm` should have no dependencies on external packages
2. Core package `github.com/mono83/arm` should have minimal dependencies on Go's
   standard packages like `io`, `fmt`, `strings`, etc.
3. Internal packages should have name prefixed with `arm`, like `armhttp`, `armsq` to 
   be automatically distinguished from corresponding packages without necessity of
   import aliasing.
4. Internal packages can require additional external dependencies to work.


Core package
------------

| Method     | Description                                                   |
|------------|---------------------------------------------------------------|
| `arm.If`   | Ternary operator implementation                               |
| `arm.Must` | Takes value and error, panic on error otherwise returns value |
| `arm.Or`   | Return first non-default value from given ones                |
| `arm.Ref`  | Constructs reference to given value                           |


Uber.Fx package `armfx`
-----------------------

| Method                | Description                                      |
|-----------------------|--------------------------------------------------|
| `armfx.ProvideStruct` | Constructs provider `fx.Option` for given struct |


Hashing tools package `armhash`
-------------------------------

| Method                | Description                                                                                             |
|-----------------------|---------------------------------------------------------------------------------------------------------|
| `armhash.Bytes`       | Calculates hash/checksum of provided byte slice using given hasher                                      |
| `armhash.CRC32`       | Produces IEEE CRC32 checksum from given reader. Use `armhash.Bytes` or `armhash.String` for simpler API | 
| `armhash.SHA256`      | Produces SHA256 hash from given reader. Use `armhash.Bytes` or `armhash.String` for simpler API         |
| `armhash.String`      | Calculates hash/checksum of provided `string` using given hasher                                        |
| `armbcrypt.NewHasher` | Creates `armhash.Hasher` using BCrypt with given cost                                                   |
| `armbcrypt.Verify`    | Compares a bcrypt hashed password with its possible plaintext equivalent.                               |
| `armbcrypt.IsValid`   | Compares a bcrypt hashed password with its possible plaintext equivalent.                               |


SQL tools package `armsql`
--------------------------

| Method       | Description                                                                               |
|--------------|-------------------------------------------------------------------------------------------|
| `armsql.One` | Returns exactly one element from slice or error if slice has different amount of elements |

String tools package `armstr`

| Method       | Description                                                                               |
|--------------|-------------------------------------------------------------------------------------------|
| `armstr.Len` | Returns count of characters in UTF8 string. Better than `len(s)` but far from being ideal |
