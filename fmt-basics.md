# FMT Basics

### General 

* %v (value in default format)
* %T (type of value)
* %% (prints an '%' character)

```Go
fmt.printf("%v", 10) // Output: "10"

fmt.printf("%v is %v years old","Dwight" 32) // Output: "Dwight is 32 years old"

```
### Boolean

* %t (true or false)

### Integer

* %b (base 2)
* %o (base 8)
* %d (base 10)
* %x (base 16) N.B. Using a capital 'X' will print out in capitals
```GoLang
fmt.Printf("Number: %b", 1025) // Output: [Binary representation of number]
fmt.Printf("Number: %o", 1025) // Output: [Ocatal representation of number]
fmt.Printf("Number: %d", 1025) // Output: [Decimal representation of number]
fmt.Printf("Number: %x", 1025) // Output: [Hexadecimal representation of number] 

```

### Floating Points

* %e (scientific notation)
* %f / %F (decimal no exponent)
* %g (for large exponents)

``` Go
fmt.Printf("Number: %e", 2.1423145624354234) // Output: "Number: 2.142314e+00"
fmt.Printf("Number: %f", 2.1423145624354234) // Output: "Number: 2.142314"
fmt.Printf("Number: %g", 2.1423145624354234) // Output: "Number: 2.1423145624354234"
```

### Strings

* %s (default)
* %q (double quoted string) N.B. prints string double quoted

```Go
fmt.Printf("Name: %s", "Jim") // Output: Name: Jim
fmt.Printf("Name: %s", "Jim") // Output: Name: "Jim"
```

### Width & Precision

* %f / %d / %g / %s / %q (default width, default precision)
* %9f / %d / %g / %s / %q (width 9, default precision)
* %.2f / %d / %g / %s / %q (default width, precision 2)
* %9.2f / %d / %g / %s / %q (width 9, precision 2)
* %9.f / %d / %g / %s / %q (width 9, precision 0)

### Padding

N.B. Padding increases the length of the string to the size of the padding value
* %09d / %f / %g / %s (pads digit to length 9 with preceeding 0's)
* %-4d / %f / %g / %s / %q (Pads with spaces (width 4, left justified))

``` Go
fmt.Printf("Name: %20s", "Michael Scott")
// Output: Name:        Michael Scott
// "Name:        Michael Scott" is now 20 characters long
fmt.Printf("Name: %-20s is the best boss", "Michael Scott") 
// Output: Name: Michael Scott        is the best boss
fmt.Printf("Name: %9s", "Jim")
// Output: Name:      Jim
fmt.Printf("Name: %-9s is a receptionist", "Pam") 
// Output: Name: Pam       is a receptionist


fmt.Printf("Number: %9f", 124.6745)
// Output: Number: 124.124.674500
fmt.Printf("Number: %9.2f", 124.6745)
// Output: Number:    124.67
fmt.Printf("Number: %.3f", 145.26252) 
// Output: Number: 145.263
fmt.Printf("Number: %9.3f", 145.26252) 
// Output: Number:   145.263

// Adding a `0` to `%09.2f` will fill the empty space with 0s
fmt.Printf("Number: %09.3f", 145.26252) 
// Output: Number: 00145.263
```

### Functions

* fmt.Sprintf()
``` Go
var out string = fmt.Sprintf("Number: %07d", 45)
fmt.Println(out)
// Output: Number: 0000045
```