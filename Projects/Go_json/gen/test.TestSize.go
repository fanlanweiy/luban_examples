
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type TestTestSize struct {
    Id int32
    X1 []int32
    X2 []int32
    X3 []int32
    X4 map[int32]int32
}

const TypeId_TestTestSize = 340006319

func (*TestTestSize) GetTypeId() int32 {
    return 340006319
}

func NewTestTestSize(_buf map[string]interface{}) (_v *TestTestSize, err error) {
    _v = &TestTestSize{}
    { var _ok_ bool; var __json_id__ interface{}; if __json_id__, _ok_ = _buf["id"]; !_ok_ || __json_id__ == nil { err = errors.New("id error"); return } else { var __x__ int32;  { var _ok_ bool; var _x_ float64; if _x_, _ok_ = __json_id__.(float64); !_ok_ { err = errors.New("__x__ error"); return }; __x__ = int32(_x_) }; _v.Id = __x__ }}
    { var _ok_ bool; var __json_x1__ interface{}; if __json_x1__, _ok_ = _buf["x1"]; !_ok_ || __json_x1__ == nil { err = errors.New("x1 error"); return } else { var __x__ []int32;  {
                    var _arr0_ []interface{}
                    var _ok0_ bool
                    if _arr0_, _ok0_ = (__json_x1__).([]interface{}); !_ok0_ { err = errors.New("__x__ error"); return }
    
                    __x__ = make([]int32, 0, len(_arr0_))
                    
                    for _, _e0_ := range _arr0_ {
                        var _list_v0_ int32
                        { var _ok_ bool; var _x_ float64; if _x_, _ok_ = _e0_.(float64); !_ok_ { err = errors.New("_list_v0_ error"); return }; _list_v0_ = int32(_x_) }
                        __x__ = append(__x__, _list_v0_)
                    }
                }
    ; _v.X1 = __x__ }}
    { var _ok_ bool; var __json_x2__ interface{}; if __json_x2__, _ok_ = _buf["x2"]; !_ok_ || __json_x2__ == nil { err = errors.New("x2 error"); return } else { var __x__ []int32;  {
                    var _arr0_ []interface{}
                    var _ok0_ bool
                    if _arr0_, _ok0_ = (__json_x2__).([]interface{}); !_ok0_ { err = errors.New("__x__ error"); return }
    
                    __x__ = make([]int32, 0, len(_arr0_))
                    
                    for _, _e0_ := range _arr0_ {
                        var _list_v0_ int32
                        { var _ok_ bool; var _x_ float64; if _x_, _ok_ = _e0_.(float64); !_ok_ { err = errors.New("_list_v0_ error"); return }; _list_v0_ = int32(_x_) }
                        __x__ = append(__x__, _list_v0_)
                    }
                }
    ; _v.X2 = __x__ }}
    { var _ok_ bool; var __json_x3__ interface{}; if __json_x3__, _ok_ = _buf["x3"]; !_ok_ || __json_x3__ == nil { err = errors.New("x3 error"); return } else { var __x__ []int32;  {
                    var _arr0_ []interface{}
                    var _ok0_ bool
                    if _arr0_, _ok0_ = (__json_x3__).([]interface{}); !_ok0_ { err = errors.New("__x__ error"); return }
    
                    __x__ = make([]int32, 0, len(_arr0_))
                    
                    for _, _e0_ := range _arr0_ {
                        var _list_v0_ int32
                        { var _ok_ bool; var _x_ float64; if _x_, _ok_ = _e0_.(float64); !_ok_ { err = errors.New("_list_v0_ error"); return }; _list_v0_ = int32(_x_) }
                        __x__ = append(__x__, _list_v0_)
                    }
                }
    ; _v.X3 = __x__ }}
    { var _ok_ bool; var __json_x4__ interface{}; if __json_x4__, _ok_ = _buf["x4"]; !_ok_ || __json_x4__ == nil { err = errors.New("x4 error"); return } else { var __x__ map[int32]int32;  {
                    var _arr0_ []interface{}
                    var _ok0_ bool
                    if _arr0_, _ok_ = (__json_x4__).([]interface{}); !_ok_ { err = errors.New("__x__ error"); return }
    
                    __x__ = make(map[int32]int32)
                    
                    for _, _e0_ := range _arr0_ {
                        var _kv0_ []interface{}
                        if _kv0_, _ok0_ = _e0_.([]interface{}); !_ok0_ || len(_kv0_) != 2 { err = errors.New("__x__ error"); return }
                        var _key0_ int32
                        { var _ok_ bool; var _x_ float64; if _x_, _ok_ = _kv0_[0].(float64); !_ok_ { err = errors.New("_key0_ error"); return }; _key0_ = int32(_x_) }
                        var _value0_ int32
                        { var _ok_ bool; var _x_ float64; if _x_, _ok_ = _kv0_[1].(float64); !_ok_ { err = errors.New("_value0_ error"); return }; _value0_ = int32(_x_) }
                        __x__[_key0_] = _value0_
                    }
                    }; _v.X4 = __x__ }}
    return
}

