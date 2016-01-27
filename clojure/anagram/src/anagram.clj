(ns anagram
  (:require [clojure.string :as str]))

(defn- anagram? [s1 s2]
  (let [lc1 (str/lower-case s1)
        lc2 (str/lower-case s2)]
    (and (= (sort lc1) (sort lc2))
         (not= lc1 lc2))))

(defn anagrams-for [s choices]
  (filter #(anagram? s %) choices))
