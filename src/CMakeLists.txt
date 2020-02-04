include(AMXConfig)
include(AddSAMPPlugin)

set(SAMP_SDK_ROOT "${PROJECT_SOURCE_DIR}/lib/samp-plugin-sdk")
find_package(SAMPSDK REQUIRED)

include_directories(
	${SAMPSDK_INCLUDE_DIR}
)

add_samp_plugin(ethereum-payments
	${SAMPSDK_DIR}/amxplugin.cpp
	${SAMPSDK_DIR}/amxplugin2.cpp
	${SAMPSDK_DIR}/amx/getch.c
	common.hpp
	main.cpp
	impl.cpp
	impl.hpp
	natives.cpp
	natives.hpp
	plugin.def
)

add_library(ethereum-payments-plugin STATIC IMPORTED)
target_link_libraries(ethereum-payments "${CMAKE_SOURCE_DIR}/gosrc/ethereum-payments-plugin.a")

