(ns crypto-square
  (:require [clojure.string :as str]))

(defn normalize-plaintext [s]
  (-> s
      .toLowerCase
      (str/replace #"\W" "")))

(defn square-size [s]
  (let [len (count s), int-sqrt (int (Math/sqrt len))]
    (if (= (* int-sqrt int-sqrt) len) int-sqrt (inc int-sqrt))))

(defn plaintext-segments [s]
  (let [norm (normalize-plaintext s)
        sq-size (square-size norm)]
    (map #(apply str %) (partition sq-size sq-size "" norm))))

(defn- ciphertext-columns [s]
  (let [segs (plaintext-segments (normalize-plaintext s))
        cols (count (first segs))
        rows (count segs)
        last-row-len (count (first (reverse segs)))]
    (for [col (range cols)]
      (apply str (flatten
                  (for [row (range rows)
                        :when (or (< row (dec rows)) (< col last-row-len))]
                    (nth (nth segs row) col)))))))

(defn ciphertext [s]
  (apply str (ciphertext-columns s)))

(defn normalize-ciphertext [s]
  (str/join " " (ciphertext-columns s)))

