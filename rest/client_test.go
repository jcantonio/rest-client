package rest

import (
	"testing"
)

func TestCall(t *testing.T) {
	type args struct {
		call Call
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Get Companies",
			args: args{
				call: Call{
					Method: "GET",
					URL:    "http://localhost:5000/companies",
					Headers: map[string]string{
						"serverSecret": "serverOwnSecretXXX",
					},
				},
			},
			want:    200,
			wantErr: false,
		},
		{
			name: "Create Company",
			args: args{
				call: Call{
					Method:      "POST",
					URL:         "http://localhost:5000/companies",
					ContentType: "application/json",
					Body: `{
					"companyName": "One-Proteus",
					"companyId": "oneproteus",
					"companyType": "shipper",
					"contactName": "Jean-Claude",
					"contactEmail": "jcantonio@xxx.com",
					"companyImage": "",
					"companyDescription": "Test",
					"companyEndpoint": "",
					"companyPin": "oneproteus"
				  }`,
				},
			},
			want:    200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Do(tt.args.call)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.StatusCode != tt.want {
				t.Errorf("Do() = %v, want %v", string(got.Body), tt.want)
			}
		})
	}
}
