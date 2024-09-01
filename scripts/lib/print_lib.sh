#!/bin/bash

init_log()
{
    if [[ ! -d "$LOG_DIR" ]];
    then 
        mkdir -p "$LOG_DIR"
    fi
    {
        echo "================================================================================"
        echo "======================== Log Started $(date +'%d %b %y %H:%M:%S') ========================"
        echo "================================================================================"
    } >> "$LOG_FILE" 2>&1
}

# print_log_message()
# {
#     echo "Timestamp : $(date +'%y%m%d_%H%M%S ')"
#     echo "Severity : $1"
#     echo "Message : $2"
#     echo "================================================================================"
# }
print_log_message()
{
    echo "$(date +'[%d %b %y %H:%M:%S]') : $1 : $2" >> "$LOG_FILE" 2>&1
}

print_par()
{
    echo "[$1 = $2]" >> "$LOG_FILE" 2>&1
}

print_log_divider()
{
    echo "================================================================================" >> "$LOG_FILE" 2>&1

}

# print_color()
# {
#     echo -e "\033[0;${color1};40m${1}\033[0m"
#     case $1 in

#     red)
#         STATEMENTS
#         ;;

#     PATTERN_2)
#         STATEMENTS
#         ;;

#     PATTERN_N)
#         STATEMENTS
#         ;;

#     *)
#         STATEMENTS
#         ;;
# esac


# }

print_red()
{
    echo -e "\033[0;31m${1}\033[0m"
}

print_with_color()
{
    echo -e "\[0;${1}m${2}\033[0m"
}

print_with_color_8bit()
{
    echo -e "\033[38;5;${1}m${2}\033[0m"
}

print_with_color_16bit()
{
    echo -e "\033[38;2;${1};${2};${3}m${4}\033[0m"
}