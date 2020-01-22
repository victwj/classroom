# JavaScript

## Sources
- [Eloquent JavaScript 3rd Edition](https://eloquentjavascript.net/)
- [JavaScript: The Good Parts](https://github.com/NorthPaulo/research/blob/master/Frontend-books%26research/JavaScript%20-%20The%20Good%20Parts%20-%20Douglas%20Crockford%20-%20May%202008.pdf)

## Notes

### Objects
- Numbers (everything in JS is float64), strings, booleans, null, undefined are simple types
- Everything else are _objects_
- Objects can simply be made using {}, pretty much JSON
- Objects can be accessed using [] or .
- Objects are always passed by reference
- All objects are linked to a _prototype object_ which it can inherit properties from
- `typeof` and `.hasOwnProperty` (check without property without delegation)is useful when working with objects

### Functions
- Crockford: "The best thing about JavaScript is its implementation of functions"
- Functions in JS are objects
- Objects produced using object literals are linked to `Object.prototype`
- Function objects are linked to `Function.prototype` which is linked to `Object.prototype`
- Creating functions:

``` javascript
var add = function (a, b) {
    return a + b;
};

```
- A function literal can appear anywhere an expression can appear; so it can be inside other functions (and has access to the outer params/vars), called _closure_
- Every function receives `this` and `arguments` as additional parameters
- `this` is determined by the invocation pattern: method, function, constructor, or apply invocations
- Invocation is a pair of parenteses that follow any expression that produces a function value
