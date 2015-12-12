from threading import Thread, Semaphore, Lock
from random import randint
from time import sleep

#semaphore = Semaphore(1)
lock = Lock()

buffer = [0,1]

def writer():
    sleep(randint(1,4))
    #semaphore.acquire()
    lock.acquire()
    buffer[0] += 1
    buffer[1] += 1
    print("Wrote to buffer, buffer[0]:", buffer[0],"buffer[1]:",buffer[1])
    #semaphore.release()
    lock.release()

def reader(id):
    sleep(randint(1,4))
    #semaphore.acquire()
    lock.acquire()
    print("Read buffer[0]:", buffer[0]," buffer[1]:", buffer[1])
    #semaphore.release()
    lock.release()

if __name__ == "__main__":
    threads = []
    for i in range(5):
        for j in range(3):
            r = Thread(target=reader, args=(j,))
            threads.append(r)
        w = Thread(target=writer)
        threads.append(w)
    for t in threads:
        t.start()
    for t in threads:
        t.join()

