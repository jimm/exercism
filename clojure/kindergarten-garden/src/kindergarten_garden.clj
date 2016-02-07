(ns kindergarten-garden
  (:require [clojure.string :as str]))

(def ^:private flower-names {\R :radishes \C :clover \G :grass \V :violets})

(defn- flowers-of [rows [key col]]
  (let [flower-chars (concat (take 2 (drop col (first rows)))
                             (take 2 (drop col (second rows))))]
    [key (map flower-names flower-chars)]))

(defn- name-to-map-entry [name col]
  [(-> name str/lower-case keyword) col])

(defn garden
  ([s] (garden s ["Alice" "Bob" "Charlie" "David" "Eve" "Fred" "Ginny"
                  "Harriet" "Ileana" "Joseph" "Kincaid" "Larry"]))
  ([s names]
   (let [cols (into {} (map #(name-to-map-entry %1 %2)
                            (sort names) (iterate #(+ % 2) 0)))
         rows (map seq (str/split-lines s))]
     (into {} (map #(flowers-of rows %) cols)))))
