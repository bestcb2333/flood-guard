package util

import "time"

var Loc *time.Location

func Init() error {
	var err error
	Loc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}
	return nil
}
