package data

import (
	"gp3/features/class"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Class string
}

func fromCore(dataCore class.Core) Class {
	return Class{
		Class: dataCore.Class,
	}

}

func (dataClass *Class) toCore() class.Core {
	return class.Core{
		ID:    dataClass.ID,
		Class: dataClass.Class,
	}
}

func toCoreList(dataClass []Class) []class.Core {
	var dataCore []class.Core

	for key := range dataClass {
		dataCore = append(dataCore, dataClass[key].toCore())

	}

	return dataCore

}
