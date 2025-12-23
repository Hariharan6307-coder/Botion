package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	// 1. HARDCODE for now to test (We will move this to an env variable later)
	// REPLACE THIS STRING with the "DB URL" you copied from 'supabase status'
	dbUrl := "postgresql://postgres:postgres@127.0.0.1:54322/postgres"

	// 2. Connect to the database
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// Close the connection when main() finishes
	defer conn.Close(context.Background())

	// 3. Test the connection with a simple Ping
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("Could not ping database")
	}

	fmt.Println("✅ Successfully connected to local Supabase Postgres!")
}