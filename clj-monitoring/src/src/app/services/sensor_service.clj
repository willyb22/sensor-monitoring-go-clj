(ns app.services.sensor-service
  (:require [app.services.services :refer [get-db-connection]]
            [app.services.utils :as sutils]
            [app.utils.misc :refer [extract-alphabets]]
            [next.jdbc :as jdbc]
            [next.jdbc.result-set :as rs]
            [app.models.sensor-model :as sensor-model]))

(defn get-variable-data [& variables]
  (let [query (apply sutils/get-query-by-columns (map keyword variables))]
    (jdbc/execute! (get-db-connection) [query] {:builder-fn rs/as-unqualified-kebab-maps})))

(defn get-sensor-names []
  (jdbc/execute! (get-db-connection) ["SELECT sensor_name FROM sensor_base"] {:builder-fn rs/as-unqualified-kebab-maps}))

(defn get-sensor-detail []
  (jdbc/execute! (get-db-connection) ["SELECT sensor_name, latitude, longitude FROM sensor_base"] {:builder-fn rs/as-unqualified-kebab-maps}))

(def sensor-names
  (for [m (get-sensor-names)] (first (vals m))))

(defn get-sensor-data [variable sensor-name]
  (let [valid-sensor-name (get (set sensor-names) sensor-name (first sensor-names))
        sensor-type (extract-alphabets valid-sensor-name)
        variables (get sensor-model/measurement-map (keyword sensor-type))
        valid-variable (get (set variables) variable (first variables))
        where-clause (str "WHERE m.sensor_name = '" valid-sensor-name "'")
        query (sutils/get-query-by-columns-sub [(keyword valid-variable)] sensor-type where-clause)]
    (jdbc/execute! (get-db-connection) [query] {:builder-fn rs/as-unqualified-kebab-maps})))