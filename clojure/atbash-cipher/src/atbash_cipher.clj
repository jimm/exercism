(ns atbash-cipher
  (:require [clojure.string :as str]))

(def ^:private charmap
  (let [chars (map char (range (int \a) (inc (int \z))))
        digits (map char (range (int \0) (inc (int \9))))]
    (zipmap (concat chars digits) (concat (reverse chars) digits))))

(defn- cleanup [s]
  (-> s
      .toLowerCase
      (str/replace #"[\W\s]" "")))

(defn- encode-chars [s]
  (clojure.core/replace charmap s))

(defn- reassemble [chars]
  (->> chars
       (partition 5)
       (interpose \space)
       flatten
       (apply str)))

(defn encode [s]
  (doall
      (->> s
           cleanup
           encode-chars
           reassemble)))
