package main

import (
	"os"
)

func main() {
	puerto := os.Args[1]
	directorioImagenes := os.Args[2]
	startServer(puerto, directorioImagenes)
}
