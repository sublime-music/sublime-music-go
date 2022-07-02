#! /usr/bin/env bash

OUTFILE=$1
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
  cat <<EOF >> $OUTFILE
    <file preprocess="xml-stripblanks">$FILE</file>
EOF
done

cat <<EOF >> $OUTFILE
  </gresource>
</gresources>
EOF
