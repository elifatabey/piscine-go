#! /bin/bash
find -name "*.sh" | rev | cut -d "/" -f1 | sort -r | rev | cut -d "." -f1