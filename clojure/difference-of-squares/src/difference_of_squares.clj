(ns difference-of-squares)

(defn square-of-sums [n]
  (let [sum (apply + (range (inc n)))]
    (* sum sum)))

(defn sum-of-squares [n]
  (apply + (map #(* % %) (range (inc n)))))

(defn difference [n] (- (square-of-sums n) (sum-of-squares n)))
