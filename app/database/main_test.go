package database

import (
	"testing"
	"time"
)

/*
default value
`user`
+-----+----------+------------+--------+------------+
| uid | username | password   | email  | createtime |
+-----+----------+------------+--------+------------+
|   1 | testapp  | helloworld | a@b.tw |        now |
+-----+----------+------------+--------+------------+
*/

var (
	uniqueName    = "uniname"
	nickNmae      = "nickname"
	uniqueName2   = "uniname2"
	nickNmae2     = "nickname2"
	superprojName = "superproject"
	superprojurl  = "urlfirst"
	projName      = "project"
	projurl       = "urlsecond"
	subprojName   = "subproject"
	subprojurl    = "urlthird"
	blogName      = "blogName"
	blogFile      = "blogFile"
	blogName2     = "blogName2"
	blogFile2     = "blogFile2"
	now           = time.Now()
	user          = &User{1, "testapp", "helloworld", "a@b.tw ", now}
	//ownerOut      = &OwnerOut{Owner{Oid: 1, Uid: 1, Nickname: nickNmae, Uniquename: uniqueName, Createtime: now, Updatetime: now}, Project{}, Blog{}, ""}
)

/*
func TestMain(m *testing.M) {
	initTable()
	code := m.Run()
	//clearTable()
	os.Exit(code)
}


func TestDatabse_PutOwner_Insert(t *testing.T) {
	err := PutOwner("0", "1", uniqueName, nickNmae)
	if err != nil {
		t.Errorf("Fail to insert new owner. apperr is %s\n", err)
	}

	err = PutOwner("0", "1", uniqueName2, nickNmae2)
	if err != nil {
		t.Errorf("Fail to insert new owner. apperr is %s\n", err)
	}

	// test name conflict
	err = PutOwner("0", "0", uniqueName, nickNmae)
	if err == nil {
		t.Errorf("Fail to detect apperr of name conflict. apperr is %s\n", err)
	}
}

func TestDatabase_GetOwner(t *testing.T) {
	// build test table
	var getTest = []struct {
		owner  string    // input
		expect *OwnerOut // expected result
	}{
		{"foo", nil},
		{uniqueName, ownerOut},
	}

	for index, value := range getTest {
		out, err := GetOwner(value.owner)
		if err != nil {
			t.Errorf("In range %d. Fail to get inserted owner. apperr is %s\n", index, err)
		} else if out != nil {
			out.Owner.Createtime = now
			out.Owner.Updatetime = now
			if out.Owner != value.expect.Owner {
				t.Errorf("In range %d. ownerOut miss match. \n\t result = %v \n\t expect = %v", index, out, value.expect)
			}
		}
	}
}

func TestDatabase_PutOwner_Update(t *testing.T) {
	// build test table
	var updateTest = []string{"foo", uniqueName}
	for index, value := range updateTest {
		if err := PutOwner(1, "1", value, nickNmae); err != nil {
			t.Errorf("In range %d. Fail to update owner. apperr is %s\n", index, err)
		}
		out, err := GetOwner(value)
		if err != nil || out == nil {
			t.Errorf("In range %d. Fail to get updated owner. apperr is %s\n", index, err)
		} else if out.Uniquename != value {
			t.Errorf("In range %d. miss match owner. out = %v, espect unique name is %s\n", index, out, value)
		}
	}

	if err := PutOwner(1, "NULL", uniqueName2, nickNmae); err == nil {
		t.Errorf("Fail to detect apperr of name conflict. apperr is %s\n", err)
	}
}

func TestDatabase_DelOwner(t *testing.T) {
	if err := DelOwner("1", "1", "test_delete_name"); err == nil {
		t.Errorf("Fail to detect apperr of no delete. apperr is %s\n", err)
	}

	if err := DelOwner("2", "1", uniqueName2); err != nil {
		t.Errorf("Fail to delete owner %s. apperr is %v\n", uniqueName2, err)
	}

	if out, err := GetOwner(uniqueName2); err != nil {
		t.Errorf("something apperr of get owner %s. apperr is %v", uniqueName2, err)
	} else if out != nil {
		t.Errorf("Fail to delete owner %s. out owner is %v\n", uniqueName2, out)
	}
}

// create a new project
func TestDatabase_PutProject_Insert(t *testing.T) {
	if err := PutProject(uniqueName, []string{superprojName}, "1", "0", "", "0", "first project", superprojurl); err != nil {
		t.Errorf("Fail to insert new project. apperr is %s\n", err)
	}

	if err := PutProject(uniqueName, []string{superprojName, projName}, "1", "1", superprojurl, "0", "second project", projurl); err != nil {
		t.Errorf("Fail to insert new project. apperr is %s\n", err)
	}

	if err := PutProject(uniqueName, []string{superprojName, projName, superprojName}, "1", "2", projurl, "0", "third project", subprojurl); err != nil {
		t.Errorf("Fail to insert new project. apperr is %s\n", err)
	}

	if err := PutProject(uniqueName, []string{superprojName, projName, superprojName}, "1", "0", projurl, "0", "fail project", superprojurl); err == nil {
		t.Errorf("Fail to detect apperr of name conflict. apperr is %s\n", err)
	}
}

func TestDatabase_GetProject(t *testing.T) {
	if projs, err := GetProject(superprojurl); err != nil {
		t.Errorf("In range d. Fail to get inserted owner. apperr is %s\n", err)
	} else {
		t.Log(projs)
	}

	if projs, err := GetProject(projurl); err != nil {
		t.Errorf("In range d. Fail to get inserted owner. apperr is %s\n", err)
	} else {
		t.Log(projs)
	}

	if projs, err := GetProject(subprojurl); err != nil {
		t.Errorf("In range d. Fail to get inserted owner. apperr is %s\n", err)
	} else {
		t.Log(projs)
	}
}

func TestDatabase_PutProject_Update(t *testing.T) {
	// build test table
	var updateTest = []string{"foo", superprojName}
	for index, value := range updateTest {
		if err := PutProject(uniqueName, []string{value}, "1", "0", "", "1", "new description", superprojurl); err != nil {
			t.Errorf("In range %d. Fail to update project. apperr is %s\n", index, err)
		}
		out, err := GetProject(superprojurl)
		if err != nil || out == nil {
			t.Errorf("In range %d. Fail to get updated project. apperr is %s\n", index, err)
		} else if out[0].Project.Name != value {
			t.Errorf("In range %d. miss match project. out = %v, espect unique name is %s\n", index, out, value)
		}
	}

	for index, value := range updateTest {
		if err := PutProject(uniqueName, []string{uniqueName, value}, "1", "2", projurl, "3", "new sub description", subprojurl); err != nil {
			t.Errorf("In range %d. Fail to update project. apperr is %s\n", index, err)
		}
		out, err := GetProject(superprojurl)
		if err != nil || out == nil {
			t.Errorf("In range %d. Fail to get updated project. apperr is %s\n", index, err)
		} else if out[0].Project.Name != value {
			t.Errorf("In range %d. miss match project. out = %v, espect unique name is %s\n", index, out, value)
		}
	}
}

func TestDatabase_DelProiect(t *testing.T) {
	if err := DelProject(uniqueName, projName, "1", "4"); err == nil {
		t.Errorf("Fail to detect apperr of no delete. apperr is %s\n", err)
	}

	if err := DelProject(uniqueName, projName, "1", "2"); err != nil {
		t.Errorf("Fail to delete project %s. apperr is %v\n", projName, err)
	}

	if err := DelProject(uniqueName, subprojName, "1", "3"); err != nil {
		t.Errorf("Fail to delete project %s. apperr is %v\n", subprojName, err)
	}

	if out, err := GetProject(projurl); err != nil {
		t.Errorf("something apperr of get project %s. apperr is %v", projName, err)
	} else if out != nil {
		t.Errorf("Fail to delete project %s. out project is %v\n", projName, out)
	}
}

func TestDatabase_PutBlog_Insert(t *testing.T) {
	if err := PutBlog("1", uniqueName, superprojurl, "0", blogName, "1", "1", "first blog", "1", blogFile, nil); err != nil {
		t.Errorf("Fail to insert new blog. apperr is %s\n", err)
	}

	if err := PutBlog("1", uniqueName, superprojurl, "0", blogName2, "1", "2", "second blog", "1", blogFile2, nil); err != nil {
		t.Errorf("Fail to insert new blog. apperr is %s\n", err)
	}

	if err := PutBlog("1", uniqueName, superprojurl, "0", blogName2, "1", "2", "second blog", "1", blogFile2, nil); err == nil {
		t.Errorf("Fail to detect apperr of name conflict. apperr is %s\n", err)
	}
}

func TestDatabase_GetBlog(t *testing.T) {
	if blog, err := GetBlog(blogFile); err != nil {
		t.Errorf("In range d. Fail to get inserted blog. apperr is %s\n", err)
	} else {
		t.Log(blog)
	}

	if blog, err := GetBlog(blogFile2); err != nil {
		t.Errorf("In range d. Fail to get inserted blog. apperr is %s\n", err)
	} else {
		t.Log(blog)
	}
}

func TestDatabase_PutBlog_Update(t *testing.T) {
	var updateTest = []string{"foo", blogName}
	for index, value := range updateTest {
		if err := PutBlog("1", uniqueName, superprojurl, "1", value, "1", "3", "new description", "1", blogFile, nil); err != nil {
			t.Errorf("In range %d. Fail to update project. apperr is %s\n", index, err)
		}
		out, err := GetBlog(blogFile)
		if err != nil || out == nil {
			t.Errorf("In range %d. Fail to get updated project. apperr is %s\n", index, err)
		} else if out.Blog.Name != value {
			t.Errorf("In range %d. miss match project. out = %v, espect unique name is %s\n", index, out, value)
		}
	}
}

func TestDatabase_DelBlog(t *testing.T) {
	if err := DelBlog("1", "1", "1", blogFile2); err == nil {
		t.Errorf("Fail to detect apperr of no delete. apperr is %s\n", err)
	}

	if err := DelBlog("1", "1", "2", blogFile2); err != nil {
		t.Errorf("Fail to delete project %s. apperr is %v\n", projName, err)
	}

	if out, err := GetBlog(blogFile2); err != nil {
		t.Errorf("something apperr of get blog %s. apperr is %v", projName, err)
	} else if out != nil {
		t.Errorf("Fail to delete blog %s. out blog is %v\n", projName, out)
	}
}

func TestDatabase_PutUser(t *testing.T) {
	PutUser("0", "newuser", "password", "email", "salt")
}

func TestDatabase_Login(t *testing.T) {
	userdata, _ := Login("testapp", "password")
	t.Log(userdata)
}


func TestDatabase_PutUser_Update(t *testing.T) {
	PutUser("2", "newuser", "newpassword", "email", "salt")
}

func TestDatabase_NewAccessToken(t *testing.T) {
	NewAccessToken("1", "token")
}

func TestDatabase_CheckAccessAuth (t *testing.T) {
	b, e := CheckAccessAuth("1", "token", "uniname")
	t.Log(b)
	t.Log(e)
}

// new a database with table
func initTable() {
	var err error
	db, err = xorm.NewEngine(setting.DBs["test"].Driver,
		fmt.Sprintf("%s:%s@/%s?%s", setting.DBs["test"].User, setting.DBs["test"].Password, setting.DBs["test"].Name, setting.DBs["test"].Param))
	if err != nil {
		panic(err)
	}
	db.ShowSQL(true)

			_, err = db.Exec("CREATE DATABASE " + setting.DBs["test"].Name)
			if err != nil {
				return err
			}

		// read schema and create
		file, err := ioutil.ReadFile(setting.WorkPath + "mysql_backup")
		if err != nil {
			return err
		}
		_ = exec.Command("/bin/sh", string(file))

		_, err = db.Exec("USE " + setting.DBs["test"].Name)
		if err != nil {
			return err
		}
}

// drop database
func clearTable() error {
	_, err := db.Exec("DROP DATABASE " + setting.DBs["test"].Name)
	return err
}
*/

func TestConnect(t *testing.T) {
	var blogs []BlogList
	err := db.SQL("call get_root(1)").Find(&blogs)
	if err != nil {
		t.Error(err)
	}
}
