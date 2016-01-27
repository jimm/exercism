(ns triangle)

(defn legal? [a b c]
  (and (> (+ a b) c)
       (> (+ b c) a)
       (> (+ c a) b)))

(defn equilateral? [a b c]
  (= a b c))

(defn isosceles? [a b c]
  (or (= a b)
      (= a c)
      (= b c)))

(defn type [a b c]
  (cond (not (legal? a b c)) :illogical
        (equilateral? a b c) :equilateral
        (isosceles? a b c)   :isosceles
        :else                :scalene))
