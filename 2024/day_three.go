// --- Mull It Over ---
// https://adventofcode.com/2024/day/3
//
// "Our computers are having issues, so I have no idea if we have any Chief
// Historians in stock! You're welcome to check the warehouse, though," says
// the mildly flustered shopkeeper at the North Pole Toboggan Rental Shop. The
// Historians head out to take a look.
//
// The shopkeeper turns to you. "Any chance you can see why our computers are
// having issues again?"
//
// The computer appears to be trying to run a program, but its memory
// (your puzzle input) is corrupted. All of the instructions have been jumbled
// up!
//
// It seems like the goal of the program is just to multiply some numbers. It
// does that with instructions like mul(X,Y), where X and Y are each 1-3 digit
// numbers. For instance, mul(44,46) multiplies 44 by 46 to get a result of
// 2024. Similarly, mul(123,4) would multiply 123 by 4.
//
// However, because the program's memory has been corrupted, there are also
// many invalid characters that should be ignored, even if they look like part
// of a mul instruction. Sequences like mul(4*, mul(6,9!, ?(12,34), or
// mul ( 2 , 4 ) do nothing.
//
// For example, consider the following section of corrupted memory:
//
// xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
//
// Only the four highlighted sections are real mul instructions. Adding up the
// result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5).
//
// Scan the corrupted memory for uncorrupted mul instructions. What do you get
// if you add up all of the results of the multiplications?

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Regular expression that matches mul(123,123).
var mulRegex = regexp.MustCompile(`mul\([-]*[0-9]{1,3},[-]*[0-9]{1,3}\)`)

func reduceMultiply(s string) (float64, error) {
	n := len(s)

	var total float64 = 0
	head := 0
	tail := 0

	for ; tail < n; tail++ {
		if s[tail] == '(' {
			head = tail + 1
		}

		if s[tail] == ',' {
			num1, err := strconv.ParseInt(s[head:tail], 10, 64)
			if err != nil {
				return total, fmt.Errorf(
					"failed to parse '%s' as integer: %w",
					s[head:tail],
					err,
				)
			}

			total = float64(num1)

			tail++
			head = tail
		}

		if s[tail] == ')' {
			num2, err := strconv.ParseInt(s[head:tail], 10, 64)
			if err != nil {
				return total, fmt.Errorf(
					"failed to parse '%s' as integer: %w",
					s[head:tail],
					err,
				)
			}

			total *= float64(num2)

			break
		}
	}

	return total, nil
}

func mullItOver(input string) (float64, error) {
	var total float64 = 0

	// Sanitize input of all white space.
	input = strings.ReplaceAll(input, " ", "")

	if len(input) == 0 {
		return total, nil
	}

	tokens := mulRegex.FindAllString(input, -1)

	for _, token := range tokens {
		v, err := reduceMultiply(token)
		if err != nil {
			return total, fmt.Errorf(
				"failed to parse multiply statement '%s': %w",
				token,
				err,
			)
		}

		total += v
	}

	return total, nil
}

