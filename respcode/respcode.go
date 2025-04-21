package respcode

// ✅ General Success Codes
const (
	SUC_CODE_200           = "200"
	SUC_CODE_200_MSG       = "Request successful."

	SUC_CODE_201           = "201"
	SUC_CODE_201_MSG       = "Resource created successfully."

	SUC_CODE_202           = "202"
	SUC_CODE_202_MSG       = "Request accepted and is being processed."

	SUC_CODE_204           = "204"
	SUC_CODE_204_MSG       = "No content returned."
)

// ✅ Specific Success Messages
const (
	SUC_CODE_INSERT        = "200"
	SUC_CODE_INSERT_MSG    = "Data inserted successfully."

	SUC_CODE_UPDATE        = "200"
	SUC_CODE_UPDATE_MSG    = "Data updated successfully."

	SUC_CODE_DELETE        = "200"
	SUC_CODE_DELETE_MSG    = "Data deleted successfully."

	SUC_CODE_FETCH         = "200"
	SUC_CODE_FETCH_MSG     = "Data fetched successfully."

	SUC_CODE_PROCESSED     = "200"
	SUC_CODE_PROCESSED_MSG = "Request processed successfully."
)

// ❌ Client Error Codes
const (
	ERR_CODE_400           = "400"
	ERR_CODE_400_MSG       = "Bad request. Check your input."

	ERR_CODE_401           = "401"
	ERR_CODE_401_MSG       = "Unauthorized access."

	ERR_CODE_403           = "403"
	ERR_CODE_403_MSG       = "Forbidden. Access denied."

	ERR_CODE_404           = "404"
	ERR_CODE_404_MSG       = "Resource not found."

	ERR_CODE_409           = "409"
	ERR_CODE_409_MSG       = "Conflict. Duplicate or already exists."
)

// ❗ Server Error Codes
const (
	ERR_CODE_500           = "500"
	ERR_CODE_500_MSG       = "Internal server error."

	ERR_CODE_502           = "502"
	ERR_CODE_502_MSG       = "Bad gateway."

	ERR_CODE_503           = "503"
	ERR_CODE_503_MSG       = "Service unavailable."
)
