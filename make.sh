#!/bin/sh
# Simple build script instead of makefile
clear

COMPILER="gcc"
CFILE="bytebeat_sdl2"
CFLAGS="-Wall -Wextra -pedantic -O3 -lSDL2main -lSDL2"

COMMAND="${COMPILER} ${CFILE}.c ${CFLAGS} -o ${CFILE}"
echo ${COMMAND}
${COMMAND}
