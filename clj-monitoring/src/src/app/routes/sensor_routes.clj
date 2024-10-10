(ns app.routes.sensor-routes
  (:require
   [compojure.core :refer [defroutes GET]]
   [app.controllers.sensor-controller :as sensor-controller]))

(defroutes sensor-routes
  (GET "/sensor/data/:variable" [variable] (sensor-controller/get-sensor-variable variable)))