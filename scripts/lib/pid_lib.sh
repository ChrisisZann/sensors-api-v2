# #!/bin/bash

# create_pid_file()
# {
#     echo "bname in pid function is $BNAME"
#     INPUT_PID=$1
#     print_log_message "INFO" "Started create_pid_file function" 
#     print_par "create_pid_file \$\$" "$$"
#     print_par "input pid" "$INPUT_PID"

#     if [ -f "$PID_FILE" ]
#     then
#         # check if old process is still running
#         PID=$(cat $PID_FILE)
#         ps -p $PID > /dev/null 2>&1
#         if [ $? -ne 0 ]
#         then
#             print_log_message "ERROR" "Daemon is already running"
#             exit 0
#         else
#             echo "$$" > "$PID_FILE"
#             if [ $? -ne 0 ]
#             then
#                 print_log_message "ERROR" "Could not update PID file"
#                 exit 1
#             else
#                 print_log_message "INFO" "PID file updated succesfully, PID=$(cat $PID_FILE)"
#             fi
#         fi
#     else
#         # Create new PID file
#         print_log_message "INFO" "Creating PID file: $PID_FILE"
#         echo "$$" > "$PID_FILE"
#         if [ $? -ne 0 ]
#         then
#             print_log_message "ERROR" "Could not create PID file"
#             exit 1
#         else
#             print_log_message "INFO" "PID file created succesfully, PID=$(cat $PID_FILE)"
#         fi
#     fi
#     print_log_message "INFO" "Finished create_pid_file function" 

#     # 

# }

# delete_pid_file()
# {
#     if [ -z $1 ]
#     then
#         PID_FILE=$$
#     else
#         PID_FILE=$1
#     fi
#     # print_par INPUT_PID $INPUT_PID
#     # TODO work with pid_id not pid_file
    
#     print_log_message "INFO" "Started delete_pid_file function"
#     print_par INPUT_PID_FILE $INPUT_PID_FILE

#     if [ -f $INPUT_PID_FILE ]
#     then
#         print_log_message "INFO" "deleting PID file: $PID_FILE"
#         rm $PID_FILE
#     else
#         print_log_message "INFO" "PID file not found: $PID_FILE "
#     fi

#     # print_log_message "INFO" "Daemon $BNAME has stopped"
# }


# stop()
# {
#     INPUT_PID=$$
#     print_log_message "INFO" "Started stop function" 
#     # print_par "PID_FILE" "$PID_FILE"
#     # print_par "input pid" "$INPUT_PID"

#     # if [ -f $PID_FILE ]
#     # then
#     #     # check if old process is still running
#     #     PID=$(cat $PID_FILE)
#     #     print_log_message "INFO" "found pid file" 

#     #     ps -p $PID > /dev/null 2>&1
#     #     if [ $? -ne 0 ]
#     #     then
#     #         print_log_message "INFO" "Daemon is running, going to stop it"
#     #         rm $PID_FILE
#     #     else
#     #         print_log_message "WARN" "Daemon is not running"
#     #     fi
#     # else
#     #     print_log_message "WARN" "Daemon is not running"
#     # fi
#     # print_log_message "INFO" "$PID_FILE not found"
# }