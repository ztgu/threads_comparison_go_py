package main

import (
    "fmt"
    "sync"
    "time"
)

var num_philosophers = 5
var philosophers = make([]int, 0)
var forks = make([]sync.Mutex, 0)
var done = make(chan bool)

func philosopher(idx int) {
    left_fork := idx
    right_fork := (idx-1)
    if right_fork < 0 {
        right_fork = num_philosophers - 1
    }
    for {
        forks[right_fork].Lock()
        forks[left_fork].Lock()
        fmt.Println("Philosopher:", idx, " eating, forks: ", 
            left_fork, right_fork)
        time.Sleep(time.Duration(2)*time.Second)
        forks[right_fork].Unlock()
        forks[left_fork].Unlock()
        time.Sleep(time.Duration(1)*time.Second)
    }
}

func main() {
    fmt.Println("ol")
    for i := 0 ; i < num_philosophers; i+=1 {
        var lock = sync.Mutex{}
        forks = append(forks, lock)
    }
    for i := 0; i < num_philosophers; i+=1 {
        go philosopher(i)
    }
    <-done
}

