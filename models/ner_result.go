package models

type NERResult struct {
	Category   string
	EntityName string
}

func (n NERResult) GetCategory() string {
	return n.Category
}

func (n NERResult) SetCategory(value string) NERResult {
	n.Category = value
	return n
}

func (n NERResult) GetEntityName() string {
	return n.EntityName
}

func (n NERResult) SetEntityName(value string) NERResult {
	n.EntityName = value
	return n
}
