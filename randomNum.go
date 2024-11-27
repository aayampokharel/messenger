package main

import (
	"math/rand"
)

//! FOR MORE RANDOM , I CAN USE UUID OF GOOGLE.
func generateRandomNumber() int {
	room_id1 := rand.Intn(50000000)
	room_id2 := rand.Intn(50000000)
	room_id3 := rand.Intn(50000000)
	room_id4 := rand.Intn(5000000000)
	return room_id1 * room_id2 + room_id3 + room_id4;
}