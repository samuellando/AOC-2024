#!/bin/sh

# Check if day argument is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <day-number>"
    exit 1
fi

# Ensure SESSION_ID is set
if [ -z "$SESSION_ID" ]; then
    echo "Error: SESSION_ID is not set. Export it with 'export SESSION_ID=your-session-id'"
    exit 1
fi

mkdir $1
curl -H "Cookie: session=$SESSION_ID" "https://adventofcode.com/2024/day/$1/input" > $1/input.txt
cp -r template/* $1/
