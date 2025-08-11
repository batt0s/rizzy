# Rizzy

![Test workflow](https://github.com/batt0s/rizzy/actions/workflows/test.yml/badge.svg)
![Build workflow](https://github.com/batt0s/rizzy/actions/workflows/build.yml/badge.svg)

## About Rizzy 

I made Rizzy as a learning project. I followed Thorsten Ball's book “Writing An Interpreter In Go” and I made a few changes, mostly changing the name of the keywords and adding some built-in functions for arrays. I just finished the book and I want to do some different things in the future, like adding some data types, some changes in language design. Before I started coding I never thought about the design because I had no idea how to make an interpreter. Now that I understand a little bit, I can think about it and make some changes. I learned a lot and I want to go beyond the book to keep learning. There are still a lot to do and a lot of problems I know about.

Rizzy is language with 6 data types (`Integer`, `Float`, `Boolean`, `String`, `Array`, `Map`) and a `NULL`. Rizzy has first-class functions with closure. Everything in Rizzy is a expression (function) except def (Define) statements and return statements. So if you type `5` it's actually a expression that evaluates 5. 

## To-Do

- [x] Floats
- [ ] Error traceback (added for parse errors, no traceback for Error objects for now)
- [x] Better terminal integration (REPL, use arrow keys for history, persistent history file, tab completion)
- [x] Multiline input
- [x] GTE ("<=") and LTE (">=")
- [x] INTEGER -> FLOAT and FLOAT -> INT
- [x] AND ("&&", "&") and OR ("||", "|")
- [x] Built-in function for formatting ("fmt()")
- [x] Range operator ("[n..n+m]")
- [x] Make runable files
- [ ] Package system



## Examples

With `def` keyword you can define variables and functions. 

```rb
def integer_var = 1;
def float_var = 1.;
def name = "Rizzler";
def arr = ["I", "am", "The", "Rizzler", "!", 1, 2];
def arr_of_nums = [1, 2, 3];
def boolean_variable = true;
```

Define functions and use call expression.

```rb
def factorial = func(x) { if (x == 0) { 1 } else { x * factorial(x - 1) } };
factorial(4);
```

Use index expression.

```rb
def arr = [1, 2, 3];
arr[0];
```

Use hashmap.

```rb
def mymap = {"name": "Rizzler", "version": 1};
mymap["name"];
```

### Built-in Functions

#### `puts` and `rizz`

They have the same definition.
Prints the arguments in their own line and returs an null. So the use is just:

```
>>> puts("The", "Name", "Is", "Rizzler", "!");
The
Name
Is
Rizzler
!
Rizzler: null
```

Notice that it says "null" in the end. It does not evaluate anything. So if you say `def a = puts("a");`, `a;` will give you `null`.

#### `exit`

Exits the interpreter, with an optional exit status code (default=0).


#### `len`

Returns the length of the input as INTEGER. Takes 1 input, ARRAY or STRING.

#### `fmt`

Format string. Use %% and replace with order. Example : 
```
>>> fmt("%% x %% = %%", 2, 2, 2*2)
Rizzler: 2 x 2 = 4
```

#### `first`

Returns first element of an array. Takes an ARRAY as argument. Same as using `array[0]`.


#### `last`

Returns last element of an array. Takes an ARRAY as argument. Same as using `array[len(array)-1]`.

#### `head`

Returns the array without the last element. Takes an ARRAY as argument. 


#### `tail`

Returns the array without the first element. Takes an ARRAY as argument. 


#### `push`

Takes 2 arguments. Takes an ARRAY as first argument and an Expression as second. Returns an ARRAY with result of the given expression as last element.


#### `pop`

Takes 1 arguments and 1 optional argument. Takes an ARRAY as first argument and an Expression as second. Returns an ARRAY without the element with index of result of the given expression.

#### `range`

Takes 2 arguments and 1 optional argument. All arguments must be INTEGERs. Takes first value (start) as first argument, last value (end) as second, and step argument as an optinal third argument. Step can be negative or positive, cannot be 0.

#### `pow`

Takes 2 arguments. Takes two INTEGER. Returns an INTEGER. `pow(2,2)` = `4`

#### `sqrt`

Takes 1 argument. Takes an INTEGER. Returns an INTEGER. `sqrt(4)` = `2`

## License

Under the MIT License.