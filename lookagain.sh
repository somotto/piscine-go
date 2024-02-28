#/bin/bash
find . -name '*.sh' -exec basename {} \; | cut -d '.' -f1 | sort -r