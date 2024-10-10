(ns app.controllers.home-controller
  (:require [ring.util.response :as response]
            [clojure.java.io :as io]))

(defn home-page [request]
  (let [index-file (io/resource "index.html")
        dbg (do
              (println index-file)
              true)]
    (if index-file
      (-> (response/file-response (io/as-file index-file))
          (response/content-type "text/html")
          (response/status 200))
      (response/response (str "404 Not Found :" index-file)))))