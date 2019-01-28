package tito

import "time"

type question struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ID          int64  `json:"id"`
}

type CheckinCreatedPayload struct {}

type TicketCreatedPayload struct {
}

type TicketCompletedPayload struct {
}

type TicketReassignedPayload struct {
}

type TicketUpdatedPayload struct {
}

type TicketUnsnoozedPayload struct {
}

type TicketUnvoidedPayload struct {
}

type TicketVoidedPayload struct {
}

type RegistrationStartedPayload struct {
}

type RegistrationFillingPayload struct {
}

type RegistrationUpdatedPayload struct {
}

type RegistrationFinishedPayload struct {
}

type RegistrationCompletedPayload struct {
	ID    int64 `jon:"id"`
	Event struct {
		ID           int64     `json:"id"`
		Title        string    `json:"title"`
		Url          string    `json:"url"`
		Account_slug string    `json:"account_slug"`
		Slug         string    `json:"slug"`
		Currency     string    `json:"currency"`
		Start_date   time.Time `json:"start_date"`
		End_date     time.Time `json:"end_date"`
	}
	Name                string    `json:"name"`
	First_name          string    `json:"first_name"`
	Last_name           string    `json:"last_name"`
	Email               string    `json:"email"`
	Phone_number        string    `json:"phone_number"`
	Company_name        string    `json:"company_name"`
	Reference           string    `json:"reference"`
	Price               string    `json:"price"`
	Tax                 string    `json:"tax"`
	Price_less_tax      string    `json:"price_less_tax"`
	Slug                string    `json:"slug"`
	State_name          string    `json:"state_name"`
	Gender              string    `json:"gender"`
	Release_price       string    `json:"release_price"`
	Discount_code_used  string    `json:"discount_code_used"`
	Total_paid          string    `json:"total_paid"`
	Total_paid_less_tax string    `json:"total_paid_less_tax"`
	Updated_at          time.Time `json:"updated_at"`
	Url                 string    `json:"url"`
	Admin_url           string    `json:"admin_url"`
	Release_title       string    `json:"release_title"`
	Release_slug        string    `json:"release_slug"`
	Release_id          string    `json:"release_id"`
	Release             struct {
		ID       int64  `json:"id"`
		Title    string `json:"title"`
		Slug     string `json:"slug"`
		Metadata string `json:"metadata"`
	} `json:"release"`
	Custom            string `json:"custom"`
	Registration_id   string `json:"registration_id"`
	Registration_slug string `json:"registration_slug"`
	Metadata          string `json:"metadata"`
	Answers           []struct {
		Question           question `json:"question"`
		Response           string   `json:"response"`
		Humanized_response string   `json:"humanized_response"`
	}
	Responses struct {
		Phone_number            string `json:"phone-number"`
		Session_1_second_choice string `json:"session-1-second-choice"`
		Country                 string `json:"country"`
		Session_1_first_choice  string `json:"session-1-first-choice"`
		Session_1_third_choice  string `json:"session-1-third-choice"`
		T_shirt_size            string `json:"t-shirt-size"`
	} `json:"responses"`
	Upgrade_ids  []int64 `json:"upgrade_ids"`
	Registration struct {
		Url               string `json:"url"`
		Admin_url         string `json:"admin_url"`
		Total             string `json:"total"`
		Currency          string `json:"currency"`
		Payment_reference string `json:"payment_reference"`
		Source            string `json:"source"`
		Name              string `json:"name"`
		Email             string `json:"email"`
		Receipt           struct {
			Total            string `json:"total"`
			Tax              string `json:"tax"`
			Payment_provider string `json:"payment_provider"`
			Paid             string `json:"paid"`
			Receipt_lines    []struct {
				Total    string `json:"total"`
				Quantity string `json:"quantity"`
				Tax      string `json:"tax"`
			} `json:"receipt_lines"`
		} `json:"receipt"`
	} `json:"registration"`
}

type RegistrationCancelledPayload struct {
}
