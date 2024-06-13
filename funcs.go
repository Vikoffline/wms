package main

import (
	"errors"
	"time"
)

//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------
//-----------------------------------------------------------------------------------------------------

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

	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (In *Instances) Get(IdNum int64) error {
	res := InGet.QueryRow(IdNum)

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
		if res == nil {
			finErr = errors.New("there are no affected rows")
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
		if res == nil {
			finErr = errors.New("there are no affected rows")
		}
	}
	Log("InstancesUpdate", finErr)

	err := In.Get(In.IdNum)
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

	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (In *Instances) GetInfo() (*instancesInfo, error) {
	res := iIGet.QueryRow(In.Id)
	iI := NewInstanceInfo()
	finErr := res.Scan(&iI.IdNum, &iI.instanceId, &iI.ContactNumber, &iI.Email, &iI.WorkingHours, &iI.Length, &iI.Width, &iI.Height, &iI.Volume, &iI.City, &iI.Adress)
	Log("InstancesGetInfo", finErr)
	return iI, finErr
}

func (In *Instances) UpdateInfo(iI *instancesInfo) (*instancesInfo, error) {
	IsCorrect, finErr := In.CheckData()

	if !IsCorrect {
		Log("InstancesUpdateInfo", finErr)
		return iI, finErr
	}

	res, finErr := iIUpdate.Exec(iI.ContactNumber, iI.Email, iI.WorkingHours, iI.Length, iI.Width, iI.Height, iI.Volume, iI.City, iI.Adress, In.Id)

	if finErr == nil {
		if res == nil {
			finErr = errors.New("there are no affected rows")
		}
	}
	Log("InstancesUpdateInfo", finErr)

	iI, err := In.GetInfo()
	if finErr == nil {
		finErr = err
	}

	return iI, finErr
}

//-----------------------------------------------------------------------------------------------------

func (In *Instances) GetParts() ([]*instanceParts, error) {
	res, finErr := InGetParts.Query(In.Id)

	iP := NewInstancePart()
	parts := []*instanceParts{}
	for res.Next() {
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
		finErr = iP.Get(iP.IdNum)
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

	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (iP *instanceParts) Get(IdNum int64) error {
	res := iPGet.QueryRow(iP.IdNum)

	finErr := res.Scan(&iP.IdNum, &iP.Id, &iP.Type, &iP.itemMaxSize, &iP.Capacity, &iP.instanceId)

	Log("instancePartsGet", finErr)
	return finErr
}

func (iP *instanceParts) Delete() error {
	res, finErr := iPDelete.Exec(iP.Id)

	if finErr == nil {
		if res == nil {
			finErr = errors.New("there are no affected rows")
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

	res, finErr := InUpdate.Exec(iP.Type, iP.itemMaxSize, iP.Capacity, iP.instanceId, iP.Id)

	if finErr == nil {
		if res == nil {
			finErr = errors.New("there are no affected rows")
		}
	}
	Log("instancePartsUpdate", finErr)

	err := iP.Get(iP.IdNum)
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

	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (It *Items) Get(IdNum int64) error {
	res := ItGet.QueryRow(It.IdNum)

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
		if res == nil {
			finErr = errors.New("there are no affected rows")
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

	res, finErr := ItUpdate.Exec(It.Size, It.vendorId, It.Name)

	if finErr == nil {
		if res == nil {
			finErr = errors.New("there are no affected rows")
		}
	}
	Log("ItemsUpdate", finErr)

	err := It.Get(It.IdNum)
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

	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (Pr *Permissions) Get(IdNum int64) error {
	res := PrGet.QueryRow(Pr.IdNum)

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
		if res == nil {
			finErr = errors.New("there are no affected rows")
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

	res, finErr := PrUpdate.Exec(Pr.Code, Pr.Name, Pr.tableName)

	if finErr == nil {
		if res == nil {
			finErr = errors.New("there are no affected rows")
		}
	}
	Log("PermissionsUpdate", finErr)

	err := Pr.Get(Pr.IdNum)
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

	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (Rl *Roles) Get(IdNum int64) error {
	res := RlGet.QueryRow(Rl.IdNum)

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
		if res == nil {
			finErr = errors.New("there are no affected rows")
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

	res, finErr := RlUpdate.Exec(Rl.Name)

	if finErr == nil {
		if res == nil {
			finErr = errors.New("there are no affected rows")
		}
	}
	Log("RolesUpdate", finErr)

	err := Rl.Get(Rl.IdNum)
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
				finErr = errors.New("there are no affected rows")
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
	return new(Managers)
}

func (Mn *Managers) CheckData() (bool, error) {
	var IsCorrect bool
	var err error

	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (Mn *Managers) Get(IdNum int64) error {
	res := MnGet.QueryRow(Mn.IdNum)

	finErr := res.Scan(&Mn.IdNum, &Mn.Id, &Mn.Name, &Mn.ContactNumber, &Mn.Email, &Mn.roleId)

	Log("ManagersGet", finErr)
	return finErr
}

func (Mn *Managers) Create() error {
	IsCorrect, finErr := Mn.CheckData()

	if !IsCorrect {
		Log("ManagersCreate", finErr)
		return finErr
	}

	res, finErr := MnCreate.Exec(Mn.Name, Mn.ContactNumber, Mn.Email, Mn.roleId)

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
		if res == nil {
			finErr = errors.New("there are no affected rows")
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

	res, finErr := MnUpdate.Exec(Mn.Name, Mn.ContactNumber, Mn.Email, Mn.roleId, Mn.Id)

	if finErr == nil {
		if res == nil {
			finErr = errors.New("there are no affected rows")
		}
	}
	Log("ManagersUpdate", finErr)

	err := Mn.Get(Mn.IdNum)
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

	IsCorrect = true
	err = nil

	return IsCorrect, err
}

func (Ac *Actions) Get(IdNum int64) error {
	res := AcGet.QueryRow(Ac.IdNum)

	date := ""
	finErr := res.Scan(&Ac.IdNum, &Ac.Id, &Ac.Type, &date, &Ac.itemId, &Ac.instId, &Ac.managerId)
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
	res, finErr := AcDo.Exec(Ac.Type, date, Ac.itemId, Ac.instId, Ac.managerId)

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
		if res == nil {
			finErr = errors.New("there are no affected rows")
		}
	}

	Log("ActionsCancel", finErr)
	*Ac = *NewAction()

	return finErr
}