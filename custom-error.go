package main

import "fmt"

type ValidationError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"ID cannot be empty"}
	}

	if id != "ucok" {
		return &NotFoundError{"Data not found"}
	}

	// anggap save data : return ok
	return nil
}

func MainCustom() {
	err := SaveData("ucok", nil)

	if err != nil {
		// if validationErr, ok := err.(*ValidationError); ok {
		// 	fmt.Println("Validation error", validationErr.Error())
		// } else if notFoundErr, ok := err.(*NotFoundError); ok {
		// 	fmt.Println("Not found error", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Unknown error", err.Error())
		// }

		switch finalErr := err.(type) {
			case *ValidationError:
				fmt.Println("Validation error", finalErr.Error())
			case *NotFoundError:
				fmt.Println("Not found error", finalErr.Error())
			default:
				fmt.Println("Unknown error", finalErr.Error())
		}
	} else {
		fmt.Println("Data saved")
	}
}