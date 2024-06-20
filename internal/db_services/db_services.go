package db_services

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var noAffectedRows = errors.New("CError: There are no affected rows")

// TODO: Реворк ролей
// TODO: Реворк ролей
// TODO: Реворк ролей

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

func ConvertStringToType(str string, t reflect.Type) (reflect.Value, error) {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return reflect.Value{}, err
		}

		return reflect.ValueOf(val).Convert(t), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(val).Convert(t), nil
	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(val).Convert(t), nil
	case reflect.Bool:
		val, err := strconv.ParseBool(str)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(val).Convert(t), nil
	case reflect.String:
		return reflect.ValueOf(str).Convert(t), nil
	default:
		return reflect.Value{}, fmt.Errorf("unsupported type: %s", t.Kind())
	}
}

func FillStruct(s any, data []string) (err error) {
	v := reflect.ValueOf(s).Elem() // Получаем отраженное значение
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if i < len(data) && field.CanSet() {
			value := reflect.ValueOf(data[i])
			if value.String() != "" {
				value, err = ConvertStringToType(value.String(), field.Type())
				field.Set(value)
			}
		}
	}
	return err
}

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

func TableGet(tableName, start, limit string) ([]map[string]any, error) {
	tableName = strings.Replace(tableName, " ", "", -1)
	start = strings.Replace(start, " ", "", -1)
	limit = strings.Replace(limit, " ", "", -1)

	if start == "" {
		start = "1"
	}

	if limit == "" {
		limit = "20"
	}

	query := "select * from " + tableName + " where IdNum >= " + start + " order by IdNum desc limit " + limit
	rows, err := con.Query(query)
	if err != nil {
		Log(tableName+"GetAll", err)
		return []map[string]any{}, err
	}

	columns, _ := rows.Columns()
	values := make([]any, len(columns))
	valuePtrs := make([]any, len(columns))

	var res []map[string]any

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		rowMap := make(map[string]any)
		for i, col := range columns {
			val := values[i]

			switch val.(type) {
			case []byte:
				val = string(val.([]byte))
			}

			rowMap[col] = val
		}

		res = append(res, rowMap)
	}

	Log(tableName+"GetAll", err)
	return res, err
}

func TableGetColumns(tableName string) ([]string, error) {
	tableName = strings.Replace(tableName, " ", "", -1)
	query := "select * from " + tableName + " limit 1"
	rows, err := con.Query(query)
	if err != nil {
		Log(tableName+"GetAllColumns", err)
		return []string{}, err
	}

	cols, _ := rows.Columns()
	if err != nil {
		Log(tableName+"GetAllColumns", err)
		return []string{}, err
	}

	return cols, err
}

func Log(funcName string, funcErr error) {
	sErr := ""
	if funcErr == nil {
		sErr = "null"
	} else {
		sErr = funcErr.Error()
	}
	query := `insert into Logging(funcName, Date, errText) values (?, ?, ?)`
	LogDate := time.Now().Format("2006-01-02 15:04:05")
	_, err := con.Exec(query, funcName, LogDate, sErr)
	if err != nil {
		panic(err)
	}
}

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

func NewInstance() *Instances {
	return &Instances{0, "In_Null", "Склад", "00.000000 00.000000", 0, true}
}

func (In *Instances) CheckData() (bool, error) {
	var IsCorrect bool
	var err error

	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (In *Instances) Get(Id any) error {
	var res *sql.Row
	switch Id.(type) {
	case int64:
		res = InGet.QueryRow(Id, "")
	case string:
		res = InGet.QueryRow(0, Id)
	}

	finErr := res.Scan(&In.IdNum, &In.Id, &In.Type, &In.Coordinates, &In.Capacity, &In.IsAvailable)

	Log("InstancesGet", finErr)
	return finErr
}

func (In *Instances) Create() error {
	IsCorrect, finErr := In.CheckData()

	if !IsCorrect {
		Log("InstancesCreate", finErr)
		return finErr
	}

	res, finErr := InCreate.Exec(In.Type, In.Coordinates, In.Capacity, In.IsAvailable)

	Log("InstancesCreate", finErr)

	if finErr == nil {
		In.IdNum, _ = res.LastInsertId()
		finErr = In.Get(In.IdNum)
	}

	return finErr
}

func (In *Instances) Delete() error {
	res, finErr := InDelete.Exec(In.Id)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}

	Log("InstancesDelete", finErr)
	*In = *NewInstance()

	return finErr
}

func (In *Instances) Update() error {
	IsCorrect, finErr := In.CheckData()

	if !IsCorrect {
		Log("InstancesUpdate", finErr)
		return finErr
	}

	res, finErr := InUpdate.Exec(In.Type, In.Coordinates, In.Capacity, In.IsAvailable, In.Id)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}
	Log("InstancesUpdate", finErr)

	err := In.Get(In.Id)
	if finErr == nil {
		finErr = err
	}

	return finErr
}

//-----------------------------------------------------------------------------------------------------

func NewInstanceInfo() *instancesInfo {
	return new(instancesInfo)
}

func (iI *instancesInfo) CheckData() (bool, error) {
	var IsCorrect bool
	var err error

	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (iI *instancesInfo) Get(Id any) error {
	var res *sql.Row
	switch Id.(type) {
	case int64:
		res = iIGet.QueryRow(Id, "")
	case string:
		res = iIGet.QueryRow(0, Id)
	}

	finErr := res.Scan(&iI.IdNum, &iI.instanceId, &iI.ContactNumber, &iI.Email, &iI.WorkingHours, &iI.Length, &iI.Width, &iI.Height, &iI.Volume, &iI.City, &iI.Adress)
	Log("instancesInfoGet", finErr)
	return finErr
}

func (iI *instancesInfo) Create() error {
	return nil
}
func (iI *instancesInfo) Delete() error {
	return nil
}

func (iI *instancesInfo) Update() error {
	IsCorrect, finErr := iI.CheckData()

	if !IsCorrect {
		Log("instancesInfoUpdate", finErr)
		return finErr
	}

	res, finErr := iIUpdate.Exec(iI.ContactNumber, iI.Email, iI.WorkingHours, iI.Length, iI.Width, iI.Height, iI.Volume, iI.City, iI.Adress, iI.instanceId)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}
	Log("instancesInfoUpdate", finErr)

	err := iI.Get(iI.instanceId)
	if finErr == nil {
		finErr = err
	}

	return finErr
}

//-----------------------------------------------------------------------------------------------------

func (In *Instances) GetParts() ([]*instanceParts, error) {
	res, finErr := InGetParts.Query(In.Id)

	parts := []*instanceParts{}
	for res.Next() {
		iP := NewInstancePart()
		err := res.Scan(&iP.IdNum, &iP.Id, &iP.Type, &iP.itemMaxSize, &iP.Capacity, &iP.instanceId)
		if err != nil {
			if finErr == nil {
				finErr = err
			}
			break
		}

		parts = append(parts, iP)
	}

	Log("InstancesGetParts", finErr)
	return parts, finErr
}

func (In *Instances) AddPart(iP *instanceParts) error {
	IsCorrect, finErr := iP.CheckData()

	if !IsCorrect {
		Log("InstancesAddPart", finErr)
		return finErr
	}

	res, finErr := iPCreate.Exec(iP.Type, iP.itemMaxSize, iP.Capacity, In.Id)

	Log("InstancesAddPart", finErr)

	if finErr == nil {
		iP.IdNum, _ = res.LastInsertId()
		finErr = iP.Get(iP.Id)
	}

	return finErr
}

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

func NewInstancePart() *instanceParts {
	return &instanceParts{Type: "Склад"}
}

func (iP *instanceParts) CheckData() (bool, error) {
	var IsCorrect bool
	var err error

	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (iP *instanceParts) Get(Id any) error {
	var res *sql.Row
	switch Id.(type) {
	case int64:
		res = iPGet.QueryRow(Id, "")
	case string:
		res = iPGet.QueryRow(0, Id)
	}

	finErr := res.Scan(&iP.IdNum, &iP.Id, &iP.Type, &iP.itemMaxSize, &iP.Capacity, &iP.instanceId)

	Log("instancePartsGet", finErr)
	return finErr
}

func (iP *instanceParts) Create() error {
	IsCorrect, finErr := iP.CheckData()

	if !IsCorrect {
		Log("instancePartsCreate", finErr)
		return finErr
	}

	res, finErr := iPCreate.Exec(iP.Type, iP.itemMaxSize, iP.Capacity, iP.instanceId, iP.Id)

	Log("InstancesCreate", finErr)

	if finErr == nil {
		iP.IdNum, _ = res.LastInsertId()
		finErr = iP.Get(iP.IdNum)
	}

	return finErr
}

func (iP *instanceParts) Delete() error {
	res, finErr := iPDelete.Exec(iP.Id)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}

	Log("InstancesDelete", finErr)
	*iP = *NewInstancePart()

	return finErr
}

func (iP *instanceParts) Update() error {
	IsCorrect, finErr := iP.CheckData()

	if !IsCorrect {
		Log("instancePartsUpdate", finErr)
		return finErr
	}

	res, finErr := iPUpdate.Exec(iP.Type, iP.itemMaxSize, iP.Capacity, iP.instanceId, iP.Id)
	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}
	Log("instancePartsUpdate", finErr)

	err := iP.Get(iP.Id)
	if finErr == nil {
		finErr = err
	}

	return finErr
}

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

func NewItem() *Items {
	return new(Items)
}

func (It *Items) CheckData() (bool, error) {
	var IsCorrect bool
	var err error

	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (It *Items) Get(Id any) error {
	var res *sql.Row
	switch Id.(type) {
	case int64:
		res = ItGet.QueryRow(Id, "")
	case string:
		res = ItGet.QueryRow(0, Id)
	}

	finErr := res.Scan(&It.IdNum, &It.Id, &It.Size, &It.vendorId, &It.Name)

	Log("ItemsGet", finErr)
	return finErr
}

func (It *Items) Create() error {
	IsCorrect, finErr := It.CheckData()

	if !IsCorrect {
		Log("ItemsCreate", finErr)
		return finErr
	}

	res, finErr := ItCreate.Exec(It.Size, It.vendorId, It.Name)

	Log("ItemsCreate", finErr)

	if finErr == nil {
		It.IdNum, _ = res.LastInsertId()
		finErr = It.Get(It.IdNum)
	}

	return finErr
}

func (It *Items) Delete() error {
	res, finErr := ItDelete.Exec(It.Id)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}

	Log("ItemsDelete", finErr)
	*It = *NewItem()

	return finErr
}

func (It *Items) Update() error {
	IsCorrect, finErr := It.CheckData()

	if !IsCorrect {
		Log("ItemsUpdate", finErr)
		return finErr
	}

	res, finErr := ItUpdate.Exec(It.Size, It.vendorId, It.Name, It.Id)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}
	Log("ItemsUpdate", finErr)

	err := It.Get(It.Id)
	if finErr == nil {
		finErr = err
	}

	return finErr
}

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

func NewPermission() *Permissions {
	return new(Permissions)
}

func (Pr *Permissions) CheckData() (bool, error) {
	var IsCorrect bool
	var err error

	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (Pr *Permissions) Get(Id any) error {
	var res *sql.Row
	switch Id.(type) {
	case int64:
		res = PrGet.QueryRow(Id, "")
	case string:
		res = PrGet.QueryRow(0, Id)
	}

	finErr := res.Scan(&Pr.IdNum, &Pr.Id, &Pr.Code, &Pr.Name, &Pr.tableName)

	Log("PermissionsGet", finErr)
	return finErr
}

func (Pr *Permissions) Create() error {
	IsCorrect, finErr := Pr.CheckData()

	if !IsCorrect {
		Log("PermissionsCreate", finErr)
		return finErr
	}

	res, finErr := PrCreate.Exec(Pr.Code, Pr.Name, Pr.tableName)

	Log("PermissionsCreate", finErr)

	if finErr == nil {
		Pr.IdNum, _ = res.LastInsertId()

		finErr = Pr.Get(Pr.IdNum)
	}

	return finErr
}

func (Pr *Permissions) Delete() error {
	res, finErr := PrDelete.Exec(Pr.Id)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}

	Log("PermissionsDelete", finErr)
	*Pr = *NewPermission()

	return finErr
}

func (Pr *Permissions) Update() error {
	IsCorrect, finErr := Pr.CheckData()

	if !IsCorrect {
		Log("PermissionsUpdate", finErr)
		return finErr
	}
	res, finErr := PrUpdate.Exec(Pr.Code, Pr.Name, Pr.tableName, Pr.Id)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}
	Log("PermissionsUpdate", finErr)

	err := Pr.Get(Pr.Id)
	if finErr == nil {
		finErr = err
	}

	return finErr
}

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

func NewRole() *Roles {
	return new(Roles)
}

func (Rl *Roles) CheckData() (bool, error) {
	var IsCorrect bool
	var err error

	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (Rl *Roles) Get(Id any) error {
	var res *sql.Row
	switch Id.(type) {
	case int64:
		res = RlGet.QueryRow(Id, "")
	case string:
		res = RlGet.QueryRow(0, Id)
	}

	finErr := res.Scan(&Rl.IdNum, &Rl.Id, &Rl.Name)

	Log("RolesGet", finErr)
	return finErr
}

func (Rl *Roles) Create() error {
	IsCorrect, finErr := Rl.CheckData()

	if !IsCorrect {
		Log("RolesCreate", finErr)
		return finErr
	}

	res, finErr := RlCreate.Exec(Rl.Name)

	Log("RolesCreate", finErr)

	if finErr == nil {
		Rl.IdNum, _ = res.LastInsertId()
		finErr = Rl.Get(Rl.IdNum)
	}

	return finErr
}

func (Rl *Roles) Delete() error {
	res, finErr := RlDelete.Exec(Rl.Id)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}

	Log("RolesDelete", finErr)
	*Rl = *NewRole()

	return finErr
}

func (Rl *Roles) Update() error {
	IsCorrect, finErr := Rl.CheckData()

	if !IsCorrect {
		Log("RolesUpdate", finErr)
		return finErr
	}

	res, finErr := RlUpdate.Exec(Rl.Name, Rl.Id)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}
	Log("RolesUpdate", finErr)

	err := Rl.Get(Rl.Id)
	if finErr == nil {
		finErr = err
	}

	return finErr
}

func (Rl *Roles) GetPerms() ([]*Permissions, error) {
	res, finErr := RlGetPerms.Query(Rl.Id)

	Pr := NewPermission()
	perms := []*Permissions{}
	for res.Next() {
		err := res.Scan(&Pr.IdNum, &Pr.Id, &Pr.Code, &Pr.Name, &Pr.tableName)
		if err != nil {
			if finErr == nil {
				finErr = err
			}
			break
		}

		perms = append(perms, Pr)
	}

	Log("RolesGetPerms", finErr)
	return perms, finErr
}

func (Rl *Roles) AddPerms(permIds []string) ([]*Permissions, error) {
	var finErr error

	for _, PrId := range permIds {
		_, err := RlAddPerm.Exec(Rl.Id, PrId)
		if err != nil {
			finErr = err
			break
		}
	}

	Log("RolesAddPerms", finErr)

	perms := []*Permissions{}
	if finErr == nil {
		perms, finErr = Rl.GetPerms()
	}

	return perms, finErr
}

func (Rl *Roles) DelPerms(permIds []string) ([]*Permissions, error) {
	var finErr error

	for _, PrId := range permIds {
		res, err := RlDelPerm.Exec(Rl.Id, PrId)
		if err == nil {
			if res == nil {
				finErr = noAffectedRows
				break
			}
		} else {
			finErr = err
			break
		}
	}

	Log("RolesDelPerms", finErr)

	perms := []*Permissions{}
	if finErr == nil {
		perms, finErr = Rl.GetPerms()
	}

	return perms, finErr
}

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

func NewManager() *Managers {
	return &Managers{Id: "Mn_Null", RoleId: "Rl_2"}
}

func (Mn *Managers) CheckData() (bool, error) {
	var IsCorrect bool
	var err error

	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (Mn *Managers) Get(Id any) error {
	var res *sql.Row
	switch Id.(type) {
	case int64:
		res = MnGet.QueryRow(Id, "")
	case string:
		res = MnGet.QueryRow(0, Id)
	}

	finErr := res.Scan(&Mn.IdNum, &Mn.Id, &Mn.Login, &Mn.Password, &Mn.Name, &Mn.ContactNumber, &Mn.Email, &Mn.RoleId)

	Log("ManagersGet", finErr)
	return finErr
}

func (Mn *Managers) Find(Login string) error {
	res := MnFind.QueryRow(Login)

	finErr := res.Scan(&Mn.IdNum, &Mn.Id, &Mn.Login, &Mn.Password, &Mn.Name, &Mn.ContactNumber, &Mn.Email, &Mn.RoleId)

	Log("ManagersFind", finErr)
	return finErr
}

func (Mn *Managers) Create() error {
	IsCorrect, finErr := Mn.CheckData()

	if !IsCorrect {
		Log("ManagersCreate", finErr)
		return finErr
	}

	res, finErr := MnCreate.Exec(Mn.Login, Mn.Password, Mn.Name, Mn.ContactNumber, Mn.Email, Mn.RoleId)

	Log("ManagersCreate", finErr)

	if finErr == nil {
		Mn.IdNum, _ = res.LastInsertId()
		finErr = Mn.Get(Mn.IdNum)
	}

	return finErr
}

func (Mn *Managers) Delete() error {
	res, finErr := MnDelete.Exec(Mn.Id)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}

	Log("ManagersDelete", finErr)
	*Mn = *NewManager()

	return finErr
}

func (Mn *Managers) Update() error {
	IsCorrect, finErr := Mn.CheckData()

	if !IsCorrect {
		Log("ManagersUpdate", finErr)
		return finErr
	}

	res, finErr := MnUpdate.Exec(Mn.Login, Mn.Password, Mn.Name, Mn.ContactNumber, Mn.Email, Mn.RoleId, Mn.Id)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}
	Log("ManagersUpdate", finErr)

	err := Mn.Get(Mn.Id)
	if finErr == nil {
		finErr = err
	}

	return finErr
}

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

func NewAction() *Actions {
	return &Actions{Type: "Связь"}
}

func (Ac *Actions) CheckData() (bool, error) {
	var IsCorrect bool
	var err error

	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (Ac *Actions) Get(Id any) error {
	var res *sql.Row
	switch Id.(type) {
	case int64:
		res = AcGet.QueryRow(Id, "")
	case string:
		res = AcGet.QueryRow(0, Id)
	}

	date := ""
	finErr := res.Scan(&Ac.IdNum, &Ac.Id, &Ac.Type, &date, &Ac.itemId, &Ac.instId, &Ac.ManagerId)
	Ac.Date, _ = time.Parse("2006-01-02 15:04:05", date)

	Log("ActionsGet", finErr)
	return finErr
}

func (Ac *Actions) Do() error {
	IsCorrect, finErr := Ac.CheckData()

	if !IsCorrect {
		Log("ManagersCreate", finErr)
		return finErr
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	res, finErr := AcDo.Exec(Ac.Type, date, Ac.itemId, Ac.instId, Ac.ManagerId)

	Log("ActionDo", finErr)

	if finErr == nil {
		Ac.IdNum, _ = res.LastInsertId()
		finErr = Ac.Get(Ac.IdNum)
	}

	return finErr
}

func (Ac *Actions) Cancel() error {
	res, finErr := AcCancel.Exec(Ac.Id)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}

	Log("ActionsCancel", finErr)
	*Ac = *NewAction()

	return finErr
}

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

func NewSession() *Sessions {
	return new(Sessions)
}

func (Sn *Sessions) CheckData() (bool, error) {
	var IsCorrect bool
	var err error

	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (Sn *Sessions) Get(token string) error {
	res := SnGet.QueryRow(token)

	finErr := res.Scan(&Sn.IdNum, &Sn.Id, &Sn.Token, &Sn.ManagerId)

	Log("SessionsGet", finErr)
	return finErr
}

func (Sn *Sessions) Create() error {
	IsCorrect, finErr := Sn.CheckData()

	if !IsCorrect {
		Log("SessionsCreate", finErr)
		return finErr
	}

	_, finErr = SnCreate.Exec(Sn.Token, Sn.ManagerId)

	Log("SessionsCreate", finErr)

	if finErr == nil {
		finErr = Sn.Get(Sn.Token)
	}

	return finErr
}

func (Sn *Sessions) Delete() error {
	res, finErr := SnDelete.Exec(Sn.Token)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}

	Log("SessionsDelete", finErr)
	*Sn = *NewSession()

	return finErr
}

func (Sn *Sessions) Update() error {
	IsCorrect, finErr := Sn.CheckData()

	if !IsCorrect {
		Log("SessionsUpdate", finErr)
		return finErr
	}

	res, finErr := SnUpdate.Exec(Sn.ManagerId, Sn.Token)

	if finErr == nil {
		ra, err := res.RowsAffected()
		finErr = err
		if ra == 0 {
			if finErr == nil {
				finErr = noAffectedRows
			}
		}
	}
	Log("SessionsUpdate", finErr)

	err := Sn.Get(Sn.Token)
	if finErr == nil {
		finErr = err
	}

	return finErr
}
