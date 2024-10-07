(ns app.config.config)

(def AppConfig {
    :BACKEND_URL (System/getenv "GO_BACKEND_URL")
    :SERVER_PORT (System/getenv "SERVER_PORT")
})