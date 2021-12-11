#!/bin/bash
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

# tool name?
DIRNAME=$(basename "${DIR}")
if [ "${DIRNAME}" == "go-project-template-cli" ]; then
    read -p "Tool name: " name
    TOOL_NAME=${name}
else
    read -p "Tool name [${DIRNAME}]: " name
    TOOL_NAME=${name:-${DIRNAME}}
fi

read -p "Go module name, like github.com/jojomi/${TOOL_NAME}: " MODULE_NAME

DEFAULT_REMOTE_URL=https://${MODULE_NAME}
read -p "Git remote URL, type \"skip\" to leave it untouched [${DEFAULT_REMOTE_URL}]: " url
GIT_URL=${url:-${DEFAULT_REMOTE_URL}}


# create required directories if needed
echo "$(tput bold)Creating empty directories...$(tput sgr0)"
mkdir --parents bin tmp

# configure go modules
echo "$(tput bold)Setting up go modules...$(tput sgr0)"
sed "s,^module .*,module ${MODULE_NAME}," -i go.mod
go mod tidy

# fix .gitignore (delete last 3 lines)
echo "$(tput bold)Fixing .gitignore...$(tput sgr0)"
sed -n -e :a -e '1,3!{P;N;D;};N;ba' -i .gitignore

# fix build info
echo "$(tput bold)Setting data in build_info.go...$(tput sgr0)"
sed "s,ToolName =.*,ToolName = \"${TOOL_NAME}\"," -i build_info.go

# set git remote url
if [ "$url" != "skip" ]; then
    set +e
    echo "$(tput bold)Setting git origin URL to ${GIT_URL}...$(tput sgr0)"
    git remote set-url origin "${GIT_URL}"
    set -e
fi

# update module dependencies
go mod tidy

# test build
make build
make clean

printf "\n"
printf "\n"
echo "$(tput bold)> If everything looks good, this script can be removed now.$(tput sgr0)"
printf "\n"
printf "\n"
