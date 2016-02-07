(ns binary-search-tree)

(def value :value)
(def left :left)
(def right :right)

(defn singleton [val] {:value val})

(defn insert [val bst]
  (cond (nil? bst) (singleton val)
        (<= val (value bst)) (assoc bst :left (insert val (left bst)))
        :else (assoc bst :right (insert val (right bst)))))

(defn from-list [coll]
  (reduce (fn [bst val] (insert val bst))
          (singleton (first coll))
          (rest coll)))

(defn to-list [bst]
  (if (nil? bst) []
      (concat (to-list (left bst))
              (list (value bst))
              (to-list (right bst)))))
