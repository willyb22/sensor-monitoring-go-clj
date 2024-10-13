(ns app.controllers.home-controller
  (:require [ring.util.response :as response]
            [selmer.parser :refer [render-file]]))

(defn home-page [request]
  (let [params {:title "Clojure Monitoring"}]
    (render-file "templates/index.html" params)))