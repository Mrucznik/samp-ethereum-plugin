// built-in include guard removal
// just in case the user has a local dependency with the same file name
#if defined _inc_ETHEREUM_PAYMENTS
	#undef _inc_ETHEREUM_PAYMENTS
#endif
// custom include-guard to ensure we don't duplicate
#if defined _ETHEREUM_PAYMENTS_included
	#endinput
#endif
#define _ETHEREUM_PAYMENTS_included


// don't forget to change the above include guard to match the file name!

native MyFunction();
