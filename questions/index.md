mkdir -p "./questions//solutions"
mkdir -p "./questions//solutions/"
for  i in kola-go; do
    mkdir -p "./questions//solutions/"
    if [ -f "./template/" ]; then
        cp -r "./template/" "./questions//solutions/"
    fi
done
