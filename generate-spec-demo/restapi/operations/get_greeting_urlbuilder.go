package operations

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GetGreetingURL generates an URL for the get greeting operation
type GetGreetingURL struct {
	Name *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetGreetingURL) WithBasePath(bp string) *GetGreetingURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetGreetingURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetGreetingURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/hello"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var nameQ string
	if o.Name != nil {
		nameQ = *o.Name
	}
	if nameQ != "" {
		qs.Set("Name", nameQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetGreetingURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetGreetingURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetGreetingURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetGreetingURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetGreetingURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetGreetingURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
