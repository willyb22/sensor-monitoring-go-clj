(ns app.middleware.logging)

(defn wrap-logging [controller]
  (fn [request]
    (let [method (-> request :request-method name)
          uri (-> request :uri)]
      (println (str "endpoint " uri " was triggered. method: " method))
      (controller request))))