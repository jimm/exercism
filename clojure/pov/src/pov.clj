(ns pov)

(defn- debug
  "To stop debug output, comment out the `apply` form using
  the #_ reader macro."
  [& args]
  (apply println args))

(defmacro debug-dump
  "Generate a call to debug with each form becoming a string and a value,
  all optionally preceded by a string. For example:

      (debug-dump x (reverse s)) ; (debug \"x\" x \"(reverse s)\" (reverse s))
      (debug-dump \"hello\" x y) ; (debug \"hello\" \"x\" x \"y\" y)"
  [init-str-or-form & forms]
  (let [prefix (if (string? init-str-or-form)
                 (list 'debug init-str-or-form)
                 (list 'debug))
        all-forms (if (string? init-str-or-form)
                   forms
                   (conj forms init-str-or-form))]
    (concat prefix
            (mapcat #(list (list 'str (list 'quote %)) %)
                    all-forms))))

(defn- node-name [node] (first node))

(defn- node-children [node] (second node))

(defn- subtree-map
  "Return a map whose keys are node names and values are those node's
  subtrees."
  [tree]
  (let [ts (tree-seq coll? next tree)]
    (zipmap (map first ts) ts)))

(defn- parent-of
  "Return the parent symbol of node in tree. Walk the tree looking for the
  node that has as one of its children the given node. This is horribly
  inefficient but good enough for the small trees we're dealing with."
  [node tree]
  (first (filter #(some #{node}
                        (set (map node-name (rest %))))
                 (tree-seq rest rest tree))))

(defn- ancestors-of [node tree]
  (let [root-name (node-name tree)]
    (loop [curr-node-name node
           parents []]
      (if (= curr-node-name root-name)
        parents
        (let [parent-name (first (parent-of curr-node-name tree))]
          (recur parent-name (conj parents parent-name)))))))

(defn- add-child
  [child-tree tree]
  "Add child-tree as a child of the root node of tree."
  (debug-dump "add-child" child-tree tree)
  (if (empty? (rest tree))
    [(first tree) child-tree]
    [(first tree) (conj (rest tree) child-tree)]))

(defn- remove-child
  [child-name tree]
  "Remove the child named child-name from the children of the root node of
  tree."
  (vec (concat [(first tree)]
               (remove #(= (first %) child-name) (rest tree)))))

(defn- ancestor-keys
  [tree ancestors]
  "Return a seq of integer keys suitable for use in assoc-in that lead from
  the root node of tree down through the given list of ancestor symbols."
  (letfn [(index-of [v val]
            (count (first (split-with #(not (= val %)) v))))]
    (loop [ancestors (rest ancestors)
           tree tree
           ks []]
      (if (empty? ancestors) ks
          (recur (rest ancestors)
                 ((first ancestors) (node-children tree))
                 (conj ks ;; index of first ancestor
                       (index-of (second tree) (first ancestors))))))))

(defn of
  [new-root tree]
  "Rewrite the tree so that node is the root."
  ;; Find target subtree
  ;; Find parent
  ;; Append parent, minus target subtree, to end of target subtree
  ;; Wash, rinse, repeat
  (debug-dump "\nof" new-root tree)
  (loop [ancestors (ancestors-of new-root tree)
         curr-node-name new-root
         tree tree]
    (debug-dump ancestors)
    (debug-dump curr-node-name)
    (if (empty? ancestors)
      tree
      (let [subtrees (subtree-map tree)
            parent-name (first ancestors)
            parent-node (parent-name subtrees)
            parent-minus-curr (remove-child curr-node-name parent-node)
            curr-node (curr-node-name subtrees)
            ks (ancestor-keys tree (conj (reverse ancestors) curr-node-name))]
        (debug-dump subtrees)
        (debug-dump parent-name)
        (debug-dump parent-minus-curr)
        (debug-dump curr-node)
        (debug "added parent as child" (add-child parent-minus-curr curr-node))
        (debug-dump (reverse ancestors))
        (recur (rest ancestors)
               (first ancestors)
               (assoc-in tree ks (add-child parent-minus-curr curr-node)))))))
