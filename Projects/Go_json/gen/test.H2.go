
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type TestH2 struct {
    Z2 int32
    Z3 int32
}

const TypeId_TestH2 = -1422503994

func (*TestH2) GetTypeId() int32 {
    return -1422503994
}

func NewTestH2(_buf map[string]interface{}) (_v *TestH2, err error) {
    _v = &TestH2{}
    { var _ok_ bool; var __json_z2__ interface{}; if __json_z2__, _ok_ = _buf["z2"]; !_ok_ || __json_z2__ == nil { err = errors.New("z2 error"); return } else { var __x__ int32;  { var _ok_ bool; var _x_ float64; if _x_, _ok_ = __json_z2__.(float64); !_ok_ { err = errors.New("__x__ error"); return }; __x__ = int32(_x_) }; _v.Z2 = __x__ }}
    { var _ok_ bool; var __json_z3__ interface{}; if __json_z3__, _ok_ = _buf["z3"]; !_ok_ || __json_z3__ == nil { err = errors.New("z3 error"); return } else { var __x__ int32;  { var _ok_ bool; var _x_ float64; if _x_, _ok_ = __json_z3__.(float64); !_ok_ { err = errors.New("__x__ error"); return }; __x__ = int32(_x_) }; _v.Z3 = __x__ }}
    return
}

