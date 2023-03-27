# Golang-practice-codes.
 ## About Go:
 * Go is a statically types language. <br>
 * tall varibles have assingned types. <br>
 * you can set types explicitly or implicitly. <br>

 ## Datatypes:
 * bool - true, false
 * string - collection of characters
 * integers - uint8, uint16, uint32, uint64
            int8, int16, int32, int64
 > aliases can be used for these datatype : byte, uint, int, uintptr
 * float - float32, float64
 * complex (real and imaginary) - complex64, complex128
 
 ## Data collections : 
 * arrays
 * slices 
 * maps
 * structs
 ## language organization : 
 * functions
 * interfaces
 * channels
 ## Data management :   
 * pointers

  ### **You can also create your own datatype.**

 ## Math Operators
* += sum
* -+ difference
* *= product
* /+ quotient
* % remainder
* & bitwise AND
* | bitwise OR
* ^ bitwise XOR
* &^ bit clear
* << left shift
* \>\> right shift 

 ## Memory management
* Go runtime is statically linked to application
* memory is allocated and deallocated automatically
* use make() or new() to initialize maps, slices and channels
    
    * `new()`
    1. allocates memory but does not initialize memory
    2. results in zeroed storage but returns a memory address

    * `make()`
    1. allocates and initializes memory
    2. allocates non-zeroed storage and returns a memory address
