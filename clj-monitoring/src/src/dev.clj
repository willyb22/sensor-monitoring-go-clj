(ns dev
  (:require [app.core :as core]
            [clojure.tools.namespace.repl :refer [refresh]]
            [clojure.java.shell :as shell]
            [clojure.string :as cstr]))

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

;; (defn kill-process []
;;   (let [pid (-> (shell/sh "lsof" "-t" (str "-i:" 3000))
;;                 :out
;;                 cstr/trim)]
;;     (when (not (empty? pid))
;;       (shell/sh "kill" "-9" pid))))

(defn restart-server []
  (stop-server)
  ;; (kill-process)
  (println "Restarting server ....")
  (refresh)
  (start-server))