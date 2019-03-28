package zammad

import "time"

type NewTicket struct {
	Title    string `json:"title"`
	Group    string `json:"group"`
	Customer string `json:"customer"`
	Article  struct {
		Subject string `json:"subject"`
		Body    string `json:"body"`
	} `json:"article"`
}

type Ticket struct {
	ID                        *int        `json:"id"`
	GroupID                   *int        `json:"group_id"`
	PriorityID                *int        `json:"priority_id"`
	StateID                   *int        `json:"state_id"`
	OrganizationID            *int        `json:"organization_id"`
	Number                    string      `json:"number"`
	Title                     string      `json:"title"`
	OwnerID                   *int        `json:"owner_id"`
	CustomerID                *int        `json:"customer_id"`
	Note                      interface{} `json:"note"`
	FirstResponseAt           *time.Time  `json:"first_response_at"`
	FirstResponseEscalationAt *time.Time  `json:"first_response_escalation_at"`
	FirstResponseInMin        interface{} `json:"first_response_in_min"`
	FirstResponseDiffInMin    interface{} `json:"first_response_diff_in_min"`
	CloseAt                   *time.Time  `json:"close_at"`
	CloseEscalationAt         *time.Time  `json:"close_escalation_at"`
	CloseInMin                interface{} `json:"close_in_min"`
	CloseDiffInMin            interface{} `json:"close_diff_in_min"`
	UpdateEscalationAt        *time.Time  `json:"update_escalation_at"`
	UpdateInMin               interface{} `json:"update_in_min"`
	UpdateDiffInMin           interface{} `json:"update_diff_in_min"`
	LastContactAt             *time.Time  `json:"last_contact_at"`
	LastContactAgentAt        *time.Time  `json:"last_contact_agent_at"`
	LastContactCustomerAt     *time.Time  `json:"last_contact_customer_at"`
	LastOwnerUpdateAt         *time.Time  `json:"last_owner_update_at"`
	CreateArticleTypeID       int         `json:"create_article_type_id"`
	CreateArticleSenderID     int         `json:"create_article_sender_id"`
	ArticleCount              int         `json:"article_count"`
	EscalationAt              *time.Time  `json:"escalation_at"`
	PendingTime               interface{} `json:"pending_time"`
	Type                      interface{} `json:"type"`
	TimeUnit                  interface{} `json:"time_unit"`
	Preferences               struct {
	} `json:"preferences"`
	UpdatedByID             int           `json:"updated_by_id"`
	CreatedByID             int           `json:"created_by_id"`
	CreatedAt               time.Time     `json:"created_at"`
	UpdatedAt               time.Time     `json:"updated_at"`
	Building                interface{}   `json:"building"`
	ArticleIds              []int         `json:"article_ids"`
	TicketTimeAccountingIds []interface{} `json:"ticket_time_accounting_ids"`
}
