package mixer

import (
	"fmt"
	"log"
	"strings"

	"github.com/manifoldco/promptui"

	"github.com/wkhub/wk/utils/types"
)

var Validators = map[string]func(string) error{
	"string": func(value string) error {
		return nil
	},
	"set": func(value string) error {
		return nil
	},
}

type Parameter struct {
	Key     string
	Type    string
	Prompt  string
	Label   string
	Default interface{}
	Choices []string
	Multi   bool
}

type Parameters []Parameter

func (params Parameters) PromptUser(ctx Context) Context {
	for _, param := range params {
		response, err := param.PromptUser(ctx)

		if err != nil {
			log.Fatal(err)
		}

		ctx[param.Key] = response
	}
	return ctx
}

func (param Parameter) PromptUser(ctx Context) (interface{}, error) {
	var (
		typ    string
		result string
		err    error
	)
	if param.Type == "" {
		typ = "string"
	} else {
		typ = param.Type
	}

	validator, ok := Validators[typ]
	if !ok {
		return nil, Error(fmt.Sprintf("Unknown type '%s'", typ))
	}

	var label string
	switch {
	case param.Prompt != "":
		label = param.Prompt
	case param.Label != "":
		label = param.Label
	default:
		label = strings.Title(param.Key)
	}

	if typ == "set" {
		if len(param.Choices) > 0 {
			choices := types.NewStringSet(param.Choices...)
			results := types.NewStringSet()
			// lastChoice := ""

			for {
				currentChoices := append(choices.Slice(), "OK")
				prompt := promptui.Select{
					Label: label,
					Items: currentChoices,
					// Templates: templates,
					Size: len(currentChoices),
					// Searcher:  searcher,
				}
				_, result, err = prompt.Run()
				if err != nil {
					return nil, err
				}
				if result == "OK" {
					break
				}
				results.Add(result)
				choices.Remove(result)
				if choices.IsEmpty() {
					break
				}
			}
			return results, nil
		}
	}

	if len(param.Choices) > 0 {
		if param.Multi {
			choices := param.Choices
			choices = append(choices, "OK")
			prompt := promptui.Select{
				Label: label,
				Items: choices,
				// Templates: templates,
				Size: len(param.Choices),
				// Searcher:  searcher,
			}
			results := []string{}
			for {
				_, result, err = prompt.Run()
				if err != nil {
					return nil, err
				}
				if result == "OK" {
					break
				}
				results = append(results, result)
				newChoices := []string{}
				for _, choice := range param.Choices {
					found := false
					for _, current := range results {
						if current == choice {
							found = true
						}
					}
					if !found {
						newChoices = append(newChoices, choice)
					}
				}
				newChoices = append(newChoices, "OK")
				prompt.Items = newChoices
			}
			if param.Type == "set" {
				return types.NewSet(results), nil
			}
			return results, nil

		}
		prompt := promptui.Select{
			Label: label,
			Items: param.Choices,
			// Templates: templates,
			Size: len(param.Choices),
			// Searcher:  searcher,
		}
		_, result, err = prompt.Run()
	} else {
		var prompt promptui.Prompt

		if param.Default != "" {
			defaultValue, err := ctx.Render(param.Default.(string))
			if err != nil {
				fmt.Printf("default: %v\n", err)
				return nil, err
			}
			prompt = promptui.Prompt{
				Label:    label,
				Validate: validator,
				Default:  defaultValue,
			}
		} else {
			prompt = promptui.Prompt{
				Label:    label,
				Validate: validator,
			}
		}
		result, err = prompt.Run()
	}

	return result, err
}
