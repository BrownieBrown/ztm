package challenges

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test 1: Reverse middle",
			args: args{
				head:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				left:  2,
				right: 4,
			},
			want: &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, nil}}}}},
		},
		{
			name: "Test 2: Reverse entire list",
			args: args{
				head:  &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
				left:  1,
				right: 3,
			},
			want: &ListNode{3, &ListNode{2, &ListNode{1, nil}}},
		},
		{
			name: "Test 3: Reverse first two nodes",
			args: args{
				head:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
				left:  1,
				right: 2,
			},
			want: &ListNode{2, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}},
		},
		{
			name: "Test 4: Single node list",
			args: args{
				head:  &ListNode{1, nil},
				left:  1,
				right: 1,
			},
			want: &ListNode{1, nil},
		},
		{
			name: "Test 5: No reverse (left == right)",
			args: args{
				head:  &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
				left:  2,
				right: 2,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
