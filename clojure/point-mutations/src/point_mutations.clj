(ns point_mutations)

(defn count-one-bits
  "Compute the number of one bits in an integer."
  [^Integer n]
  (loop [n n
         sum 0]
    (cond (zero? n) sum
          (even? n) (recur (bit-shift-right n 1) sum)
          :else     (recur (bit-shift-right n 1) (inc sum)))))

(def count-one-bits (memoize count-one-bits))

(defn hamming-distance
  "Compute the Hamming distance between two bit sequences."
  [bytes0 bytes1]
  (reduce + (map (comp count-one-bits bit-xor) bytes0 bytes1)))
