package uzer

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"go_lab/lab/gomock_test/mocks"
	"testing"
)

func TestUser_Use(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}

	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	fmt.Println(testUser.Use())
	println()
}

func TestUser_SaySomething(t *testing.T) {

	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	user := &User{Doer: mockDoer}

	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		expect  func()
		wantErr bool
	}{
		{
			name: "Positive test case 1",
			expect: func() {
				mockDoer.EXPECT().SaySomething("spam").Return("bad")
			},
			args:    args{num: 3},
			wantErr: false,
			want:    "bad",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tt.expect()
			if got := user.SaySomething(tt.args.num); (got != tt.want) != tt.wantErr {

				fmt.Println("gott :" + got)
				t.Errorf("User SaySomething() = %v, want %v", got, tt.want)
			}

		})
	}
}
