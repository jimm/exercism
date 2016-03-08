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
	return shiftString(s, c.distance)
}

func (c shift) Decode(s string) string {
	return shiftString(s, -c.distance)
}

func shiftString(s string, distance int) string {
	s = normalize(s)
	encoded := make([]rune, len(s))
	for i, r := range s {
		ord := int(r) + distance
		if ord > int('z') {
			ord -= 26
		} else if ord < int('a') {
			ord += 26
		}
		encoded[i] = rune(ord)
	}
	return string(encoded)
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
	s = normalize(s)
	encoded := make([]rune, len(s))
	keyOffset := 0
	for i, r := range s {
		distance := direction * (int(key[keyOffset]) - int('a'))
		ord := int(r) + distance
		if ord > int('z') {
			ord -= 26
		} else if ord < int('a') {
			ord += 26
		}
		encoded[i] = rune(ord)
		keyOffset++
		if keyOffset >= len(key) {
			keyOffset = 0
		}
	}
	return string(encoded)
}

// ================ common ================

func normalize(s string) string {
	re := regexp.MustCompile("[^a-z]+")
	return string(re.ReplaceAll([]byte(strings.ToLower(s)), []byte("")))
}
