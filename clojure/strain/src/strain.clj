(ns strain)

(defn retain [f coll]
  (for [elem coll, :when (f elem)] elem))

(defn discard [f coll]
  (let [not-f (comp not f)]
    (for [elem coll, :when (not-f elem)] elem)))
