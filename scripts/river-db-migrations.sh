#!/bin/sh

# Dump river queue migrations and format them to dbmate standards

# Get the directory of the current script
script_dir=$(dirname "$(realpath "$0")")

MIGRATIONS_FILE="${script_dir}/last-river.version"
echo $MIGRATIONS_FILE
#Get the last version of migrations
LAST_MIGRATION_VERSION=$(cat "$MIGRATIONS_FILE")

NEXT_MIGRATION_VERSION=$(($LAST_MIGRATION_VERSION+1))

# If there are no migrations, exit
river migrate-get --line main --version $NEXT_MIGRATION_VERSION --up > /dev/null 2>&1
if [ $? -ne 0 ]; then
  echo "Migration v$NEXT_MIGRATION_VERSION not found"
  exit 1
fi

timestamp=$(date +%Y%m%d%H%M%S)
filename="${DBMATE_MIGRATIONS_DIR}/${timestamp}_${NEXT_MIGRATION_VERSION}_river_queue.sql"

# Create/recreate empty file
echo > "$filename"

# Create comment to indicate it's the up command
echo "-- migrate:up\n" > "$filename"

# Dump the up queries for that version
river migrate-get --line main --version $NEXT_MIGRATION_VERSION --up >> "$filename"

# Create comment to indicate it's the down command
echo "\n\n-- migrate:down\n" >> "$filename"

# Dump the down queries for that version
river migrate-get --line main --version $NEXT_MIGRATION_VERSION --down >> "$filename"

# Store next version
echo $NEXT_MIGRATION_VERSION >"$MIGRATIONS_FILE"
