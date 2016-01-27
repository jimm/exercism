(ns word-count
  (:require [clojure.string :as str]))

(defn word-count [s]
  (frequencies (-> s
                   str/lower-case
                   (str/replace #"\W" " ")
                   (str/replace #"  +" " ")
                   (str/split #" "))))
