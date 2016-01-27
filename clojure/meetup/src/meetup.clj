(ns meetup
  (import java.util.Calendar))

(def days {:sunday Calendar/SUNDAY
           :monday Calendar/MONDAY
           :tuesday Calendar/TUESDAY
           :wednesday Calendar/WEDNESDAY
           :thursday Calendar/THURSDAY
           :friday Calendar/FRIDAY
           :saturday Calendar/SATURDAY})

;; for each week name, value is [start-day, nth-day, day inc/dec]
(def week-ops {:first  [ 1 0  1]
               :second [ 1 1  1]
               :third  [ 1 2  1]
               :fourth [ 1 3  1]
               :last   [ 1 0 -1]
               :teenth [13 0  1]})

(defn meetup
  [month year day-name week]
  (let [d (days day-name)
        [start-day nth-day incdec-days] (week-ops week)
        c (Calendar/getInstance)]
    (.set c year (dec month) start-day 0 0 0)
    (when (= week :last)
      (.set c Calendar/DAY_OF_MONTH (.getActualMaximum c Calendar/DAY_OF_MONTH)))

    (loop [nth-day nth-day]
      (let [right-day-of-week? (= (.get c Calendar/DAY_OF_WEEK) d)]
      (cond (and right-day-of-week? (zero? nth-day))
            [year month (.get c Calendar/DAY_OF_MONTH)]

            right-day-of-week?
            (do (.add c Calendar/DAY_OF_MONTH (* 7 incdec-days))
                (recur (dec nth-day)))

            :else
            (do (.add c Calendar/DAY_OF_MONTH incdec-days)
                (recur nth-day)))))))
