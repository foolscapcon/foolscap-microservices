Will be replaced with the ToC, excluding the “Contents” header
Webhooks
Tito provides the option to add a webhook to be notified when certain things happen.

You can specify a webhook in the Webhooks tab in the Customize section.

Webhooks will POST a JSON payload to the endpoint you specify, with a X-Webhook-Name header. This will be either:

checkin.created
Fired when an attendee is checked-in via one of the Tito apps (web, iOS or Android).

ticket.created
Fired when tickets are created immediately after a registration is confirmed. Tickets can be unassigned, incomplete or complete at this stage depending on whether it is a single or multiple ticket registration or whether the ticket has required fields or not.

ticket.completed
Fired when all required fields like name, email, etc and required questions have been answered.

ticket.reassigned
Fired specifically when a ticket is reassigned to someone else.

ticket.updated
Fired anytime ticket details are changed, including changes to the email address.

ticket.unsnoozed
Fired when a “snoozed” ticket is assigned to someone. When people register tickets but do not want to or can’t assign them straight away they can be “snoozed”.

ticket.unvoided
Fired when a voided ticket is returned to a valid state.

ticket.voided
Fired when a ticket is changed to a voided state. Free tickets that are cancelled by the attendee will also cause this webhook to fire.

registration.started
Fired when a registration is started e.g. after they have selected tickets but before they have filled in any details.

registration.filling
Fired as soon as the person registering starts to filling in details like name, email, etc.

registration.updated
Fired when the registration is updated in the Tito admin area or via the API.

registration.finished
Fired as soon as the registration is confirmed, e.g when payment is complete.

registration.completed
Fired when all tickets on a registration have either been assigned or “snoozed”.

registration.cancelled
Fired when the registration is cancelled in the Tito admin area or via the API. Free tickets that are cancelled by the attendee will also cause this webhook to fire.

Basic Ticket Payload
The payload will look something like this for a basic ticket:

{
  "id": 795,
  "event": {
    "id": 64,
    "title": "AwesomeConf ?",
    "url": "https://ti.to/paulca/awesomeconf",
    "account_slug": "paulca",
    "slug": "awesomeconf",
    "currency": "EUR",
    "start_date": "2020-12-01",
    "end_date": null
  },
  "name": "Peter Winter",
  "first_name": "Peter",
  "last_name": "Winter",
  "email": "winter@example.com",
  "phone_number": null,
  "company_name": "Peter Winter",
  "reference": "1i5nfwf2t",
  "price": "1.00",
  "tax": "0.19",
  "price_less_tax": "0.81",
  "slug": "1i5nfwf2t24",
  "state_name": "incomplete",
  "gender": "male",
  "release_price": "1.00",
  "discount_code_used": "",
  "total_paid": "1.00",
  "total_paid_less_tax": "0.81",
  "updated_at": "2017-08-27T13:07:47.000Z",
  "url": "https://ti.to/tickets/1i5nfwf2t24",
  "admin_url": "https://ti.to/paulca/awesomeconf/admin/tickets/1i5nfwf2t24",
  "release_title": "AwesomeConf ticket",
  "release_slug": "31az2uvoh4",
  "release_id": 90,
  "release": {
    "id": 90,
    "title": "AwesomeConf ticket",
    "slug": "31az2uvoh4",
    "metadata": null
  },
  "custom": null,
  "registration_id": "bdtyap3hguq",
  "registration_slug": "bdtyap3hguq",
  "metadata": null,
  "answers": [
    {
      "question": {
        "title": "Phone Number",
        "description": "",
        "id": 200
      },
      "response": "555-555-0199@example.com",
      "humanized_response": "555-555-0199@example.com"
    },
    {
      "question": {
        "title": "Session 1 – Second Choice",
        "description": "",
        "id": 1293
      },
      "response": "One",
      "humanized_response": "One"
    },
    {
      "question": {
        "title": "Session 1 - First Choice",
        "description": "",
        "id": 1292
      },
      "response": "555-555-0199@example.com",
      "humanized_response": "555-555-0199@example.com"
    },
    {
      "question": {
        "title": "Country",
        "description": "",
        "id": 1005357
      },
      "response": "AU",
      "humanized_response": "AU"
    },
    {
      "question": {
        "title": "Session 1 — Third Choice",
        "description": "",
        "id": 1295
      },
      "response": "One",
      "humanized_response": "One"
    },
    {
      "question": {
        "title": "T-Shirt Size",
        "description": "",
        "id": 1009604
      },
      "response": "Women's S",
      "humanized_response": "Women's S"
    }
  ],
  "responses": {
    "phone-number": "555-555-0199@example.com",
    "session-1-second-choice": "One",
    "session-1-first-choice": "555-555-0199@example.com",
    "country": "AU",
    "session-1-third-choice": "One",
    "t-shirt-size": "Women's S"
  },
  "upgrade_ids": [],
  "registration": {
    "url": "https://ti.to/registrations/bdtyap3hguq",
    "admin_url": "https://ti.to/paulca/awesomeconf/admin/registrations/bdtyap3hguq",
    "total": "1.00",
    "currency": "EUR",
    "payment_reference": null,
    "source": null,
    "name": "Paul Campbell",
    "email": "funconf@gmail.com",
    "receipt": {
      "total": "1.00",
      "tax": "0.19",
      "payment_provider": "Unknown",
      "paid": true,
      "receipt_lines": [
        {
          "total": "1.00",
          "quantity": 1,
          "tax": "0.19"
        }
      ]
    }
  }
}
Payload with Responses to Questions
If the ticket has answers to questions, it will look a bit like this:

{
  "id": 1052678,
  "event": {
    "id": 64,
    "title": "AwesomeConf ?",
    "url": "https://ti.to/paulca/awesomeconf",
    "account_slug": "paulca",
    "slug": "awesomeconf",
    "currency": "EUR",
    "start_date": "2020-12-01",
    "end_date": null
  },
  "name": "Peter Winter",
  "first_name": "Peter",
  "last_name": "Winter",
  "email": "winter@example.com",
  "phone_number": null,
  "company_name": "Peter Winter",
  "reference": "A8FV-3",
  "price": "1.00",
  "tax": "0.19",
  "price_less_tax": "0.81",
  "slug": "pGZ3De1X3wLslG7yQwO7QFg",
  "state_name": "incomplete",
  "gender": "male",
  "release_price": "1.00",
  "discount_code_used": "",
  "total_paid": "1.00",
  "total_paid_less_tax": "0.81",
  "updated_at": "2017-08-27T13:07:50.000Z",
  "url": "https://ti.to/tickets/pGZ3De1X3wLslG7yQwO7QFg",
  "admin_url": "https://ti.to/paulca/awesomeconf/admin/tickets/pGZ3De1X3wLslG7yQwO7QFg",
  "release_title": "AwesomeConf ticket",
  "release_slug": "31az2uvoh4",
  "release_id": 90,
  "release": {
    "id": 90,
    "title": "AwesomeConf ticket",
    "slug": "31az2uvoh4",
    "metadata": null
  },
  "custom": null,
  "registration_id": "pbidGD4FWx7r47py0dYdDeg",
  "registration_slug": "pbidGD4FWx7r47py0dYdDeg",
  "metadata": null,
  "answers": [
    {
      "question": {
        "title": "Phone Number",
        "description": "",
        "id": 200
      },
      "response": "12345",
      "humanized_response": "12345"
    },
    {
      "question": {
        "title": "Session 1 – Second Choice",
        "description": "",
        "id": 1293
      },
      "response": "One",
      "humanized_response": "One"
    },
    {
      "question": {
        "title": "Country",
        "description": "",
        "id": 1005357
      },
      "response": "AU",
      "humanized_response": "AU"
    },
    {
      "question": {
        "title": "Session 1 - First Choice",
        "description": "",
        "id": 1292
      },
      "response": "555-555-0199@example.com",
      "humanized_response": "555-555-0199@example.com"
    },
    {
      "question": {
        "title": "Session 1 — Third Choice",
        "description": "",
        "id": 1295
      },
      "response": "One",
      "humanized_response": "One"
    },
    {
      "question": {
        "title": "T-Shirt Size",
        "description": "",
        "id": 1009604
      },
      "response": "Women's S",
      "humanized_response": "Women's S"
    }
  ],
  "responses": {
    "phone-number": "12345",
    "session-1-second-choice": "One",
    "country": "AU",
    "session-1-first-choice": "555-555-0199@example.com",
    "session-1-third-choice": "One",
    "t-shirt-size": "Women's S"
  },
  "upgrade_ids": [],
  "registration": {
    "url": "https://ti.to/registrations/pbidGD4FWx7r47py0dYdDeg",
    "admin_url": "https://ti.to/paulca/awesomeconf/admin/registrations/pbidGD4FWx7r47py0dYdDeg",
    "total": "3.00",
    "currency": "EUR",
    "payment_reference": "ch_4MkyHRycVDKfCp",
    "source": null,
    "name": "Paul Campbell",
    "email": "paul@rslw.com",
    "receipt": {
      "total": "3.00",
      "tax": "0.56",
      "payment_provider": "Stripe (live)",
      "paid": true,
      "receipt_lines": [
        {
          "total": "3.00",
          "quantity": 3,
          "tax": "0.56"
        }
      ]
    }
  }
}
Verifying the payload
Each payload comes with a signed HMAC signature so that you can verify the origin of the webhook request.

In Customize -> Webhooks, you’ll find a shared security token. You can use the to sign the payload.

The security token is used as a key to sign the payload data with an HMAC key.

The HMAC key is your security token, the HMAC digest is SHA2, and the data is the raw payload JSON that is sent. The key is sent Base64 encoded via the Tito-Signature HTTP header.

If you’re using Ruby, for example, you can verify the authenticity of your payload with the following code:

key = 'YOUR EVENT SECURITY TOKEN'
hash = OpenSSL::Digest.new('sha256')
data = 'THE WEBHOOK REQUEST'

Base64.encode64(OpenSSL::HMAC.digest(hash, key, data)).strip
If you ever need some help, don't hesitate to contact us…
Email: support@tito.io
