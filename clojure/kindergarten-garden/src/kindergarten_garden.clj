(ns kindergarten-garden
  (:require [clojure.string :as str]))

(def ^:private cols {:alice 1, :bob 3, :charlie 5, :david 7, :eve 9, :fred 11,
                     :ginny 13, :harriet 15, :ileana 17, :joseph 19,
                     :kinkaid 21, :larry 23})
(def ^:private flower-names {"R" :radishes "C" :clover "G" :grass "V" :violets})

(defn garden [s]
  (let [rows (str/split-lines s)]
    