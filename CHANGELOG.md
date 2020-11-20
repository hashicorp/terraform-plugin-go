# 0.2.0 (November 20, 2020)

ENHANCEMENTS:

* `tftypes.NewValue` can now accept a wider array of standard library types which will be automatically converted to their standard representation ([#46](https://github.com/hashicorp/terraform-plugin-go/issues/46)] [[#47](https://github.com/hashicorp/terraform-plugin-go/issues/47))
* `tfprotov5.RawState` now has an `Unmarshal` method, just like `tfprotov5.DynamicValue`, yielding a `tftypes.Value`. ([#42](https://github.com/hashicorp/terraform-plugin-go/issues/42))

# 0.1.0 (November 02, 2020)

Initial release.
