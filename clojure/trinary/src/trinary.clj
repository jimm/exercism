(ns trinary)

(defn to-decimal [s]
  (loop [n 0, chars (seq s)]
    (if (empty? chars) n
        (let [c (first chars)]
          (if (some #{c} #{\0 \1 \2})
            (recur (+ (* 3 n) (cond (= c \1) 1 (= c \2) 2 :else 0))
                   (rest chars))
            0)))))
