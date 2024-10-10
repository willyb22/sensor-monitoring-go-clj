(ns app.utils.math)

(defn factorial [n]
  (if-not (integer? n)
    (throw (ex-info "n is not a non-negative integer" {:type :not-integer-input}))
    (cond
      (neg? n) (throw (ex-info "n is not a non-negative integer" {:type :not-integer-input}))
      (zero? n) 1
      :else (* n (factorial (dec n))))))