package database

func NewSqlStorage(config Post) {
	pgDB, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=your_password dbname=your_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer pgDB.Close()

	
	err = pgDB.Ping()
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("Successfully connected to PostgreSQL database!")
}