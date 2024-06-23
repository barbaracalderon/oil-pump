\c oil;

CREATE TABLE EquipmentData (
    id SERIAL PRIMARY KEY,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    suction_pressure FLOAT NOT NULL,
    discharge_pressure FLOAT NOT NULL,
    flow_rate FLOAT NOT NULL,
    fluid_temperature FLOAT NOT NULL,
    bearing_temperature FLOAT NOT NULL,
    vibration FLOAT NOT NULL,
    impeller_speed INT NOT NULL,
    lubrication_oil_level FLOAT NOT NULL,
    npsh FLOAT NOT NULL
);
