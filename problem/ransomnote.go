package problem

func canConstruct(ransomNote string, magazine string) bool {
	hasWord := map[byte]int{}
	for i := 0; i < len(magazine); i++ {
		hasWord[magazine[i]]++
	}
	for z := 0; z < len(ransomNote); z++ {
		num, ok := hasWord[ransomNote[z]]
		if !ok {
			return false
		}
		if num == 0 {
			return false
		}
		hasWord[ransomNote[z]]--
	}
	return true
}
