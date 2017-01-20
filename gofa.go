
package gofa

import (
	"math/rand"
	"time"
)

func Generatecode() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(8999) + 1000
}

func Createauth() {

}

func Verifyauth() {

}
