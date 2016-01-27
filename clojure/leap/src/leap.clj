(ns leap)

(defn leap-year? [n]
  (and (zero? (mod n 4))
       (or (zero? (mod n 400))
           (pos? (mod n 100)))))
