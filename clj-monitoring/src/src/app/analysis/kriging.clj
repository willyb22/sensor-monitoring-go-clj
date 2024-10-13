(ns app.analysis.kriging
  (:require [app.controllers.sensor-service :as sensor-service]
            [app.utils.math :as umath]))

(defn distance [x1 y1 x2 y2]
  (Math/sqrt (+ (Math/pow (- x2 x1) 2) (Math/pow (- y2 y1) 2))))

(defn empirical-variogram [data lag-distance]
  (let [data (sensor-service/get-data-variable )]))

(defn kriging-spatio-temporal [lat lon time data]
  nil)