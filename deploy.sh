#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

echo "Type your commit comment, followed by [ENTER]:"
read comment
echo "${GREEN}Commit comment has been loaded${NC}"
git add .
echo "${GREEN}git add . was performed${NC}"
eval "git commit -am \"${comment}\""
echo "${GREEN}git commit -am was performed${NC}"
git push
echo "${GREEN}git push was performed${NC}"
echo "${GREEN}deployment complete!${NC}"
