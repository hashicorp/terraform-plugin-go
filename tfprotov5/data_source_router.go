package tfprotov5

import "context"

type DataSourceRouter map[string]DataSourceServer

func (d DataSourceRouter) ValidateDataSourceConfig(ctx context.Context, req *ValidateDataSourceConfigRequest) (*ValidateDataSourceConfigResponse, error) {
	ds, ok := d[req.TypeName]
	if !ok {
		// TODO: return appropriate error
	}
	return ds.ValidateDataSourceConfig(ctx, req)
}

func (d DataSourceRouter) ReadDataSource(ctx context.Context, req *ReadDataSourceRequest) (*ReadDataSourceResponse, error) {
	ds, ok := d[req.TypeName]
	if !ok {
		// TODO: return appropriate error
	}
	return ds.ReadDataSource(ctx, req)
}
