(ns app.core
  (:require 
   [ring.adapter.jetty :refer [run-jetty]] 
   [app.services.services :refer [ping-db close-db]]
   [app.routes.routes :refer [app-routes]]
   [app.config.config :refer [app-config]]
   [app.utils.misc :refer [str-to-int]]))

(defn -main []
  (println "App starting ...")
  (println "Trying to connect to database")
  (ping-db 5)
  (run-jetty app-routes
             {:port (str-to-int (:server-port app-config))
              :join false})
  (close-db)
  (print "App has stopped"))
