package main

import (
    "fmt"
    "math/rand"
    "time"
)

var buffer_max = 4
var buffer_chan = make(chan int, buffer_max)
var buffer_copy = make([]int, 0)

func producer() {
    for {
        write_this := rand.Intn(50)
        buffer_chan <-write_this
        buffer_copy = append(buffer_copy, write_this)
        fmt.Println(buffer_copy, " produced: ", write_this)
        time.Sleep(time.Duration(rand.Intn(3))*time.Second)
    }
}

func consumer(id int) {
    time.Sleep(time.Duration(1)*time.Second)
    for {
        read_this := <-buffer_chan
        buffer_copy = buffer_copy[1:]
        fmt.Println(buffer_copy, " consumer: ", id , " consumed: ", read_this)
        time.Sleep(time.Duration(rand.Intn(7-2)+2)*time.Second)
    }
}

func main() {
    done := make(chan bool)

    go producer()
    go consumer(1)
    go consumer(2)

    <-done
}
