# musl
echo pleval-static
src/musl/pleval-static.exe >src/musl/pleval-static.err || echo pleval-static failed

# functional static
echo argv-static
src/functional/argv-static.exe >src/functional/argv-static.err || echo argv-static failed
echo basename-static
src/functional/basename-static.exe >src/functional/basename-static.err || echo basename-static failed
echo clocale_mbfuncs-static
src/functional/clocale_mbfuncs-static.exe >src/functional/clocale_mbfuncs-static.err || echo clocale_mbfuncs-static failed
echo clock_gettime-static
src/functional/clock_gettime-static.exe >src/functional/clock_gettime-static.err || echo clock_gettime-static failed
echo crypt-static
src/functional/crypt-static.exe >src/functional/crypt-static.err || echo crypt-static failed
echo dirname-static
src/functional/dirname-static.exe >src/functional/dirname-static.err || echo dirname-static failed
echo env-static
src/functional/env-static.exe >src/functional/env-static.err || echo env-static failed
echo fdopen-static
src/functional/fdopen-static.exe >src/functional/fdopen-static.err || echo fdopen-static failed
echo fnmatch-static
src/functional/fnmatch-static.exe >src/functional/fnmatch-static.err || echo fnmatch-static failed
echo fscanf-static
src/functional/fscanf-static.exe >src/functional/fscanf-static.err || echo fscanf-static failed
echo fwscanf-static
src/functional/fwscanf-static.exe >src/functional/fwscanf-static.err || echo fwscanf-static failed
echo iconv_open-static
src/functional/iconv_open-static.exe >src/functional/iconv_open-static.err || echo iconv_open-static failed
echo inet_pton-static
src/functional/inet_pton-static.exe >src/functional/inet_pton-static.err || echo inet_pton-static failed
echo ipc_sem-static
src/functional/ipc_sem-static.exe >src/functional/ipc_sem-static.err || echo ipc_sem-static failed
echo mbc-static
src/functional/mbc-static.exe >src/functional/mbc-static.err || echo mbc-static failed
echo memstream-static
src/functional/memstream-static.exe >src/functional/memstream-static.err || echo memstream-static failed
echo pthread_cancel-points-static
src/functional/pthread_cancel-points-static.exe >src/functional/pthread_cancel-points-static.err || echo pthread_cancel-points-static failed
echo pthread_mutex-static
src/functional/pthread_mutex-static.exe >src/functional/pthread_mutex-static.err || echo pthread_mutex-static failed
echo pthread_tsd-static
src/functional/pthread_tsd-static.exe >src/functional/pthread_tsd-static.err || echo pthread_tsd-static failed
echo qsort-static
src/functional/qsort-static.exe >src/functional/qsort-static.err || echo qsort-static failed
echo random-static
src/functional/random-static.exe >src/functional/random-static.err || echo random-static failed
echo search_hsearch-static
src/functional/search_hsearch-static.exe >src/functional/search_hsearch-static.err || echo search_hsearch-static failed
echo search_insque-static
src/functional/search_insque-static.exe >src/functional/search_insque-static.err || echo search_insque-static failed
echo search_lsearch-static
src/functional/search_lsearch-static.exe >src/functional/search_lsearch-static.err || echo search_lsearch-static failed
echo search_tsearch-static
src/functional/search_tsearch-static.exe >src/functional/search_tsearch-static.err || echo search_tsearch-static failed
echo sem_init-static
src/functional/sem_init-static.exe >src/functional/sem_init-static.err || echo sem_init-static failed
echo sem_open-static
src/functional/sem_open-static.exe >src/functional/sem_open-static.err || echo sem_open-static failed
echo setjmp-static
src/functional/setjmp-static.exe >src/functional/setjmp-static.err || echo setjmp-static failed
echo snprintf-static
src/functional/snprintf-static.exe >src/functional/snprintf-static.err || echo snprintf-static failed
echo sscanf-static
src/functional/sscanf-static.exe >src/functional/sscanf-static.err || echo sscanf-static failed
echo sscanf_long-static
src/functional/sscanf_long-static.exe >src/functional/sscanf_long-static.err || echo sscanf_long-static failed
echo stat-static
src/functional/stat-static.exe >src/functional/stat-static.err || echo stat-static failed
echo strftime-static
src/functional/strftime-static.exe >src/functional/strftime-static.err || echo strftime-static failed
echo string-static
src/functional/string-static.exe >src/functional/string-static.err || echo string-static failed
echo string_memcpy-static
src/functional/string_memcpy-static.exe >src/functional/string_memcpy-static.err || echo string_memcpy-static failed
echo string_memmem-static
src/functional/string_memmem-static.exe >src/functional/string_memmem-static.err || echo string_memmem-static failed
echo string_memset-static
src/functional/string_memset-static.exe >src/functional/string_memset-static.err || echo string_memset-static failed
echo string_strchr-static
src/functional/string_strchr-static.exe >src/functional/string_strchr-static.err || echo string_strchr-static failed
echo string_strcspn-static
src/functional/string_strcspn-static.exe >src/functional/string_strcspn-static.err || echo string_strcspn-static failed
echo string_strstr-static
src/functional/string_strstr-static.exe >src/functional/string_strstr-static.err || echo string_strstr-static failed
echo strtod-static
src/functional/strtod-static.exe >src/functional/strtod-static.err || echo strtod-static failed
echo strtod_long-static
src/functional/strtod_long-static.exe >src/functional/strtod_long-static.err || echo strtod_long-static failed
echo strtod_simple-static
src/functional/strtod_simple-static.exe >src/functional/strtod_simple-static.err || echo strtod_simple-static failed
echo strtof-static
src/functional/strtof-static.exe >src/functional/strtof-static.err || echo strtof-static failed
echo strtol-static
src/functional/strtol-static.exe >src/functional/strtol-static.err || echo strtol-static failed
echo strtold-static
src/functional/strtold-static.exe >src/functional/strtold-static.err || echo strtold-static failed
echo swprintf-static
src/functional/swprintf-static.exe >src/functional/swprintf-static.err || echo swprintf-static failed
echo tgmath-static
src/functional/tgmath-static.exe >src/functional/tgmath-static.err || echo tgmath-static failed
echo time-static
src/functional/time-static.exe >src/functional/time-static.err || echo time-static failed
echo tls_align-static
src/functional/tls_align-static.exe >src/functional/tls_align-static.err || echo tls_align-static failed
echo tls_init-static
src/functional/tls_init-static.exe >src/functional/tls_init-static.err || echo tls_init-static failed
echo tls_local_exec-static
src/functional/tls_local_exec-static.exe >src/functional/tls_local_exec-static.err || echo tls_local_exec-static failed
echo udiv-static
src/functional/udiv-static.exe >src/functional/udiv-static.err || echo udiv-static failed
echo ungetc-static
src/functional/ungetc-static.exe >src/functional/ungetc-static.err || echo ungetc-static failed
echo utime-static
src/functional/utime-static.exe >src/functional/utime-static.err || echo utime-static failed
echo wcsstr-static
src/functional/wcsstr-static.exe >src/functional/wcsstr-static.err || echo wcsstr-static failed
echo wcstol-static
src/functional/wcstol-static.exe >src/functional/wcstol-static.err || echo wcstol-static failed

# regression static
echo dn_expand-empty-static
src/regression/dn_expand-empty-static.exe >src/regression/dn_expand-empty-static.err || echo dn_expand-empty-static failed
echo dn_expand-ptr-0-static
src/regression/dn_expand-ptr-0-static.exe >src/regression/dn_expand-ptr-0-static.err || echo dn_expand-ptr-0-static failed
echo execle-env-static
src/regression/execle-env-static.exe >src/regression/execle-env-static.err || echo execle-env-static failed
echo fgets-eof-static
src/regression/fgets-eof-static.exe >src/regression/fgets-eof-static.err || echo fgets-eof-static failed
echo fgetwc-buffering-static
src/regression/fgetwc-buffering-static.exe >src/regression/fgetwc-buffering-static.err || echo fgetwc-buffering-static failed
echo fpclassify-invalid-ld80-static
src/regression/fpclassify-invalid-ld80-static.exe >src/regression/fpclassify-invalid-ld80-static.err || echo fpclassify-invalid-ld80-static failed
echo ftello-unflushed-append-static
src/regression/ftello-unflushed-append-static.exe >src/regression/ftello-unflushed-append-static.err || echo ftello-unflushed-append-static failed
echo getpwnam_r-crash-static
src/regression/getpwnam_r-crash-static.exe >src/regression/getpwnam_r-crash-static.err || echo getpwnam_r-crash-static failed
echo getpwnam_r-errno-static
src/regression/getpwnam_r-errno-static.exe >src/regression/getpwnam_r-errno-static.err || echo getpwnam_r-errno-static failed
echo iconv-roundtrips-static
src/regression/iconv-roundtrips-static.exe >src/regression/iconv-roundtrips-static.err || echo iconv-roundtrips-static failed
echo inet_ntop-v4mapped-static
src/regression/inet_ntop-v4mapped-static.exe >src/regression/inet_ntop-v4mapped-static.err || echo inet_ntop-v4mapped-static failed
echo inet_pton-empty-last-field-static
src/regression/inet_pton-empty-last-field-static.exe >src/regression/inet_pton-empty-last-field-static.err || echo inet_pton-empty-last-field-static failed
echo iswspace-null-static
src/regression/iswspace-null-static.exe >src/regression/iswspace-null-static.err || echo iswspace-null-static failed
echo lrand48-signextend-static
src/regression/lrand48-signextend-static.exe >src/regression/lrand48-signextend-static.err || echo lrand48-signextend-static failed
echo lseek-large-static
src/regression/lseek-large-static.exe >src/regression/lseek-large-static.err || echo lseek-large-static failed
echo malloc-0-static
src/regression/malloc-0-static.exe >src/regression/malloc-0-static.err || echo malloc-0-static failed
echo mbsrtowcs-overflow-static
src/regression/mbsrtowcs-overflow-static.exe >src/regression/mbsrtowcs-overflow-static.err || echo mbsrtowcs-overflow-static failed
echo memmem-oob-read-static
src/regression/memmem-oob-read-static.exe >src/regression/memmem-oob-read-static.err || echo memmem-oob-read-static failed
echo memmem-oob-static
src/regression/memmem-oob-static.exe >src/regression/memmem-oob-static.err || echo memmem-oob-static failed
echo mkdtemp-failure-static
src/regression/mkdtemp-failure-static.exe >src/regression/mkdtemp-failure-static.err || echo mkdtemp-failure-static failed
echo mkstemp-failure-static
src/regression/mkstemp-failure-static.exe >src/regression/mkstemp-failure-static.err || echo mkstemp-failure-static failed
echo printf-1e9-oob-static
src/regression/printf-1e9-oob-static.exe >src/regression/printf-1e9-oob-static.err || echo printf-1e9-oob-static failed
echo pthread_cancel-sem_wait-static
src/regression/pthread_cancel-sem_wait-static.exe >src/regression/pthread_cancel-sem_wait-static.err || echo pthread_cancel-sem_wait-static failed
echo printf-fmt-g-round-static
src/regression/printf-fmt-g-round-static.exe >src/regression/printf-fmt-g-round-static.err || echo printf-fmt-g-round-static failed
echo printf-fmt-g-zeros-static
src/regression/printf-fmt-g-zeros-static.exe >src/regression/printf-fmt-g-zeros-static.err || echo printf-fmt-g-zeros-static failed
echo printf-fmt-n-static
src/regression/printf-fmt-n-static.exe >src/regression/printf-fmt-n-static.err || echo printf-fmt-n-static failed
echo pthread-robust-detach-static
src/regression/pthread-robust-detach-static.exe >src/regression/pthread-robust-detach-static.err || echo pthread-robust-detach-static failed
echo pthread_condattr_setclock-static
src/regression/pthread_condattr_setclock-static.exe >src/regression/pthread_condattr_setclock-static.err || echo pthread_condattr_setclock-static failed
echo pthread_exit-cancel-static
src/regression/pthread_exit-cancel-static.exe >src/regression/pthread_exit-cancel-static.err || echo pthread_exit-cancel-static failed
echo pthread_once-deadlock-static
src/regression/pthread_once-deadlock-static.exe >src/regression/pthread_once-deadlock-static.err || echo pthread_once-deadlock-static failed
echo pthread_rwlock-ebusy-static
src/regression/pthread_rwlock-ebusy-static.exe >src/regression/pthread_rwlock-ebusy-static.err || echo pthread_rwlock-ebusy-static failed
echo putenv-doublefree-static
src/regression/putenv-doublefree-static.exe >src/regression/putenv-doublefree-static.err || echo putenv-doublefree-static failed
echo regex-backref-0-static
src/regression/regex-backref-0-static.exe >src/regression/regex-backref-0-static.err || echo regex-backref-0-static failed
echo regex-bracket-icase-static
src/regression/regex-bracket-icase-static.exe >src/regression/regex-bracket-icase-static.err || echo regex-bracket-icase-static failed
echo regex-ere-backref-static
src/regression/regex-ere-backref-static.exe >src/regression/regex-ere-backref-static.err || echo regex-ere-backref-static failed
echo regex-escaped-high-byte-static
src/regression/regex-escaped-high-byte-static.exe >src/regression/regex-escaped-high-byte-static.err || echo regex-escaped-high-byte-static failed
echo regex-negated-range-static
src/regression/regex-negated-range-static.exe >src/regression/regex-negated-range-static.err || echo regex-negated-range-static failed
echo regexec-nosub-static
src/regression/regexec-nosub-static.exe >src/regression/regexec-nosub-static.err || echo regexec-nosub-static failed
echo sigreturn-static
src/regression/sigreturn-static.exe >src/regression/sigreturn-static.err || echo sigreturn-static failed
echo scanf-bytes-consumed-static
src/regression/scanf-bytes-consumed-static.exe >src/regression/scanf-bytes-consumed-static.err || echo scanf-bytes-consumed-static failed
echo scanf-match-literal-eof-static
src/regression/scanf-match-literal-eof-static.exe >src/regression/scanf-match-literal-eof-static.err || echo scanf-match-literal-eof-static failed
echo scanf-nullbyte-char-static
src/regression/scanf-nullbyte-char-static.exe >src/regression/scanf-nullbyte-char-static.err || echo scanf-nullbyte-char-static failed
echo setvbuf-unget-static
src/regression/setvbuf-unget-static.exe >src/regression/setvbuf-unget-static.err || echo setvbuf-unget-static failed
echo sigaltstack-static
src/regression/sigaltstack-static.exe >src/regression/sigaltstack-static.err || echo sigaltstack-static failed
echo sigprocmask-internal-static
src/regression/sigprocmask-internal-static.exe >src/regression/sigprocmask-internal-static.err || echo sigprocmask-internal-static failed
echo sscanf-eof-static
src/regression/sscanf-eof-static.exe >src/regression/sscanf-eof-static.err || echo sscanf-eof-static failed
echo strverscmp-static
src/regression/strverscmp-static.exe >src/regression/strverscmp-static.err || echo strverscmp-static failed
echo syscall-sign-extend-static
src/regression/syscall-sign-extend-static.exe >src/regression/syscall-sign-extend-static.err || echo syscall-sign-extend-static failed
echo uselocale-0-static
src/regression/uselocale-0-static.exe >src/regression/uselocale-0-static.err || echo uselocale-0-static failed
echo wcsncpy-read-overflow-static
src/regression/wcsncpy-read-overflow-static.exe >src/regression/wcsncpy-read-overflow-static.err || echo wcsncpy-read-overflow-static failed
echo wcsstr-false-negative-static
src/regression/wcsstr-false-negative-static.exe >src/regression/wcsstr-false-negative-static.err || echo wcsstr-false-negative-static failed
