#define a_ll a_ll
int a_ll(volatile int *p);

#define a_ll_p a_ll_p
void *a_ll_p(volatile void *p);

#define a_sc a_sc
int a_sc(volatile int *p, int v);

#define a_sc_p a_sc_p
int a_sc_p(volatile int *p, void *v);

#define a_cas a_cas
int a_cas(volatile int *p, int t, int s);

#define a_cas_p a_cas_p
void *a_cas_p(volatile void *p, void *t, void *s);

#define a_swap a_swap
int a_swap(volatile int *p, int v);
