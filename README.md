# varibles and types

> **Go** is a statically type language. All varibles have assingned types
you can set types explicitly or implicitly

## Datatypes :

- ### bool 
  - true
  - false

- ### string 
  - collection of characters

- ### integers :
  - uint8 
  - uint16
  - uint32
  - uint64
  - int8
  - int16 
  - int32 
  - int64

- ### aliases : 
  - byte
  - uint
  - int
  - uintptr

- ### float 
   - float32
   - float64

- ### complex 
  - complex64
  - complex128

- ### Data collections :
  - arrays
  - slices
  - maps
  - structs
- ### language organization :
  - functions
  - interfaces
  - channels
 - ### Data management :
   - pointers
> You can also create your own datatype.

- ### Math Operators
  - += sum

  - -+ difference

  - *= product

  - /+ quotient

  - % remainder

  - & bitwise AND

  - | bitwise OR

  - ^ bitwise XOR

  - &^ bit clear

  - << left shift

  - \>\> right shift

- ### Memory management

- Go runtime is statically linked to application

- memory is allocated and deallocated automatically

- use make() or new() to initialize maps, slices and channels
 
 - ####  new()

    - allocates memory but does not initialize memory

    - results in zeroed storage but returns a memory address

  - #### make()

    - allocates and initializes memory

    - allocates non-zeroed storage and returns a memory address
