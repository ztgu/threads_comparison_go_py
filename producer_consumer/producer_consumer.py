from threading import Thread, Semaphore
from random import randint
from time import sleep

buffer_max = 4
buffer = []

sema_dont_below = Semaphore(0)
sema_dont_above = Semaphore(buffer_max)

def producer():
    while (1):
        sema_dont_above.acquire()
        write_this = randint(0,50)
        buffer.append(write_this)
        sema_dont_below.release()
        print(buffer, " produced: ", write_this)
        sleep( randint(0,3) )

def consumer(id):
    sleep(1)
    while (1):
        sema_dont_below.acquire()
        read_this = buffer.pop(0)
        sema_dont_above.release()
        print(buffer, " consumer: ", id , " consumed: ", read_this)
        sleep( randint(2,7) )

if __name__ == "__main__":
    p = Thread(target=producer)
    c1 = Thread(target=consumer, args=(1,))
    c2 = Thread(target=consumer, args=(2,))

    p.start()
    c1.start()
    c2.start()

    p.join()
    c1.join()
    c2.join()
