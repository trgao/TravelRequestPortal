package dao

import (
	"fmt"
	"time"

	"database/sql"

	"go-backend/db"
	"go-backend/models"
)

type FullRequest struct {
	ID            uint      `json:"id"`
	StartLocation string    `json:"start_location"`
	Destination   string    `json:"destination"`
	DateTime      time.Time `json:"date_time"`
	Purpose       string    `json:"purpose"`
	Status        string    `json:"status"`
	AdminRemarks  string    `json:"admin_remarks"`
	UserID        uint      `json:"user_id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
}

func GetPending(userId uint) (pending int64, err error) {
	result := db.MasterConn.Raw("SELECT COUNT(*) FROM requests INNER JOIN users ON user_id = users.id WHERE status = 'pending' AND company_id = (SELECT company_id FROM users WHERE id = @id)", sql.Named("id", userId)).Scan(&pending)
	return pending, result.Error
}

func GetApproved(userId uint) (approved int64, err error) {
	result := db.MasterConn.Raw("SELECT COUNT(*) FROM requests INNER JOIN users ON user_id = users.id WHERE status = 'approved' AND company_id = (SELECT company_id FROM users WHERE id = @id)", sql.Named("id", userId)).Scan(&approved)
	return approved, result.Error
}

func GetRejected(userId uint) (rejected int64, err error) {
	result := db.MasterConn.Raw("SELECT COUNT(*) FROM requests INNER JOIN users ON user_id = users.id WHERE status = 'rejected' AND company_id = (SELECT company_id FROM users WHERE id = @id)", sql.Named("id", userId)).Scan(&rejected)
	return rejected, result.Error
}

func GetRequestsForEmployee(userId uint, limit int64, page int64) (requests []models.Request, err error) {
	requests = make([]models.Request, 0)
	queryString := "SELECT * FROM requests WHERE user_id = @id ORDER BY id"
	queryString += fmt.Sprintf(" LIMIT %d", max(1, limit))
	queryString += fmt.Sprintf(" OFFSET %d", max(0, (page-1)*limit))

	result := db.MasterConn.Raw(
		queryString,
		sql.Named("id", userId),
	).Scan(&requests)

	return requests, result.Error
}

func GetRequestsCountForEmployee(userId uint) (count int64, err error) {
	queryString := "SELECT COUNT(*) FROM requests WHERE user_id = @id"

	result := db.MasterConn.Raw(
		queryString,
		sql.Named("id", userId),
	).Scan(&count)

	return count, result.Error
}

func GetRequestsForAdmin(userId uint, employee string, limit int64, page int64) (requests []FullRequest, err error) {
	requests = make([]FullRequest, 0)
	queryString := "SELECT requests.id, start_location, destination, date_time, purpose, status, admin_remarks, first_name, last_name FROM requests INNER JOIN users ON user_id = users.id WHERE company_id = (SELECT company_id FROM users WHERE id = @id)"

	if employee != "" {
		queryString += " AND (first_name || ' ' || last_name ILIKE @employee OR first_name || ' ' || last_name ILIKE @employee)"
	}

	queryString += " ORDER BY requests.id"
	queryString += fmt.Sprintf(" LIMIT %d", max(1, limit))
	queryString += fmt.Sprintf(" OFFSET %d", max(0, (page-1)*limit))

	result := db.MasterConn.Raw(
		queryString,
		sql.Named("id", userId),
		sql.Named("employee", "%"+employee+"%"),
	).Scan(&requests)

	return requests, result.Error
}

func GetRequestsCountForAdmin(userId uint, employee string) (count int64, err error) {
	queryString := "SELECT COUNT(*) FROM requests INNER JOIN users ON user_id = users.id WHERE company_id = (SELECT company_id FROM users WHERE id = @id)"

	if employee != "" {
		queryString += " AND (first_name || ' ' || last_name ILIKE @employee OR first_name || ' ' || last_name ILIKE @employee)"
	}

	result := db.MasterConn.Raw(
		queryString,
		sql.Named("id", userId),
		sql.Named("employee", "%"+employee+"%"),
	).Scan(&count)

	return count, result.Error
}

func CreateRequest(startLocation string, destination string, dateTime time.Time, purpose string, userId uint) (err error) {
	request := models.Request{
		StartLocation: startLocation,
		Destination:   destination,
		DateTime:      dateTime,
		Purpose:       purpose,
		Status:        "pending",
		AdminRemarks:  "",
		UserID:        userId,
	}

	result := db.MasterConn.Create(&request)
	return result.Error
}

func UpdateRequestStatus(requestId uint64, status string) (err error) {
	result := db.MasterConn.Model(&models.Request{}).Where("id = ?", requestId).Update("status", status)
	return result.Error
}

func UpdateRequestRemarks(requestId uint64, remarks string) (err error) {
	result := db.MasterConn.Model(&models.Request{}).Where("id = ?", requestId).Update("admin_remarks", remarks)
	return result.Error
}
