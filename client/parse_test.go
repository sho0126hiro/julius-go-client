package client

import (
	"reflect"
	"testing"
)

func Test_parseMessage(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		args    args
		want    *Result
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				msg: "<ENDRECOG/>\n" +
					".\n" +
					"<INPUTPARAM FRAMES=\"81\" MSEC=\"810\"/>\n" +
					".\n" +
					"<RECOGOUT>\n" +
					"  <SHYPO RANK=\"1\" SCORE=\"-2134.392090\" GRAM=\"0\">\n" +
					"    <WHYPO WORD=\"電気オン\" CLASSID=\"電気オン\" PHONE=\"silB d e N k i o N silE\" CM=\"0.937\"/>\n" +
					"  </SHYPO>\n" +
					"</RECOGOUT>\n" +
					".\n" +
					"<INPUT STATUS=\"LISTEN\" TIME=\"1621422699\"/>\n" +
					".\n",
			},
			want: &Result{
				Rank:  "1",
				Score: "-2134.392090",
				Gram:  "0",
				Details: []Detail{
					{
						Word:    "電気オン",
						ClassID: "電気オン",
						Phone:   "silB d e N k i o N silE",
						CM:      "0.937",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMessage(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
