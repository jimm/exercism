(ns prime-factors)

;;; This works, but is extremely slow. Finding the third prime factor of
;;; 93819012551 takes much too long.

(defn- prime-test
  [^long n ^long i ^long ceiling]
  (if (= i ceiling)
    true
    (loop [i i]
      (cond (= n i) (>= i ceiling)
            (zero? (rem n i)) (= i ceiling)
            :else (recur (unchecked-inc i))))))

(defn- prime? [^long n]
  (let [ceiling (unchecked-inc (int (Math/sqrt n)))]
    (if (= n (* ceiling ceiling))
      false
      (prime-test n 2 ceiling))))

;;; Assumes n is either 2 or odd.
(defn- next-prime ^long [^long n]
  (if (= n 2)
    3
    (first
     (take 1
           (drop-while (comp not prime?)
                       (iterate #(+ % 2) (+ n 2)))))))

(defn of [^long n]
  (loop [n n, prime 2, prime-factors ()]
    (cond
      (= n 1) (reverse prime-factors)
      (zero? (rem n prime)) (recur (/ n prime) prime (conj prime-factors prime))
      :else (recur n (next-prime prime) prime-factors))))
