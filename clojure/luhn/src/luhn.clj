(ns luhn
  (:require [clojure.string :as str]))

(defn digits [n]
  (map #(- (int %) (int \0)) (str/replace (str n) #"\D" "")))

(defn- wrap [n] (if (>= n 10) (- n 9) n))

(defn checksum[n]
  (let [ds (partition 2 2 [0] (reverse (digits n)))]
    (rem (apply + (map #(+ (wrap (first %)) (wrap (* 2 (second %))))
                       ds))
         10)))

(defn valid? [n]
  (zero? (checksum n)))

(defn add-check-digit [n]
  (let [num (* n 10)]
    (first (for [check-digit (range 10)
                 :when (valid? (+ num check-digit))]
             (+ num check-digit)))))
