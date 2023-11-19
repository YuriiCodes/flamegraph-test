#!/bin/bash

# Start your application in the background
./bin/unoptimised-main &
APP_PID=$!


# Run the sample command with the captured PID
sample $APP_PID -f unoptimised-output.prof

# Additional commands if needed
# ...
