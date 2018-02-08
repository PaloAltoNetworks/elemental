package main

import (
	"fmt"
	"os"
	"path"

	"github.com/Sirupsen/logrus"
	"github.com/aporeto-inc/elemental/cmd/elegen/versions"
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

	version := fmt.Sprintf("%s - %s", versions.ProjectVersion, versions.ProjectSha)
	cmd := regolithe.NewCommand(generatorName, generatorDescription, version, attributeNameConverter, attributeTypeConverter, generationName, generator)

	if err := cmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Error during generation")
	}
}

func generator(sets []*spec.SpecificationSet, out string) error {

	set := sets[0]

	outFolder := path.Join(out, "elemental")
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
