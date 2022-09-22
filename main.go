package main

import (
    "fmt"

    "github.com/stretchr/testify/assert"
)

func main() {
    watcher := keyboard.NewWatcher()
    status := watcher.States()

    fmt.Println(status)
}

// func main() {
//     x, err := strconv.Atoi(os.Args[1])
//     if err != nil {
//         panic(err)
//     }

//     y, err := strconv.Atoi(os.Args[2])
//     if err != nil {
//         panic(err)
//     }

//     d, err := strconv.Atoi(os.Args[2])
//     if err != nil {
//         panic(err)
//     }

//     robot := domain.Robot {
//         X: x,
//         Y: y,
//         D: d,

//     }

//     fmt.Println(robot)
// }
