(ns app.config.config)

(def GoBackendURL (System/getenv "GO_BACKEND_URL"))

(def AppConfig {
    :SERVER_PORT (System/getenv "SERVER_PORT")
})

(def DBConfig {
    :DB_PORT (System/getenv "DB_PORT")
    :DB_HOST (System/getenv "DB_HOST")
    :DB_NAME (System/getenv "DB_NAME")
    :DB_USER (System/getenv "DB_USER")
    :DB_PASSWORD (System/getenv "DB_PASSWORD")
})