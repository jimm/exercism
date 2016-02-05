(ns sieve)

(defn sieve [n]
  (loop [nums (range 2 (inc n))
         primes []]
    (let [p (first nums)]
      (if (empty? nums) primes
          (recur (filter #(pos? (rem % p)) nums)
                 (conj primes p))))))
