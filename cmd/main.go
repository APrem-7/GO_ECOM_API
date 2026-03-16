package main

type application struct{
//methods that are going to run and moiunt your api and you can add graceful shutdown also

config config
//later on add logger and db drvier

}


type config struct {
	addr string //address to listen on
	db dbConfig


}

type dbConfig struct {
	dsn string //data source name its like the username password dbname and stuff
}
