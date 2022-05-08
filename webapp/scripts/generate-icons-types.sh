#!/usr/bin/env bash

icons=$(grep \\.icon < ./src/icomoon/style.css | cut -d: -f1 | cut -d ' ' -f1 | cut -c7- | uniq | sort)
# converted_icons=$(echo $icons | xargs -n 1 -I {} echo {} | perl -pe 's/(^|-)./uc($&)/ge;s/_//g' | awk -F'[/-]' '{print $1$2}' | xargs -n 1 -I {} echo {})

echo export enum Icons {
for icon in $icons
do
  echo \ \ $(echo $icon | perl -pe 's/(^|_)./uc($&)/ge;s/_//g' | awk -F- '{gsub(FS,"")} 1' | xargs -n 1 -I {} echo {}) = \"$icon\",
done
echo }
