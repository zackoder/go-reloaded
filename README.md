# This program is a simple text compilation, editing, and auto-correction tool.*

# How to run the program:

# Clone the folder go-reloaded:

# git clone https://learn.zone01oujda.ma/git/zbessado/go-reloaded.git

# Navigate to the go-reloaded folder:

# cd go-reloaded/corrector_project

# Run the command:
# go run main.go <input_file_name> <output_file_name>
# Replace <input_file_name> with the name of the file containing the text you want to modify, and <output_file_name> with the name of the file where you want to store the modified text.

# EX : 
# ubuntu: go-reloaded/corrector_project$ echo "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair." > sample.txt

# ubuntu: go-reloaded/corrector_project$ go run exe.go sample.txt res.txt 

# ubuntu: go-reloaded/corrector_project$ cat result.txt

# It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.