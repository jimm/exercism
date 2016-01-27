(ns nucleotide-count)

(defn nucleotide-counts [s]
  (merge {\A 0 \C 0 \G 0 \T 0}
         (frequencies s)))

(defn count [ch s]
  (or ((nucleotide-counts s) ch)
      (throw "ch must be \\A, \\C, \\G, or \\T")))
