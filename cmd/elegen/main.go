// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

//go:generate go-bindata -pkg static -o static/bindata.go templates

import (
	"fmt"
	"os"
	"path"

	"go.aporeto.io/elemental/cmd/elegen/versions"
	"go.aporeto.io/regolithe"
	"go.aporeto.io/regolithe/spec"
	"golang.org/x/sync/errgroup"

	"go.aporeto.io/elemental/cmd/internal/genopenapi3"
)

const (
	generatorName        = "elegen"
	generatorDescription = "Generate a Go model based on elemental."
	generationName       = "elemental"
)

func main() {

	// will be initialized later
	var (
		genType    string
		publicMode bool
	)

	generator := func(sets []spec.SpecificationSet, out string) error {
		switch genType {
		case "openapi3":
			return genopenapi3.GeneratorFunc(sets, out, publicMode)
		case "", "elemental":
			return genElemental(sets, out, publicMode)
		default:
			return fmt.Errorf("unhandled generation type: '%s'", genType)
		}
	}

	version := fmt.Sprintf("%s - %s", versions.ProjectVersion, versions.ProjectSha)
	cmd := regolithe.NewCommand(
		generatorName,
		generatorDescription,
		version,
		attributeNameConverter,
		attributeTypeConverter,
		generationName,
		generator,
	)

	cmd.PersistentFlags().BoolVar(
		&publicMode,
		"public",
		false,
		"If set to true, only exposed attributes and public objects will be generated",
	)
	cmd.PersistentFlags().StringVarP(
		&genType,
		"gen-type",
		"g",
		"elemental",
		"The desired type of what needs to be generated. Possible choices are: [elemental openapi3]",
	)

	if err := cmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err) // nolint
		os.Exit(1)
	}
}

func genElemental(sets []spec.SpecificationSet, out string, publicMode bool) error {

	set := sets[0]
	outFolder := path.Join(out, "elemental")
	if err := os.MkdirAll(outFolder, 0750); err != nil && !os.IsExist(err) {
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
