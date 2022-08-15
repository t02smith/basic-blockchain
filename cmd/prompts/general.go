package prompts

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

func IntegerPrompt(label string, min, max int) (*promptui.Prompt, error) {
	if min >= max {
		return &promptui.Prompt{}, errors.New("min must be smaller than max")
	}

	return &promptui.Prompt{
		Label: fmt.Sprintf("%s (max %d): ", label, max),
		Validate: func(s string) error {
			res, err := strconv.ParseInt(s, 10, 32)
			x := int(res)
			if err != nil || x < min || x > max {
				return errors.New("invalid number")
			}
			return nil
		},
	}, nil
}
