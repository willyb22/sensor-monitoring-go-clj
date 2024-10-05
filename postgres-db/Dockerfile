FROM postgres:alpine

# Copy custom scripts for database initialization
COPY ./init-db.sh ./init-db.sql ./sensor.csv /docker-entrypoint-initdb.d/

# Ensure the file is executable
RUN chmod +x /docker-entrypoint-initdb.d/init-db.sh

ENTRYPOINT [ "docker-entrypoint.sh" ]

CMD ["postgres"]