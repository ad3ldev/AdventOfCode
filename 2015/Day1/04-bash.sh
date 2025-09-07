#!/bin/bash

foo="$(cat "input.txt")"
count=0

for ((i = 0; i < ${#foo}; i++)); do
    char="${foo:$i:1}"
    if [ "$char" == "(" ]; then
        ((count += 1))
    elif [ "$char" == ")" ]; then
        ((count -= 1))
    fi
done

echo $count
