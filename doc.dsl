#!/usr/bin/env redis-test

# redis-test DSL
# Format is: COMMAND EXPR EXPR...

# This is a comment.

# Do Redis command.
SET a 1

# Use RET to check last command result.
# Must OK, or "OK", we can ignore quotes if possible.
RET OK

GET a
RET "1"

INCR a
# Must integer 2.
RET 2

ECHO abc
# RET_PRINT prints return data for debugging.
RET_PRINT

GET nil_value
# nil is a special token.
RET nil 

MGET a nil_value
# Use [] for RESP array type.
RET ["2", nil]

# Check return data's size
# If return type is array, length is the array size.
# If return type is simple string or bulk, length is the string/bulk size.
# If return type is nil, length is 0.
# If return type is integer, error.
RET_LEN 2 
