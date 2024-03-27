# input="reverse_iter3_132" 
# while read -r line
# do
#   echo "$line"
# done < "$input"


#!/bin/bash

# Example: Nested loops to print multiplication table



# Outer loop for rows
for (( row = 1; row <= 5; row++ )); do
    # Inner loop for columns
    for (( col = 1; col <= 5; col++ )); do
        # Calculate and print multiplication
        product=$((row * col))
        echo -n "$product "
    done
    # Move to the next line after each row
    echo
done
