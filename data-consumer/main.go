package main

import (
    "database/sql/driver"
    "encoding/json"
    "log"
    "net/http"
    "time"
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    "github.com/streadway/amqp"
)

type Config struct {
    RabbitMQURL  string
    QueueName    string
    PostgresURL  string
}

type CustomTime struct {
    time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
    str := string(b)
    if str == "null" {
        ct.Time = time.Time{}
        return
    }
    // Remove the quotes around the date string
    str = str[1 : len(str)-1]
    ct.Time, err = time.Parse("2006-01-02T15:04:05.999999", str)
    return
}

func (ct CustomTime) Value() (driver.Value, error) {
    return ct.Time, nil
}

func (ct *CustomTime) Scan(value interface{}) error {
    v, ok := value.(time.Time)
    if !ok {
        return fmt.Errorf("can only scan time.Time, got %T", value)
    }
    ct.Time = v
    return nil
}

type EquipmentData struct {
    ID                  int        `json:"id" db:"id"`
    Timestamp           CustomTime `json:"timestamp" db:"timestamp"`
    SuctionPressure     float64    `json:"suction_pressure" db:"suction_pressure"`
    DischargePressure   float64    `json:"discharge_pressure" db:"discharge_pressure"`
    FlowRate            float64    `json:"flow_rate" db:"flow_rate"`
    FluidTemperature    float64    `json:"fluid_temperature" db:"fluid_temperature"`
    BearingTemperature  float64    `json:"bearing_temperature" db:"bearing_temperature"`
    Vibration           float64    `json:"vibration" db:"vibration"`
    ImpellerSpeed       int        `json:"impeller_speed" db:"impeller_speed"`
    LubricationOilLevel float64    `json:"lubrication_oil_level" db:"lubrication_oil_level"`
    Npsh                float64    `json:"npsh" db:"npsh"`
}


var db *sqlx.DB

func main() {
    config := Config{
        RabbitMQURL: "amqp://guest:guest@rabbitmq:5672/",
        QueueName:   "oil_queue",
        PostgresURL: "postgresql://postgres:example@postgres:5432/oil?sslmode=disable",
    }

    var err error
    db, err = sqlx.Connect("postgres", config.PostgresURL)
    if err != nil {
        log.Fatalf("Failed to connect to PostgreSQL: %v", err)
    }

    go consumeRabbitMQ(config)

    router := gin.Default()
    router.POST("/oil-data", postOilData)
    router.GET("/pressure", getPressureData)
    router.GET("/material", getMaterialData)
    router.GET("/fluid", getFluidData)
    router.Run(":8080")
}

func consumeRabbitMQ(config Config) {
    conn, err := amqp.Dial(config.RabbitMQURL)
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }
    defer ch.Close()

    msgs, err := ch.Consume(
        config.QueueName,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to register a consumer: %v", err)
    }

    for d := range msgs {
        log.Printf("Received a message: %s", d.Body)
        var data EquipmentData
        if err := json.Unmarshal(d.Body, &data); err != nil {
            log.Printf("Error decoding JSON: %v", err)
            continue
        }

        if _, err := db.NamedExec(`INSERT INTO EquipmentData (timestamp, suction_pressure, discharge_pressure, flow_rate, fluid_temperature, bearing_temperature, vibration, impeller_speed, lubrication_oil_level, npsh) VALUES (:timestamp, :suction_pressure, :discharge_pressure, :flow_rate, :fluid_temperature, :bearing_temperature, :vibration, :impeller_speed, :lubrication_oil_level, :npsh)`, data); err != nil {
            log.Printf("Error inserting data into PostgreSQL: %v", err)
        }
    }
}

func postOilData(c *gin.Context) {
    var data EquipmentData
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if _, err := db.NamedExec(`INSERT INTO EquipmentData (timestamp, suction_pressure, discharge_pressure, flow_rate, fluid_temperature, bearing_temperature, vibration, impeller_speed, lubrication_oil_level, npsh) VALUES (:timestamp, :suction_pressure, :discharge_pressure, :flow_rate, :fluid_temperature, :bearing_temperature, :vibration, :impeller_speed, :lubrication_oil_level, :npsh)`, data); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "data inserted"})
}

func getPressureData(c *gin.Context) {
    var data []EquipmentData
    err := db.Select(&data, "SELECT id, timestamp, suction_pressure, discharge_pressure, npsh FROM EquipmentData ORDER BY timestamp")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
        return
    }

    pressures := make([]gin.H, len(data))
    for i, d := range data {
        pressures[i] = gin.H{
            "id":                  d.ID,
            "timestamp":           d.Timestamp,
            "suction_pressure":    d.SuctionPressure,
            "discharge_pressure":  d.DischargePressure,
            "npsh":                d.Npsh,
        }
    }

    c.JSON(http.StatusOK, gin.H{"pressures": pressures})

}

func getMaterialData(c *gin.Context) {
    var data []EquipmentData
    err := db.Select(&data, "SELECT id, timestamp, vibration, bearing_temperature, impeller_speed FROM EquipmentData ORDER BY timestamp")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
        return
    }

    materials := make([]gin.H, len(data))
    for i, d := range data {
        materials[i] = gin.H{
            "id":                  d.ID,
            "timestamp":           d.Timestamp,
            "vibration":           d.Vibration,
            "bearing_temperature": d.BearingTemperature,
            "impeller_speed":      d.ImpellerSpeed,
        }
    }

    c.JSON(http.StatusOK, gin.H{"materials": materials})
}

func getFluidData(c *gin.Context) {
    var data []EquipmentData
    err := db.Select(&data, "SELECT id, timestamp, flow_rate, fluid_temperature, lubrication_oil_level FROM EquipmentData ORDER BY timestamp")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
        return
    }

    fluids := make([]gin.H, len(data))
    for i, d := range data {
        fluids[i] = gin.H{
            "id":                       d.ID,
            "timestamp":                d.Timestamp,
            "flow_rate":                d.FlowRate,
            "lubrication_oil_level":    d.LubricationOilLevel,
            "fluid_temperature":        d.FluidTemperature,
        }
    }

    c.JSON(http.StatusOK, gin.H{"fluids": fluids})

}
