(ns grade-school)

(defn add [db name grade]
  (assoc db grade (conj (or (db grade) []) name)))

(defn grade [db grade]
  (or (db grade) []))

(defn sorted [db]
  (->> db
       (into [])
       (map #(identity [(first %)
                        (sort (second %))]))
       (into (sorted-map))))
