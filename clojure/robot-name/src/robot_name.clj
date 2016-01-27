(ns robot-name)

(def used-names (ref #{}))

(defn- rand-letter []
  (char (+ (int \A) (rand 26))))

(defn- unique-name []
  (let [name (str (rand-letter) (rand-letter) (rand-int 1000000))]
    (if-not (some #{name} @used-names)
      (dosync
        (ref-set used-names (conj @used-names name))
        name)
      (recur))))

(defn robot [] (ref (unique-name)))

(defn reset-name [robot]
  (let [new-name (unique-name)]
    (dosync
     (ref-set used-names (remove #(= name %) @used-names))
     (ref-set robot new-name))))

(defn robot-name [robot] @robot)
