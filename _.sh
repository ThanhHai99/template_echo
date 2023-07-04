#!/bin/bash
thisBranch=$(git branch | sed -n -e 's/^\* \(.*\)/\1/p')
git fetch --all
git add .
git commit -m "$(date +%F-%H:%M:%S)"
git push origin $thisBranch