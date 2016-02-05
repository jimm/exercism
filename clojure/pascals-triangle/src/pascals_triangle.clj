(ns pascals-triangle)

(defn- first-half [n]
  (reduce (fn [row-so-far k]
            (conj row-so-far
                  (* (first (reverse row-so-far))
                     (/ (- (+ n 1) k) k))))
          [1]
          (range 1 (int (/ (+ 2 n) 2)))))

(defn row [n]
  (let [fh (first-half (dec n))
        odd-row (odd? n)]
    (concat fh
            (if odd-row
              (drop 1 (reverse fh))
              (reverse fh)))))

(def triangle (map row (iterate inc 1)))
