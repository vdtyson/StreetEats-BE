package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func main() {
	con, err := pgx.Connect(context.Background(), pgConfigs["test"].toString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer con.Close(context.Background())
	fmt.Println("Connected!")
}
