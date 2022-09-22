package domain

import (
    "fmt"
)

type Robot struct {
    Name    string
    X       int
    Y       int
    D       int
}

const ymax = 10
const xmax = 10

func (r Robot) Hello() string {
    return fmt.Sprintf("Hello, %s", r.Name)
}

func (r Robot) Move(s int) {

}
