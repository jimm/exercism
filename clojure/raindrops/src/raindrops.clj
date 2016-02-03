(ns raindrops)

(defn convert [n]
  (let [s (apply str
                 (when (zero? (rem n 3)) "Pling")
                 (when (zero? (rem n 5)) "Plang")
                 (when (zero? (rem n 7)) "Plong"))]
    (if (= s "")
      (str n)
      s)))
