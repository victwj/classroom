# JavaScript

## Sources
- [Eloquent JavaScript 3rd Edition](https://eloquentjavascript.net/)
- [JavaScript: The Good Parts](https://github.com/NorthPaulo/research/blob/master/Frontend-books%26research/JavaScript%20-%20The%20Good%20Parts%20-%20Douglas%20Crockford%20-%20May%202008.pdf)

## Notes
- JavaScript is loosely typed; it never casts
- "What matters about an object is what it can do, not what it is descended from"
- JavaScript is a _prototypal_ language; objects inherit directly from other objects (not classes)
- "The pseudoclassical form [of JavaScript] can provide comfort to programmers who are unfamiliar with JavaScript, but it also hides the true nature of the language. Much of the complexity of class hierarchies is motivated by the constraints of static type checking. JavaScript is completely free of those constraints. In classical languages, class inheritance is the only form of code reuse. JavaScript has more and better options."

### Objects
- Numbers (everything in JS is float64), strings, booleans, null, undefined are simple types
- Everything else are _objects_
- Objects can simply be made using {}, pretty much JSON
- Objects can be accessed using [] or .
- Objects are always passed by reference
- All objects are linked to a _prototype object_ which it can inherit properties from
- `typeof` and `.hasOwnProperty` (check without property without delegation)is useful when working with objects
- Objects produced using object literals are linked to `Object.prototype`
- Function objects are linked to `Function.prototype` which is linked to `Object.prototype`
- Javascript _does not_ have block scope, unlike C-style languages
- It only has function scope
- Best style is therefore to declare variables at the top, as opposed to right-before-usage, to avoid unintentional variable hoisting

### Functions
- Crockford: "The best thing about JavaScript is its implementation of functions"
- Functions in JS are objects
- Creating functions:

``` javascript
var add = function (a, b) {
    return a + b;
};

```
- A function literal can appear anywhere an expression can appear; so it can be inside other functions (and has access to the outer params/vars), called _closure_
- Every function receives `this` and `arguments` as additional parameters
- `this` is determined by the invocation pattern: method, function, constructor, or apply invocations

#### Invocations
- Invocation is a pair of parentheses that follow any expression that produces a function value
    - There is no type checking to the given arguments
    - Too many arguments will be ignored, too few will result in `undefined`
- Method invocation:
    - Function stored as a property of an object
    - When invoked, the function is bound to that object
    - Methods can use `this` to access/modify the object
- Function invocation: when the function is not the property of an object
    - `this` is bound to the global object
    - Error in language design; it should be bound to the outer function `this`
    - Can be worked around, by convention a variable called `that` set to `this`
- Constructor invocation:
    - JS is a _prototypal_ inheritance language; objects can inherit properties directly from other objects, and the language is class-free
    - Invoke with a `new` prefix; but this is not recommended

    ``` javascript
    var Quo = function (string) { // Capitalize first letter for constructors
        this.status = string; // "this" is special for constructors
    }
    
    Quo.prototype.get_status = function () { return this.status; };
    var myQuo = new Quo("hello")
    document.writeln(myQuo.get_status()); // "hello"
    ```
- Apply invocation:
    - JS is a _functional_ object-oriented language; functions can have methods
    - `apply` is a method to call a function with a given `this` value 
    - First argument of `apply` is the `this`, the second argument is an array of arguments to call the function with
- `arguments` is a bonus parameter available to functions, and it is an array that contains _all_ arguments including those in excess
    - `arguments` is array-like but does not contain any array methods
- The `exception` object contains a `name` and a `message`; other properties can be added

## Prototypes
- 


## The Good Parts
- JavaScript makes it easy to define global variables, but globals increase the chance of bad interactions and reduces readability
- To reduce globals, use global abatement:

``` javascript
// Limit global footprint to a single name
var MYAPP = {};
MYAPP.flight = {
    airline: "Oceanic",
    number: 815,
};
```

- Closures are also a good way to reduce globals:

``` javascript
// Invoke function, store result in myObject
var myObject = (function() {
    var value = 0;
    // Return an object, working with a value that is scoped off from others
    return {
        increment: function (inc) {
            value += typeof inc === "number" ? inc : 1;
        },
        getValue: function () {
            return value;
        }
    };
}());
```
- Remember that closures have read/write access to actual variables, no copies
- When a function has a lot of arguments, use object specifiers:

``` javascript
// Don't do this
var myObject = myFunction(f, l, m, c, s);
// Instead do this
var myObject = myFunction({
    first: f,
    last: l,
    middle: m,
    state: s,
    city: c
});
// Arguments can be in different order, more clear, less errors
```

## The Bad Parts
- Don't use `==` or `!=`; Always use `===` or `!==`
- Don't use `with`
- Don't use bit operators (unless you have to); they're very slow
- Don't use `eval()`; it's slow and unclear
- Don't use `continue`; bad coding practice
- Don't use `new`
- Don't use `void`; it's confusingly an operator that takes and operand and returns undefined
- Don't create functions in a loop
