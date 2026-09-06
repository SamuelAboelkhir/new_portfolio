package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/a-h/templ"
)

func GenerateHtml(component templ.Component, fileName string) error {
	file, err := os.Create(fileName + ".html")
	fmt.Println("FILENAME ISSSSSS", fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := component.Render(context.Background(), file); err != nil {
		return err
	}
	return nil
}
