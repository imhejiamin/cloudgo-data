package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {

    var user UserInfo
    counts, err := engine.Count(&user)
    checkErr(err)
    u.UID = int(counts)
    engine.Insert(u)

    return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
    //dao := userInfoDao{mydb}
    var users []UserInfo
    err := engine.Find(&users)
    checkErr(err)
    return users//dao.FindAll()
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
    //dao := userInfoDao{mydb}
    var user UserInfo
    user.UID = id
    _, err := engine.Get(&user)
    checkErr(err)
    return &user//dao.FindByID(id)
}
