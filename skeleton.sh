#!/bin/bash

# Function to display usage
usage() {
    echo "Usage: $0 --dir <directory_name>"
    exit 1
}

# Check if arguments are provided
if [ $# -eq 0 ]; then
    usage
fi

# Parse arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
        --dir) DIR="$2"; shift ;;
        *) echo "Unknown parameter passed: $1"; usage ;;
    esac
    shift
done

# Check if directory name is provided
if [ -z "$DIR" ]; then
    usage
fi

# Create directory
mkdir -p "$DIR"

# Change to the new directory
cd "$DIR" || exit

# Create empty files
touch main.go example.txt input.txt

# Print success message
echo "Created aoc skeleton files in directory: $DIR"
echo "Files created:"
echo "- main.go"
echo "- example.txt"
echo "- input.txt"