(ns app.controllers.sensor-controller
  (:require [app.services.sensor-service :as sensor-service]
            [app.models.sensor-model :as sensor-model]
            [ring.util.response :as response]
            [selmer.parser :refer [render-file]]
            [cheshire.core :as json]))
(def -test (->> {:sensor-names (for [m (sensor-service/get-sensor-names)] (first (vals m)))
                :measurement-map sensor-model/measurement-map}
               (json/generate-string)))

(defn get-sensor [request]
  (render-file "templates/sensor.html" {:sensor-names sensor-service/sensor-names}))

(defn get-measurement-types [sensor-type]
   (-> (get sensor-model/measurement-map (keyword sensor-type) ["temperature" "humidity"]) 
       json/generate-string 
       response/response 
       (response/content-type "application/json") 
       (response/status 200)))

(defn get-sensor-variable [variable]
  (try
    (let [data (sensor-service/get-variable-data variable)]
      (-> (response/response (json/generate-string data))
          (response/content-type "application/json")
          (response/status 200)))
    (catch Exception e
      (response/response (str "404 Not Found Error: " e)))))

(defn get-sensor-data [variable sensor-name]
  (let [data (sensor-service/get-sensor-data variable sensor-name)]
    (-> (json/generate-string data)
        response/response
        (response/content-type "application/json")
        (response/status 200))))