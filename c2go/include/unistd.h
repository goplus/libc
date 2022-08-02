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

char *getlogin(void);

off_t __lseek(int, off_t, int);

pid_t fork(void);

int execl(const char *path, const char *argv0, ...);

pid_t getpid(void);
// int setpgid(pid_t pid, pid_t pgid);

pid_t setsid(void);
pid_t setpgrp(void);

_Noreturn void _exit(int status);

char *getcwd(char *buf, size_t size);
ssize_t readlink(const char *restrict path, char *restrict buf, size_t bufsize);

int ioctl(int fd, int req, ...);
int dup2(int old, int new);

int pipe(int fd[2]);
int pipe2(int fd[2], int flag);

ssize_t read(int fd, void *buf, size_t count);
ssize_t write(int, const void *, size_t);
ssize_t pwrite(int, const void *, size_t, off_t);
int close(int fd);

extern char **__environ;

#endif // _C2GO_UNISTD_H
