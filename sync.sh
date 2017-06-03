#!/bin/bash

echo "Sync script started"
echo

SRC=$(pwd)
DEST=root@107.174.69.165:/root/go/src/github.com/gnanakeethan/
if [ "$SRC" != "" ] ; then
    echo "Syncing $SRC to $DEST"

    /usr/local/bin/fswatch -r --latency=0.1 -0 $SRC | xargs -0 -n 1 -I{} /usr/local/bin/rsync -rav -S --force --delete -e "/usr/bin/ssh -o StrictHostKeyChecking=yes" --exclude="conf" --exclude=".git" --exclude=".env" --exclude="*.tmp" $SRC $DEST


fi
