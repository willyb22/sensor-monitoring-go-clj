DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE schemaname='public' AND tablename='sensor')
    THEN
        -- create sensor table
        CREATE TABLE sensor (
            sensorname VARCHAR(20) PRIMARY KEY,
            latitude DECIMAL(10,6) NOT NULL,
            longitude DECIMAL(10,6) NOT NULL
        );

        COPY sensor(sensorname, latitude, longitude)
        FROM '/docker-entrypoint-initdb.d/sensor.csv'
        DELIMITER ','
        CSV HEADER;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE schemaname='public' AND tablename='measurement')
    THEN
        -- create measurement table
        CREATE TABLE measurement (
            id SERIAL PRIMARY KEY,
            sensorname VARCHAR(20),
            temperature DECIMAL(10,6),
            humidity DECIMAL(10,6),
            "timestamp" TIMESTAMP,
            FOREIGN KEY (sensorname) REFERENCES sensor(sensorname) ON DELETE CASCADE
        );
    END IF;
END $$;