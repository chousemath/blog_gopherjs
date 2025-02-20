#!/bin/bash

GREEN='\033[0;32m'
LIGHTCYAN='\033[1;36m'
NC='\033[0m'

echo "Type your commit comment, followed by [ENTER]:"
# the read command stores the user input as a string
read comment
# the -e command enables the backslash necessary for colored text
echo -e "${GREEN}commit comment has been loaded${NC}"
/Users/jo/go/src/github.com/gopherjs/gopherjs/gopherjs build main/main.go -o main.js
echo -e "${GREEN}Go code compiled to Javascript${NC}"
# just a spacer
echo -e "${LIGHTCYAN}............${NC}"
git add .
echo -e "${GREEN}git add . was performed${NC}"
echo -e "${LIGHTCYAN}............${NC}"
eval "git commit -am \"${comment}\""
echo -e "${GREEN}git commit -am was performed${NC}"
echo -e "${LIGHTCYAN}............${NC}"
git push
echo -e "${GREEN}git push was performed${NC}"
echo -e "${LIGHTCYAN}............${NC}"
echo -e "${GREEN}deployment complete!${NC}"
