package utils

// errがnilでなければerrを返す。errがnilならnilを返す。
func CheckError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
