#include <stdio.h>
#include <pthread.h>
#include <semaphore.h>

int s[] = {7, 2, 8, -9, 4, 0};
int results[3];
sem_t sem_write, sem_read;

void* sum(void* arg) {
    int* slice = (int*) arg;
    int sum = 0;
    for (int i = 0; i < 3; i++) {
        sum += slice[i];
    }

    sem_wait(&sem_write);
    results[2] = sum;
    sem_post(&sem_read);

    return NULL;
}

int main() {
    pthread_t t1, t2;

    sem_init(&sem_write, 0, 3); // buffer size of 3
    sem_init(&sem_read, 0, 0);

    pthread_create(&t1, NULL, sum, s);
    pthread_create(&t2, NULL, sum, s + 3);

    sem_wait(&sem_write);
    results[0] = 66;
    sem_post(&sem_read);

    sem_wait(&sem_read);
    int x = results[0];
    sem_wait(&sem_read);
    int y = results[1];
    sem_wait(&sem_read);
    int z = results[2];

    printf("%d %d %d\n", x, y, x + y);
    printf("%d\n", z);

    pthread_join(t1, NULL);
    pthread_join(t2, NULL);

    sem_destroy(&sem_write);
    sem_destroy(&sem_read);

    return 0;
}
