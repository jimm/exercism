(ns prime-factors)

(defn- prime-test
  [n i j h]
  (loop [i i]
    (cond (= n i) (>= i h)
          (zero? (rem n i)) (= i h)
          :else (recur (inc i)))))

(defn- prime? [n]
  (let [ceiling (inc (int (Math/sqrt n)))]
    (if (= n (* ceiling ceiling)) false
        (prime-test n 2 ceiling ceiling))))

;;; Assumes n is either 2 or odd.
(defn- next-prime [n]
  (if (= n 2) 3
      (first
       (take 1
             (drop-while (comp not prime?)
                         (iterate #(+ % 2) (+ n 2)))))))

(defn of [n]
  (loop [n n, prime 2, prime-factors ()]
    (cond (= n 1) (reverse prime-factors)
          (zero? (rem n prime)) (recur (/ n prime) prime (conj prime-factors prime))
          :else (recur n (next-prime prime) prime-factors))))
