package elemental

import "net/http"

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
	r.Header.Add("X-Namespace", namespace)
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
