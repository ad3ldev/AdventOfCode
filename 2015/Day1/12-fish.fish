#!/usr/bin/env fish

set file "input.txt"

if test -f $file
    set open_count (cat $file | sed 's/[^(]//g' | tr -d '\n' | wc -c)
    set close_count (cat $file | sed 's/[^)]//g' | tr -d '\n' | wc -c)
    set count (math $open_count - $close_count)
    echo $count
else
    echo "Error: File not found"
end
