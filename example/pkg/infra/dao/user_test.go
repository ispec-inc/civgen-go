package dao

import "testing"

import "github.com/stretchr/testify/assert"

func TestUserDao_Get(t *testing.T) {
	t.Helper()
	d := NewUser(db)

	cases := []struct {
		name string
		id   int64
		want int64
		err  bool
	}{
		{
			name: "Found",
			id:   int64(1),
			want: int64(1),
			err:  false,
		},
		{
			name: "NotFound",
			id:   int64(2),
			want: int64(0),
			err:  true,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			if err := prepareTestData("./testdata/User/get.sql"); err != nil {
				t.Error(err)
			}

			opt, aerr := d.Get(tc.id)

			assert.Exactly(t, tc.want, opt.ID)
			if tc.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
			}
		})
	}

}

func TestUserDao_List(t *testing.T) {
	t.Helper()
	d := NewUser(db)

	cases := []struct {
		name string
		ids  []int64
		want int
		err  bool
	}{
		{
			name: "ByIDs",
			ids:  []int64{1},
			want: 1,
			err:  false,
		},
		{
			name: "All",
			ids:  nil,
			want: 1,
			err:  false,
		},
		{
			name: "NotFound",
			ids:  []int64{2},
			want: 0,
			err:  false,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			if err := prepareTestData("./testdata/User/list.sql"); err != nil {
				t.Error(err)
			}

			opt, aerr := d.List(tc.ids)

			assert.Exactly(t, tc.want, len(opt))
			if tc.err {
				assert.Error(t, aerr)
			} else {
				assert.NoError(t, aerr)
			}
		})
	}

}
