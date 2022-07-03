#! /usr/bin/env bash

OUTFILE=$1
shift
PREPROCESS=$1
shift
ALIAS=$1
shift
INCLUDE_FILES=$@

cat <<EOF > $OUTFILE
<?xml version="1.0" encoding="UTF-8"?>
<gresources>
  <gresource prefix="/app/sublimemusic/SublimeMusicNext">
EOF

for FILE in $INCLUDE_FILES; do
  # Remove the "resources/" part of the path from FILE
  FILE=${FILE#*/}
  echo $FILE
  alias=""
  if [ "$ALIAS" = "true" ]; then
    # get the filename of ALIAS
    alias="alias=\"$(basename $FILE)\""
  fi
  cat <<EOF >> $OUTFILE
    <file $PREPROCESS $alias>$FILE</file>
EOF
done

cat <<EOF >> $OUTFILE
  </gresource>
</gresources>
EOF
