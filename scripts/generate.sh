#!/bin/zsh

DAY=$1
YEAR=2023

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

T_PATH="$(pwd)/$TEMPLATE_FOLDER/$TEMPLATE_PREFIX$TEMPLATE_POSTFIX"
G_PATH="$(pwd)/$YEAR/$TEMPLATE_PREFIX$DAY"

if [ ! -d "$G_PATH" ]; then
  cp -R "$T_PATH" "$G_PATH"                                                       #copy template folder in a new day folder
  find "$G_PATH" -name "*.go" -print0 | xargs -0 sed -i "s/day_/day$DAY/g"     #replace day placeholders with the current day
  find "$G_PATH" -name "*.go" -print0 | xargs -0 sed -i "s/year_/$YEAR/g"     #replace year placeholders with the current year
  touch "$G_PATH/$INPUT_FILE"                                                     #add input file
  touch "$G_PATH/$TEST_INPUT_FILE"                                                #add test input file
else
  echo "Year $YEAR Day $DAY already exists."
fi
