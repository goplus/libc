#include <stdlib.h>

void __stdio_exit();
void *__memrchr(const void *m, int c, size_t n);

void __stdio_exit_needed() {
	__stdio_exit();
}

void *memrchr(const void *m, int c, size_t n) {
	return __memrchr(m, c, n);
}
