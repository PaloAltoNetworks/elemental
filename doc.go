// Package elemental provides a set of low level structure and interfaces
// to manage a model generated from some Monolithe specifications set.
//
// If you are not familiar with with Monolithe, please read https://github.com/aporeto-inc/monolithe
//
// Elementa also provide various structures to handle multiple tings like an errors, events, requests/responses
// authentication, validation and much more.
// Most of those structure and interface are useless by themselves, but can be used by various systems, like
// https://github.com/aporeto-inc/bahamut, that provides a way to write API servers in no time, or https://github.com/aporeto-inc/manipulate
// which provides interface to store an elemental model in a database, or to send them to a API Server and so on.
package elemental
