package main

//go:generate go-bindata -pkg static -o static/bindata.go templates

import (
	"fmt"
	"os"
	"path"

	"github.com/aporeto-inc/elemental/cmd/elegen/versions"
	"github.com/aporeto-inc/regolithe"
	"github.com/aporeto-inc/regolithe/spec"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

const (
	generatorName        = "elegen"
	generatorDescription = "Generate a Go model based on elemental."
	generationName       = "elemental"
)

func main() {

	version := fmt.Sprintf("%s - %s", versions.ProjectVersion, versions.ProjectSha)
	cmd := regolithe.NewCommand(generatorName, generatorDescription, version, attributeNameConverter, attributeTypeConverter, generationName, generator)

	cmd.PersistentFlags().Bool("public", false, "If set to true, only exposed attributes and public objects will be generated")

	if err := cmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func generator(sets []spec.SpecificationSet, out string) error {

	set := sets[0]
	publicMode := viper.GetBool("public")
	outFolder := path.Join(out, "elemental")
	if err := os.MkdirAll(outFolder, 0755); err != nil && !os.IsExist(err) {
		return err
	}

	var g errgroup.Group

	g.Go(func() error { return writeIdentitiesRegistry(set, outFolder, publicMode) })
	g.Go(func() error { return writeRelationshipsRegistry(set, outFolder, publicMode) })

	for _, s := range set.Specifications() {
		func(restName string) {
			g.Go(func() error { return writeModel(set, restName, outFolder, publicMode) })
		}(s.Model().RestName)
	}

	return g.Wait()
}
