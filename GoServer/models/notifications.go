package models

type Notification struct {
	AUserId            *uint32 `json:"user_id"`
	UserName           *string `json:"user_name"`
	LItemId            *uint32 `json:"item_id"`
	ItemName           *string `json:"item_name"`
	NotificationType   *uint32 `json:"notification_type"`
	NotificationString *string `json:"notification_string"`
	Ctime              *int64  `json:"ctime"`
}

type GetUserNotificationsRequest struct {
	UserID *uint32 `json:"user_id"`
}

func (r *GetUserNotificationsRequest) GetUserID() uint32 {
	if r != nil && r.UserID != nil {
		return *r.UserID
	}
	return 0
}

type GetUserNotificationsResponse struct {
	NotificationsCount uint32         `json:"notification_count"`
	NotificationsList  []Notification `json:"notification_list"`
}
