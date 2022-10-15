// +build ignore

package main

import (
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/shurcooL/vfsgen"
)

func main() {
	var fs http.FileSystem = rice.MustFindBox("res").HTTPBox()

	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "main",
		BuildTags:    "vfs",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
