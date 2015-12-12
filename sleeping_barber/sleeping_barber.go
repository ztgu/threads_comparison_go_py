package main


import (
    "fmt"
    "sync"
    "time"
    "math/rand"
)

var condition = &sync.Cond{L: &sync.Mutex{}}
var barber_available = false

var customers_waiting = make([]int, 0)
var seats = 3
var done = make(chan bool)

var threads sync.WaitGroup

func cutting_hair(id int) {
    fmt.Println(id, ": Having a haircut")
    time.Sleep(time.Duration(rand.Intn(6-1)+1)*time.Second)
    fmt.Println(id, ": Done")
}

func barber() {
    print("Barber: Opening shop")
    for {
        condition.L.Lock()
        if len(customers_waiting) > 0 {
            id := customers_waiting[0]
            customers_waiting = customers_waiting[1:]
            condition.L.Unlock()
            barber_available = false
            cutting_hair(id)
            done <-true
        } else {
            fmt.Println("Barber: Going to sleep")
            for ; barber_available == false ; {
                condition.Wait()
            }
            condition.L.Unlock()
            fmt.Println("Barber: Woke up")
        }
    }
}

func customer(id int) {
    condition.L.Lock()
    if len(customers_waiting) == seats {
        fmt.Println(id, ": Leaving")
        done <-true
        condition.L.Unlock()
    } else {
        customers_waiting = append(customers_waiting,id)
        fmt.Println(id, ": In waiting-room: ", customers_waiting)
        barber_available = true
        condition.Signal()
        condition.L.Unlock()
    }
}

func main() {
    go barber()

    time.Sleep(time.Duration(1)*time.Second)
    var num_customers = 6
    for i := num_customers-1; i >= 0; i-=1 {
        go customer(i)
        time.Sleep(time.Duration(rand.Intn(3-1)+1)*time.Second)
    }
    for i:=0;i<num_customers;i+=1 {
        <-done
    }
}
