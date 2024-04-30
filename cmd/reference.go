/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/arkan501/wut/reference"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

func createReferenceCmd() *cobra.Command {
	// make sure the proto directory exists
	entries, err := os.ReadDir("./reference")
	// throw an error if it doesn't
	if err != nil {
		log.Fatal("Error reading directory", err)
	}

	var referenceCmd *cobra.Command

	// iterate through the entries in the directory
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".pb") {
			referenceName := strings.TrimSuffix(entry.Name(), ".pb")
			referenceCmd = generateReferenceCmd(referenceName)
		}
	}

	// if there are no references, create a default command that tells the user
	// how to create a new reference.
	if referenceCmd == nil {
		referenceCmd = &cobra.Command{
			Use:   "reference",
			Short: "reference",
			Long: `    This is the default reference command, please create a new reference
    using:

    wut new-ref [referenceName]`,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("This is the reference command, create a new " +
					"reference with:\n\twut new-ref [referenceName]")
			},
		}
	}

	return referenceCmd
}

func generateReferenceCmd(referenceName string) *cobra.Command {
	referenceCmd := &cobra.Command{
		Use:   referenceName,
		Short: fmt.Sprintf("%s reference", referenceName),
		Run: func(cmd *cobra.Command, args []string) {

			add, err := cmd.Flags().GetBool("add")
			if err != nil {
				log.Fatal("Error getting flag", err)
			}

            list, err := cmd.Flags().GetBool("list")
            if err != nil {
                log.Fatal("Error getting flag", err)
            }

			if add {
				addTopic(referenceName)
			} else if list {
				listReference(referenceName)
			}
		},
	}
	// The flags for this command are defined here so that the default command
	// has no flags.

	// Flags that I would like to implement would be the following:
	// 1. --string, -s, "string to search for"
	// 2. --tag, -t, "tag to search for"
	// 3. --list, -l, "list of all topics stored in the current reference"
	referenceCmd.Flags().BoolP("add", "a", false, "add a new topic to the current reference")
    referenceCmd.Flags().BoolP("list", "l", false, "list all topics under the current reference")
	// 5. --remove, -r, "remove a topic from the current reference"
	return referenceCmd
}

func addTopic(referenceName string) {
	fileName := "/reference/" + referenceName + ".pb"

	// Open up the template file
	original, err := os.Open("template.txt")
	if err != nil {
		log.Fatal("could not open file", err)
	}
	defer original.Close()

	// Create a new temporary file from the template
	tempFile, err := os.CreateTemp("", "template")
	if err != nil {
		log.Fatal("could not create temp file", err)
	}
	defer tempFile.Close()
    defer os.Remove(tempFile.Name())

	// Copy over the original template to the temporary one
	_, err = io.Copy(tempFile, original)
	if err != nil {
		log.Fatal("could not copy to temp file", err)
	}
	fmt.Println("copied template to temp file")

	// Open temp file in system text editor for editing
	err = editTopic(tempFile.Name())
	if err != nil {
		log.Fatal("There was an error closing the editor: ", err)
	}
	fmt.Println("tempfile opened by system editor")

	categories := readTemp(tempFile.Name())

    data := serialize(categories)

    appendReference(fileName, data)
}

// TODO: fix the editTopic function to work with editors other than neovim
func editTopic(tempFile string) error {
	editorCmd := exec.Command("nvim", tempFile)
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr

	return editorCmd.Run()
}

func readTemp(fileName string) []strings.Builder {
	category := make([]strings.Builder, 0)
	categoryIndex := -1

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("could not open temp file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			categoryIndex++
			category = append(category, strings.Builder{})
			continue
		}

		if categoryIndex >= 0 {
			if category[categoryIndex].Len() != 0 {
				category[categoryIndex].WriteString("\n")
			}
			category[categoryIndex].WriteString(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error reading file", err)
	}

	return category
}

func serialize(category []strings.Builder) []byte {
    // convert the tags category to a slice of strings
    tagField := strings.Split(category[3].String(), ",")

    // convert the categories into a topic message
	topic := reference.Topic{
		Comment:     category[0].String(),
		Snippet:     category[1].String(),
		Description: category[2].String(),
		Tags:        tagField,
	}

    // add the topic to the reference
    ref := reference.Reference{}
    ref.Topics = append(ref.Topics, &topic)

    // Marshal the data
	data, err := proto.Marshal(&ref)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}
    
    return data
}

// TODO: move things around so that serialize is done after appendReference.
func appendReference(fileName string, data []byte) {

    currentData, err := os.ReadFile(fileName)
    if err != nil {
        log.Fatal("could not read file", err)
    }

    var currentReference reference.Reference

    err = proto.Unmarshal(currentData, &currentReference)
    if err != nil {
        log.Fatal("unmarshaling error:", err)
    }

    // var newTopic reference.Topic
}

func listReference(referenceName string) {
    fileName := "/reference/" + referenceName + ".pb"

    data, err := os.ReadFile(fileName)
    if err != nil {
        log.Fatal("could not read file", err)
    }

    var currentReference reference.Reference
    err = proto.Unmarshal(data, &currentReference)
    if err != nil {
        log.Fatal("unmarshaling error:", err)
    }

    for _, topic := range currentReference.Topics {
        fmt.Println("// Comment\n", topic.GetComment())
        fmt.Println("// Snippet\n", topic.GetSnippet())
        fmt.Println("// Description\n", topic.GetDescription())
        fmt.Println("// Tags\n", strings.Join(topic.GetTags(), ", "))
    }
}

func init() {
	rootCmd.AddCommand(createReferenceCmd())
}
