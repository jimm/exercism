(ns phone-number
  (:require [clojure.string :as str]))

(defn digits [s]
  (str/replace s #"\D" ""))

(defn number [s]
  (let [m (re-matcher #"^1?(.{10})$" (digits s))]
    (if (.find m)
      (.group m 1)
      "0000000000")))

(defn area-code [s]
  (subs (number s) 0 3))

(defn pretty-print [s]
  (let [num (number s)]
    (str "(" (subs num 0 3) ") " (subs num 3 6) "-" (subs num 6))))