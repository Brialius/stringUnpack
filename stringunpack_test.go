package stringUnpack

import "testing"

func TestStringUnpack(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "a4bc2d5e",
			args: args{
				s: "a4bc2d5e",
			},
			want:    "aaaabccddddde",
			wantErr: false,
		}, {
			name: "abcd",
			args: args{
				s: "abcd",
			},
			want:    "abcd",
			wantErr: false,
		}, {
			name: "45",
			args: args{
				s: "45",
			},
			want:    "",
			wantErr: true,
		}, {
			name: `qwe\4\5`,
			args: args{
				s: `qwe\4\5`,
			},
			want:    `qwe45`,
			wantErr: false,
		}, {
			name: `qwe\45`,
			args: args{
				s: `qwe\45`,
			},
			want:    "qwe44444",
			wantErr: false,
		}, {
			name: `qwe\\5`,
			args: args{
				s: `qwe\\5`,
			},
			want:    `qwe\\\\\`,
			wantErr: false,
		}, {
			name: `qwe1f`,
			args: args{
				s: `qwe1f`,
			},
			want:    `qwef`,
			wantErr: false,
		}, {
			name: `qwef11`,
			args: args{
				s: `qwe1f11`,
			},
			want:    `qwefffffffffff`,
			wantErr: false,
		}, {
			name: `qwe1f`,
			args: args{
				s: `qwe1f`,
			},
			want:    `qwef`,
			wantErr: false,
		}, {
			name: `2qwe2f\5`,
			args: args{
				s: `2qwe2f\5`,
			},
			want:    "",
			wantErr: true,
		}, {
			name: `\2qwe2f\5`,
			args: args{
				s: `\2qwe2f\5`,
			},
			want:    `2qweef5`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringUnpack(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringUnpack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringUnpack() = %v, want %v", got, tt.want)
			}
		})
	}
}
