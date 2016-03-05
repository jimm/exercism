(ns point-mutations)

;; This isn't REALLY Hamming distance, since it simply counts changed
;; letters and doesn't care about insertions or deletions.
(defn hamming-distance [s1 s2]
  (when (= (count s1) (count s2))
    (count (filter (complement identity) (map = s1 s2)))))