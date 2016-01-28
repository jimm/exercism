(ns binary)

(defn to-decimal [s]
  (loop [n 0, chars (seq s)]
    (if (empty? chars) n
        (let [c (first chars)]
          (recur (+ (* 2 n) (if (= c \1) 1 0))
                 (rest chars))))))
