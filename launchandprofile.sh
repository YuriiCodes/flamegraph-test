#!/bin/bash

# Start your application in the background
./main &
APP_PID=$!


# Run the sample command with the captured PID
sample $APP_PID -f output.prof

# Additional commands if needed
# ...
