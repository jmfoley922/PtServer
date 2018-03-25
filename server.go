package main

import (
	pb "PtServer/api"
	"PtServer/db"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	_ "database/sql/driver"

	_ "github.com/denisenkom/go-mssqldb"
)

//Exported app settings
type Settings struct {
	Port         string `json:"port"`
	CertName     string `json:"certname"`
	KeyName      string `json:"keyname"`
	SqlServer    string `json:"sqlserver"`
	Db           string `json:"db"`
	DbServerPort int    `json:"dbServerPort"`
	DbUser       string `json:"dbUser"`
	DbPassword   string `json:"dbPassword"`
}

type playerTrackingServer struct {
}

var appSettings = Settings{}
var SqlDb *sql.DB

//Read settings file and populate structure
func getSettings() (err error) {

	settingsStr, err := ioutil.ReadFile("./settings.json")

	err = json.Unmarshal(settingsStr, &appSettings)

	if err != nil {
		log.Printf("Error = %s\n", err)
		return err
	}

	return err

}

//Handle the CardInEvent
func (playerTrackingServer) CardInEvent(ctx context.Context, in *pb.CardIn) (*pb.Status, error) {

	message, points, displayname, err := db.CardIn(in.MachNumber, in.CardNumber, in.StartTime)

	return &pb.Status{Message: message, Points: points, DisplayName: displayname}, err

}

//Handle the MeterUpdateEvent
func (playerTrackingServer) MeterUpdateEvent(ctx context.Context, in *pb.MeterUpdate) (*pb.PointUpdate, error) {

	newBalance, err := db.MeterUpdate(in.MachNumber, in.CardNumber, in.CoinIn, in.CoinOut, in.Jackpot)

	return &pb.PointUpdate{MachNumber: in.MachNumber, CardNumber: in.CardNumber,
		NewPointBalance: newBalance}, err
}

//Handle the CardOutEvent
func (playerTrackingServer) CardOutEvent(ctx context.Context, in *pb.CardOut) (*pb.Status, error) {

	layout := "2006-01-02 15:04:05.000"
	date := time.Now().Format(layout)

	message, points, err := db.CardOut(in.MachNumber, in.CardNumber, in.CoinIn,
		in.CoinOut, in.Jackpot, string(date))

	return &pb.Status{Message: message, Points: points, DisplayName: ""}, err

}

//Get the Sql Server connection string from the settings file
func GetConnectionString() string {

	var connString = fmt.Sprintf("server=%s;user id=%s;password=%s;encrypt=%s;database=%s", appSettings.SqlServer, appSettings.DbUser, appSettings.DbPassword,
		"disable", appSettings.Db)

	return connString
}

func main() {

	err := getSettings()
	if err != nil {
		log.Fatal("Error reading settings file: " + err.Error())
	}

	db.Connstring = GetConnectionString()

	//Start the SSL server
	lis, err := net.Listen("tcp", ":"+appSettings.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(appSettings.CertName, appSettings.KeyName)
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}
	var players playerTrackingServer

	// Creates a new gRPC server
	s := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterPlayerTrackingServer(s, players)
	s.Serve(lis)
}
