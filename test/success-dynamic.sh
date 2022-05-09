# functional
echo argv
src/functional/argv.exe >src/functional/argv.err || echo argv failed
echo basename
src/functional/basename.exe >src/functional/basename.err || echo basename failed
echo clocale_mbfuncs
src/functional/clocale_mbfuncs.exe >src/functional/clocale_mbfuncs.err || echo clocale_mbfuncs failed
echo clock_gettime
src/functional/clock_gettime.exe >src/functional/clock_gettime.err || echo clock_gettime failed
echo crypt
src/functional/crypt.exe >src/functional/crypt.err || echo crypt failed
echo dirname
src/functional/dirname.exe >src/functional/dirname.err || echo dirname failed
echo dlopen
src/functional/dlopen.exe >src/functional/dlopen.err || echo dlopen failed
echo env
src/functional/env.exe >src/functional/env.err || echo env failed
echo fdopen
src/functional/fdopen.exe >src/functional/fdopen.err || echo fdopen failed
echo fnmatch
src/functional/fnmatch.exe >src/functional/fnmatch.err || echo fnmatch failed
echo fscanf
src/functional/fscanf.exe >src/functional/fscanf.err || echo fscanf failed
echo fwscanf
src/functional/fwscanf.exe >src/functional/fwscanf.err || echo fwscanf failed
echo iconv_open
src/functional/iconv_open.exe >src/functional/iconv_open.err || echo iconv_open failed
echo inet_pton
src/functional/inet_pton.exe >src/functional/inet_pton.err || echo inet_pton failed
echo ipc_sem
src/functional/ipc_sem.exe >src/functional/ipc_sem.err || echo ipc_sem failed
echo mbc
src/functional/mbc.exe >src/functional/mbc.err || echo mbc failed
echo memstream
src/functional/memstream.exe >src/functional/memstream.err || echo memstream failed
echo pthread_cancel
src/functional/pthread_cancel.exe >src/functional/pthread_cancel.err || echo pthread_cancel failed
echo pthread_cancel-points
src/functional/pthread_cancel-points.exe >src/functional/pthread_cancel-points.err || echo pthread_cancel-points failed
echo pthread_mutex
src/functional/pthread_mutex.exe >src/functional/pthread_mutex.err || echo pthread_mutex failed
echo pthread_tsd
src/functional/pthread_tsd.exe >src/functional/pthread_tsd.err || echo pthread_tsd failed
echo qsort
src/functional/qsort.exe >src/functional/qsort.err || echo qsort failed
echo random
src/functional/random.exe >src/functional/random.err || echo random failed
echo search_hsearch
src/functional/search_hsearch.exe >src/functional/search_hsearch.err || echo search_hsearch failed
echo search_insque
src/functional/search_insque.exe >src/functional/search_insque.err || echo search_insque failed
echo search_lsearch
src/functional/search_lsearch.exe >src/functional/search_lsearch.err || echo search_lsearch failed
echo search_tsearch
src/functional/search_tsearch.exe >src/functional/search_tsearch.err || echo search_tsearch failed
echo sem_init
src/functional/sem_init.exe >src/functional/sem_init.err || echo sem_init failed
echo sem_open
src/functional/sem_open.exe >src/functional/sem_open.err || echo sem_open failed
echo setjmp
src/functional/setjmp.exe >src/functional/setjmp.err || echo setjmp failed
echo snprintf
src/functional/snprintf.exe >src/functional/snprintf.err || echo snprintf failed
echo sscanf
src/functional/sscanf.exe >src/functional/sscanf.err || echo sscanf failed
echo sscanf_long
src/functional/sscanf_long.exe >src/functional/sscanf_long.err || echo sscanf_long failed
echo stat
src/functional/stat.exe >src/functional/stat.err || echo stat failed
echo strftime
src/functional/strftime.exe >src/functional/strftime.err || echo strftime failed
echo string
src/functional/string.exe >src/functional/string.err || echo string failed
echo string_memcpy
src/functional/string_memcpy.exe >src/functional/string_memcpy.err || echo string_memcpy failed
echo string_memmem
src/functional/string_memmem.exe >src/functional/string_memmem.err || echo string_memmem failed
echo string_memset
src/functional/string_memset.exe >src/functional/string_memset.err || echo string_memset failed
echo string_strchr
src/functional/string_strchr.exe >src/functional/string_strchr.err || echo string_strchr failed
echo string_strcspn
src/functional/string_strcspn.exe >src/functional/string_strcspn.err || echo string_strcspn failed
echo string_strstr
src/functional/string_strstr.exe >src/functional/string_strstr.err || echo string_strstr failed
echo strtod
src/functional/strtod.exe >src/functional/strtod.err || echo strtod failed
echo strtod_long
src/functional/strtod_long.exe >src/functional/strtod_long.err || echo strtod_long failed
echo strtod_simple
src/functional/strtod_simple.exe >src/functional/strtod_simple.err || echo strtod_simple failed
echo strtof
src/functional/strtof.exe >src/functional/strtof.err || echo strtof failed
echo strtol
src/functional/strtol.exe >src/functional/strtol.err || echo strtol failed
echo strtold
src/functional/strtold.exe >src/functional/strtold.err || echo strtold failed
echo swprintf
src/functional/swprintf.exe >src/functional/swprintf.err || echo swprintf failed
echo tgmath
src/functional/tgmath.exe >src/functional/tgmath.err || echo tgmath failed
echo time
src/functional/time.exe >src/functional/time.err || echo time failed
echo tls_align
src/functional/tls_align.exe >src/functional/tls_align.err || echo tls_align failed
echo tls_align_dlopen
src/functional/tls_align_dlopen.exe >src/functional/tls_align_dlopen.err || echo tls_align_dlopen failed
echo tls_init
src/functional/tls_init.exe >src/functional/tls_init.err || echo tls_init failed
echo tls_init_dlopen
src/functional/tls_init_dlopen.exe >src/functional/tls_init_dlopen.err || echo tls_init_dlopen failed
echo tls_local_exec
src/functional/tls_local_exec.exe >src/functional/tls_local_exec.err || echo tls_local_exec failed
echo udiv
src/functional/udiv.exe >src/functional/udiv.err || echo udiv failed
echo ungetc
src/functional/ungetc.exe >src/functional/ungetc.err || echo ungetc failed
echo utime
src/functional/utime.exe >src/functional/utime.err || echo utime failed
echo wcsstr
src/functional/wcsstr.exe >src/functional/wcsstr.err || echo wcsstr failed
echo wcstol
src/functional/wcstol.exe >src/functional/wcstol.err || echo wcstol failed

# math
echo acos
src/math/acos.exe >src/math/acos.err || echo acos failed
echo acosf
src/math/acosf.exe >src/math/acosf.err || echo acosf failed
echo acosh
src/math/acosh.exe >src/math/acosh.err || echo acosh failed
echo acoshf
src/math/acoshf.exe >src/math/acoshf.err || echo acoshf failed
echo acoshl
src/math/acoshl.exe >src/math/acoshl.err || echo acoshl failed
echo acosl
src/math/acosl.exe >src/math/acosl.err || echo acosl failed
echo asin
src/math/asin.exe >src/math/asin.err || echo asin failed
echo asinf
src/math/asinf.exe >src/math/asinf.err || echo asinf failed
echo asinh
src/math/asinh.exe >src/math/asinh.err || echo asinh failed
echo asinhl
src/math/asinhl.exe >src/math/asinhl.err || echo asinhl failed
echo asinhf
src/math/asinhf.exe >src/math/asinhf.err || echo asinhf failed
echo asinl
src/math/asinl.exe >src/math/asinl.err || echo asinl failed
echo atan
src/math/atan.exe >src/math/atan.err || echo atan failed
echo atan2
src/math/atan2.exe >src/math/atan2.err || echo atan2 failed
echo atan2f
src/math/atan2f.exe >src/math/atan2f.err || echo atan2f failed
echo atan2l
src/math/atan2l.exe >src/math/atan2l.err || echo atan2l failed
echo atanf
src/math/atanf.exe >src/math/atanf.err || echo atanf failed
echo atanh
src/math/atanh.exe >src/math/atanh.err || echo atanh failed
echo atanhf
src/math/atanhf.exe >src/math/atanhf.err || echo atanhf failed
echo atanhl
src/math/atanhl.exe >src/math/atanhl.err || echo atanhl failed
echo atanl
src/math/atanl.exe >src/math/atanl.err || echo atanl failed
echo cbrt
src/math/cbrt.exe >src/math/cbrt.err || echo cbrt failed
echo cbrtf
src/math/cbrtf.exe >src/math/cbrtf.err || echo cbrtf failed
echo cbrtl
src/math/cbrtl.exe >src/math/cbrtl.err || echo cbrtl failed
echo ceil
src/math/ceil.exe >src/math/ceil.err || echo ceil failed
echo ceilf
src/math/ceilf.exe >src/math/ceilf.err || echo ceilf failed
echo ceill
src/math/ceill.exe >src/math/ceill.err || echo ceill failed
echo copysign
src/math/copysign.exe >src/math/copysign.err || echo copysign failed
echo copysignf
src/math/copysignf.exe >src/math/copysignf.err || echo copysignf failed
echo copysignl
src/math/copysignl.exe >src/math/copysignl.err || echo copysignl failed
echo cos
src/math/cos.exe >src/math/cos.err || echo cos failed
echo cosf
src/math/cosf.exe >src/math/cosf.err || echo cosf failed
echo cosh
src/math/cosh.exe >src/math/cosh.err || echo cosh failed
echo coshf
src/math/coshf.exe >src/math/coshf.err || echo coshf failed
echo coshl
src/math/coshl.exe >src/math/coshl.err || echo coshl failed
echo cosl
src/math/cosl.exe >src/math/cosl.err || echo cosl failed
echo drem
src/math/drem.exe >src/math/drem.err || echo drem failed
echo dremf
src/math/dremf.exe >src/math/dremf.err || echo dremf failed
echo erf
src/math/erf.exe >src/math/erf.err || echo erf failed
echo erfc
src/math/erfc.exe >src/math/erfc.err || echo erfc failed
echo erfcf
src/math/erfcf.exe >src/math/erfcf.err || echo erfcf failed
echo erfcl
src/math/erfcl.exe >src/math/erfcl.err || echo erfcl failed
echo erff
src/math/erff.exe >src/math/erff.err || echo erff failed
echo erfl
src/math/erfl.exe >src/math/erfl.err || echo erfl failed
echo exp
src/math/exp.exe >src/math/exp.err || echo exp failed
echo exp10
src/math/exp10.exe >src/math/exp10.err || echo exp10 failed
echo exp10f
src/math/exp10f.exe >src/math/exp10f.err || echo exp10f failed
echo exp10l
src/math/exp10l.exe >src/math/exp10l.err || echo exp10l failed
echo exp2
src/math/exp2.exe >src/math/exp2.err || echo exp2 failed
echo exp2f
src/math/exp2f.exe >src/math/exp2f.err || echo exp2f failed
echo exp2l
src/math/exp2l.exe >src/math/exp2l.err || echo exp2l failed
echo expf
src/math/expf.exe >src/math/expf.err || echo expf failed
echo expl
src/math/expl.exe >src/math/expl.err || echo expl failed
echo expm1
src/math/expm1.exe >src/math/expm1.err || echo expm1 failed
echo expm1l
src/math/expm1l.exe >src/math/expm1l.err || echo expm1l failed
echo expm1f
src/math/expm1f.exe >src/math/expm1f.err || echo expm1f failed
echo fabs
src/math/fabs.exe >src/math/fabs.err || echo fabs failed
echo fabsf
src/math/fabsf.exe >src/math/fabsf.err || echo fabsf failed
echo fabsl
src/math/fabsl.exe >src/math/fabsl.err || echo fabsl failed
echo fdim
src/math/fdim.exe >src/math/fdim.err || echo fdim failed
echo fdimf
src/math/fdimf.exe >src/math/fdimf.err || echo fdimf failed
echo fdiml
src/math/fdiml.exe >src/math/fdiml.err || echo fdiml failed
echo fenv
src/math/fenv.exe >src/math/fenv.err || echo fenv failed
echo floor
src/math/floor.exe >src/math/floor.err || echo floor failed
echo floorf
src/math/floorf.exe >src/math/floorf.err || echo floorf failed
echo floorl
src/math/floorl.exe >src/math/floorl.err || echo floorl failed
echo fma
src/math/fma.exe >src/math/fma.err || echo fma failed
echo fmaf
src/math/fmaf.exe >src/math/fmaf.err || echo fmaf failed
echo fmax
src/math/fmax.exe >src/math/fmax.err || echo fmax failed
echo fmaxf
src/math/fmaxf.exe >src/math/fmaxf.err || echo fmaxf failed
echo fmaxl
src/math/fmaxl.exe >src/math/fmaxl.err || echo fmaxl failed
echo fmin
src/math/fmin.exe >src/math/fmin.err || echo fmin failed
echo fminf
src/math/fminf.exe >src/math/fminf.err || echo fminf failed
echo fminl
src/math/fminl.exe >src/math/fminl.err || echo fminl failed
echo fmod
src/math/fmod.exe >src/math/fmod.err || echo fmod failed
echo fmodf
src/math/fmodf.exe >src/math/fmodf.err || echo fmodf failed
echo fmodl
src/math/fmodl.exe >src/math/fmodl.err || echo fmodl failed
echo fpclassify
src/math/fpclassify.exe >src/math/fpclassify.err || echo fpclassify failed
echo frexp
src/math/frexp.exe >src/math/frexp.err || echo frexp failed
echo frexpf
src/math/frexpf.exe >src/math/frexpf.err || echo frexpf failed
echo frexpl
src/math/frexpl.exe >src/math/frexpl.err || echo frexpl failed
echo hypot
src/math/hypot.exe >src/math/hypot.err || echo hypot failed
echo hypotf
src/math/hypotf.exe >src/math/hypotf.err || echo hypotf failed
echo hypotl
src/math/hypotl.exe >src/math/hypotl.err || echo hypotl failed
echo ilogb
src/math/ilogb.exe >src/math/ilogb.err || echo ilogb failed
echo ilogbl
src/math/ilogbl.exe >src/math/ilogbl.err || echo ilogbl failed
echo ilogbf
src/math/ilogbf.exe >src/math/ilogbf.err || echo ilogbf failed
echo isless
src/math/isless.exe >src/math/isless.err || echo isless failed
echo j0
src/math/j0.exe >src/math/j0.err || echo j0 failed
echo j0f
src/math/j0f.exe >src/math/j0f.err || echo j0f failed
echo j1
src/math/j1.exe >src/math/j1.err || echo j1 failed
echo j1f
src/math/j1f.exe >src/math/j1f.err || echo j1f failed
echo jn
src/math/jn.exe >src/math/jn.err || echo jn failed
echo jnf
src/math/jnf.exe >src/math/jnf.err || echo jnf failed
echo ldexp
src/math/ldexp.exe >src/math/ldexp.err || echo ldexp failed
echo ldexpf
src/math/ldexpf.exe >src/math/ldexpf.err || echo ldexpf failed
echo ldexpl
src/math/ldexpl.exe >src/math/ldexpl.err || echo ldexpl failed
echo lgamma
src/math/lgamma.exe >src/math/lgamma.err || echo lgamma failed
echo lgamma_r
src/math/lgamma_r.exe >src/math/lgamma_r.err || echo lgamma_r failed
echo lgammaf
src/math/lgammaf.exe >src/math/lgammaf.err || echo lgammaf failed
echo lgammaf_r
src/math/lgammaf_r.exe >src/math/lgammaf_r.err || echo lgammaf_r failed
echo lgammal
src/math/lgammal.exe >src/math/lgammal.err || echo lgammal failed
echo lgammal_r
src/math/lgammal_r.exe >src/math/lgammal_r.err || echo lgammal_r failed
echo llrint
src/math/llrint.exe >src/math/llrint.err || echo llrint failed
echo llrintf
src/math/llrintf.exe >src/math/llrintf.err || echo llrintf failed
echo llrintl
src/math/llrintl.exe >src/math/llrintl.err || echo llrintl failed
echo llround
src/math/llround.exe >src/math/llround.err || echo llround failed
echo llroundf
src/math/llroundf.exe >src/math/llroundf.err || echo llroundf failed
echo llroundl
src/math/llroundl.exe >src/math/llroundl.err || echo llroundl failed
echo log
src/math/log.exe >src/math/log.err || echo log failed
echo log10
src/math/log10.exe >src/math/log10.err || echo log10 failed
echo log10f
src/math/log10f.exe >src/math/log10f.err || echo log10f failed
echo log10l
src/math/log10l.exe >src/math/log10l.err || echo log10l failed
echo log1p
src/math/log1p.exe >src/math/log1p.err || echo log1p failed
echo log1pf
src/math/log1pf.exe >src/math/log1pf.err || echo log1pf failed
echo log1pl
src/math/log1pl.exe >src/math/log1pl.err || echo log1pl failed
echo log2
src/math/log2.exe >src/math/log2.err || echo log2 failed
echo log2f
src/math/log2f.exe >src/math/log2f.err || echo log2f failed
echo log2l
src/math/log2l.exe >src/math/log2l.err || echo log2l failed
echo logb
src/math/logb.exe >src/math/logb.err || echo logb failed
echo logbf
src/math/logbf.exe >src/math/logbf.err || echo logbf failed
echo logbl
src/math/logbl.exe >src/math/logbl.err || echo logbl failed
echo logf
src/math/logf.exe >src/math/logf.err || echo logf failed
echo logl
src/math/logl.exe >src/math/logl.err || echo logl failed
echo lrint
src/math/lrint.exe >src/math/lrint.err || echo lrint failed
echo lrintf
src/math/lrintf.exe >src/math/lrintf.err || echo lrintf failed
echo lrintl
src/math/lrintl.exe >src/math/lrintl.err || echo lrintl failed
echo lround
src/math/lround.exe >src/math/lround.err || echo lround failed
echo lroundf
src/math/lroundf.exe >src/math/lroundf.err || echo lroundf failed
echo lroundl
src/math/lroundl.exe >src/math/lroundl.err || echo lroundl failed
echo modf
src/math/modf.exe >src/math/modf.err || echo modf failed
echo modff
src/math/modff.exe >src/math/modff.err || echo modff failed
echo modfl
src/math/modfl.exe >src/math/modfl.err || echo modfl failed
echo nearbyint
src/math/nearbyint.exe >src/math/nearbyint.err || echo nearbyint failed
echo nearbyintf
src/math/nearbyintf.exe >src/math/nearbyintf.err || echo nearbyintf failed
echo nearbyintl
src/math/nearbyintl.exe >src/math/nearbyintl.err || echo nearbyintl failed
echo nextafter
src/math/nextafter.exe >src/math/nextafter.err || echo nextafter failed
echo nextafterf
src/math/nextafterf.exe >src/math/nextafterf.err || echo nextafterf failed
echo nextafterl
src/math/nextafterl.exe >src/math/nextafterl.err || echo nextafterl failed
echo nexttoward
src/math/nexttoward.exe >src/math/nexttoward.err || echo nexttoward failed
echo nexttowardf
src/math/nexttowardf.exe >src/math/nexttowardf.err || echo nexttowardf failed
echo nexttowardl
src/math/nexttowardl.exe >src/math/nexttowardl.err || echo nexttowardl failed
echo pow
src/math/pow.exe >src/math/pow.err || echo pow failed
echo pow10
src/math/pow10.exe >src/math/pow10.err || echo pow10 failed
echo pow10f
src/math/pow10f.exe >src/math/pow10f.err || echo pow10f failed
echo pow10l
src/math/pow10l.exe >src/math/pow10l.err || echo pow10l failed
echo powf
src/math/powf.exe >src/math/powf.err || echo powf failed
echo powl
src/math/powl.exe >src/math/powl.err || echo powl failed
echo remainder
src/math/remainder.exe >src/math/remainder.err || echo remainder failed
echo remainderf
src/math/remainderf.exe >src/math/remainderf.err || echo remainderf failed
echo remainderl
src/math/remainderl.exe >src/math/remainderl.err || echo remainderl failed
echo remquo
src/math/remquo.exe >src/math/remquo.err || echo remquo failed
echo remquof
src/math/remquof.exe >src/math/remquof.err || echo remquof failed
echo remquol
src/math/remquol.exe >src/math/remquol.err || echo remquol failed
echo round
src/math/round.exe >src/math/round.err || echo round failed
echo roundf
src/math/roundf.exe >src/math/roundf.err || echo roundf failed
echo roundl
src/math/roundl.exe >src/math/roundl.err || echo roundl failed
echo rint
src/math/rint.exe >src/math/rint.err || echo rint failed
echo rintf
src/math/rintf.exe >src/math/rintf.err || echo rintf failed
echo rintl
src/math/rintl.exe >src/math/rintl.err || echo rintl failed
echo scalb
src/math/scalb.exe >src/math/scalb.err || echo scalb failed
echo scalbf
src/math/scalbf.exe >src/math/scalbf.err || echo scalbf failed
echo scalbln
src/math/scalbln.exe >src/math/scalbln.err || echo scalbln failed
echo scalblnf
src/math/scalblnf.exe >src/math/scalblnf.err || echo scalblnf failed
echo scalblnl
src/math/scalblnl.exe >src/math/scalblnl.err || echo scalblnl failed
echo scalbn
src/math/scalbn.exe >src/math/scalbn.err || echo scalbn failed
echo scalbnf
src/math/scalbnf.exe >src/math/scalbnf.err || echo scalbnf failed
echo scalbnl
src/math/scalbnl.exe >src/math/scalbnl.err || echo scalbnl failed
echo sin
src/math/sin.exe >src/math/sin.err || echo sin failed
echo sincos
src/math/sincos.exe >src/math/sincos.err || echo sincos failed
echo sincosf
src/math/sincosf.exe >src/math/sincosf.err || echo sincosf failed
echo sincosl
src/math/sincosl.exe >src/math/sincosl.err || echo sincosl failed
echo sinf
src/math/sinf.exe >src/math/sinf.err || echo sinf failed
echo sinh
src/math/sinh.exe >src/math/sinh.err || echo sinh failed
echo sinhf
src/math/sinhf.exe >src/math/sinhf.err || echo sinhf failed
echo sinhl
src/math/sinhl.exe >src/math/sinhl.err || echo sinhl failed
echo sinl
src/math/sinl.exe >src/math/sinl.err || echo sinl failed
echo sqrt
src/math/sqrt.exe >src/math/sqrt.err || echo sqrt failed
echo sqrtf
src/math/sqrtf.exe >src/math/sqrtf.err || echo sqrtf failed
echo sqrtl
src/math/sqrtl.exe >src/math/sqrtl.err || echo sqrtl failed
echo tan
src/math/tan.exe >src/math/tan.err || echo tan failed
echo tanf
src/math/tanf.exe >src/math/tanf.err || echo tanf failed
echo tanh
src/math/tanh.exe >src/math/tanh.err || echo tanh failed
echo tanhf
src/math/tanhf.exe >src/math/tanhf.err || echo tanhf failed
echo tanhl
src/math/tanhl.exe >src/math/tanhl.err || echo tanhl failed
echo tanl
src/math/tanl.exe >src/math/tanl.err || echo tanl failed
echo tgamma
src/math/tgamma.exe >src/math/tgamma.err || echo tgamma failed
echo tgammaf
src/math/tgammaf.exe >src/math/tgammaf.err || echo tgammaf failed
echo tgammal
src/math/tgammal.exe >src/math/tgammal.err || echo tgammal failed
echo trunc
src/math/trunc.exe >src/math/trunc.err || echo trunc failed
echo truncf
src/math/truncf.exe >src/math/truncf.err || echo truncf failed
echo truncl
src/math/truncl.exe >src/math/truncl.err || echo truncl failed
echo y0
src/math/y0.exe >src/math/y0.err || echo y0 failed
echo y0f
src/math/y0f.exe >src/math/y0f.err || echo y0f failed
echo y1
src/math/y1.exe >src/math/y1.err || echo y1 failed
echo y1f
src/math/y1f.exe >src/math/y1f.err || echo y1f failed
echo yn
src/math/yn.exe >src/math/yn.err || echo yn failed
echo ynf
src/math/ynf.exe >src/math/ynf.err || echo ynf failed

# regression
echo dn_expand-empty
src/regression/dn_expand-empty.exe >src/regression/dn_expand-empty.err || echo dn_expand-empty failed
echo dn_expand-ptr-0
src/regression/dn_expand-ptr-0.exe >src/regression/dn_expand-ptr-0.err || echo dn_expand-ptr-0 failed
echo execle-env
src/regression/execle-env.exe >src/regression/execle-env.err || echo execle-env failed
echo fgets-eof
src/regression/fgets-eof.exe >src/regression/fgets-eof.err || echo fgets-eof failed
echo fgetwc-buffering
src/regression/fgetwc-buffering.exe >src/regression/fgetwc-buffering.err || echo fgetwc-buffering failed
echo fpclassify-invalid-ld80
src/regression/fpclassify-invalid-ld80.exe >src/regression/fpclassify-invalid-ld80.err || echo fpclassify-invalid-ld80 failed
echo ftello-unflushed-append
src/regression/ftello-unflushed-append.exe >src/regression/ftello-unflushed-append.err || echo ftello-unflushed-append failed
echo getpwnam_r-errno
src/regression/getpwnam_r-errno.exe >src/regression/getpwnam_r-errno.err || echo getpwnam_r-errno failed
echo getpwnam_r-crash
src/regression/getpwnam_r-crash.exe >src/regression/getpwnam_r-crash.err || echo getpwnam_r-crash failed
echo iconv-roundtrips
src/regression/iconv-roundtrips.exe >src/regression/iconv-roundtrips.err || echo iconv-roundtrips failed
echo inet_ntop-v4mapped
src/regression/inet_ntop-v4mapped.exe >src/regression/inet_ntop-v4mapped.err || echo inet_ntop-v4mapped failed
echo inet_pton-empty-last-field
src/regression/inet_pton-empty-last-field.exe >src/regression/inet_pton-empty-last-field.err || echo inet_pton-empty-last-field failed
echo iswspace-null
src/regression/iswspace-null.exe >src/regression/iswspace-null.err || echo iswspace-null failed
echo lrand48-signextend
src/regression/lrand48-signextend.exe >src/regression/lrand48-signextend.err || echo lrand48-signextend failed
echo lseek-large
src/regression/lseek-large.exe >src/regression/lseek-large.err || echo lseek-large failed
echo malloc-0
src/regression/malloc-0.exe >src/regression/malloc-0.err || echo malloc-0 failed
echo mbsrtowcs-overflow
src/regression/mbsrtowcs-overflow.exe >src/regression/mbsrtowcs-overflow.err || echo mbsrtowcs-overflow failed
echo memmem-oob-read
src/regression/memmem-oob-read.exe >src/regression/memmem-oob-read.err || echo memmem-oob-read failed
echo memmem-oob
src/regression/memmem-oob.exe >src/regression/memmem-oob.err || echo memmem-oob failed
echo mkdtemp-failure
src/regression/mkdtemp-failure.exe >src/regression/mkdtemp-failure.err || echo mkdtemp-failure failed
echo mkstemp-failure
src/regression/mkstemp-failure.exe >src/regression/mkstemp-failure.err || echo mkstemp-failure failed
echo printf-1e9-oob
src/regression/printf-1e9-oob.exe >src/regression/printf-1e9-oob.err || echo printf-1e9-oob failed
echo printf-fmt-g-round
src/regression/printf-fmt-g-round.exe >src/regression/printf-fmt-g-round.err || echo printf-fmt-g-round failed
echo printf-fmt-g-zeros
src/regression/printf-fmt-g-zeros.exe >src/regression/printf-fmt-g-zeros.err || echo printf-fmt-g-zeros failed
echo printf-fmt-n
src/regression/printf-fmt-n.exe >src/regression/printf-fmt-n.err || echo printf-fmt-n failed
echo pthread_cancel-sem_wait
src/regression/pthread_cancel-sem_wait.exe >src/regression/pthread_cancel-sem_wait.err || echo pthread_cancel-sem_wait failed
echo pthread-robust-detach
src/regression/pthread-robust-detach.exe >src/regression/pthread-robust-detach.err || echo pthread-robust-detach failed
echo pthread_condattr_setclock
src/regression/pthread_condattr_setclock.exe >src/regression/pthread_condattr_setclock.err || echo pthread_condattr_setclock failed
echo pthread_exit-cancel
src/regression/pthread_exit-cancel.exe >src/regression/pthread_exit-cancel.err || echo pthread_exit-cancel failed
echo pthread_once-deadlock
src/regression/pthread_once-deadlock.exe >src/regression/pthread_once-deadlock.err || echo pthread_once-deadlock failed
echo pthread_rwlock-ebusy
src/regression/pthread_rwlock-ebusy.exe >src/regression/pthread_rwlock-ebusy.err || echo pthread_rwlock-ebusy failed
echo putenv-doublefree
src/regression/putenv-doublefree.exe >src/regression/putenv-doublefree.err || echo putenv-doublefree failed
echo regex-backref-0
src/regression/regex-backref-0.exe >src/regression/regex-backref-0.err || echo regex-backref-0 failed
echo regex-bracket-icase
src/regression/regex-bracket-icase.exe >src/regression/regex-bracket-icase.err || echo regex-bracket-icase failed
echo regex-ere-backref
src/regression/regex-ere-backref.exe >src/regression/regex-ere-backref.err || echo regex-ere-backref failed
echo regex-escaped-high-byte
src/regression/regex-escaped-high-byte.exe >src/regression/regex-escaped-high-byte.err || echo regex-escaped-high-byte failed
echo regex-negated-range
src/regression/regex-negated-range.exe >src/regression/regex-negated-range.err || echo regex-negated-range failed
echo regexec-nosub
src/regression/regexec-nosub.exe >src/regression/regexec-nosub.err || echo regexec-nosub failed
echo scanf-bytes-consumed
src/regression/scanf-bytes-consumed.exe >src/regression/scanf-bytes-consumed.err || echo scanf-bytes-consumed failed
echo scanf-match-literal-eof
src/regression/scanf-match-literal-eof.exe >src/regression/scanf-match-literal-eof.err || echo scanf-match-literal-eof failed
echo scanf-nullbyte-char
src/regression/scanf-nullbyte-char.exe >src/regression/scanf-nullbyte-char.err || echo scanf-nullbyte-char failed
echo setvbuf-unget
src/regression/setvbuf-unget.exe >src/regression/setvbuf-unget.err || echo setvbuf-unget failed
echo sigaltstack
src/regression/sigaltstack.exe >src/regression/sigaltstack.err || echo sigaltstack failed
echo sigprocmask-internal
src/regression/sigprocmask-internal.exe >src/regression/sigprocmask-internal.err || echo sigprocmask-internal failed
echo sigreturn
src/regression/sigreturn.exe >src/regression/sigreturn.err || echo sigreturn failed
echo sscanf-eof
src/regression/sscanf-eof.exe >src/regression/sscanf-eof.err || echo sscanf-eof failed
echo strverscmp
src/regression/strverscmp.exe >src/regression/strverscmp.err || echo strverscmp failed
echo syscall-sign-extend
src/regression/syscall-sign-extend.exe >src/regression/syscall-sign-extend.err || echo syscall-sign-extend failed
echo uselocale-0
src/regression/uselocale-0.exe >src/regression/uselocale-0.err || echo uselocale-0 failed
echo wcsncpy-read-overflow
src/regression/wcsncpy-read-overflow.exe >src/regression/wcsncpy-read-overflow.err || echo wcsncpy-read-overflow failed
echo wcsstr-false-negative
src/regression/wcsstr-false-negative.exe >src/regression/wcsstr-false-negative.err || echo wcsstr-false-negative failed

