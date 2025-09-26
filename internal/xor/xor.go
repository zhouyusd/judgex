package xor

func normalizeKey(key []byte) []byte {
	if len(key) == 0 {
		return []byte("default-32-byte-secret-key-here!!")
	}
	if len(key) >= 32 {
		return key[:32]
	}
	result := make([]byte, 32)
	for i := range result {
		result[i] = key[i%len(key)]
	}
	return result
}
