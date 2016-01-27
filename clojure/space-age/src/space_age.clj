(ns space-age)

(def secs-in-earth-year (* 365.25 24 60 60))

(defn- days-in-year [period]
  (* secs-in-earth-year period))

(defn- on-planet [period seconds]
  (/ seconds (days-in-year period)))

(defn on-earth [seconds] (on-planet 1.0 seconds))
(defn on-mercury [seconds] (on-planet 0.2408467 seconds))
(defn on-venus [seconds] (on-planet 0.61519726 seconds))
(defn on-mars [seconds] (on-planet 1.8808158 seconds))
(defn on-jupiter [seconds] (on-planet 11.862615 seconds))
(defn on-saturn [seconds] (on-planet 29.447498 seconds))
(defn on-uranus [seconds] (on-planet 84.016846 seconds))
(defn on-neptune [seconds] (on-planet 164.79132 seconds))
