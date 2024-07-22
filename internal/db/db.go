package db

import (
	"database/sql"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/Sohail-9098/vehicle-data-processing-analytics/internal/protobufs/vehicle"
	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DB struct {
	Conn *sql.DB
}

func NewDB() *DB {
	return &DB{}
}

func (d *DB) Connect() {
	serviceURI := os.Getenv("AIVEN_CREDENTIALS")
	if serviceURI == "" {
		log.Fatalf("DB Credentails Not Set")
	}
	conn, _ := url.Parse(serviceURI)
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"

	db, err := sql.Open("postgres", conn.String())
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	d.Conn = db
}

func (d *DB) Disconnect() {
	d.Conn.Close()
}

func (d *DB) CreateTable(tableName string, columns []string) {
	// Create table users (username VARCHAR, )
	query := "CREATE TABLE " + tableName + " ("
	for i := 0; i < len(columns)-1; i++ {
		query += columns[i] + ","
	}
	query += columns[len(columns)-1] + ");"
	_, err := d.Conn.Exec(query)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
	}
	log.Println(tableName, " created")
}

func (d *DB) DropTable(tableName string) {
	query := "DROP TABLE " + tableName + ";"

	_, err := d.Conn.Exec(query)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
	}
	log.Println(tableName, " deleted")
}

func (d *DB) InsertTelemetryData(tableName string, data *vehicle.Telemetry) {
	query := "INSERT INTO " + tableName + " VALUES ($1,$2,$3,$4,$5,$6);"
	_, err := d.Conn.Exec(query, data.VehicleId, data.Timestamp.AsTime(), data.Latitude, data.Longitude, data.Speed, data.FuelLevel)
	if err != nil {
		log.Println("failed to insert data: ", err)
		return
	}
	log.Println("data inserted: ", tableName)
}

func (d *DB) GetTelemetryData() []*vehicle.Telemetry {
	query := "SELECT * FROM telemetry;"
	rows, err := d.Conn.Query(query)
	if err != nil {
		log.Println("failed to fetch data: ", err)
	}
	telemetryData := &vehicle.Telemetry{}
	result := []*vehicle.Telemetry{}
	var timestamp time.Time
	for rows.Next() {
		err := rows.Scan(&telemetryData.VehicleId,
			&timestamp,
			&telemetryData.Latitude,
			&telemetryData.Longitude,
			&telemetryData.Speed,
			&telemetryData.FuelLevel,
		)
		telemetryData.Timestamp = timestamppb.New(timestamp)
		if err != nil {
			log.Println("failed to read row: ", err)
		}
		result = append(result, telemetryData)
	}
	return result
}
