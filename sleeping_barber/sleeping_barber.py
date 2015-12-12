from threading import Thread, Lock, Condition
from time import sleep
from random import randint

mutex = Lock()
condition = Condition(mutex)
barber_available = 0

customers_waiting = []
seats = 3

def cutting_hair(id):
    print(id, ": Having a haircut")
    sleep(randint(1,6))
    print(id, ": Done")

def barber():
    global barber_available, customers_waiting, mutex
    print("Barber: Opening shop")
    while (1):
        mutex.acquire()
        if len(customers_waiting) > 0:
            id = customers_waiting.pop(0)
            mutex.release()
            barber_available = 0
            cutting_hair(id)
        else:
            print("Barber: Going to sleep")
            while barber_available == 0:
                condition.wait()
            mutex.release()
            print("Barber: Woke up")

def customer(id):
    global barber_available, customers_waiting, mutex, seats
    mutex.acquire()
    if len(customers_waiting) == seats:
        print(id, ": Leaving")
        mutex.release()
    else:
        customers_waiting.append(id)
        print(id, ": In waiting-room: ", customers_waiting)
        barber_available = 1
        condition.notify()
        mutex.release()

if __name__ == "__main__":
    b = Thread(target=barber)
    b.start()

    sleep(1)
    customers = []
    for i in range(6):
        c = Thread(target=customer, args=(i,))
        customers.append(c)
    while len(customers) > 0:
        c = customers.pop()
        c.start()
        sleep(randint(1,3))




