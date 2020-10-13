package architecture_class

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPut(t *testing.T) {
	ctrl := gomock.NewController(t)
	acc := NewMockAccessor(ctrl)

	p := Person{First: "Jokke"}

	acc.EXPECT().Save(1, p).MinTimes(1).MaxTimes(1)

	Put(acc, 1, p)

	ctrl.Finish()

}
