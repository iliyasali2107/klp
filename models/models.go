package models

type Line struct {
	Sport string
	Cf    float64
}

type Env struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBName     string
	DBPassword string
}
