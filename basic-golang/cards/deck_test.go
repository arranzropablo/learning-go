package main

import (
	"os"
	"testing"
)

//We want to do every test with the same deck
var d = newDeck()

const FILE_NAME = "deck_testing"

func TestNewDeckSize(t *testing.T) {
	//Without a library we havent got assertions but we check with if
	if len(d) != 48 {
		t.Errorf("Length of deck should be 48, but got %d", len(d))
	}
}

func TestNewDeckFirstCard(t *testing.T) {
	if d[0].number != 1 || d[0].suite != "Spades" {
		t.Errorf("First card should be 1 of Spades, but got %s", d[0].String())
	}
}

func TestNewDeckLastCard(t *testing.T) {
	if d[len(d)-1].number != 12 || d[len(d)-1].suite != "Hearts" {
		t.Errorf("First card should be 12 of Hearts, but got %s", d[len(d)-1].String())
	}
}

func TestSaveAndRead(t *testing.T) {
	os.Remove(FILE_NAME)
	success := d.save(FILE_NAME)

	if !success {
		t.Errorf("Something went wrong when saving")
	}

	testDeck := read(FILE_NAME)

	if testDeck.String() != d.String() {
		t.Errorf("Error when saving or loading, result Strings are not the same")
	}
	os.Remove(FILE_NAME)
}

//Benchmark test (performance test), it executes a func (in this case newDeck) N times and prints how many times it's executed and
//the mean of nanoseconds it took each time

//go test -bench='.*'
func BenchmarkNewDeck (b *testing.B){
	for i := 0; i < b.N; i++ {
		newDeck()
	}
}