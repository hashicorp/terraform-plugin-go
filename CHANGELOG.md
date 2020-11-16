# 0.2.0 (Unreleased)

ENHANCEMENTS:

* `tftypes.NewValue` can now accept a wider array of standard library types which will be automatically converted to their standard representation [GH-46]
* `tfprotov5.RawState` now has an `Unmarshal` method, just like `tfprotov5.DynamicValue`, yielding a `tftypes.Value`. [GH-42]

# 0.1.0 (November 02, 2020)

Initial release.
