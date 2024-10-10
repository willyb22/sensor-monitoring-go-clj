(ns app.services.sensor-service
  (:require [app.services.services :refer [db-connection]]
            [app.services.utils :as sutils]
            [next.jdbc :as jdbc]
            [next.jdbc.result-set :as rs]))

(defn get-variable-data [variable]
  (let [query (sutils/get-query-by-column (keyword variable))]
    (jdbc/execute! db-connection [query] {:builder-fn rs/as-unqualified-kebab-maps})))