# Pov

Reparent a graph on a selected node

# Tree Reparenting

This exercise is all about re-orientating a graph to see things from a different
point of view. For example family trees are usually presented from the
ancestor's perspective:

```
          O
          |
    1-----2-----3
    |     |     |
  4-+-5 6-+-7 8-+-9
```

But the same information can be presented from the perspective of any other node
in the graph, by pulling it up to the root and dragging its relationships along
with it. So the same graph from 6's perspective would look like:

```
          6
          |
    7-----+-----2
                |
                0
          1-----+-----3
          |           |
        4-+-5       8-+-9
```

This lets us more simply describe the paths between two nodes. So for example
the path from 6-9 (which in the first graph goes up to the root and then down to
a different leaf node) can be seen to follow the path 6-2-0-3-9

This exercise involves taking an input graph and re-orientating it one one of
the nodes. 

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

Adaptation of exercise from 4clojure [view source](https://www.4clojure.com/)
