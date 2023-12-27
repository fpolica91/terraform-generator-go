
for dir in cmd pkg templates internal api; do
    if [ ! -d $dir ]; then
        mkdir -p $dir
    fi
done