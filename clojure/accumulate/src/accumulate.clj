(ns accumulate)

(defn accumulate [f coll]
  (for [elem coll] (f elem)))
