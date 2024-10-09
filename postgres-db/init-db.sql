DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE schemaname='public' AND tablename='sensor_base')
    THEN
        -- create sensor table
        CREATE TABLE sensor_base (
            sensor_name VARCHAR(10) PRIMARY KEY,
            sensor_type VARCHAR(5) NOT NULL,
            latitude DECIMAL(10,6) NOT NULL,
            longitude DECIMAL(10,6) NOT NULL
        );

        COPY sensor_base(sensor_name, sensor_type, latitude, longitude)
        FROM '/docker-entrypoint-initdb.d/sensor.csv'
        DELIMITER ','
        CSV HEADER;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE schemaname='public' AND tablename='measurements')
    THEN
        -- create measurements table
        CREATE TABLE measurements (
            id SERIAL PRIMARY KEY,
            sensor_name VARCHAR(10),
            sensor_type VARCHAR(5),
            "timestamp" TIMESTAMP,
            FOREIGN KEY (sensor_name) REFERENCES sensor_base(sensor_name) ON DELETE CASCADE
        );
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE schemaname='public' AND tablename='msi_measurements')
    THEN
        -- create multi-sensors instrument table
        CREATE TABLE msi_measurements (
            id INT REFERENCES measurements(id),
            temperature DECIMAL(10,6),
            humidity DECIMAL(10,6),
            air_pressure DECIMAL(10,6),
            wind_speed DECIMAL(10,6)
        );
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE schemaname='public' AND tablename='aqs_measurements')
    THEN
        -- create air quality sensor table
        CREATE TABLE aqs_measurements (
            id INT REFERENCES measurements(id),
            temperature DECIMAL(10,6),
            humidity DECIMAL(10,6),
            co2_level DECIMAL(10,6)
        );
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE schemaname='public' AND tablename='bs_measurements')
    THEN
        -- create basic sensor table
        CREATE TABLE bs_measurements (
            id INT REFERENCES measurements(id),
            temperature DECIMAL(10,6),
            humidity DECIMAL(10,6)
        );
    END IF; 
END $$;