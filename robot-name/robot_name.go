package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var usedNames = make(map[string]bool)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomLetter() string {
	return string(rune(rand.Intn(26)) + 'A')
}

func createName() string {
	return fmt.Sprintf("%s%s%03d", randomLetter(), randomLetter(), rand.Intn(1000))
}

func createUniqueName() string {
	for {
		if name := createName(); !usedNames[name] {
			usedNames[name] = true
			return name
		}
	}
}

func (r *Robot) Name() string {
	if r.name == "" {
		r.name = createUniqueName()
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = ""
}
