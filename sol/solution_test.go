package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	nums := []int{3, 2, 1, 0, 4}
	for idx := 0; idx < b.N; idx++ {
		canJump(nums)
	}
}
func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nums = [2,3,1,1,4]",
			args: args{nums: []int{2, 3, 1, 1, 4}},
			want: true,
		},
		{
			name: "nums = [3,2,1,0,4]",
			args: args{nums: []int{3, 2, 1, 0, 4}},
			want: false,
		},
		{
			name: "nums = [2,0,0]",
			args: args{nums: []int{2, 0, 0}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
