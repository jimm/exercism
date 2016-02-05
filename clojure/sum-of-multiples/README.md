# Sum Of Multiples

Write a program that, given a number, can find the sum of all the multiples of particular numbers up to but not including that number.

If we list all the natural numbers up to but not including 20 that are
multiples of either 3 or 5, we get 3, 5, 6 and 9, 10, 12, 15, and 18.

The sum of these multiples is 78.

Write a program that can find the sum of the multiples of a given set of
numbers. If no set of numbers is given, default to 3 and 5.

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

A variation on Problem 1 at Project Euler [view source](http://projecteuler.net/problem=1)
