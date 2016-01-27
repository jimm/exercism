(ns gigasecond
  (:import [java.util Calendar]))

(defn from [y m d]
  (let [c (Calendar/getInstance)]
    (.set c Calendar/YEAR y)
    (.set c Calendar/MONTH (dec m))
    (.set c Calendar/DATE d)
    (.add c Calendar/SECOND 1000000000)
    (vector (.get c Calendar/YEAR)
            (inc (.get c Calendar/MONTH))
            (.get c Calendar/DATE))))
