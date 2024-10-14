(ns dev
  (:require [app.core :as core]
            [clojure.tools.namespace.repl :refer [refresh]]))

(defonce server (atom nil))

(defn start-server []
  (when-not @server
    (reset! server (core/-main))
    (println "server started")))

(defn stop-server []
  (when @server
    (.stop @server)
    (reset! server nil)
    (println "server stopped")))

(defn restart-server []
  (stop-server)
  ;; (kill-process)
  (println "Restarting server ....")
  (refresh)
  (start-server))