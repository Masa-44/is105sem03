package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyz      0123456789ABCDEFGHIJKLMONPQRSTUWXYZ ^f ^x ^e.,:; ")

func Dekrypter(melding []rune, alphabet []rune, chiffer int) []rune {
        modCipher := chiffer % len(ALF_SEM03)
        shiftedAlphabet := shiftAlphabet(ALF_SEM03, -modCipher) // Reverse the direction of shift

        lookup := make(map[rune]rune)

        for i := 0; i < len(ALF_SEM03); i++ {
                lookup[shiftedAlphabet[i]] = ALF_SEM03[i] // Reverse the lookup map
        }

        dekryptertMelding := make([]rune, len(melding))
        for i := 0; i < len(melding); i++ {
                dekryptertMelding[i] = lookup[melding[i]]
        }
        return dekryptertMelding
}

func shiftAlphabet(alphabet []rune, chiffer int) []rune {
        return append(alphabet[chiffer:], alphabet[:chiffer]...)
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
        for i := 0; i < len(alfabet); i++ {
                if symbol == alfabet[i] {
                        return i
                        break
                }
        }
        return -1
}
