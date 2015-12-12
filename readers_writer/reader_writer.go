package main

import (
    "fmt"
    "time"
    "math/rand"
)

var channel = make(chan int, 1)
var buffer = []int{0,1}
var done = make(chan bool)

func writer() {
    time.Sleep(time.Duration(rand.Intn(4-1)+1)*time.Second)
    <-channel   // Lock
    buffer[0] += 1
    buffer[1] += 1
    fmt.Println("Wrote to buffer, buffer[0]:", buffer[0],"buffer[1]:",buffer[1])
    channel <-1 // Unlock
    done <-true
}

func reader() {
    time.Sleep(time.Duration(rand.Intn(4-1)+1)*time.Second)
    <-channel   // Lock
    fmt.Println("Read buffer[0]:", buffer[0]," buffer[1]:", buffer[1])
    channel <-1 // Unlock
    done <-true
}

func main() {

    fmt.Println("lol")
    channel <-1

    for i := 0; i < 5; i++ {
        go reader()
        go reader()
        go reader()
        go writer()
    }
    for i := 0; i < 5; i++ {
        <-done
        <-done
        <-done
        <-done
    }
}
