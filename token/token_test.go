package token

import "testing"

func TestLookupIdent(t *testing.T) {
	t.Parallel()
	type args struct {
		ident string
	}
	tests := []struct {
		name string
		args args
		want TokenType
	}{
		{
			name: "Test_LookupIdent_fn",
			args: args{
				ident: "funk",
			},
			want: FUNCTION,
		},
		{
			name: "Test_LookupIdent_IDENT",
			args: args{
				ident: "foo",
			},
			want: IDENT,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LookupIdent(tt.args.ident); got != tt.want {
				t.Errorf("LookupIdent() = %v, want %v", got, tt.want)
			}
		})
	}
}
