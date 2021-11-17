package elemental

import (
	"fmt"
	"net/http"
	"strings"
)

// Namespacer is the interface that any namespace extraction/injection
// implementation should support.
type Namespacer interface {
	Extract(r *http.Request) (string, error)
	Inject(r *http.Request, namespace string) error
}

var (
	namespacer = Namespacer(&defaultNamespacer{})
)

type defaultNamespacer struct{}

// defaultExtractor will retrieve the namespace value from the header X-Namespace.
func (d *defaultNamespacer) Extract(r *http.Request) (string, error) {
	return r.Header.Get("X-Namespace"), nil
}

// defaultInjector will set the namespace as an HTTP header.
func (d *defaultNamespacer) Inject(r *http.Request, namespace string) error {
	if r.Header == nil {
		r.Header = http.Header{}
	}
	r.Header.Set("X-Namespace", namespace)
	return nil
}

// SetNamespacer will configure the package. It must be only called once
// and it is global for the package.
func SetNamespacer(custom Namespacer) {
	namespacer = custom
}

// GetNamespacer retrieves the configured namespacer.
func GetNamespacer() Namespacer {
	return namespacer
}

// ParentNamespaceFromString returns the parent namespace of a namespace
// It returns empty it the string is invalid
func ParentNamespaceFromString(namespace string) (string, error) {

	if namespace == "" {
		return "", fmt.Errorf("invalid empty namespace name")
	}

	if namespace == "/" {
		return "", nil
	}

	index := strings.LastIndex(namespace, "/")

	switch index {
	case -1:
		return "", fmt.Errorf("invalid namespace name")
	case 0:
		return namespace[:index+1], nil
	default:
		return namespace[:index], nil
	}
}

// IsNamespaceRelatedToNamespace returns true if the given namespace is related to the given parent
func IsNamespaceRelatedToNamespace(ns string, parent string) bool {
	return IsNamespaceParentOfNamespace(ns, parent) ||
		IsNamespaceChildrenOfNamespace(ns, parent) ||
		(ns == parent && ns != "" && parent != "")
}

// IsNamespaceParentOfNamespace returns true if the given namespace is a parent of the given parent
func IsNamespaceParentOfNamespace(ns string, child string) bool {

	if ns == "" || child == "" {
		return false
	}

	if ns == child {
		return false
	}

	if ns[len(ns)-1] != '/' {
		ns = ns + "/"
	}

	return strings.HasPrefix(child, ns)
}

// IsNamespaceChildrenOfNamespace returns true of the given ns is a children of the given parent.
func IsNamespaceChildrenOfNamespace(ns string, parent string) bool {

	if parent == "" || ns == "" {
		return false
	}

	if ns == parent {
		return false
	}

	if parent[len(parent)-1] != '/' {
		parent = parent + "/"
	}

	return strings.HasPrefix(ns, parent)
}

// NamespaceAncestorsNames returns the list of fully qualified namespaces
// in the hierarchy of a given namespace. It returns an empty
// array for the root namespace
func NamespaceAncestorsNames(namespace string) []string {

	if namespace == "/" || namespace == "" {
		return []string{}
	}

	parts := strings.Split(namespace, "/")
	sep := "/"
	namespaces := []string{}

	for i := len(parts) - 1; i >= 2; i-- {
		namespaces = append(namespaces, sep+strings.Join(parts[1:i], sep))
	}

	namespaces = append(namespaces, sep)

	return namespaces
}
