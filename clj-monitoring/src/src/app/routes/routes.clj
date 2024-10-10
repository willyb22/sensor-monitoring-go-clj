(ns app.routes.routes
  (:require 
   [compojure.core :refer [routes]]
   [compojure.route :as route]
   [app.routes.sensor-routes :refer [sensor-routes]]
   [app.routes.home-routes :refer [home-routes]]))


(def app-routes
  (routes home-routes sensor-routes (route/not-found "Not Found")))