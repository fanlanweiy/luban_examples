
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

#include "schema.h"

namespace cfg {


bool test::TestScriptableObject::deserialize(::luban::ByteBuf& _buf)
{

    if(!_buf.readInt(id)) return false;
    if(!_buf.readString(desc)) return false;
    if(!_buf.readFloat(rate)) return false;
    if(!_buf.readInt(num)) return false;
    if(!vec2::deserializevec2(_buf, v2)) return false;
    if(!vec3::deserializevec3(_buf, v3)) return false;
    if(!vec4::deserializevec4(_buf, v4)) return false;

    return true;
}

bool test::TestScriptableObject::deserializeTestScriptableObject(::luban::ByteBuf& _buf, ::luban::SharedPtr<test::TestScriptableObject>& _out)
{
    _out.reset(new test::TestScriptableObject());
    if (_out->deserialize(_buf))
    {
        return true;
    }
    else
    { 
        _out.reset();
        return false;
    }
}


bool test::Path::deserialize(::luban::ByteBuf& _buf)
{

    if(!_buf.readInt(id)) return false;
    if(!_buf.readString(res)) return false;

    return true;
}

bool test::Path::deserializePath(::luban::ByteBuf& _buf, ::luban::SharedPtr<test::Path>& _out)
{
    _out.reset(new test::Path());
    if (_out->deserialize(_buf))
    {
        return true;
    }
    else
    { 
        _out.reset();
        return false;
    }
}


bool test::TestMapper::deserialize(::luban::ByteBuf& _buf)
{

    if(!_buf.readInt(id)) return false;
    {int __enum_temp__; if(!_buf.readInt(__enum_temp__)) return false; audioType = AudioType(__enum_temp__); }
    if(!vec2::deserializevec2(_buf, v2)) return false;

    return true;
}

bool test::TestMapper::deserializeTestMapper(::luban::ByteBuf& _buf, ::luban::SharedPtr<test::TestMapper>& _out)
{
    _out.reset(new test::TestMapper());
    if (_out->deserialize(_buf))
    {
        return true;
    }
    else
    { 
        _out.reset();
        return false;
    }
}


bool DefineFromExcel2::deserialize(::luban::ByteBuf& _buf)
{

    if(!_buf.readInt(id)) return false;
    if (!_buf.readBool(x1)) return false;
    if(!_buf.readLong(x5)) return false;
    if(!_buf.readFloat(x6)) return false;
    if(!_buf.readInt(x8)) return false;
    if(!_buf.readString(x10)) return false;
    {int __enum_temp__; if(!_buf.readInt(__enum_temp__)) return false; x13 = test::DemoEnum(__enum_temp__); }
    {int __enum_temp__; if(!_buf.readInt(__enum_temp__)) return false; x132 = test::DemoFlag(__enum_temp__); }
    if(!test::DemoDynamic::deserializeDemoDynamic(_buf, x14)) return false;
    if(!test::Shape::deserializeShape(_buf, x15)) return false;
    if(!vec2::deserializevec2(_buf, v2)) return false;
    if(!_buf.readLong(t1)) return false;
    {::luban::int32 n; if(!_buf.readSize(n)) return false; n = std::min(n, ::luban::int32(_buf.size()));k1.reserve(n);for(int i = 0 ; i < n ; i++) { ::luban::int32 _e;if(!_buf.readInt(_e)) return false; k1.push_back(_e);}}
    {::luban::int32 n; if(!_buf.readSize(n)) return false; n = std::min(n, ::luban::int32(_buf.size()));k2.reserve(n);for(int i = 0 ; i < n ; i++) { ::luban::int32 _e;if(!_buf.readInt(_e)) return false; k2.push_back(_e);}}
    {::luban::int32 n; if(!_buf.readSize(n)) return false; n = std::min(n, (::luban::int32)_buf.size()); k8.reserve(n * 3 / 2);for(int i = 0 ; i < n ; i++) { ::luban::int32 _k;  if(!_buf.readInt(_k)) return false; ::luban::int32 _v;  if(!_buf.readInt(_v)) return false;     k8[_k] = _v;}}
    {::luban::int32 n; if(!_buf.readSize(n)) return false; n = std::min(n, ::luban::int32(_buf.size())); k9.reserve(n);for(int i = 0 ; i < n ; i++) { ::luban::SharedPtr<test::DemoE2> _e;  if(!test::DemoE2::deserializeDemoE2(_buf, _e)) return false; k9.push_back(_e);}}
    {::luban::int32 n; if(!_buf.readSize(n)) return false; n = std::min(n, ::luban::int32(_buf.size())); k10.reserve(n);for(int i = 0 ; i < n ; i++) { ::luban::SharedPtr<vec3> _e;  if(!vec3::deserializevec3(_buf, _e)) return false; k10.push_back(_e);}}
    {::luban::int32 n; if(!_buf.readSize(n)) return false; n = std::min(n, ::luban::int32(_buf.size())); k11.reserve(n);for(int i = 0 ; i < n ; i++) { ::luban::SharedPtr<vec4> _e;  if(!vec4::deserializevec4(_buf, _e)) return false; k11.push_back(_e);}}

    return true;
}

bool DefineFromExcel2::deserializeDefineFromExcel2(::luban::ByteBuf& _buf, ::luban::SharedPtr<DefineFromExcel2>& _out)
{
    _out.reset(new DefineFromExcel2());
    if (_out->deserialize(_buf))
    {
        return true;
    }
    else
    { 
        _out.reset();
        return false;
    }
}


bool test::Shape::deserialize(::luban::ByteBuf& _buf)
{


    return true;
}

bool test::Shape::deserializeShape(::luban::ByteBuf& _buf, ::luban::SharedPtr<test::Shape>& _out)
{
    int id;
    if (!_buf.readInt(id)) return false;
    switch (id)
    {
        case test::Circle::__ID__: { _out.reset(new test::Circle()); if (_out->deserialize(_buf)) { return true; } else { _out.reset(); return false;} }
        case test2::Rectangle::__ID__: { _out.reset(new test2::Rectangle()); if (_out->deserialize(_buf)) { return true; } else { _out.reset(); return false;} }
        default: { _out = nullptr; return false;}
    }
}


bool test::Circle::deserialize(::luban::ByteBuf& _buf)
{
    if (!test::Shape::deserialize(_buf))
    {
        return false;
    }

    if(!_buf.readFloat(radius)) return false;

    return true;
}

bool test::Circle::deserializeCircle(::luban::ByteBuf& _buf, ::luban::SharedPtr<test::Circle>& _out)
{
    _out.reset(new test::Circle());
    if (_out->deserialize(_buf))
    {
        return true;
    }
    else
    { 
        _out.reset();
        return false;
    }
}


bool test2::Rectangle::deserialize(::luban::ByteBuf& _buf)
{
    if (!test::Shape::deserialize(_buf))
    {
        return false;
    }

    if(!_buf.readFloat(width)) return false;
    if(!_buf.readFloat(height)) return false;

    return true;
}

bool test2::Rectangle::deserializeRectangle(::luban::ByteBuf& _buf, ::luban::SharedPtr<test2::Rectangle>& _out)
{
    _out.reset(new test2::Rectangle());
    if (_out->deserialize(_buf))
    {
        return true;
    }
    else
    { 
        _out.reset();
        return false;
    }
}

}
