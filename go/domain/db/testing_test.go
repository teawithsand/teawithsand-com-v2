package db_test

/*
// this is de-facto tested by all stuff that uses it
// so no need to test db per se

// TODO(teawithsand): move this test to somewhere else, since it's not unit test but rather integration test
func Test_WithDBUtil(t *testing.T) {
	testutil.SetTestENV()

	var cfg db.Config
	util.ReadConfig(&cfg)

	tdb, err := db.MakeTestingDB(context.TODO(), cfg)
	if err != nil {
		t.Error(err)
		return
	}

	err = tdb.Close()
	if err != nil {
		t.Error(err)
		return
	}
}
*/
