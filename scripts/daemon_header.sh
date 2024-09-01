#!/bin/bash

STARTTIME=$(date +'%Y%m%d_%H%M%S')

# Setup Directories
export LIB_DIR="./lib"
export LOG_DIR="${HOME}/logs/${PROC_NAME}_${STARTTIME}"
export WORK_DIR="./work/${STARTTIME}"

export BNAME="$(basename -s ".sh" $0)"

# Setup Filenames 
export LOG_FILE="$LOG_DIR/${BNAME}_$$.log"

# Include Libs
. "$LIB_DIR/print_lib.sh"
. "$LIB_DIR/work_lib.sh"



