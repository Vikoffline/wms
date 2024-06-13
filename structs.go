package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var con, _ = sql.Open("mysql", "site:DevPassword@/wms")

type Instances struct {
	IdNum       int64
	Id          string
	Type        string
	Coordinates string
	Capacity    int
	IsAvailable bool
}

var InGet, _ = con.Prepare("select * from Instances where IdNum = ?")
var InGetParts, _ = con.Prepare("select * from instanceParts where instanceId = ?")
var InCreate, _ = con.Prepare("insert into Instances(Type, Coordinates, Capacity, IsAvailable) values (?, ?, ?, ?);")
var InUpdate, _ = con.Prepare("update Instances set Type = ?, Coordinates = ?, Capacity = ?, IsAvailable = ? where Id = ?;")
var InDelete, _ = con.Prepare("delete from Instances where Id = ?;")

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

type instancesInfo struct {
	IdNum         int64
	instanceId    string
	ContactNumber string
	Email         string
	WorkingHours  string
	Length        int
	Width         int
	Height        int
	Volume        int
	City          string
	Adress        string
}

var iIGet, _ = con.Prepare("select * from instancesInfo where instanceId = ?")
var iIUpdate, _ = con.Prepare("update instancesInfo set ContactNumber = ?, Email = ?, WorkingHours = ?, Length = ?, Width = ?, Height = ?, Volume = ?, City = ?, Adress = ? where instanceId = ?;")

// var iIDelete, _ = con.Prepare("delete from instancesInfo where instanceId = ?);")
// var iICreate, _ = con.Prepare("insert into instancesInfo(ContactNumber, Email, WorkingHours, Length, Width, Height, Volume, City, Adress) values (?, ?, ?, ?, ?, ?, ?, ?, ?);")

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

type instanceParts struct {
	IdNum       int64
	Id          string
	Type        string
	itemMaxSize int
	Capacity    int
	instanceId  string
}

var iPGet, _ = con.Prepare("select * from instanceParts where IdNum = ?")
var iPCreate, _ = con.Prepare("insert into instanceParts(Type, itemMaxSize, Capacity, instanceId) values (?, ?, ?, ?);")
var iPUpdate, _ = con.Prepare("update instanceParts set Type = ?, itemMaxSize = ?, Capacity = ?, instanceId = ? where Id = ?;")
var iPDelete, _ = con.Prepare("delete from instanceParts where Id = ?;")

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

type Items struct {
	IdNum    int64
	Id       string
	Size     int
	vendorId string
	Name     string
}

var ItGet, _ = con.Prepare("select * from Items where IdNum = ?")
var ItCreate, _ = con.Prepare("insert into Items(Size, VendorId, Name) values (?, ?, ?);")
var ItUpdate, _ = con.Prepare("update Items set Size = ?, VendorId = ?, Name = ? where Id = ?;")
var ItDelete, _ = con.Prepare("delete from Items where Id = ?;")

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

type Permissions struct {
	IdNum     int64
	Id        string
	Code      int
	Name      string
	tableName string
}

var PrGet, _ = con.Prepare("select * from Permissions where IdNum = ?")
var PrCreate, _ = con.Prepare("insert into Permissions(Code, Name, tableName) values (?, ?, ?);")
var PrUpdate, _ = con.Prepare("update Permissions set Code = ?, Name = ?, tableName = ? where Id = ?;")
var PrDelete, _ = con.Prepare("delete from Permissions where Id = ?;")

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

type Roles struct {
	IdNum int64
	Id    string
	Name  string
}

var RlGet, _ = con.Prepare("select * from Roles where IdNum = ?")
var RlGetPerms, _ = con.Prepare("select * from Roles_Perms where roleId = ?")
var RlCreate, _ = con.Prepare("insert into Roles(Name) values (?);")
var RlUpdate, _ = con.Prepare("update Roles set Name = ? where Id = ?;")
var RlDelete, _ = con.Prepare("delete from Roles where Id = ?;")
var RlAddPerm, _ = con.Prepare("insert into Roles_Perms(roleId, permId) values (?, ?);")
var RlDelPerm, _ = con.Prepare("delete from Roles_Perms where permId = ?;")

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

type Roles_Perms struct {
	IdNum  int64
	Id     string
	roleId string
	permId string
}

var RPGet, _ = con.Prepare("select * from Roles_Perms where IdNum = ?")
var RPCreate, _ = con.Prepare("insert into Roles_Perms(roleId, permId) values (?, ?);")
var RPUpdate, _ = con.Prepare("update Roles_Perms set roleId = ?, permId = ? where Id = ?;")
var RPDelete, _ = con.Prepare("delete from Roles_Perms where Id = ?;")

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

type Managers struct {
	IdNum         int64
	Id            string
	Name          string
	ContactNumber string
	Email         string
	roleId        string
}

var MnGet, _ = con.Prepare("select * from Managers where IdNum = ?")
var MnCreate, _ = con.Prepare("insert into Managers(Name, ContactNumber, Email, roleId) values (?, ?, ?, ?);")
var MnUpdate, _ = con.Prepare("update Managers set Name = ?, ContactNumber = ?, Email = ?, roleId = ? where Id = ?;")
var MnDelete, _ = con.Prepare("delete from Managers where Id = ?;")

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

type Actions struct {
	IdNum     int64
	Id        string
	Type      string
	Date      time.Time
	itemId    string
	instId    string
	managerId string
}

var AcGet, _ = con.Prepare("select * from Actions where IdNum = ?")
var AcDo, _ = con.Prepare("insert into Actions(Type, Date, itemId, instId, managerId) values (?, ?, ?, ?, ?);")
var AcCancel, _ = con.Prepare("delete from Actions where Id = ?;")
var AcUpdate, _ = con.Prepare("update Actions set Type = ?, Date = ?, itemId = ?, instId = ?, managerId = ? where Id = ?;")

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

type Sessions struct {
	IdNum      int64
	Id         string
	secret_key string
	managerId  string
}

var SnGet, _ = con.Prepare("select * from Sessions where IdNum = ?")
var SnCreate, _ = con.Prepare("insert into Sessions(secret_key, managerId) values (?, ?);")
var SnUpdate, _ = con.Prepare("update Sessions set secret_key = ?, managerId = ? where Id = ?;")
var SnDelete, _ = con.Prepare("delete from Sessions where Id = ?;")

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

type Logging struct {
	IdNum    int64
	Id       string
	funcName string
	Date     time.Time
	errText  string
}

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
