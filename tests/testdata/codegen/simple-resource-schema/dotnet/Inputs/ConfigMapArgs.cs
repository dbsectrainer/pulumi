// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example.Inputs
{

    public sealed class ConfigMapArgs : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        public Input<string>? Config { get; set; }

        public ConfigMapArgs()
        {
        }
        public static new ConfigMapArgs Empty => new ConfigMapArgs();
    }
}