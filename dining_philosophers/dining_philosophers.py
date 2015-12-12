from threading import Thread, Lock
from time import sleep

num_philosophers = 5
philosophers = []
forks = []

def philosopher(idx):
    left_fork = idx
    right_fork = (idx-1)%num_philosophers

    while (1):
        forks[right_fork].acquire()
        forks[left_fork].acquire()
        print("Philosopher: ", idx, " eating, forks: ", left_fork, right_fork)
        sleep(2)
        forks[right_fork].release()
        forks[left_fork].release()
        sleep(1)

if __name__ == "__main__":
    for i in range(num_philosophers):
        P = Thread(target=philosopher, args=(i,))
        philosophers.append(P)
        forks.append(Lock())
    for P in philosophers:
        P.start()

