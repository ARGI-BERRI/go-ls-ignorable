package main

import (
	"context"
	ignore "github.com/sabhiram/go-gitignore"
	"github.com/urfave/cli/v3"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var gitignorePath string
	var walkDir string

	cmd := &cli.Command{
		Name:  "go-ls-ignorable",
		Usage: "Simple ls command reimplementation, but filterable with .gitignore",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "gitignore",
				Aliases:     []string{"g"},
				Usage:       "Path to .gitignore file",
				Destination: &gitignorePath,
			},
			&cli.StringFlag{
				Name:        "dir",
				Aliases:     []string{"d"},
				Value:       ".",
				Usage:       "Path to directory to list files",
				Destination: &walkDir,
			},
		},
		Action: func(ctx context.Context, command *cli.Command) error {
			files, err := listFiles(gitignorePath, walkDir)

			if err != nil {
				log.Fatal(err)
				return err
			}

			for _, file := range files {
				println(file)
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		return
	}
}

func listFiles(gitignorePath string, walkDir string) ([]string, error) {
	gitignore, err := ignore.CompileIgnoreFile(gitignorePath)

	if err != nil {
		return nil, err
	}

	var files []string

	err = filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !gitignore.MatchesPath(path) {
			files = append(files, filepath.ToSlash(path))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, err
}
