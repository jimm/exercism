(ns octal)

(defn to-decimal [s]
  (loop [n 0
         chars (seq s)]
    (let [ch (first chars)]
      (cond
        (empty? chars) n
        (some #{ch} #{\0 \1 \2 \3 \4 \5 \6 \7})
          (recur (+ (bit-shift-left n 3) (- (int ch) (int \0)))
                 (rest chars))
        :else 0))))

            
    