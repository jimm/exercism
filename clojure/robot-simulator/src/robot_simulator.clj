(ns robot-simulator)

(def turn-right {:north :east, :east :south, :south :west, :west :north})

(def turn-left {:north :west, :east :north, :south :east, :west :south})

(defn robot [loc dir] {:bearing dir :coordinates loc})

;; Maps instruction to function that takes a robot and returns a new robot.
(def ^:private instructions
  {\R #(update % :bearing (fn [dir] (turn-right dir)))
   \L #(update % :bearing (fn [dir] (turn-left dir)))
   \A #(update % :coordinates
                  (fn [coords bearing]
                    (let [x (:x coords) y (:y coords)]
                      (cond (= bearing :north) {:x x :y (inc y)}
                            (= bearing :east) {:x (inc x) :y y}
                            (= bearing :south) {:x x :y (dec y)}
                            (= bearing :west) {:x (dec x) :y y})))
                  (:bearing %))})

(defn- move [robot instruction]
  ((instructions instruction) robot))

(defn simulate [str robot]
  (reduce #(move %1 %2) robot (seq str)))
