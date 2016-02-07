(ns largest-series-product)

(defn digits [s]
  (map #(- (int %) (int \0)) s))

(defn slices [n s] (partition n 1 (digits s)))

(defn largest-product [n s]
  (when (> n (count s))
    (throw (Throwable. "Give me something to work with, here!")))
  (let [prods (map #(apply * %) (slices n s))]
    (if (empty? prods)
      1
      (apply max prods))))
