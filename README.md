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