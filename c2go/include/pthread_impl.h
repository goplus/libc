#ifndef _C2GO_PTHREAD_IMPL_H
#define _C2GO_PTHREAD_IMPL_H

#include <pthread.h>
#include "atomic_arch.h"

#define pthread __pthread

struct pthread {
	int tid;
	int errno_val;
};

pthread_t __pthread_self();

void __wake(volatile void *addr, int cnt, int priv);
void __futexwait(volatile void *addr, int val, int priv);

#endif // _C2GO_PTHREAD_IMPL_H
