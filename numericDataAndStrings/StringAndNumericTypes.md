### Booleans  
* Can be declared with `var x bool`  
* Booleans can be set by testing a conditional `var x := 1 == 1`  
* Default 0 value for booleans is `false`  

### Integers
* 0 value for numeric data is `0`  
* Declaring an integer without a specific integer type will result in a varying integer type (depending on the system)  
* Two integers of different types can't be combined  
    * If two integers of different types need to be combined, change the type of one variable. For example,`fmt.Println(y + int(z))`  
* Different types of integers can be declared by declaring a variable with `var x int` and replacing `int` with a specific integer type  
#### Types of integers  
* int8: -128 - 127  
* int16: -32,768 - 32,767  
* int32 (most common type of integer): -2,147,483,648 - 2,147,483,647  
* int64: -9,223,372,036,854,775,808 - 9,223,372,036,854,775,807  
#### Types of unsigned integers  
* uint8: 0 - 255  
* uint16: 0 - 65,535  
* uint32: 0 - 4,294,967,295  
#### Integer Operations  
* Addition: `a + b`  
* Subtraction: `a - b`  
* Multiplication: `a * b`  
* Division: `a / b`  
    * All division results must be an integer. Thus, dividing 11 by 2 will result in 5, not 5.5
* Remainder `a % b`  

### Floating Point Numbers  
* Floating point numbers allow for decimals  
* Float32 can allow for very small numbers to very large numbers
* Float64 can allow for even more precision as it allows for more digits past the decimal point and a larger numbers as well  
* Floating point numbers allow for the same operations as integers, except for the remainder operator and bit shifting operators  


### Bit Operators  
##### Example binary values  
For these examples, 3 and 10 are used. The binary representation of 3 is 0011 while the binary representation of 10 is 1010  
##### Types of bit operators  
* `&` (and) - Check what bits are set in the first number **and** second number  
    * For example, `fmt.Println(3 & 10)` will return 0010 (2)
* `|` (or) - Checks which bits are set in **either** of the numbers  
    * For example, `fmt.Println(3 & 10)` will return 1011 (11)  
* `^` (exclusive or) - Checks which bits are set in one of the numbers, but not both  
    * For example, `fmt.Println(3 ^ 10)` will return 1001 (9)  
* `&^` (and not) - Checks which bits are not set in neither of the numbers  
    * For example, `fmt.Println(3 &^ 11)` will return 0100 (8)  

### Bit shifting  
Bit shifting changes the exponent of a number  
If `x := 8`, `x == 2^3` is true since that's also equal to 8
Bit shifting, can change the exponent  
For example, `fmt.Println(x << 4)` will return 128 since it's adding 4 to the exponent. Thus, the value of 2^7 is printed (128)  
Bit shifting can also decrease the value  
For example, `fmt.Println(x >> 4)` will return 0.5 since it's equivalent to 2^-1  

### Strings  
* Strings in Go consist of any UTF-8 characters  
* Often, `fmt`, a Go package, is used to output data such as strings  
* Strings characters can be retrived like an array item. For example, `fmt.Println(x[1])` will return the 2nd character (as a byte) of `x`  
    * To output the byte as a string, `fmt.Println(string(x[1]))` could be used instead  
* An example of a string literal is `x := "Hello, world!"`  
* An example of outputting a formatting string is `fmt.Printf("%v, %T\n", x, x)` which would output the value and type of `x`. It would also output a newline  
* String concatenation (combining strings) can be accomplished by using a plus sign `+`  
    * For example, if `x := "h"` and `y := "i"`, `fmt.Println(x+y)` would output "hi"   
    * By default, using a plus for string concatenation doesn't add a space  
    * Using a comma for string concatenation adds a space by default. For example, `fmt.Printf("one plus one equals", "two")` would return "one plus one equals two"  
* Strings can be converted to byte slices using `y := []byte(x)`  
    * This code declares `y` as a byte slice of `x` (assuming `x` is a string)  
    * `y` can be converted back to a string with `string(y)`  

#### One more thing  
Often, the default data type is fine. For example, not specifying an integer type will often be fine.  