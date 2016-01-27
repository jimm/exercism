(ns rna-transcription)

(def rna-mapping {\G \C, \C \G, \T \A, \A \U})

(defn to-rna [s]
  (apply str (map #(let [c (rna-mapping %)]
                     (assert c)         ; required by tests
                     c)
                  s)))
