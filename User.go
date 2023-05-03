package main

import "fmt"

func (data *UserDataBodyParams) AddUserInDB() NormalFunctionResponse {
	tnx, err := DbConnection.Begin()

	if err != nil {
		tnx.Rollback()
		return NormalFunctionResponse{Success: false, Message: "Unable to add user -1", Data: struct{ error string }{error: fmt.Sprintf("%v", err)}, Code: 500}
	}
	qRes, err := tnx.Exec(fmt.Sprintf(`INSERT INTO users (name,age) VALUES ("%v",%v)`, data.Name, data.Age))
	if err != nil {
		tnx.Rollback()
		return NormalFunctionResponse{Success: false, Message: "Unable to add user -2", Data: struct{ error string }{error: fmt.Sprintf("%v", err)}, Code: 500}

	}

	insertedId, err := qRes.LastInsertId()

	if err != nil {
		tnx.Rollback()
		return NormalFunctionResponse{Success: false, Message: "Unable to add user -3", Data: struct{ error string }{error: fmt.Sprintf("%v", err)}, Code: 500}

	}
	tnx.Commit()
	return NormalFunctionResponse{Success: true, Message: "Successfully added user", Data: struct{ Id int }{Id: int(insertedId)}, Code: 200}

}
