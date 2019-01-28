package tito

import "time"

type question struct {
	title       string `json:"title"`
	description string `json:"description"`
	ID          int64  `json:"id"`
}

type CheckinCreatedPayload struct {
	ID    int64 `jon:"id"`
	Event struct {
		ID           int64     `json:"id"`
		title        string    `json:"title"`
		url          string    `json:"url"`
		account_slug string    `json:"account_slug"`
		slug         string    `json:"slug"`
		currency     string    `json:"currency"`
		start_date   time.Time `json:"start_date"`
		end_date     time.Time `json:"end_date"`
	}
	name                string    `json:"name"`
	first_name          string    `json:"first_name"`
	last_name           string    `json:"last_name"`
	email               string    `json:"email"`
	phone_number        string    `json:"phone_number"`
	company_name        string    `json:"company_name"`
	reference           string    `json:"reference"`
	price               string    `json:"price"`
	tax                 string    `json:"tax"`
	price_less_tax      string    `json:"price_less_tax"`
	slug                string    `json:"slug"`
	state_name          string    `json:"state_name"`
	gender              string    `json:"gender"`
	release_price       string    `json:"release_price"`
	discount_code_used  string    `json:"discount_code_used"`
	total_paid          string    `json:"total_paid"`
	total_paid_less_tax string    `json:"total_paid_less_tax"`
	updated_at          time.Time `json:"updated_at"`
	url                 string    `json:"url"`
	admin_url           string    `json:"admin_url"`
	release_title       string    `json:"release_title"`
	release_slug        string    `json:"release_slug"`
	release_id          string    `json:"release_id"`
	release             struct {
		ID       int64  `json:"id"`
		title    string `json:"title"`
		slug     string `json:"slug"`
		metadata string `json:"metadata"`
	} `json:"release"`
	custom            string `json:"custom"`
	registration_id   string `json:"registration_id"`
	registration_slug string `json:"registration_slug"`
	metadata          string `json:"metadata"`
	answers           []struct {
		question           question `json:"question"`
		response           string   `json:"response"`
		humanized_response string   `json:"humanized_response"`
	}
	responses struct {
		phone_number            string `json:"phone-number"`
		session_1_second_choice string `json:"session-1-second-choice"`
		country                 string `json:"country"`
		session_1_first_choice  string `json:"session-1-first-choice"`
		session_1_third_choice  string `json:"session-1-third-choice"`
		t_shirt_size            string `json:"t-shirt-size"`
	} `json:"responses"`
	upgrade_ids  []int64 `json:"upgrade_ids"`
	registration struct {
		url               string `json:"url"`
		admin_url         string `json:"admin_url"`
		total             string `json:"total"`
		currency          string `json:"currency"`
		payment_reference string `json:"payment_reference"`
		source            string `json:"source"`
		name              string `json:"name"`
		email             string `json:"email"`
		receipt           struct {
			total            string `json:"total"`
			tax              string `json:"tax"`
			payment_provider string `json:"payment_provider"`
			paid             string `json:"paid"`
			receipt_lines    []struct {
				total    string `json:"total"`
				quantity string `json:"quantity"`
				tax      string `json:"tax"`
			} `json:"receipt_lines"`
		} `json:"receipt"`
	} `json:"registration"`
}

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
}

type RegistrationCancelledPayload struct {
}
