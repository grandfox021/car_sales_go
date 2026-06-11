package helper

import "car_sales_mod/src/api/validations"

type BaseHttpResponse struct {
	Result           any  `json:"result"`
	Success          bool `json:"success"`
	ResultCode       int  `json:"resultcode"`
	ValidationErrors *[]validations.ValidationError `json:"validationErrors"`
	Errors           any  `json:"errors"`
}


func GenerateBaseRespone(result any , succeess bool , resultcode int ,) *BaseHttpResponse {

	return &BaseHttpResponse{
		Result: result,
		Success: succeess,
		ResultCode: resultcode, 
	
	}

}
func GenerateBaseResponeWithError(result any , succeess bool , resultcode int ,err error) *BaseHttpResponse {

	return &BaseHttpResponse{
		Result: result,
		Success: succeess,
		ResultCode: resultcode,
		Errors: err.Error(), 
	}

}
func GenerateBaseResponeWithValidationError(result any , succeess bool , resultcode int ,err error) *BaseHttpResponse {

	return &BaseHttpResponse{
		Result: result,
		Success: succeess,
		ResultCode: resultcode, 
		ValidationErrors: validations.Get_validation_errors(err),
	
	}

}