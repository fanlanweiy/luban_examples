
//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

using Luban;


namespace cfg.ai
{
    public sealed partial class UeCooldown : ai.Decorator
    {
        public UeCooldown(ByteBuf _buf)  : base(_buf) 
        {
            CooldownTime = _buf.ReadFloat();
        }

        public static UeCooldown DeserializeUeCooldown(ByteBuf _buf)
        {
            return new ai.UeCooldown(_buf);
        }

        public readonly float CooldownTime;
   
        public const int __ID__ = -951439423;
        public override int GetTypeId() => __ID__;

        public override void ResolveRef(Tables tables)
        {
            base.ResolveRef(tables);
            
        }

        public override string ToString()
        {
            return "{ "
            + "id:" + Id + ","
            + "nodeName:" + NodeName + ","
            + "flowAbortMode:" + FlowAbortMode + ","
            + "cooldownTime:" + CooldownTime + ","
            + "}";
        }
    }

}