(ns queen-attack)

(defn board-string [{[w-row w-col] :w [b-row b-col] :b}]
  (let [board-chars (->>
                     (for [row (range 8), col (range 8)]
                       (cond (and (= row w-row) (= col w-col)) \W
                             (and (= row b-row) (= col b-col)) \B
                             :true \_)))]
    (->> board-chars
         (partition 8)
         (map #(concat (interpose \space %) (list \newline)))
         flatten
         (apply str))))

(defn can-attack [{[w-row w-col] :w [b-row b-col] :b}]
  (or (= w-row b-row)
      (= w-col b-col)
      (= (Math/abs (- w-row b-row)) (Math/abs (- w-col b-col)))))
