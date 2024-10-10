(ns app.config.config)

(def go-backend-url (System/getenv "GO_BACKEND_URL"))

(def app-config {:server-port (System/getenv "SERVER_PORT")})

(def db-config {:classname "org.postgresql.Driver"
                :dbtype "postgresql" 
                :port (System/getenv "DB_PORT") 
                :host (System/getenv "DB_HOST") 
                :dbname (System/getenv "DB_NAME") 
                :user (System/getenv "DB_USER") 
                :password (System/getenv "DB_PASS")})