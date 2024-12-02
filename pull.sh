mkdir $1
curl -H "Cookie: session=$SESSION_ID" "https://adventofcode.com/2024/day/$1/input" > $1/input.txt
cp -r template/* $1/
