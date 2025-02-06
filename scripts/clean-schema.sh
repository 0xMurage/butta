#!/usr/bin/env sh

# Script to remove schema namespaces from the exported pg_dump emitted by db-mate
# This script requires $DB_SCHEMA and $DBMATE_SCHEMA_FILE to be defined on environment


dir1=$(cd "$(dirname "$0")" && cd .. && pwd) #the parent of current script directory
dir2=$DBMATE_SCHEMA_FILE

if [ -z "${dir2}" ]; then
  dir2='./db/schema.sql'
fi


#check if db-mate schema file path is absolute
if [ "${dir2%"${dir2#?}"}" = "/" ]; then
    # if yes, that will be it
    schema_path=$dir2
else
    # else, join with the current script directory.
    # first ensure dir1 ends with a single slash if it doesn't already
    if [ "${dir1}" != "${dir1%/}" ]; then
        dir1="${dir1%/}"
    fi
    schema_path="$dir1/$dir2"
fi

schema_path=$(realpath ${schema_path})

#THE ACTUAL CLEANUP COMMANDS

#remove the headers, we only need schema for reference
sed '1,/SET default_table_access_method = heap;/d'  "$schema_path" > schema.tmp && mv schema.tmp "$schema_path"

# remove schema comments
sed -r -i '' "s/Schema: ${DB_SCHEMA};//g"  "$schema_path"

# remove schema prefix from queries
sed -r -i '' "s|(${DB_SCHEMA}).||"  "$schema_path"
