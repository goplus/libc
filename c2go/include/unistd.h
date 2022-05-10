#ifndef _C2GO_UNISTD_H
#define _C2GO_UNISTD_H

#include <features.h>

#define __NEED_size_t
#define __NEED_ssize_t
#define __NEED_uid_t
#define __NEED_gid_t
#define __NEED_off_t
#define __NEED_pid_t
#define __NEED_intptr_t
#define __NEED_useconds_t

#include <bits/alltypes.h>

#define STDIN_FILENO  0
#define STDOUT_FILENO 1
#define STDERR_FILENO 2

#define SEEK_SET 0
#define SEEK_CUR 1
#define SEEK_END 2

off_t __lseek(int, off_t, int);

ssize_t write(int, const void *, size_t);
ssize_t pwrite(int, const void *, size_t, off_t);

extern char **__environ;

#endif // _C2GO_UNISTD_H
