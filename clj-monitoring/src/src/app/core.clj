(ns app.core
  (:require [app.config.config :refer [AppConfig]]))

(defn -main []
  (println (str "Server will be run at port " (:SERVER_PORT AppConfig))))
