(ns bob
  (:require [clojure.string :as str]))

(defn response-for [str]
  (cond (str/blank? str)          "Fine. Be that way!"
        (and (=    str (str/upper-case str))
             (not= str (str/lower-case str)))
                                  "Whoa, chill out!"
        (.endsWith str "?")       "Sure." ; #str/ends-with? is Clojure 1.8
        :else                     "Whatever."))
