package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	EMPTY_SPACE  	= ""
	TWO_SPACES		= "  "
)


type PromptPackage struct {
	Name			string			`json:"name"`
	Description		string			`json:"description"`
	Tags			map[string]int		`json:"tags"`
	Version			string			`json:"version"`
	Author			string			`json:"author"`
	License			string			`json:"license"`
	Repository		string			`json:"repository"`
	Dependencies	map[string]int		`json:"dependencies"`
}


var initCmd = &cobra.Command{
	Use: "init",
	Short: "creates a skeleton prompt package",
	Long: "creates a skeleton prompt package",
	Run: func(cmd *cobra.Command, args []string) {

		initPromptPackage()

	},
}


func init() {

} // init


func checkExists() bool {
  return false
} // checkExists


func initPromptPackage() {

  if checkExists() {
	fmt.Println("Prompt package already exists, use -f to override.")
  } else {

	p := PromptPackage{}

	j, err := json.MarshalIndent(p, TWO_SPACES, EMPTY_SPACE)

	if err != nil {
		log.Println(err)
	} else {

		os.WriteFile("prompt.json", j, 0644)

	}

  }

} // initPromptPackage
