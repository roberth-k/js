#!/usr/bin/env bash
set -euo pipefail

out=$1


types="Null Bool Number String Object Array"
target_types="bool float64 int64 string Object Array"

tmp=$(mktemp)

function write {
    printf "$@\n" >> $tmp
}

write "package js"

for type in $types; do
    write ""
    write "func ($type) sealed() {}"
    write ""
    write "func (v $type) Type() Type {"
    write "\treturn Type$type"
    write "}"

    for target_type in $types; do
        write ""
        write "func (v $type) Is$target_type() bool {"
        if [[ $type == $target_type ]]; then
            write "\treturn true"
        else
            write "\treturn false"
        fi
        write "}"
    done

    for target_type in $target_types; do
        write ""
        write "func (v $type) ${target_type^}() $target_type {"
        if [[ $type == ${target_type^} || ($type == "Number" && ($target_type == "int64" || $target_type == "float64")) ]]; then
            write "\treturn $target_type(v)"
        else
            write "\tpanic(\"impossible\")"
        fi
        write "}"
    done
done

for target_type in $target_types; do
        write ""
        write "func (v Object) Get${target_type^}(field string) $target_type {"
        write "\treturn v.Get(field).${target_type^}()"
        write "}"
done

for target_type in $target_types; do
        write ""
        write "func (v Array) Get${target_type^}(i int) $target_type {"
        write "\treturn v.Get(i).${target_type^}()"
        write "}"
done

mv $tmp $out

