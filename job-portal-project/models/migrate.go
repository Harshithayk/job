package models

func (s *Conn) AutoMigrate() error {

	// err := s.db.Migrator().DropTable(&User{}, &Company{}, &Job{})
	// if err != nil {
	// 	return err
	// }

	err := s.db.Migrator().AutoMigrate(&User{}, &Company{}, &Job{})
	if err != nil {
		return err
	}
	return nil
}
