(ns binary-search)

(defn middle
  ([coll] (middle coll 0 (dec (count coll))))
  ([coll start end] (int (+ start (/ (- end start) 2)))))

;; error: not taking offset of subseq into account
(defn search-for [e coll]
  (when (not= coll (sort coll))
    (throw (Throwable. "must be sorted")))
  (loop [start 0
         end (dec (count coll))]
    (when (< end start)
      (throw (Throwable. "not found")))
    (let [mid (middle coll start end)
          nth-elem (nth coll mid)]
      (cond (= e nth-elem) mid
            (< e nth-elem) (recur start (dec mid))
            :else (recur (inc mid) end)))))
