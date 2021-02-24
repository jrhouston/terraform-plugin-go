package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func ValidateStorageConfigRequest(in *tfplugin5.ValidateStorageConfig_Request) (*tfprotov5.ValidateStorageConfigRequest, error) {
	resp := &tfprotov5.ValidateStorageConfigRequest{
		StorageName: in.StorageName,
	}
	if in.Config != nil {
		resp.Config = DynamicValue(in.Config)
	}
	return resp, nil
}

func ValidateStorageConfigResponse(in *tfplugin5.ValidateStorageConfig_Response) (*tfprotov5.ValidateStorageConfigResponse, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	resp := &tfprotov5.ValidateStorageConfigResponse{
		Diagnostics: diags,
	}

	return resp, nil
}

func ConfigureStorageRequest(in *tfplugin5.ConfigureStorage_Request) (*tfprotov5.ConfigureStorageRequest, error) {
	resp := &tfprotov5.ConfigureStorageRequest{
		StorageName: in.StorageName,
	}
	if in.Config != nil {
		resp.Config = DynamicValue(in.Config)
	}
	return resp, nil
}

func ConfigureStorageResponse(in *tfplugin5.ConfigureStorage_Response) (*tfprotov5.ConfigureStorageResponse, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	return &tfprotov5.ConfigureStorageResponse{
		Diagnostics: diags,
	}, nil
}

func ReadStateRequest(in *tfplugin5.ReadState_Request) (*tfprotov5.ReadStateRequest, error) {
	resp := &tfprotov5.ReadStateRequest{
		StateName: in.StateName,
	}
	return resp, nil
}

func ReadStateResponse(in *tfplugin5.ReadState_Response) (*tfprotov5.ReadStateResponse, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	return &tfprotov5.ReadStateResponse{
		State:       in.State,
		Diagnostics: diags,
	}, nil
}

func WriteStateRequest(in *tfplugin5.WriteState_Request) (*tfprotov5.WriteStateRequest, error) {
	resp := &tfprotov5.WriteStateRequest{
		StateName: in.StateName,
		State:     in.State,
	}
	return resp, nil
}

func WriteStateResponse(in *tfplugin5.WriteState_Response) (*tfprotov5.WriteStateResponse, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	return &tfprotov5.WriteStateResponse{
		Diagnostics: diags,
	}, nil
}

func LockStateRequest(in *tfplugin5.LockState_Request) (*tfprotov5.LockStateRequest, error) {
	resp := &tfprotov5.LockStateRequest{
		StateName: in.StateName,
	}
	return resp, nil
}

func LockStateResponse(in *tfplugin5.LockState_Response) (*tfprotov5.LockStateResponse, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	return &tfprotov5.LockStateResponse{
		Diagnostics: diags,
	}, nil
}

func UnlockStateRequest(in *tfplugin5.UnlockState_Request) (*tfprotov5.UnlockStateRequest, error) {
	resp := &tfprotov5.UnlockStateRequest{
		StateName: in.StateName,
		Force:     in.Force,
	}
	return resp, nil
}

func UnlockStateResponse(in *tfplugin5.UnlockState_Response) (*tfprotov5.UnlockStateResponse, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	return &tfprotov5.UnlockStateResponse{
		Diagnostics: diags,
	}, nil
}

func GetStatesRequest(in *tfplugin5.GetStates_Request) (*tfprotov5.GetStatesRequest, error) {
	resp := &tfprotov5.GetStatesRequest{}
	return resp, nil
}

func GetStatesResponse(in *tfplugin5.GetStates_Response) (*tfprotov5.GetStatesResponse, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	return &tfprotov5.GetStatesResponse{
		Diagnostics: diags,
	}, nil
}

func DeleteStateRequest(in *tfplugin5.DeleteState_Request) (*tfprotov5.DeleteStateRequest, error) {
	resp := &tfprotov5.DeleteStateRequest{
		StateName: in.StateName,
	}
	return resp, nil
}

func DeleteStateResponse(in *tfplugin5.DeleteState_Response) (*tfprotov5.DeleteStateResponse, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	return &tfprotov5.DeleteStateResponse{
		Diagnostics: diags,
	}, nil
}
