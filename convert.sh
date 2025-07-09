#!/bin/bash

# Set the directory to search
TARGET_DIR="$1"

if [ -z "$TARGET_DIR" ]; then
  echo "Usage: $0 <directory>"
  exit 1
fi

# Find all JSON files recursively
find "$TARGET_DIR" -type f -name "*.json" | while read -r file; do
  # Transform crafting_shaped recipes
  jq '
    if .type == "minecraft:crafting_shaped" and (.key | type == "object") then
      .key |= with_entries(
        if (.value | type == "object") and (.value.item) then
          .value = .value.item
        else
          .
        end
      )
    # Transform crafting_shapeless recipes
    elif .type == "minecraft:crafting_shapeless" and (.ingredients | type == "array") then
      .ingredients |= map(
        if (type == "object") and (.item) then
          .item
        else
          .
        end
      )
    else
      .
    end
  ' "$file" > "$file.tmp" && mv "$file.tmp" "$file"
done

