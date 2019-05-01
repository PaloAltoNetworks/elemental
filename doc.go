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

// Package elemental provides a set of interfaces and structures used to manage a model generated from a
// Regolithe Specifications Set.
//
// If you are not familiar with with Regolithe, please read https://github.com/aporeto-inc/regolithe.
//
// Elemental is the basis of Bahamut (https://github.com/aporeto-inc/bahamut) and Manipulate
// (https://github.com/aporeto-inc/manipulate).
//
// The main interface it provides is the Identifiable. This interface must be implemented by all object of a model.
// It allows to identify an object from its Identity (which is a name and category) and by its identifier. It also
// embeds the Versionable interface that allows to retrieve the current version of the model. The Identifiables
// interface must be implemented by lists managing a collection of Identifiable entities.
//
// The ModelManager is an interface to perform lookup on Identities, Relationships between them and also allow to
// instantiate objects based on their Identity.
//
// Elemental also contains some Request/Response structures representing various Operation on Identifiable or
// Identifiables as well as a bunch of validators to enforce specification constraints on attributes like max length,
// pattern etc.
// There is also an Event structure that can be used to notify clients of the the result of an Operation sent through
// a Request.
//
// Elemental is mainly an abstract package and cannot really be used by itself. You must use the provided command
// (elegen) to generate an Elemental Model from a Regolithe Specification Set.
package elemental // import "go.aporeto.io/elemental"
