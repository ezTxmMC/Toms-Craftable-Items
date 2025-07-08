#!/bin/bash

read -p "Gib die Version ein: " version
read -p "Handelt es sich um eine Dev-Version? (y/n): " isdev

filename="Toms-Craftable-Items-${version}"
if [[ "$isdev" == "y" || "$isdev" == "Y" ]]; then
  filename="${filename}-dev"
fi
filename="${filename}.zip"

zip -r "$filename" . -x "build.sh"

echo "Archiv erstellt: $filename"

