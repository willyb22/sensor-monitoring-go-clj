(ns app.utils.misc
  (:require [clojure.string :as cstr]))

(defn str-to-int [s] (Integer/parseInt s))

(defn str-to-bool [s]
  (condp = (cstr/lower-case s) 
    "true" true 
    "false" false 
    :else false))

(defn extract-alphabets [s]
  (re-find #"[a-zA-Z]+" s))

(defn contain-keys? 
  "m is a map and ks is keys"
  [m & [k & ks]]
  (let [c? (fn [c k1]
             (and c (contains? m k1)))]
    (if (empty? ks)
      (contains? m k)
      (reduce c? k ks))))

(defn update-map [orig new-map]
  (merge-with (fn [_ new-value] new-value) orig (select-keys new-map (keys orig))))