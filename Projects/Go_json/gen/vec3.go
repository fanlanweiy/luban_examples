
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type vec3 struct {
    X float32
    Y float32
    Z float32
}

const TypeId_vec3 = 3615519

func (*vec3) GetTypeId() int32 {
    return 3615519
}

func Newvec3(_buf map[string]interface{}) (_v *vec3, err error) {
    _v = &vec3{}
    { var _ok_ bool; var __json_x__ interface{}; if __json_x__, _ok_ = _buf["x"]; !_ok_ || __json_x__ == nil { err = errors.New("x error"); return } else { var __x__ float32;  { var _ok_ bool; var _x_ float64; if _x_, _ok_ = __json_x__.(float64); !_ok_ { err = errors.New("__x__ error"); return }; __x__ = float32(_x_) }; _v.X = __x__ }}
    { var _ok_ bool; var __json_y__ interface{}; if __json_y__, _ok_ = _buf["y"]; !_ok_ || __json_y__ == nil { err = errors.New("y error"); return } else { var __x__ float32;  { var _ok_ bool; var _x_ float64; if _x_, _ok_ = __json_y__.(float64); !_ok_ { err = errors.New("__x__ error"); return }; __x__ = float32(_x_) }; _v.Y = __x__ }}
    { var _ok_ bool; var __json_z__ interface{}; if __json_z__, _ok_ = _buf["z"]; !_ok_ || __json_z__ == nil { err = errors.New("z error"); return } else { var __x__ float32;  { var _ok_ bool; var _x_ float64; if _x_, _ok_ = __json_z__.(float64); !_ok_ { err = errors.New("__x__ error"); return }; __x__ = float32(_x_) }; _v.Z = __x__ }}
    return
}

