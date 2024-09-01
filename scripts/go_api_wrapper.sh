#!/bin/bash

export PROC_NAME="SENSORS_API_V2"

. "./daemon_header.sh"

init_log "$LOG_DIR" "$LOG_FILE"
echo "LOG: ${LOG_FILE}"

{

    go work use .

    go run ../api


} >> "$LOG_FILE" 2>&1
