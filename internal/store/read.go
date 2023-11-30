package store

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/xor111xor/gutenberg-ts/internal/models"
)

// Read file and parse to models
func ReadDict(fileName string) ([]models.Word, error) {
	var word models.Word
	var words []models.Word

	dic_file, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer dic_file.Close()

	scanner := bufio.NewScanner(dic_file)

	for scanner.Scan() {
		s := strings.Fields(scanner.Text())

		// Scan through line
		// if !scanner.Scan() {
		// 	break
		// }

		// Parse line
		word.Number, err = strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}

		word.Original = s[1]

		float_num, err := strconv.ParseFloat(s[2], 64)
		if err != nil {
			panic(err)
		}
		word.Frequentcy = int64(float_num)

		words = append(words, word)
	}

	return words, nil
}
