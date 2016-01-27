(ns beer-song
  (:require [clojure.string :as str]))

(defn- bob [n]
  (str "bottle" (if (= n 1) "" "s") " of beer"))

(defn verse [n]
  (cond (= n 1)
        (str "1 bottle of beer on the wall, 1 bottle of beer.\n"
             "Take it down and pass it around,"
             " no more bottles of beer on the wall.\n")
        (zero? n)
        (str "No more bottles of beer on the wall, no more bottles of beer.\n"
             "Go to the store and buy some more,"
             " 99 bottles of beer on the wall.\n")
        :else
        (str n " bottles of beer on the wall, " n " bottles of beer.\n"
             "Take one down and pass it around, " (dec n) " " (bob (dec n))
             " on the wall.\n")))

(defn sing
  ([n] (sing n 0))
  ([n downto]
   (str/join "\n"
             (map verse (range n (dec downto) -1)))))
