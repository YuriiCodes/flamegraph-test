#!/bin/bash

# Start your application in the background
./bin/optimised-main &
APP_PID=$!


# Run the sample command with the captured PID
sample $APP_PID -f optimised-output.prof

# Additional commands if needed
# ...
