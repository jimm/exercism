(ns grains)

;;; Left-bit-shifting doesn't handle bigints, so use bit shifting until we
;;; can't any more, then start multiplying by 2.
(defn square [n]
  (if (< n 63)
    (bit-shift-left 1 (dec n))
    (loop [answer (bigint (bit-shift-left 1 62))
           n (- n 63)]
      (if (zero? n) answer
          (recur (* 2 answer) (dec n))))))

(defn total [] (dec (square 65)))
