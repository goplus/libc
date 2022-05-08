#ifndef _C2GO_SYSCALL_H
#define _C2GO_SYSCALL_H

#ifndef _C2GO_SYSCALL_ARCH_H
#include "syscall_arch.h"
#endif

#undef _INTERNAL_SYSCALL_H
#include "../../src/internal/syscall.h"

// -----------------------------------------------------------------------

#define __SYSCALL_R1_DISP(b,...) \
    __SYSCALL_CONCAT(__SYSCALL_CONCAT(b,__SYSCALL_NARGS(__VA_ARGS__)),_r1)(__VA_ARGS__)

#define __syscall1_r1(n,a) __syscall1_r1(n,__scc(a))
#define __syscall2_r1(n,a,b) __syscall2_r1(n,__scc(a),__scc(b))
#define __syscall3_r1(n,a,b,c) __syscall3_r1(n,__scc(a),__scc(b),__scc(c))
#define __syscall4_r1(n,a,b,c,d) __syscall4_r1(n,__scc(a),__scc(b),__scc(c),__scc(d))
#define __syscall5_r1(n,a,b,c,d,e) __syscall5_r1(n,__scc(a),__scc(b),__scc(c),__scc(d),__scc(e))
#define __syscall6_r1(n,a,b,c,d,e,f) __syscall6_r1(n,__scc(a),__scc(b),__scc(c),__scc(d),__scc(e),__scc(f))
#define __syscall7_r1(n,a,b,c,d,e,f,g) __syscall7_r1(n,__scc(a),__scc(b),__scc(c),__scc(d),__scc(e),__scc(f),__scc(g))

#define __syscall_r1(...) __SYSCALL_R1_DISP(__syscall,__VA_ARGS__)

#undef syscall
#define syscall(...) __syscall_r1(__VA_ARGS__)

// -----------------------------------------------------------------------

#ifdef SYS_open
#define __sys_open2_r1(x,pn,fl) __syscall2_r1(SYS_open, pn, (fl)|O_LARGEFILE)
#define __sys_open3_r1(x,pn,fl,mo) __syscall3_r1(SYS_open, pn, (fl)|O_LARGEFILE, mo)
#define __sys_open_cp2_r1(x,pn,fl) __syscall_cp2_r1(SYS_open, pn, (fl)|O_LARGEFILE)
#define __sys_open_cp3_r1(x,pn,fl,mo) __syscall_cp3_r1(SYS_open, pn, (fl)|O_LARGEFILE, mo)
#else
#define __sys_open2_r1(x,pn,fl) __syscall3_r1(SYS_openat, AT_FDCWD, pn, (fl)|O_LARGEFILE)
#define __sys_open3_r1(x,pn,fl,mo) __syscall4_r1(SYS_openat, AT_FDCWD, pn, (fl)|O_LARGEFILE, mo)
#define __sys_open_cp2_r1(x,pn,fl) __syscall_cp3_r1(SYS_openat, AT_FDCWD, pn, (fl)|O_LARGEFILE)
#define __sys_open_cp3_r1(x,pn,fl,mo) __syscall_cp4_r1(SYS_openat, AT_FDCWD, pn, (fl)|O_LARGEFILE, mo)
#endif

#define __sys_open_r1(...) __SYSCALL_R1_DISP(__sys_open,,__VA_ARGS__)

#undef sys_open
#define sys_open(...) __sys_open_r1(__VA_ARGS__)

// -----------------------------------------------------------------------

#endif
