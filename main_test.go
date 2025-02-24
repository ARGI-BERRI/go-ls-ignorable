package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func Test_listFiles(t *testing.T) {
	mainTextFile, err := os.Open("resources/a.txt")
	if err != nil {
		log.Fatal(err)
	}
	mainBinFile, err := os.Open("resources/b.bin")
	if err != nil {
		log.Fatal(err)
	}
	subTextFile, err := os.Open("resources/sub/a.txt")
	if err != nil {
		log.Fatal(err)
	}
	subBinFile, err := os.Open("resources/sub/b.bin")
	if err != nil {
		log.Fatal(err)
	}
	type args struct {
		gitignore string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "all files",
			args: args{
				gitignore: "",
			},
			want: []string{
				mainTextFile.Name(), mainBinFile.Name(), subTextFile.Name(), subBinFile.Name(),
			},
		},
		{
			name: "only .txt files",
			args: args{
				gitignore: "*.bin\n",
			},
			want: []string{
				mainTextFile.Name(), subTextFile.Name(),
			},
		},
	}

	for _, tt := range tests {
		gitignore, _ := os.CreateTemp("", "")
		_, _ = gitignore.WriteString(tt.args.gitignore)

		t.Run(tt.name, func(t *testing.T) {
			got, _ := listFiles(gitignore.Name(), "./resources")

			if !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("listFiles() got = %v, want %v", got, tt.want)
			}
		})

		_ = gitignore.Close()
	}
}
