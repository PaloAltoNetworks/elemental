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
// embeds the Versionable interface that allows to retrieve the current version of the model.
//
// Elemental also contains an Identifiables interface that must be implemented to manage a collection of Identifiable
// entities.
package elemental
