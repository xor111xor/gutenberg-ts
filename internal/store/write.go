package store

import (
	"fmt"
	"os"

	"github.com/xor111xor/gutenberg-ts/internal/models"
)

// Write result of translate to file
func WriteFile(fileName string, words []models.Word) error {

	dic_file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer dic_file.Close()

	for i := range words {
		str := fmt.Sprintf("%v  %v  %v  %v", words[i].Number, words[i].Original, words[i].Translate, words[i].Frequentcy)
		fmt.Println(str)
		_, err := dic_file.WriteString(str + "\n")
		// err := os.(fileName, []byte(str), 0600)
		if err != nil {
			fmt.Println("this")
			return err
		}

	}
	return nil

}
