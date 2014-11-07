# no idea
cmake_minimum_required (VERSION 2.8)

project (tmux-status-helpers C)
set(tmux-status-helpers_VERSION_MAJOR 0)
set(tmux-status-helpers_VERSION_MINOR 1)

set(CMAKE_C_FLAGS "-std=c11 -Wall -Wextra")

# load
add_executable(load load.c display.c)
install(TARGETS load DESTINATION bin)

add_executable(mem mem.c display.c)
install(TARGETS mem DESTINATION bin)

add_executable(io io.c display.c)
install(TARGETS io DESTINATION bin)

enable_testing()
add_test( NAME load_usage
    COMMAND load --help )
add_test( NAME load
    COMMAND load )
