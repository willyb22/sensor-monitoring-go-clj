(ns app.controllers.sensor-controller
  (:require [app.services.sensor-service :as sensor-service]
            [ring.util.response :as response]
            [cheshire.core :as json]))

(defn get-sensor-variable [variable]
  (try
    (let [data (sensor-service/get-variable-data variable)]
      (-> (response/response (json/generate-string data))
          (response/content-type "application/json")
          (response/status 200)))
    (catch Exception e
      (response/response (str "404 Not Found Error: " e)))))