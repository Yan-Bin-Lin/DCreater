package database

import (
	"database/sql"
)

func GetRoot(page string) (blogs []BlogList, err error) {
	err = db.SQL("call get_root(?)", page).Find(&blogs)
	if err != nil {
		return nil, err
	} else if len(blogs) == 0 {
		return nil, nil
	}
	return blogs, nil
}

// return owner page
func GetOwner(owner string) (*OwnerOut, error) {
	ownerData := &OwnerOut{}
	has, err := db.SQL("call get_owner(?)", owner).Get(ownerData)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return ownerData, nil
}

// Update an owner
func CreateOwner(uid, uniquename, nickname, descript string) error {
	return checkAffect(db.Exec("call create_owner(?, ?, ?, ?)", uid, uniquename, nickname, descript))
}

// Update an owner if no need to update uniquename, newuniname should be ""
func UpdateOwner(uid, oid, uniquename, newuniname, nickname, descript string) error {
	return checkResult(db.Exec("call update_owner(?, ?, ?, ?, ?, ?)", uid, oid, uniquename, newuniname, nickname, descript))
}

// delete an owner
func DelOwner(uid, oid, owner string) error {
	return checkAffect(db.Exec("call del_owner(?, ?, ?)", uid, oid, owner))
}

// get blog with owner super project and category data
func GetBlog(path string) (*BlogProj, error) {
	blogData := &BlogProj{}
	has, err := db.SQL("call get_blog(?)", path).Get(blogData)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return blogData, nil
}

// Update an blog
func CreateBlog(oid, superid, blog, descript, typeid, superUrl string) error {
	return checkAffect(db.Exec("call create_blog(?, ?, ?, ?, ?, ?)", oid, superid, blog, descript, typeid, superUrl))
}

// Update an owner if no need to update uniquename, newuniname should be ""
func UpdateBlog(oid, superid, newsuperid, bid, blog, newblog, descript, originUrl, newsuperUrl string) error {
	return checkResult(db.Exec("call update_blog(?, ?, ?, ?, ?, ?, ?, ?, ?)", oid, superid, newsuperid, bid, blog, newblog, descript, originUrl, newsuperUrl))
}

// delete a blog
func DelBlog(oid, bid, url string) error {
	return checkAffect(db.Exec("call del_blog(?, ?, ?)", oid, bid, url))
}

// check affect row is > 0 or not
func checkAffect(res sql.Result, err error) error {
	if err == nil {
		count, err := res.RowsAffected()
		if err == nil {
			if count > 0 {
				// affect success
				return nil
			} else {
				// affect fail
				return ERR_TASK_FAIL
			}
		}
		return err
	}
	return err
}

// check database return rror or not
func checkResult(res sql.Result, err error) error {
	if err == nil {
		return nil
	}
	return err
}
