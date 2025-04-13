package kea_ca_client

import (
	"kea-ca-client/model"
	"net/http"
	"os"
	"testing"
)

func TestKeaCaClient_Lease4GetAll(t *testing.T) {
	type fields struct {
		httpClient *http.Client
		url        string
	}
	type args struct {
		arguments model.Lease4GetAllArguments
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Not Empty Result",
			fields: fields{
				httpClient: &http.Client{},
				url:        "http://" + os.Getenv("KEA_CA_HOST") + ":" + os.Getenv("KEA_CA_PORT"),
			},
			args: args{
				arguments: model.Lease4GetAllArguments{Subnets: []int{1}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KeaCaClient{
				httpClient: tt.fields.httpClient,
				url:        tt.fields.url,
			}
			got, err := k.Lease4GetAll(tt.args.arguments)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lease4GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got.Leases) == 0 {
				t.Errorf("Leases is empty in response")
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		config KeaCaClientConfig
	}
	tests := []struct {
		name string
		args args
		url  string
	}{
		{
			name: "Correct Url",
			args: args{
				config: KeaCaClientConfig{
					Host: "host-name.domain",
					Port: 1234,
				},
			},
			url: "http://host-name.domain:1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.config); got.url != tt.url {
				t.Errorf("New().url = %v, want %v", got, tt.url)
			}
		})
	}
}
