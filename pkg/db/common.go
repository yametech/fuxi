package db

func FindRole(roleId int) (*Role, error) {
	var role Role
	if err := DB.Find(role).Where("ID = ?", roleId).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func FindPermissionByName(roleName string) (*Permission, error) {
	var permission Permission
	if err := DB.Find(permission).
		Where("Name = ?", roleName).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (c *User) FindCurrentUserPermissionValue() (uint32, error) {
	role, err := FindRole(c.RoleId)
	if err != nil {
		return 0, err
	}
	p, err := FindPermissionByName(role.RoleName)
	if err != nil {
		return 0, err
	}
	return p.Value, nil
}

func FindUserByName(name string) (*User, error) {
	var user User
	if err := DB.Find(&user).
		Where("name = ?", name).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
