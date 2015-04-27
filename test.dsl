#!/usr/bin/env redis-test

# This is a comment.

# Do Redis command.
SET a 1

# Use RET to check last command result.
# Must OK, or "OK", we can ignore quotes if possible.
RET OK

GET a

# $(resp) is the variable storing last command result.
# Below equals: `RET "1"` or `RET 1` if the string/bulk result can be converted to integer.
ASSERT $(resp) "1"

INCR a
# Must integer 2.
RET 2

ECHO abc
# PRINT prints variable in debugging.
PRINT $(resp)

GET nil_value
# nil is a special token.
RET nil 

MGET a nil_value
# Use [] for RESP array type.
RET ["1", nil]

# If return type is array, LEN returns the array size.
# If return type is simple string or bulk, LEN returns the string/bulk size.
# If return type is nil, LEN returns 0.
# If return type is integer, error.
ASSERT LEN($(resp)) 2 

# We can use MATCH_LEN to check length too.
ASSERT_LEN $(resp) 2

# Check first element in the array result.
ASSERT $(resp[0]) "1"

# Check slice elements in the array result.
ASSERT $(resp[0:1]) ["1", nil]

SET a 123
GET a

# Check second character in the bulk result.
ASSERT $(resp[1]) "2"
# Check slice characters in the bulk result.
ASSERT $(resp[0:2]) "12"