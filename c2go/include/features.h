#ifndef _C2GO_FEATURES_H
#define _C2GO_FEATURES_H

#include "../../include/features.h"
#include "../../src/include/features.h"

#undef weak_alias
#define weak_alias(old, new)

#define ___errno_location __errno_location

#endif // _C2GO_FEATURES_H
