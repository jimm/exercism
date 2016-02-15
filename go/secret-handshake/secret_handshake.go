package secret

func Handshake(n int) []string {
	if n <= 0 {
		return nil
	}
	if n&0x1ff == 0 {
		return nil
	}

	// Yes, I could use a lookup table. Personally I think that's overkill
	// for this.
	var instructions []string
	if n&1 == 1 {
		instructions = append(instructions, "wink")
	}
	if n&2 == 2 {
		instructions = append(instructions, "double blink")
	}
	if n&4 == 4 {
		instructions = append(instructions, "close your eyes")
	}
	if n&8 == 8 {
		instructions = append(instructions, "jump")
	}
	if n&16 == 16 {
		instructions = reverse(instructions)
	}
	return instructions
}

func reverse(ss []string) []string {
	for i := len(ss)/2 - 1; i >= 0; i-- {
		opp := len(ss) - 1 - i
		ss[i], ss[opp] = ss[opp], ss[i]
	}
	return ss
}
