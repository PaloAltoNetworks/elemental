package main

import (
	"os"
	"path"

	"github.com/Sirupsen/logrus"
	"github.com/aporeto-inc/regolithe"
	"github.com/aporeto-inc/regolithe/spec"
	"golang.org/x/sync/errgroup"
)

const (
	generatorName        = "elegen"
	generatorDescription = "Generate a Go model based on elemental."
	generationName       = "elemental"
)

func main() {

	cmd := regolithe.NewCommand(generatorName, generatorDescription, attributeNameConverter, attributeTypeConverter, generationName, generator)
	if err := cmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Error during generation")
	}
}

func generator(set *spec.SpecificationSet) error {

	outFolder := path.Join(set.Configuration.Output, "elemental")
	if err := os.MkdirAll(outFolder, 0755); err != nil && !os.IsExist(err) {
		return err
	}

	var g errgroup.Group

	g.Go(func() error { return writeIdentitiesRegistry(set, outFolder) })
	g.Go(func() error { return writeRelationshipsRegistry(set, outFolder) })

	for _, s := range set.Specifications() {
		func(restName string) {
			g.Go(func() error { return writeModel(set, restName, outFolder) })
		}(s.RestName)
	}

	return g.Wait()
}
