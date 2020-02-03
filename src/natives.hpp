/*
# natives.hpp

Contains all the `PAWN_NATIVE_DECL` for native function declarations.
*/

#ifndef ETHEREUM_PAYMENTS_NATIVES_H
#define ETHEREUM_PAYMENTS_NATIVES_H

#include <string>

#include <amx/amx2.h>

#include "common.hpp"

namespace Natives {
cell MyFunction(AMX* amx, cell* params);
}

#endif
