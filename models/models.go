package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"

	//_ "github.com/mattn/go-sqlite3"
	"github.com/unknwon/com"
	"os"
	"path"
	"time"
)

const (
	_DB_NAME      = "data/beeblog.db"
	_MYSQL_DRIVER = "mysql"
	//_SQLITE3_DRIVER="sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64 //最后操作分类的用户
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_MYSQL_DRIVER, orm.DRMySQL)
	//orm.RegisterDataBase("default",_MYSQL_DRIVER,_DB_NAME,10)
	maxIdle := 10
	maxConn := 10
	orm.RegisterDataBase("default", _MYSQL_DRIVER, "root:root@tcp(localhost:3306)/orm_test?charset=utf8", maxIdle, maxConn)
}

func AddCategory(name string) error {
	// 获取orm对象
	newOrm := orm.NewOrm()
	cate := &Category{Title: name, Created: time.Now().Local(), TopicTime: time.Now().Local()}
	qs := newOrm.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	_, err = newOrm.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTopics() ([]*Topic, error) {
	torm := orm.NewOrm()
	topics := make([]*Topic, 0)
	tqs := torm.QueryTable("topic")
	_, err := tqs.All(&topics)
	return topics, err
}

func GetAllCategories() ([]*Category, error) {
	newOrm := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := newOrm.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func AddTopic(title, content string) error {
	ormer := orm.NewOrm()

	topic := &Topic{
		Title:     title,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}
	_, err := ormer.Insert(topic)
	if err != nil {
		return err
	}
	return nil
}
