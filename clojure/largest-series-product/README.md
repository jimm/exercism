# Largest Series Product

Write a program that, when given a string of digits, can calculate the largest product for a series of consecutive digits of length n.

For example, for the input `'0123456789'`, the largest product for a
series of 3 digits is 504 (7 * 8 * 9), and the largest product for a
series of 5 digits is 15120 (5 * 6 * 7 * 8 * 9).

For the input `'73167176531330624919225119674426574742355349194934'`,
the largest product for a series of 6 digits is 23520.

* * * *

For learning resources and help with installation, refer to the
[Exercism help page][].

To run the tests provided, you will need to install [Leiningen][].

To install Leiningen on Mac OS X using [Homebrew][], run the following command:

    brew install leiningen

For help installing on Linux, Windows or without Homebrew see:
[Leiningen installation][].

[Exercism help page]: http://exercism.io/languages/clojure
[Leiningen]: http://leiningen.org
[Homebrew]: http://brew.sh
[Leiningen installation]: https://github.com/technomancy/leiningen#installation

In an exercise directory, create a `src` directory and a file therein to hold
your solution. The name of the file should be the exercise name with dashes `-`
replaced by underscores `_`.  For example, if the exercise is called
`hello-world`, name the solution file `hello_world.clj`.

Your resulting file tree should look something like this:

    /path/to/hello-world
    ├── project.clj
    ├── src
    │   └── hello_world.clj
    └── test
        └── hello_world_test.clj


To run the tests, navigate to the exercise directory and run the following
command:

    lein test

## Source

A variation on Problem 8 at Project Euler [view source](http://projecteuler.net/problem=8)
