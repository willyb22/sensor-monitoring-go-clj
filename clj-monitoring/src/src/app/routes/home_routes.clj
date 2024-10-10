(ns app.routes.home-routes
  (:require
   [compojure.core :refer [defroutes GET]]
   [app.controllers.home-controller :as home-controller]))

(defroutes home-routes
  (GET "/" [] home-controller/home-page))