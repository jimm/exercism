(ns etl)

(defn transform [m]
  (into {}
        (for [[n letters] m, letter letters]
          [(.toLowerCase letter) n])))
