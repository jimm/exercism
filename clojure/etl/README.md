# Etl

We are going to do the `Transform` step of an Extract-Transform-Load.

### ETL
Extract-Transform-Load (ETL) is a fancy way of saying, "We have some crufty, legacy data over in this system, and now we need it in this shiny new system over here, so
we're going to migrate this."

(Typically, this is followed by, "We're only going to need to run this
once." That's then typically followed by much forehead slapping and
moaning about how stupid we could possibly be.)

### The goal
We're going to extract some scrabble scores from a legacy system.

The old system stored a list of letters per score:

- 1 point: "A", "E", "I", "O", "U", "L", "N", "R", "S", "T",
- 2 points: "D", "G",
- 3 points: "B", "C", "M", "P",
- 4 points: "F", "H", "V", "W", "Y",
- 5 points: "K",
- 8 points: "J", "X",
- 10 points: "Q", "Z",

The shiny new scrabble system instead stores the score per letter, which
makes it much faster and easier to calculate the score for a word. It
also stores the letters in lower-case regardless of the case of the
input letters:

- "a" is worth 1 point.
- "b" is worth 3 points.
- "c" is worth 3 points.
- "d" is worth 2 points.
- Etc.

Your mission, should you choose to accept it, is to write a program that
transforms the legacy data format to the shiny new format.

### Notes
Note that both the old and the new system use strings to represent
letters, even in languages that have a separate data type for
characters.

A final note about scoring, Scrabble is played around the world in a
variety of languages, each with its own unique scoring table. For
example, an "A" is scored at 14 in the Basque-language version of the
game while being scored at 9 in the Latin-language version.

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

The Jumpstart Lab team [view source](http://jumpstartlab.com)
