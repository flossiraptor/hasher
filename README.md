# Hasher

Hasher is a golang implementation of Java's `hashCode()` feature.

It provides an implementation for the following primitive types:

- int and int(8|16|32|64)
- uint and uint(8|16|32|64)
- float32 and float64
- string

```go
var i float32 := 1.5
hash := hasher.HashCode(i) // -1090519040
```

It's most useful for comparing structs which have implemented the Hasher
interface.

```go
type Person struct {
  name string
  dob string
  favouriteColor string
}
func (person *Person) HashCode() int32 {
  // Uniquely identify a Person by their name and date of birth.
  // Ignore their favourite color.
  hash := int64(HashCode(person.name))
  hash += int64(HashCode(person.dob))
  return HashCode(hash)
}
```
