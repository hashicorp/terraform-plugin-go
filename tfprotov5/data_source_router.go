package tfprotov5

import "context"

// ErrUnsupportedDataSource is returned when a `DataSourceRouter` receives a
// request for a data source that is not found in its map. It will be set to
// the data source that was not found.
type ErrUnsupportedDataSource string

func (e ErrUnsupportedDataSource) Error() string {
	return "unsupported data source: " + string(e)
}

// DataSourceRouter is an implementation of the `DataSourceServer` interface.
// Requests will have their TypeName property matched to the keys in the
// DataSourceRouter, and the equivalent method of the `DataSourceServer` for
// that key will be called, passing in the request.
//
// This is used for giving each type of data source its own
// ValidateDataSourceConfig and ReadDataSource implementation, rather than
// needing to implement all data sources in the same functions.
type DataSourceRouter map[string]DataSourceServer

// ValidateDataSourceConfig is called when Terraform is attempting to validate
// the configuration for a data source. ValidateDataSourceConfig uses the
// TypeName property of `req` to look up a `DataSourceServer` in `d`, then
// calls that `DataSourceServer`'s `ValidateDataSourceConfig` method, passing
// in `req`.
//
// If no key matches `req`'s TypeName property, an `ErrUnsupportedDataSource`
// error is returned, with its value set to `req`'s TypeName.
func (d DataSourceRouter) ValidateDataSourceConfig(ctx context.Context, req *ValidateDataSourceConfigRequest) (*ValidateDataSourceConfigResponse, error) {
	ds, ok := d[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedDataSource(req.TypeName)
	}
	return ds.ValidateDataSourceConfig(ctx, req)
}

// ReadDataSource is called when Terraform is attempting to retrieve the latest
// state for a data source. ReadDataSource uses the TypeName property of `req`
// to look up a `DataSourceServer` in `d`, then calls that `DataSourceServer`'s
// `ReadDataSource` method, passing in `req`.
//
// If no key matches `req`'s TypeName property, an `ErrUnsupportedDataSource`
// error is returned, with its value set to `req`'s TypeName.
func (d DataSourceRouter) ReadDataSource(ctx context.Context, req *ReadDataSourceRequest) (*ReadDataSourceResponse, error) {
	ds, ok := d[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedDataSource(req.TypeName)
	}
	return ds.ReadDataSource(ctx, req)
}
