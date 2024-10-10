(ns app.services.services
  (:require 
   [app.config.config :refer [db-config]]
   [app.services.utils :as sutils]
   [next.jdbc :as jdbc]))

(defn ping-db-sub []
  (try
    ;; Execute simple query
    (jdbc/execute! db-config ["SELECT 1"])
    (println "Database is reachable.")
    true
    (catch Exception e
      (println "Error connection to the database:"  (.getMessage e))
      false)))

(defn ping-db [max-retry]
  (if (false? (ping-db-sub))
    (do
      (println "Retrying ...")
      (loop [cond false
             retry 1]
        (if (or (> retry max-retry) cond)
          cond
          (do
            (Thread/sleep 1000) ;; sleep for 5s
            (println (str "Retry ... " retry)) 
            (recur (ping-db-sub) (+ retry 1))))))
    true))

(def datasource (jdbc/get-datasource db-config))

(def db-connection (jdbc/get-connection datasource))

(defn close-db [] (.close db-connection))