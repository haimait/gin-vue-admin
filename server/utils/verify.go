package utils

var (
	IdVerify               = Rules{"ID": []string{NotEmpty()}}
	ApiVerify              = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify             = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
	LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify         = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	RegisterCasUserVerify  = Rules{"Username": {NotEmpty()}, "Phone": {RegexpMatch(mobileRegex)}, "Password": {NotEmpty()}, "RePassword": {NotEmpty()}, "Captcha": {NotEmpty()}, "CaptchaId": {NotEmpty()}}
	PageInfoVerify         = Rules{"Type": {NotEmpty()}, "Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify         = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify         = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	AutoPackageVerify      = Rules{"PackageName": {NotEmpty()}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify     = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}

	ArticleVoteVerify   = Rules{"CategoryId": {NotEmpty()}, "IDs": {NotEmpty()}}
	ArticleUpdateVerify = Rules{
		"Title":       {NotEmpty()},
		"Description": {NotEmpty()},
		"Company":     {NotEmpty()},
		"Phone":       {NotEmpty()},
		"Author":      {NotEmpty()},
		"CategoryID":  {NotEmpty()},
		"Content":     {NotEmpty()},
	}
)

var mobileRegex = `^1[3456789]\d{9}$` // [3456789] 表示第二个字符为3、4、5、6、7、8、9中的一个； \d{9} 表示接下来的9个字符均为数字
