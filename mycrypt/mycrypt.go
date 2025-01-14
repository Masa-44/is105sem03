package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyz      0123456789ABCDEFGHIJKLMONPQRSTUWXYZ ^f ^x ^e.,:; ")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
        kryptertMelding := make([]rune, len(melding))
        for i := 0; i < len(melding); i++ {
                indeks := sokIAlfabetet(melding[i], alphabet)
                if indeks+chiffer >= len(alphabet) {
                        kryptertMelding[i] = alphabet[indeks+chiffer-len(alphabet)]
                } else {
                        kryptertMelding[i] = alphabet[indeks+chiffer]
                }
        }
        return kryptertMelding
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
func Dekrypter(melding []rune, alphabet []rune, chiffer int) []rune {
        dekryptertMelding := make([]rune, len(melding))
        for i := 0; i < len(melding); i++ {
                indeks := sokIAlfabetet(melding[i], alphabet)
                if indeks-chiffer < 0 {
                        dekryptertMelding[i] = alphabet[indeks-chiffer+len(alphabet)%len(alphabet)]
                } else {
                        dekryptertMelding[i] = alphabet[(indeks-chiffer+len(alphabet))%len(alphabet)]
                }
        }
        return dekryptertMelding
}

