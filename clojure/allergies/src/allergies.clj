(ns allergies)

(def ^:private allergy-data
  {:eggs 1,
   :peanuts 2,
   :shellfish 4,
   :strawberries 8,
   :tomatoes 16,
   :chocolate 32,
   :pollen 64,
   :cats 128})

(defn allergies [n]
  (for [[allergy score] allergy-data
        :when (pos? (bit-and n score))]
    allergy))

(defn allergic-to? [n allergy]
  (pos? (bit-and n (allergy-data allergy))))
