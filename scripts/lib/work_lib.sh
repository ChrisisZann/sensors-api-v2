#!/bin/bash

init_work()
{
    if [[ ! -d "$WORK_DIR" ]];
    then 
        mkdir -p "$WORK_DIR"
    fi
}