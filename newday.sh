#!/bin/bash

# ---- Args ----
PROBLEM_LINK="$1"

if [ -z "$PROBLEM_LINK" ]; then
  echo "Usage: ./new_day.sh <problem_link>"
  exit 1
fi

# ---- Date & Naming ----
DATE=$(date +"%d_%m_%Y")
DIR="dsa_${DATE}"
PKG="dsa$(date +"%d%m%Y")"

# ---- Guard ----
if [ -d "$DIR" ]; then
  echo "Folder $DIR already exists"
  exit 1
fi

# ---- Create Folder ----
mkdir "$DIR"

# ---- solution.go ----
cat <<EOF > "$DIR/solution.go"
// DSA Daily — $(date +"%Y-%m-%d")
// Problem: $PROBLEM_LINK

package $PKG

func solution() {
	// implement here
}

func main() {}
EOF

# ---- problem.md ----
cat <<EOF > "$DIR/problem.md"
## Problem
<!-- Short summary -->

## Source
$PROBLEM_LINK

## Tags
<!-- array, hash-map, etc -->
EOF

# ---- notes.md ----
cat <<EOF > "$DIR/notes.md"
### Key Insight

### Time & Space Complexity

### Trade-offs
EOF

echo "Created $DIR"
echo "Problem link: $PROBLEM_LINK"
