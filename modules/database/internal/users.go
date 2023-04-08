package internal

import (
	"fmt"

	"contacts_api/modules/model"
)

func (u *usersRepo) CreateUser(user model.User) (int, error) {
	ins := fmt.Sprintf("insert into users (first_name,last_name,address,email,phone) values(%s);", user.String())
	res, err := u.db.Exec(ins)
	if err != nil {
		return 0, err
	}
	lastID, _ := res.LastInsertId()
	return int(lastID), nil
}

func (u *usersRepo) UpdateUser(user model.User) error {
	_, err := u.db.Exec(
		fmt.Sprintf("update users set first_name = '%s', last_name = '%s', address = '%s', email = '%s', phone = '%s' where id = %d;",
			user.Name, user.Surname, user.Address, user.Email, user.Phone, user.ID),
	)
	return err
}

func (u *usersRepo) DeleteUser(userID int) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback() // nolint:errcheck
		}
	}()

	_, err = tx.Exec(fmt.Sprintf("delete from user_skills where user_id = %d;", userID))
	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf("delete from users where id = %d;", userID))
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (u *usersRepo) GetUsers() ([]*model.User, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback() // nolint:errcheck
		}
	}()

	userRows, err := tx.Query("select * from users;")
	if err != nil {
		return nil, err
	}
	defer userRows.Close()

	users := make([]*model.User, 0)
	for userRows.Next() {
		var user model.User
		err = userRows.Scan(&user.ID, &user.Name, &user.Surname, &user.Address, &user.Email, &user.Phone)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	skillRows, err := tx.Query("select skills.name, user_skills.level, user_skills.user_id from user_skills left join skills on user_skills.skill_id = skills.id ")
	if err != nil {
		return nil, err
	}
	defer skillRows.Close()

	userSkills := make(map[int][]model.UserSkill)
	for skillRows.Next() {
		var skill model.UserSkill
		var userID int
		err = skillRows.Scan(&skill.Name, &skill.Level, &userID)
		if err != nil {
			return nil, err
		}

		if _, ok := userSkills[userID]; ok {
			userSkills[userID] = append(userSkills[userID], skill)
		} else {
			userSkills[userID] = []model.UserSkill{skill}
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if skills, ok := userSkills[user.ID]; ok {
			user.Skills = skills
		}
	}

	return users, nil
}

func (u *usersRepo) GetUser(userID int) (*model.User, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback() // nolint:errcheck
		}
	}()

	row := tx.QueryRow(fmt.Sprintf("select * from users where id = %d;", userID))
	err = row.Err()
	if err != nil {
		return nil, err
	}

	var user model.User
	err = row.Scan(&user.ID, &user.Name, &user.Surname, &user.Address, &user.Email, &user.Phone)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(fmt.Sprintf("select skills.name, user_skills.level from user_skills left join skills on user_skills.skill_id = skills.id where user_skills.user_id = %d;", user.ID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	skills := make([]model.UserSkill, 0)
	for rows.Next() {
		var skill model.UserSkill
		err = rows.Scan(&skill.Name, &skill.Level)
		if err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}
	user.Skills = skills

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *usersRepo) AddSkill(userID int, skill model.UserSkill) error {
	_, err := u.db.Exec(
		fmt.Sprintf("insert into user_skills values(%d, (select id from skills where name = '%s'), %d);",
			userID, skill.Name, skill.Level),
	)
	return err
}

func (u *usersRepo) UpdateSkill(userID int, skill model.UserSkill) error {
	_, err := u.db.Exec(
		fmt.Sprintf("update user_skills set level = %d where user_id = %d and skill_id = (select id from skills where name = '%s');",
			skill.Level, userID, skill.Name),
	)
	return err
}

func (u *usersRepo) DeleteSkill(userID int, skillName string) error {
	_, err := u.db.Exec(
		fmt.Sprintf("delete from user_skills where user_id = %d and skill_id = (select id from skills where name = '%s');",
			userID, skillName),
	)
	return err
}
