package tfprotov5

import "context"

type ErrUnsupportedDataSource string

func (e ErrUnsupportedDataSource) Error() string {
	return "unsupported data source: " + string(e)
}

type DataSourceRouter map[string]DataSourceServer

func (d DataSourceRouter) ValidateDataSourceConfig(ctx context.Context, req *ValidateDataSourceConfigRequest) (*ValidateDataSourceConfigResponse, error) {
	ds, ok := d[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedDataSource(req.TypeName)
	}
	return ds.ValidateDataSourceConfig(ctx, req)
}

func (d DataSourceRouter) ReadDataSource(ctx context.Context, req *ReadDataSourceRequest) (*ReadDataSourceResponse, error) {
	ds, ok := d[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedDataSource(req.TypeName)
	}
	return ds.ReadDataSource(ctx, req)
}
