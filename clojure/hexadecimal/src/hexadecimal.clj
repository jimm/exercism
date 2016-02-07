(ns hexadecimal)

(def ^:private cmap {\0 0 \1 1 \2 2 \3 3 \4 4 \5 5 \6 6 \7 7
                     \8 8 \9 9 \a 10 \b 11 \c 12 \d 13 \e 14 \f 15})

(defn hex-to-int [s]
  (println "hex-to-int" s)
  (let [digits (keys cmap)]
    (loop [n 0
           chars (seq (.toLowerCase s))]
      (let [ch (first chars)]
        (println "ch" ch)
        (cond (or (nil? ch) (nil? (some #{ch} digits))) 0
              (empty? chars) n
              :else (recur (+ (* 16 n) (cmap ch)) (rest chars)))))))
