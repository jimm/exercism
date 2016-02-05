(ns sum-of-multiples)

(defn- any-multiple? [i nums]
  (reduce (fn [found num] (or found (zero? (rem i num))))
          false nums))

(defn sum-of-multiples
  ([n] (sum-of-multiples [3 5] n))
  ([nums n]
   (apply +
          (for [i (range 1 n)
                :when (any-multiple? i nums)]
            i))))
