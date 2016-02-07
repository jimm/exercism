(ns queen-attack)

(defn board-string [{[w-row w-col] :w [b-row b-col] :b}]
  (->>
   (for [row (range 8)
         col (range 8)]
     (cond (and (= row w-row) (= col w-col)) \W
           (and (= row b-row) (= col b-col)) \B
           :true \-))
   (interpose \space)
   (apply str)
   (interpose \n)
   (apply str)))



