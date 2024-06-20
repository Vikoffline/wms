package db_services

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

var InGet, _ = con.Prepare("select * from Instances where IdNum = ? or Id = ?")
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

var iIGet, _ = con.Prepare("select * from instancesInfo where IdNum = ? or instanceId = ?")
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

var iPGet, _ = con.Prepare("select * from instanceParts where IdNum = ? or Id = ?")
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

var ItGet, _ = con.Prepare("select * from Items where IdNum = ? or Id = ?")
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

var PrGet, _ = con.Prepare("select * from Permissions where IdNum = ? or Id = ?")
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

var RlGet, _ = con.Prepare("select * from Roles where IdNum = ? or Id = ?")
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

var RPGet, _ = con.Prepare("select * from Roles_Perms where IdNum = ? or Id = ?")
var RPCreate, _ = con.Prepare("insert into Roles_Perms(roleId, permId) values (?, ?);")
var RPUpdate, _ = con.Prepare("update Roles_Perms set roleId = ?, permId = ? where Id = ?;")
var RPDelete, _ = con.Prepare("delete from Roles_Perms where Id = ?;")

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

type Managers struct {
	IdNum         int64
	Id            string
	Login         string
	Password      string
	Name          string
	ContactNumber string
	Email         string
	RoleId        string
}

var MnGet, _ = con.Prepare("select * from Managers where IdNum = ? or Id = ?")
var MnGetRoles, _ = con.Prepare("select R.* from Roles R inner join Roles_Perms RP on RP.roleId = R.Id inner join Manages M on M.Id = RP.managerId where M.Id = ?")
var MnFind, _ = con.Prepare("select * from Managers where Login = ?")
var MnCreate, _ = con.Prepare("insert into Managers(Login, Password, Name, ContactNumber, Email, roleId) values (?, ?, ?, ?, ?, ?);")
var MnUpdate, _ = con.Prepare("update Managers set Login = ?, Password = ?, Name = ?, ContactNumber = ?, Email = ?, roleId = ? where Id = ?;")
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
	ManagerId string
}

var AcGet, _ = con.Prepare("select * from Actions where IdNum = ? or Id = ?")
var AcDo, _ = con.Prepare("insert into Actions(Type, Date, itemId, instId, managerId) values (?, ?, ?, ?, ?);")
var AcCancel, _ = con.Prepare("delete from Actions where Id = ?;")
var AcUpdate, _ = con.Prepare("update Actions set Type = ?, Date = ?, itemId = ?, instId = ?, managerId = ? where Id = ?;")

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

type Sessions struct {
	IdNum     int64
	Id        string
	Token     string
	ManagerId string
}

var SnGet, _ = con.Prepare("select * from Sessions where Token = ?")
var SnCreate, _ = con.Prepare("insert into Sessions(Token, managerId) values (?, ?);")
var SnUpdate, _ = con.Prepare("update Sessions set managerId = ? where Token = ?;")
var SnDelete, _ = con.Prepare("delete from Sessions where Token = ?;")
var SnCheckRights, _ = con.Prepare(`select Sn.Id from Sessions Sn inner join Managers Mn on Mn.Id = Sn.managerId inner join Roles Rl on Rl.Id = Mn.roleId inner join Roles_Perms RP on RP.roleId = Rl.Id inner join Permissions Pr on Pr.Id = RP.permId where Sn.Token = ? and ((Pr.Name = "AllRights") or (Pr.Name = ?)) and ((Pr.tableName = "AllTables") or (Pr.tableName = ?))`)

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
