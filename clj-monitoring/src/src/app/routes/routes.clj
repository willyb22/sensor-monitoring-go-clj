(ns app.routes.routes
  (:require 
   [compojure.core :refer [routes]]
   [compojure.route :as route]
   [ring.middleware.resource :refer [wrap-resource]]
   [app.routes.sensor-routes :refer [sensor-routes]]
   [app.routes.home-routes :refer [home-routes]]
   [app.middleware.logging :refer [wrap-logging]]))


(def app-routes
  (-> (routes home-routes sensor-routes (route/not-found "Not Found"))
      (wrap-logging)
      (wrap-resource "public")))