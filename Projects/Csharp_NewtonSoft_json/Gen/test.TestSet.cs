
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

using Luban;
using Newtonsoft.Json.Linq;



namespace cfg.test
{

public sealed partial class TestSet : Luban.BeanBase
{
    public TestSet(JToken _buf) 
    {
        JObject _obj = _buf as JObject;
        Id = (int)_obj.GetValue("id");
        X0 = (string)_obj.GetValue("x0");
        { var __json0 = _obj.GetValue("x1"); X1 = new System.Collections.Generic.List<int>((__json0 as JArray).Count); foreach(JToken __e0 in __json0) { int __v0;  __v0 = (int)__e0;  X1.Add(__v0); }   }
        { var __json0 = _obj.GetValue("x2"); X2 = new System.Collections.Generic.List<long>((__json0 as JArray).Count); foreach(JToken __e0 in __json0) { long __v0;  __v0 = (long)__e0;  X2.Add(__v0); }   }
        { var __json0 = _obj.GetValue("x3"); X3 = new System.Collections.Generic.List<string>((__json0 as JArray).Count); foreach(JToken __e0 in __json0) { string __v0;  __v0 = (string)__e0;  X3.Add(__v0); }   }
        { var __json0 = _obj.GetValue("x4"); X4 = new System.Collections.Generic.List<test.DemoEnum>((__json0 as JArray).Count); foreach(JToken __e0 in __json0) { test.DemoEnum __v0;  __v0 = (test.DemoEnum)(int)__e0;  X4.Add(__v0); }   }
    }

    public static TestSet DeserializeTestSet(JToken _buf)
    {
        return new test.TestSet(_buf);
    }

    public readonly int Id;
    public readonly string X0;
    public readonly System.Collections.Generic.List<int> X1;
    public readonly System.Collections.Generic.List<long> X2;
    public readonly System.Collections.Generic.List<string> X3;
    public readonly System.Collections.Generic.List<test.DemoEnum> X4;


    public const int __ID__ = -543221516;
    public override int GetTypeId() => __ID__;

    public  void ResolveRef(Tables tables)
    {
    }

    public override string ToString()
    {
        return "{ "
        + "id:" + Id + ","
        + "x0:" + X0 + ","
        + "x1:" + Luban.StringUtil.CollectionToString(X1) + ","
        + "x2:" + Luban.StringUtil.CollectionToString(X2) + ","
        + "x3:" + Luban.StringUtil.CollectionToString(X3) + ","
        + "x4:" + Luban.StringUtil.CollectionToString(X4) + ","
        + "}";
    }
}
}
