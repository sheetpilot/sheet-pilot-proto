for entry in "$PWD"/*
do
    
    if [ -d $entry ];
    then
        echo "Generating Proto files for $entry...."
        cd $entry && protoc --go_out=plugins=grpc:. *.proto
    fi
done
