package internal

import (
	"fmt"

	"contacts_api/modules/model"
)

func (s *skillsRepo) GetSkills() ([]model.Skill, error) {
	rows, err := s.db.Query("select * from skills;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	skills := make([]model.Skill, 0)
	for rows.Next() {
		var skill model.Skill
		var id int
		err = rows.Scan(&id, &skill.Name)
		if err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	return skills, nil
}

func (s *skillsRepo) SkillsAsMap() (map[string]struct{}, error) {
	skills, err := s.GetSkills()
	if err != nil {
		return nil, err
	}

	asMap := make(map[string]struct{})
	for _, skill := range skills {
		asMap[skill.Name] = struct{}{}
	}

	return asMap, nil
}

func (s *skillsRepo) CreateSkills(skills []model.Skill) error {
	for _, skill := range skills {
		_, err := s.db.Exec(fmt.Sprintf("insert into skills (name) values('%s');", skill.Name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *skillsRepo) DeleteSkill(name string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback() // nolint:errcheck
		}
	}()

	_, err = tx.Exec(fmt.Sprintf("delete from user_skills where skill_id = (select id from skills where name = '%s');", name))
	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf("delete from skills where name = '%s';", name))
	if err != nil {
		return err
	}

	return tx.Commit()
}
