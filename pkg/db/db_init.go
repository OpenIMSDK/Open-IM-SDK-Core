package db

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"open_im_sdk/pkg/constant"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/utils"
	"sync"
)

type DataBase struct {
	loginUserID string
	dbDir       string
	conn        *gorm.DB
	mRWMutex    sync.RWMutex
}

func NewDataBase(loginUserID string, dbDir string) (*DataBase, error) {
	dataBase := &DataBase{loginUserID: loginUserID, dbDir: dbDir}
	err := dataBase.initDB()
	if err != nil {
		return dataBase, utils.Wrap(err, "initDB failed")
	}
	dataBase.setChatLogFailedStatus()
	return dataBase, nil
}

func (d *DataBase) setChatLogFailedStatus() {
	msgList, err := d.GetSendingMessageList()
	if err != nil {
		log.Error("", "GetSendingMessageList failed ", err.Error())
		return
	}
	for _, v := range msgList {
		v.Status = constant.MsgStatusSendFailed
		err := d.UpdateMessage(v)
		if err != nil {
			log.Error("", "UpdateMessage failed ", err.Error(), v)
			continue
		}
	}
}

func (d *DataBase) initDB() error {
	if d.loginUserID == "" {
		return errors.New("no uid")
	}
	d.mRWMutex.Lock()
	defer d.mRWMutex.Unlock()

	dbFileName := d.dbDir + "OpenIM_" + constant.BigVersion + "_" + d.loginUserID + ".db"
	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	log.Info("open db:", dbFileName)
	if err != nil {
		return utils.Wrap(err, "open db failed")
	}
	d.conn = db
	//db, err := sql.Open("sqlite3", SvrConf.DbDir+"OpenIM_"+uid+".db")
	//sdkLog("open db:", SvrConf.DbDir+"OpenIM_"+uid+".db")
	//if err != nil {
	//	sdkLog("failed open db:", SvrConf.DbDir+"OpenIM_"+uid+".db", err.Error())
	//	return err
	//}
	//u.db = db

	db.AutoMigrate(&LocalFriend{},
		&LocalFriendRequest{},
		&LocalGroup{},
		&LocalGroupMember{},
		&LocalGroupRequest{},
		&LocalErrChatLog{},
		&LocalUser{},
		&LocalBlack{},
		&LocalSeqData{},
		&LocalConversation{},
		&LocalChatLog{},
		&LocalAdminGroupRequest{})
	if !db.Migrator().HasTable(&LocalFriend{}) {
		//log.NewInfo("CreateTable Friend")
		db.Migrator().CreateTable(&LocalFriend{})
	}

	if !db.Migrator().HasTable(&LocalFriendRequest{}) {
		//log.NewInfo("CreateTable FriendRequest")
		db.Migrator().CreateTable(&LocalFriendRequest{})
	}

	if !db.Migrator().HasTable(&LocalGroup{}) {
		//log.NewInfo("CreateTable Group")
		db.Migrator().CreateTable(&LocalGroup{})
	}

	if !db.Migrator().HasTable(&LocalGroupMember{}) {
		//log.NewInfo("CreateTable GroupMember")
		db.Migrator().CreateTable(&LocalGroupMember{})
	}

	if !db.Migrator().HasTable(&LocalGroupRequest{}) {
		//log.NewInfo("CreateTable GroupRequest")
		db.Migrator().CreateTable(&LocalGroupRequest{})
	}

	if !db.Migrator().HasTable(&LocalUser{}) {
		//log.NewInfo("CreateTable User")
		db.Migrator().CreateTable(&LocalUser{})
	}

	if !db.Migrator().HasTable(&LocalBlack{}) {
		//log.NewInfo("CreateTable Black")
		db.Migrator().CreateTable(&LocalBlack{})
	}

	if !db.Migrator().HasTable(&LocalSeqData{}) {
		db.Migrator().CreateTable(&LocalSeqData{})
	}
	if !db.Migrator().HasTable(&LocalConversation{}) {
		db.Migrator().CreateTable(&LocalConversation{})
	}
	if !db.Migrator().HasTable(&LocalChatLog{}) {
		db.Migrator().CreateTable(&LocalChatLog{})
	}
	if !db.Migrator().HasTable(&LocalAdminGroupRequest{}) {
		db.Migrator().CreateTable(&LocalAdminGroupRequest{})
	}
	return nil
}
