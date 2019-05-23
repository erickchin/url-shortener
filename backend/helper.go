package main

import (
	"math/rand"
	"time"
  )

// Taken from https://www.calhoun.io/creating-random-strings-in-go/
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
	  b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
  }