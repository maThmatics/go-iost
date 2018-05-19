package state

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/iost-official/prototype/db/mocks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPoolImpl(t *testing.T) {
	Convey("Test of state pool", t, func() {
		k1 := Key("key1")
		v1 := MakeVInt(123)
		k2 := Key("key2")
		v2 := MakeVFloat(3.14)

		ctl := gomock.NewController(t)
		mockDB := db_mock.NewMockDatabase(ctl)
		mockDB.EXPECT().Get(gomock.Any()).Return(nil, errors.New("not found"))

		db := NewDatabase(mockDB)

		Convey("copy", func() {
			sp1 := NewPool(db)
			sp2 := sp1.Copy()
			So(sp2.(*PoolImpl).parent, ShouldEqual, sp1)
		})
		Convey("put, get, has", func() {
			sp1 := NewPool(db)
			sp2 := sp1.Copy()
			sp1.Put(k1, v1)
			So(sp2.Has(k1), ShouldBeTrue)
			vt, _ := sp2.Get(k1)
			So(vt.EncodeString(), ShouldEqual, v1.EncodeString())
			sp2.Put(k2, v2)
			mockDB.EXPECT().Has(gomock.Any()).Return(false, nil)
			So(sp1.Has(k2), ShouldBeFalse)
			So(sp2.Has(k2), ShouldBeTrue)
			sp2.Put(k1, VNil)
			So(sp2.Has(k1), ShouldBeFalse)
			So(sp1.Has(k1), ShouldBeTrue)
			mockDB.EXPECT().Get(gomock.Any()).Return(nil, errors.New("not found"))
			v, err := sp2.Get(k2)
			So(err, ShouldBeNil)
			So(v, ShouldEqual, v2)
		})
		Convey("hash map", func() {
			sp1 := NewPool(db)
			sp2 := sp1.Copy()
			mockDB.EXPECT().Has(gomock.Any()).AnyTimes().Return(false, nil)
			mockDB.EXPECT().GetHM(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, errors.New("not found"))
			mockDB.EXPECT().Get(gomock.Any()).AnyTimes().Return(nil, errors.New("not found"))
			err := sp1.PutHM(k1, k2, v1)
			vvv, err := sp1.Get(k1)
			So(err, ShouldBeNil)
			So(vvv.EncodeString(), ShouldEqual, "{key2:i123,")
			v, err := sp2.GetHM(k1, k2)
			So(err, ShouldBeNil)
			So(v, ShouldEqual, v1)

			sp2.PutHM(k1, k1, v2)
			vv, err := sp2.Get(k1)
			So(err, ShouldBeNil)
			So(vv.Type(), ShouldEqual, Map)
			So(len(vv.EncodeString()), ShouldEqual, 39)

			sp2.PutHM(k1, k2, v2)
			v3, err := sp2.GetHM(k1, k1)
			So(err, ShouldBeNil)
			So(v3, ShouldEqual, v2)
			v4, err := sp2.GetHM(k1, k2)
			So(err, ShouldBeNil)
			So(v4, ShouldEqual, v2)

			//sp2.PutHM(Key("iost"), Key("a"), MakeVFloat(1000000))
			//vxx, err := sp2.GetHM(Key("iost"), Key("a"))
			//So(err, ShouldBeNil)
			//So(vxx.Type(), ShouldEqual, Float)
		})
	})

}