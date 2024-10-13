(ns app.services.utils
  (:require [app.models.sensor-model :as sensor-model]
            [app.utils.misc :refer [contain-keys?]]
            [clojure.string :as cstr]))

(defn get-query-by-columns-sub [column-keys sensor-type & [where-clause]]
  (let [columns (->> (map #(str "s." (name %)) column-keys)
                     ((fn [x] (-> (interleave x (repeat (dec (count x)) ", ")) 
                                  (concat (last x)))))
                     (apply str))
        sensor_table (str sensor-type "_measurements")]
    (str "SELECT m.timestamp, " columns " FROM measurements m JOIN " sensor_table " s ON m.id = s.id"
         (when where-clause
           (str " " where-clause)))))

(defn get-query-by-columns [& column-keys]
  (loop [[k & ks] (->> [[sensor-model/bs-sensor-data-schema "bs"]
                        [sensor-model/aqs-sensor-data-schema "aqs"]
                        [sensor-model/msi-sensor-data-schema "msi"]]
                       (filter (fn [x] (apply contain-keys? (concat [(first x)] column-keys)))))
         i 0
         result ""]
    (if (nil? k)
      result
      (recur ks (+ i 1) (str result
                             (when (pos? i) " UNION ALL ")
                             (get-query-by-columns-sub column-keys (last k)))))))

(defn get-where-clause
  "columns is a collection of variable
   expression-str-f is a function that return str
   for example:
   let columns = ['m.timestamp' 's.temperature']
   estf = #(str %1 '<=`2024-10-11' ' and ' %2 '>0')"
  [columns-str expression-str-f]
  (str "where " (apply expression-str-f columns-str)))

(defn get-query-by-map [{:keys [measurements-str expression-str-f]}]
  (let [where-clause (get-where-clause measurements-str expression-str-f)
        measurements-other (filter #(not= "timestamp" %) measurements-str)
        query (->> (map keyword measurements-other)
                   (apply get-query-by-columns)
                   #(cstr/replace % "m.timestamp" "m.timestamp, m.latitude, m.longitude"))]
    (str query " " where-clause)))