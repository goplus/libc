#ifndef _C2GO_SYSCALL_ARCH_H
#define _C2GO_SYSCALL_ARCH_H

long __syscall0(long n);
long __syscall1(long n, long a);
long __syscall2(long n, long a, long b);
long __syscall3(long n, long a, long b, long c);
long __syscall4(long n, long a, long b, long c, long d);
long __syscall5(long n, long a, long b, long c, long d, long e);
long __syscall6(long n, long a, long b, long c, long d, long e, long f);

long __syscall0_r1(long n);
long __syscall1_r1(long n, long a);
long __syscall2_r1(long n, long a, long b);
long __syscall3_r1(long n, long a, long b, long c);
long __syscall4_r1(long n, long a, long b, long c, long d);
long __syscall5_r1(long n, long a, long b, long c, long d, long e);
long __syscall6_r1(long n, long a, long b, long c, long d, long e, long f);

#endif
