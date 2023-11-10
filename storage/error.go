package storage

func CheckError(err error) error {
    if err != nil {
        panic(err)
    }
    
    return err
}
