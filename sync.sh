#!/bin/bash

echo "Sync script started"
echo

SRC=$(pwd)
DEST=ubuntu@52.70.31.64:/home/ubuntu/src/
if [ "$SRC" != "" ] ; then
    echo "Syncing $SRC to $DEST"

    /usr/local/bin/fswatch -r --latency=0.1 -0 $SRC | xargs -0 -n 1 -I{} /usr/local/bin/rsync -rav -S --force --delete -e "/usr/bin/ssh -i /Users/gnanakeethan/.ssh/id_aws_use.pem -o StrictHostKeyChecking=yes" --exclude="vendor" --exclude="node_modules" --exclude="bower_components"  --exclude ".idea" --exclude="storage" --exclude=".env" --exclude="*.tmp" $SRC $DEST


fi
