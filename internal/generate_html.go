package internal

import (
	"context"
	"os"

	"github.com/a-h/templ"
)

func GenerateHtml(component templ.Component, fileName string) error {
	file, err := os.Create(fileName + ".html")
	if err != nil {
		return err
	}
	defer file.Close()

	if err := component.Render(context.Background(), file); err != nil {
		return err
	}
	return nil
}
