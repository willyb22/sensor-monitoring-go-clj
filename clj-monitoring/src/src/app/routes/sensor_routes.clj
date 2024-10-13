(ns app.routes.sensor-routes
  (:require
   [compojure.core :refer [defroutes GET]]
   [app.controllers.sensor-controller :as sensor-controller]))

(defroutes sensor-routes
  (GET "/sensor" [] sensor-controller/get-sensor)
  (GET "/sensor/measurements/:sensortype" [sensortype] (sensor-controller/get-measurement-types sensortype))
  (GET "/sensor/data/:variable/:sensor-name" [variable sensor-name] (sensor-controller/get-sensor-data variable sensor-name))
  (GET "/sensor/data-var/:variable" [variable] (sensor-controller/get-sensor-variable variable)))