#!/bin/zsh

DAY=$1
YEAR=2015
if [ -n "$2" ]; then
  YEAR=$2

  if [ ! -d "$YEAR" ]; then
    mkdir "$YEAR"
  fi
fi

TEMPLATE_PREFIX=day
TEMPLATE_POSTFIX=_
TEMPLATE_FOLDER=template
INPUT_FILE=part1.input
TEST_INPUT_FILE=part1.test

cpath="$(pwd)/$YEAR/$TEMPLATE_PREFIX"
template_path="$(pwd)/$TEMPLATE_FOLDER/$TEMPLATE_PREFIX"

if [ ! -d "$cpath$DAY" ]; then
  cp -R "$template_path$TEMPLATE_POSTFIX" "$cpath$DAY"
  find "$cpath$DAY" -name "*.go" | xargs sed -i '' "s/day_/day$DAY/g"
  find "$cpath$DAY" -name "*.go" | xargs sed -i '' "s/year_/$YEAR/g"
  touch "$cpath$DAY/$INPUT_FILE"
  touch "$cpath$DAY/$TEST_INPUT_FILE"
else
  echo "Year $YEAR Day $DAY already exists."
fi
