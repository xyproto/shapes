# shapes

Deal with points, rectangles and polygons.

This is one of the cases where there is a lot of repition of code, unless using `interface{}` everywhere.

## TODO

* Use `interface{}` everywhere, instead of one function for `int`, one function for `float64`, one function for `int` that returns `float64` etc.
* Alternatively, try out a code generation tool.
* Decide if the goal should be performance or ease of use, or create a package for each one.
