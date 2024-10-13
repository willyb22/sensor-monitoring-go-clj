(ns app.models.sensor-model
  (:require [clojure.spec.alpha :as s]))

(def measurement-map 
  {:bs ["temperature" "humidity"]
   :aqs ["temperature" "humidity" "co2_level"]
   :msi ["temperature" "humidity" "air_pressure" "wind_speed"]})

(def bs-sensor-data-schema
  {:id "int"
   :sensor_name "string"
   :sensor_type "string"
   :temperature "float"
   :humidity "float"
   :timestamp "datetime"})

(def aqs-sensor-data-schema
  {:id "int"
   :sensor_name "string"
   :sensor_type "string"
   :temperature "float"
   :humidity "float"
   :co2_level "float"
   :timestamp "datetime"})

(def msi-sensor-data-schema 
  {:id "int"
   :sensor_name "string"
   :sensor_type "string"
   :temperature "float"
   :humidity "float"
   :air_pressure "float"
   :wind_speed "float"
   :timestamp "datetime"})

;; Specs for individual field
(s/def ::id int?)
(s/def ::sensor_name string?)
(s/def ::sensor_type string?)
(s/def ::timestamp inst?)
(s/def ::temperature #(and (float? %) (pos? %)))
(s/def ::humidity #(and (float? %) (pos? %)))
(s/def ::co2_level #(and (float? %) (pos? %)))
(s/def ::air_pressure #(and (float? %) (pos? %)))
(s/def ::wind_speed #(and (float? %) (pos? %)))

;; Specs for sensor data
(s/def ::bs-sensor-data
  (s/keys :req [::id ::sensor_name ::sensor_type ::timestamp ::temperature ::humidity]))

(s/def ::aqs-sensor-data
  (s/keys :req [::id ::sensor_name ::sensor_type ::timestamp ::temperature ::humidity ::co2_level]))

(s/def ::msi-sensor-data
  (s/keys :req [::id ::sensor_name ::sensor_type ::timestamp ::temperature ::humidity ::air_pressure ::wind_speed]))

;; 
(defn is-bs? [data]
  (s/valid? ::bs-sensor-data data))

(defn is-aqs? [data]
  (s/valid? ::aqs-sensor-data data))

(defn is-msi? [data]
  (s/valid? ::msi-sensor-data data))