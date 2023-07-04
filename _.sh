#!/bin/bash
git fetch --all
git add .
git commit -m "$(date +%F-%H:%M:%S)"
git push
