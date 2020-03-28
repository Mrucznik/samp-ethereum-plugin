#pragma warning disable 239
#pragma warning disable 214

#define RUN_TESTS

#include <a_samp>
#include <YSI\y_testing>

#include "../../ethereum-payments.inc"

main() {
    //
}

Test:RunTest() {
    new ret = MyFunction();
    printf("ret: %d", ret);
    ASSERT(ret == 0);
}
