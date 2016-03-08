package cipher

import (
	"regexp"
	"strings"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type shift struct {
	distance int
}

type vigenere struct {
	key string
}

// ================ Caesar ================

func NewCaesar() Cipher {
	return shift{3}
}

// ================ Shift ================

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}
	return shift{distance}
}

func (c shift) Encode(s string) string {
	return munge(s, func() int { return c.distance })
}

func (c shift) Decode(s string) string {
	return munge(s, func() int { return -c.distance })
}

// ================ Vigenere ================

func NewVigenere(key string) Cipher {
	if key == "" || normalize(key) != key || len(key) < 3 {
		return nil
	}
	return vigenere{key}
}

func (c vigenere) Encode(s string) string {
	return vigenereize(s, c.key, 1)
}

func (c vigenere) Decode(s string) string {
	return vigenereize(s, c.key, -1)
}

func vigenereize(s, key string, direction int) string {
	keyOffset := 0
	return munge(s, func() int {
		distance := direction * (int(key[keyOffset]) - int('a'))
		keyOffset++
		if keyOffset >= len(key) {
			keyOffset = 0
		}
		return distance
	})
}

// ================ common ================

// munge encodes/decodes s. For each rune in s, the distance func is called
// to see how far to shift that rune.
func munge(s string, distance func() int) string {
	s = normalize(s)
	encoded := make([]rune, len(s))
	for i, r := range s {
		encoded[i] = shiftRune(r, distance())
	}
	return string(encoded)
}

// normalize returns s stripped of all but a-z and made lower case.
func normalize(s string) string {
	re := regexp.MustCompile("[^a-z]+")
	return string(re.ReplaceAll([]byte(strings.ToLower(s)), []byte("")))
}

// shiftRune shifts r by distance and clamps it within the range ['a', 'z'].
func shiftRune(r rune, distance int) rune {
	ord := int(r) + distance
	if ord > int('z') {
		ord -= 26
	} else if ord < int('a') {
		ord += 26
	}
	return rune(ord)
}
