package main

import "fmt"

func dbdebug() {
	var err error

	Wh := NewInstance()
	GoodWh := NewInstance()
	WrongWh := NewInstance()
	GoodWh.Coordinates = "0.324234 0.3423423"
	WrongWh.Coordinates = "hgvldfbgasdjfbgskldfbgsjdfbgsjdfbgsljdfbglsjdbfgksbdflgsbdfgsdffbgsjhdfgbfgbsdkjfgb"

	err = Wh.Create()
	fmt.Println(err)
	err = GoodWh.Create()
	fmt.Println(err)
	err = WrongWh.Create()
	fmt.Println(err)
	WrongWh = NewInstance()
	WrongWh.Create()

	GoodWh.Coordinates = "0.1928743 0.123123"
	WrongWh.Coordinates = "hgvldfbgasdjfbgskldfbgsjdfbgsjdfbgsljdfbglsjdbfgksbdflgsbdfgsdffbgsjhdfgbfgbsdkjfgb"

	err = Wh.Update()
	fmt.Println(err)
	err = GoodWh.Update()
	fmt.Println(err)
	err = WrongWh.Update()
	fmt.Println(err)

	fmt.Println("-----------------------------------------------")
	check := NewInstance()
	for i := Wh.IdNum; i <= WrongWh.IdNum; i++ {
		check.Get(int64(i))
		fmt.Println(check)
	}
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")

	err = Wh.AddPart(NewInstancePart())
	fmt.Println(err)
	GoodPart := NewInstancePart()
	WrongPart := NewInstancePart()
	Wh.AddPart(GoodPart)
	Wh.AddPart(WrongPart)

	GoodPart.Capacity = 10000
	err = GoodPart.Update()
	fmt.Println(err)
	WrongPart.Capacity = 10000
	WrongPart.Type = "asfasmfk"
	err = WrongPart.Update()
	fmt.Println(err)

	fmt.Println("-----------------------------------------------")
	Parts, err := Wh.GetParts()
	if err != nil {
		panic(err)
	}
	for _, part := range Parts {
		fmt.Println(part)
	}
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")

	WhInfo, err := Wh.GetInfo()
	fmt.Println(err)
	WhInfo.Adress = "sadasd"
	WhInfo, err = Wh.UpdateInfo(WhInfo)
	fmt.Println(err)

	WrongInfo, _ := WrongWh.GetInfo()
	WrongInfo.ContactNumber = "823109238102830129"
	WrongInfo, err = WrongWh.UpdateInfo(WrongInfo)
	fmt.Println(err)

	fmt.Println("-----------------------------------------------")
	fmt.Println(WhInfo)
	fmt.Println(WrongInfo)
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")

	It := NewItem()
	GoodIt := NewItem()
	WrongIt := NewItem()

	err = It.Create()
	fmt.Println(err)
	GoodIt.Name = "Item"
	err = GoodIt.Create()
	fmt.Println(err)
	WrongIt.Name = "112343242353464576457"
	err = WrongIt.Create()
	fmt.Println(err)
	WrongIt.Name = "Wrong"
	_ = WrongIt.Create()
	fmt.Println("-----------------------------------------------")

	GoodIt.Name = "asda"
	err = GoodIt.Update()
	fmt.Println(err)
	WrongIt.Name = "fskdgskdasdasdgndgsdg"
	err = WrongIt.Update()
	fmt.Println(err)
	fmt.Println("-----------------------------------------------")
	res, _ := con.Query("select * from Items where IdNum >= ? order by IdNum", It.IdNum)
	for res.Next() {
		it := NewItem()
		res.Scan(&it.IdNum, &it.Id, &it.Size, &it.vendorId, &it.Name)
		fmt.Println(it)
	}
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")

	Pr := NewPermission()
	GoodPr := NewPermission()
	WrongPr := NewPermission()

	err = Pr.Create()
	fmt.Println(err)
	GoodPr.Code = 15
	err = GoodPr.Create()
	fmt.Println(err)
	WrongPr.Code = 15
	err = WrongPr.Create()
	fmt.Println(err)
	WrongPr.Code = 1
	WrongPr.Create()

	GoodPr.Code = 17
	err = GoodPr.Update()
	fmt.Println(err)
	WrongPr.Code = 17
	err = WrongPr.Update()
	fmt.Println(err)
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")

	Rl := NewRole()
	GoodRl := NewRole()
	WrongRl := NewRole()

	err = Rl.Create()
	fmt.Println(err)
	GoodRl.Name = "test"
	err = GoodRl.Create()
	fmt.Println(err)
	WrongRl.Name = "tesdfsdfsdfsdfsfsdfst"
	err = WrongRl.Create()
	fmt.Println(err)
	WrongRl.Name = "name"
	WrongRl.Create()

	GoodRl.Name = "test1"
	err = GoodRl.Update()
	fmt.Println(err)
	WrongRl.Name = "testtesttesttesttesttesttesttesttesttesttesttesttesttest"
	err = WrongRl.Update()
	fmt.Println(err)
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")

	Mn := NewManager()
	GoodMn := NewManager()
	WrongMn := NewManager()

	Mn.roleId = "Rl_1"
	GoodMn.roleId = "Rl_1"
	WrongMn.roleId = "Rl_1"

	err = Mn.Create()
	fmt.Println(err)
	GoodMn.ContactNumber = "sadasdas"
	GoodMn.Create()
	fmt.Println(err)
	WrongMn.ContactNumber = "asufhasjfksdfksldfksdfsd"
	err = WrongMn.Create()
	fmt.Println(err)
	WrongMn.ContactNumber = ""
	WrongMn.Create()

	GoodMn.ContactNumber = "test1"
	err = GoodMn.Update()
	fmt.Println(err)
	WrongMn.ContactNumber = "testtesttesttesttesttesttesttesttesttesttesttesttesttest"
	err = WrongMn.Update()
	fmt.Println(err)
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")

	Ac := NewAction()
	GoodAc := NewAction()
	WrongAc := NewAction()

	Ac.instId = "In_1"
	GoodAc.instId = "In_1"
	WrongAc.instId = "In_1"

	Ac.itemId = "It_1"
	GoodAc.itemId = "It_1"
	WrongAc.itemId = "It_1"

	Ac.managerId = "Mn_1"
	GoodAc.managerId = "Mn_1"
	WrongAc.managerId = "Mn_1"

	err = Ac.Do()
	fmt.Println(err)
	fmt.Println(Ac.Id)
	err = GoodAc.Do()
	fmt.Println(err)
	err = WrongAc.Do()
	fmt.Println(err)
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")
	err = Wh.Delete()
	fmt.Println(err)
	err = GoodPart.Delete()
	fmt.Println(err)
	err = It.Delete()
	fmt.Println(err)
	err = Pr.Delete()
	fmt.Println(err)
	err = Rl.Delete()
	fmt.Println(err)
	err = Mn.Delete()
	fmt.Println(err)
	fmt.Println(GoodAc.Id)
	err = GoodAc.Cancel()
	fmt.Println(err)
}
