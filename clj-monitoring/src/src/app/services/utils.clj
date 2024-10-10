(ns app.services.utils
  (:require [app.models.sensor-model :as sensor-model]))

(defn get-query-by-column-sub [column_key sensor_type]
  (let [col (name column_key)
        sensor_table (str sensor_type "_measurements")]
    (str "SELECT m.timestamp, s." col " FROM measurements m JOIN " sensor_table " s ON m.id = s.id")))

(defn get-query-by-column [column_key]
  (loop [[k & ks] (->> [[sensor-model/bs-sensor-data-schema "bs"]
                        [sensor-model/aqs-sensor-data-schema "aqs"]
                        [sensor-model/msi-sensor-data-schema "msi"]]
                       (filter #(contains? (first %) column_key)))
         i 0
         result ""]
    (if (nil? k)
      result
      (recur ks (+ i 1) (str result
                             (when (pos? i) " UNION ALL ")
                             (get-query-by-column-sub column_key (last k)))))))