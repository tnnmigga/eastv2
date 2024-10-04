package account

type Account struct {
	ID        int64  `gorm:"column:id;primary_key;"`
	Account   string `gorm:"column:account;uniqueIndex:account_IDX;"`
	PwdDigest string `gorm:"column:pwd_digest;"`
}
