(ns queen-attack)

(defn board-string [{[w-row w-col] :w [b-row b-col] :b}]
  (let [board-chars (->>
                     (for [row (range 8)
                           col (range 8)
                           :let [_ (println "row" row "col" col)]
                           ]
                       (cond (and (= row w-row) (= col w-col)) \W
                             (and (= row b-row) (= col b-col)) \B
                             :true \-)))]
    (println (interpose \space board-chars))
    (-> (conj (interpose \space board-chars) \space)
        (partition 16))))
))
   (interpose \newline)
   flatten
   (apply str)))
