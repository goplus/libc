#ifndef _C2GO_FEATURES_H
#define _C2GO_FEATURES_H

#include "../../include/features.h"
#include "../../src/include/features.h"

#undef weak_alias
#define weak_alias(old, new)

#define ___errno_location __errno_location

#define _PTHREAD_IMPL_H // don't include old pthread_impl.h
#include "pthread_impl.h"

#endif // _C2GO_FEATURES_H
