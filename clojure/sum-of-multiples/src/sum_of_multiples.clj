(ns sum-of-multiples)

(defn- any-multiple? [i nums]
  (some #(zero? (rem i %)) nums))

(defn sum-of-multiples
  ([n] (sum-of-multiples [3 5] n))
  ([nums n]
   (apply + (filter #(any-multiple? % nums) (range 1 n)))))
