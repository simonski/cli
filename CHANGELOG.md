# CHANGELOG

## 2022-09-19
Bugfix in getFileExistsOrDefault - incorrectly exit 1 when shoudl return the default filename.

## 2022-09-11

Added IS_INTERACTIVE, IS_VERBOSE, IS_EXIT_ON_ERROR instance variables.
Modifed GetCommand() to take from 0 (previously was 1)
Added Shift()
Added Flatten()

## 2022-08-30

GetIntOrEnvOrDefault(...)
GetStringOrEnvOrDefault(...)

## 2022-08-03

go1.19
staticcheck

## 2022-05-08

Added `GetStringFromSetOrDefeault`, `GetStringFromSetOrDie`

## 2022-04-24

Moved to own repository.
