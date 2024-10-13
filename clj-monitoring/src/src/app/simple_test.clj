(ns app.simple-test
  (:require [app.utils.misc :as misc]
            [app.services.utils :refer [get-query-by-columns]]))

(defn -main []
  (let [m {:a 1, :b 2, :c 3}]
    (do
      (println (contains? m :a))
      (println (misc/contain-keys? m :a))
      (println (misc/contain-keys? m :a :b :d))
      (println (get-query-by-columns :temperature))
      (println (get-query-by-columns :temperature :humidity))
      (println (get-query-by-columns :temperature :co2_level :air_pressure))
      (println (get-query-by-columns :temperature :humidity :air_pressure))
      true)))