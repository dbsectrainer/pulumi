// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

var _ = internal.GetEnvOrDefault

type Sandwich struct {
	Bread   *string  `pulumi:"bread"`
	Veggies []string `pulumi:"veggies"`
}

// SandwichInput is an input type that accepts SandwichArgs and SandwichOutput values.
// You can construct a concrete instance of `SandwichInput` via:
//
//	SandwichArgs{...}
type SandwichInput interface {
	pulumi.Input

	ToSandwichOutput() SandwichOutput
	ToSandwichOutputWithContext(context.Context) SandwichOutput
}

type SandwichArgs struct {
	Bread   pulumi.StringPtrInput   `pulumi:"bread"`
	Veggies pulumi.StringArrayInput `pulumi:"veggies"`
}

func (SandwichArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sandwich)(nil)).Elem()
}

func (i SandwichArgs) ToSandwichOutput() SandwichOutput {
	return i.ToSandwichOutputWithContext(context.Background())
}

func (i SandwichArgs) ToSandwichOutputWithContext(ctx context.Context) SandwichOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SandwichOutput)
}

// SandwichArrayInput is an input type that accepts SandwichArray and SandwichArrayOutput values.
// You can construct a concrete instance of `SandwichArrayInput` via:
//
//	SandwichArray{ SandwichArgs{...} }
type SandwichArrayInput interface {
	pulumi.Input

	ToSandwichArrayOutput() SandwichArrayOutput
	ToSandwichArrayOutputWithContext(context.Context) SandwichArrayOutput
}

type SandwichArray []SandwichInput

func (SandwichArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Sandwich)(nil)).Elem()
}

func (i SandwichArray) ToSandwichArrayOutput() SandwichArrayOutput {
	return i.ToSandwichArrayOutputWithContext(context.Background())
}

func (i SandwichArray) ToSandwichArrayOutputWithContext(ctx context.Context) SandwichArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SandwichArrayOutput)
}

type SandwichOutput struct{ *pulumi.OutputState }

func (SandwichOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sandwich)(nil)).Elem()
}

func (o SandwichOutput) ToSandwichOutput() SandwichOutput {
	return o
}

func (o SandwichOutput) ToSandwichOutputWithContext(ctx context.Context) SandwichOutput {
	return o
}

func (o SandwichOutput) Bread() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sandwich) *string { return v.Bread }).(pulumi.StringPtrOutput)
}

func (o SandwichOutput) Veggies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Sandwich) []string { return v.Veggies }).(pulumi.StringArrayOutput)
}

type SandwichArrayOutput struct{ *pulumi.OutputState }

func (SandwichArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Sandwich)(nil)).Elem()
}

func (o SandwichArrayOutput) ToSandwichArrayOutput() SandwichArrayOutput {
	return o
}

func (o SandwichArrayOutput) ToSandwichArrayOutputWithContext(ctx context.Context) SandwichArrayOutput {
	return o
}

func (o SandwichArrayOutput) Index(i pulumi.IntInput) SandwichOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Sandwich {
		return vs[0].([]Sandwich)[vs[1].(int)]
	}).(SandwichOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SandwichInput)(nil)).Elem(), SandwichArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SandwichArrayInput)(nil)).Elem(), SandwichArray{})
	pulumi.RegisterOutputType(SandwichOutput{})
	pulumi.RegisterOutputType(SandwichArrayOutput{})
}