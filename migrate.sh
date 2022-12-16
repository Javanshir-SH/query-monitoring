#!/bin/sh

echo "migrate the database at startup of app"

# Wait for few minute and run db migraiton
# while ! MY_GOOSE_COMMAND_TO_MIGRATE  2>&1; do
#    echo "migration is in progress status"
#    sleep 3
# done

cd /app/migrations/

/app/goose postgres "host=postgres user=postgres password=mysecretpassword dbname=postgres sslmode=disable" up
exec "$@"