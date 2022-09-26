package main

import (
	"fmt"

	"github.com/Theetat/cinema/movie"
	"github.com/Theetat/cinema/ticket"
)

func init() {
	fmt.Println("init: main")
}

func main() {
	moviename := movie.FindMovieName("tt4154796")
	fmt.Println(moviename)
	movie.Review(moviename, 8.4)
	ticket.BuyTicket(moviename)
}
