(ns roman-numerals)

(def ^:private digits-for-power
  {1 [\I \V \X]
   10 [\X \L \C]
   100 [\C \D \M]
   1000 [\M \M \M]})                    ; wrong but good enough

(def ^:private digit-counts
  {0 []
   1 [0]
   2 [0 0]
   3 [0 0 0]
   4 [0 1]
   5 [1]
   6 [1 0]
   7 [1 0 0]
   8 [1 0 0 0]
   9 [0 2]})

(defn- roman-digit [n power]
  (let [digits (digits-for-power power)
        indexes (digit-counts n)]
    (apply str (map #(nth digits %) indexes))))

(defn numerals [n]
  (loop [n n
         power 1
         numerals ()]
    (if (zero? n) (apply str numerals)
        (let [remaining-digits (int (/ n 10))
              digit (- n (* remaining-digits 10))]
          (recur remaining-digits
                 (* power 10)
                 (conj numerals (roman-digit digit power)))))))
