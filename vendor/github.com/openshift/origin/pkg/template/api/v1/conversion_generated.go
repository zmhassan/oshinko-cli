// +build !ignore_autogenerated

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	template_api "github.com/openshift/origin/pkg/template/api"
	api "k8s.io/kubernetes/pkg/api"
	conversion "k8s.io/kubernetes/pkg/conversion"
	reflect "reflect"
)

func init() {
	if err := api.Scheme.AddGeneratedConversionFuncs(
		Convert_v1_Parameter_To_api_Parameter,
		Convert_api_Parameter_To_v1_Parameter,
		Convert_v1_Template_To_api_Template,
		Convert_api_Template_To_v1_Template,
		Convert_v1_TemplateList_To_api_TemplateList,
		Convert_api_TemplateList_To_v1_TemplateList,
	); err != nil {
		// if one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}

func autoConvert_v1_Parameter_To_api_Parameter(in *Parameter, out *template_api.Parameter, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*Parameter))(in)
	}
	out.Name = in.Name
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	out.Value = in.Value
	out.Generate = in.Generate
	out.From = in.From
	out.Required = in.Required
	return nil
}

func Convert_v1_Parameter_To_api_Parameter(in *Parameter, out *template_api.Parameter, s conversion.Scope) error {
	return autoConvert_v1_Parameter_To_api_Parameter(in, out, s)
}

func autoConvert_api_Parameter_To_v1_Parameter(in *template_api.Parameter, out *Parameter, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*template_api.Parameter))(in)
	}
	out.Name = in.Name
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	out.Value = in.Value
	out.Generate = in.Generate
	out.From = in.From
	out.Required = in.Required
	return nil
}

func Convert_api_Parameter_To_v1_Parameter(in *template_api.Parameter, out *Parameter, s conversion.Scope) error {
	return autoConvert_api_Parameter_To_v1_Parameter(in, out, s)
}

func autoConvert_v1_TemplateList_To_api_TemplateList(in *TemplateList, out *template_api.TemplateList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*TemplateList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]template_api.Template, len(*in))
		for i := range *in {
			if err := Convert_v1_Template_To_api_Template(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_TemplateList_To_api_TemplateList(in *TemplateList, out *template_api.TemplateList, s conversion.Scope) error {
	return autoConvert_v1_TemplateList_To_api_TemplateList(in, out, s)
}

func autoConvert_api_TemplateList_To_v1_TemplateList(in *template_api.TemplateList, out *TemplateList, s conversion.Scope) error {
	if defaulting, found := s.DefaultingInterface(reflect.TypeOf(*in)); found {
		defaulting.(func(*template_api.TemplateList))(in)
	}
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Template, len(*in))
		for i := range *in {
			if err := Convert_api_Template_To_v1_Template(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_TemplateList_To_v1_TemplateList(in *template_api.TemplateList, out *TemplateList, s conversion.Scope) error {
	return autoConvert_api_TemplateList_To_v1_TemplateList(in, out, s)
}
