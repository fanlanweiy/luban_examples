
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

package cfg;


import "errors"

type AiDistanceLessThan struct {
    Id int32
    NodeName string
    FlowAbortMode int32
    Actor1Key string
    Actor2Key string
    Distance float32
    ReverseResult bool
}

const TypeId_AiDistanceLessThan = -1207170283

func (*AiDistanceLessThan) GetTypeId() int32 {
    return -1207170283
}

func NewAiDistanceLessThan(_buf map[string]interface{}) (_v *AiDistanceLessThan, err error) {
    _v = &AiDistanceLessThan{}
    { var _ok_ bool; var __json_id__ interface{}; if __json_id__, _ok_ = _buf["id"]; !_ok_ || __json_id__ == nil { err = errors.New("id error"); return } else { var __x__ int32;  { var _ok_ bool; var _x_ float64; if _x_, _ok_ = __json_id__.(float64); !_ok_ { err = errors.New("__x__ error"); return }; __x__ = int32(_x_) }; _v.Id = __x__ }}
    { var _ok_ bool; var __json_node_name__ interface{}; if __json_node_name__, _ok_ = _buf["node_name"]; !_ok_ || __json_node_name__ == nil { err = errors.New("node_name error"); return } else { var __x__ string;  {  if __x__, _ok_ = __json_node_name__.(string); !_ok_ { err = errors.New("__x__ error"); return } }; _v.NodeName = __x__ }}
    { var _ok_ bool; var __json_flow_abort_mode__ interface{}; if __json_flow_abort_mode__, _ok_ = _buf["flow_abort_mode"]; !_ok_ || __json_flow_abort_mode__ == nil { err = errors.New("flow_abort_mode error"); return } else { var __x__ int32;  { var _ok_ bool; var _x_ float64; if _x_, _ok_ = __json_flow_abort_mode__.(float64); !_ok_ { err = errors.New("__x__ error"); return }; __x__ = int32(_x_) }; _v.FlowAbortMode = __x__ }}
    { var _ok_ bool; var __json_actor1_key__ interface{}; if __json_actor1_key__, _ok_ = _buf["actor1_key"]; !_ok_ || __json_actor1_key__ == nil { err = errors.New("actor1_key error"); return } else { var __x__ string;  {  if __x__, _ok_ = __json_actor1_key__.(string); !_ok_ { err = errors.New("__x__ error"); return } }; _v.Actor1Key = __x__ }}
    { var _ok_ bool; var __json_actor2_key__ interface{}; if __json_actor2_key__, _ok_ = _buf["actor2_key"]; !_ok_ || __json_actor2_key__ == nil { err = errors.New("actor2_key error"); return } else { var __x__ string;  {  if __x__, _ok_ = __json_actor2_key__.(string); !_ok_ { err = errors.New("__x__ error"); return } }; _v.Actor2Key = __x__ }}
    { var _ok_ bool; var __json_distance__ interface{}; if __json_distance__, _ok_ = _buf["distance"]; !_ok_ || __json_distance__ == nil { err = errors.New("distance error"); return } else { var __x__ float32;  { var _ok_ bool; var _x_ float64; if _x_, _ok_ = __json_distance__.(float64); !_ok_ { err = errors.New("__x__ error"); return }; __x__ = float32(_x_) }; _v.Distance = __x__ }}
    { var _ok_ bool; var __json_reverse_result__ interface{}; if __json_reverse_result__, _ok_ = _buf["reverse_result"]; !_ok_ || __json_reverse_result__ == nil { err = errors.New("reverse_result error"); return } else { var __x__ bool;  { var _ok_ bool; if __x__, _ok_ = __json_reverse_result__.(bool); !_ok_ { err = errors.New("__x__ error"); return } }; _v.ReverseResult = __x__ }}
    return
}

