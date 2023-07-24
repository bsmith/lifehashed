#!/bin/bash

declare -a CFLAGS LDFLAGS

CC=clang++
LDFLAGS=()
CFLAGS=(-Wall -W -O3 -mcpu=native -flto)

read -ra libpng_cflags < <(pkg-config --cflags libpng)
read -ra libpng_libs < <(pkg-config --libs libpng)

CFLAGS+=("${libpng_cflags[@]}")
LDFLAGS+=("${libpng_libs[@]}")

echo "CFLAGS=${CFLAGS[*]}"
echo "LDFLAGS=${LDFLAGS[*]}"

runcmd() {
    echo "$*"
    "$@"
}

runcmd ${CC} "${CFLAGS[@]}" "${LDFLAGS[@]}" -o life_to_png life_to_png.cc

