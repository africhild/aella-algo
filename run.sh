#!/bin/bash
projects=("kola-go" "kola-js" "kola-rs")
if [ -z "$1" ]; then
    echo "question name required"
    exit 1
fi
export question=$(echo "$1" | iconv -t ascii//TRANSLIT | sed -E 's/[~\^]+//g' | sed -E 's/[^a-zA-Z0-9]+/-/g' | sed -E 's/^-+\|-+$//g' | sed -E 's/^-+//g' | sed -E 's/-+$//g' | tr A-Z a-z)

# if [ ! -e "/run.sh" ] > /dev/null; then
#     echo "you must be in the root directory of the project"
#     exit 1
# fi
# mkdir -p "./questions/$question"
mkdir -p "./questions/$question/solutions"
echo "## $1" > "./questions/$question/index.md"
  if [ "$2" ]; then
    echo -e "\n$2" >> "./questions/$question/index.md"
  fi
for  i in $projects; do
    mkdir -p "./questions/$question/solutions/$i"
    if [ -d "./templates/$i" ]; then
    # echo "./templates/$i"
    cp -n -R "./templates/$i/" "./questions/$question/solutions/$i"
    fi
done
echo "cd ./questions/$question/solutions/${projects[0]}"
cd "./questions/$question/solutions/${projects[0]}"
# gotests -all -w -parallel .