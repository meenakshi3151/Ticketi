# GoConference Booking Application

GoConference is a simple and efficient ticket booking application written in Go. It handles user validations, ticket management, and booking functionalities while ensuring a smooth and concurrent user experience.

---

## Features

- **Dynamic Ticket Management**: Keeps track of booked and remaining tickets.
- **User Validations**:
  - Validates first and last names.
  - Ensures the email is a valid Gmail address.
  - Verifies the number of tickets requested.
- **City Selection**: Validates conference cities, with a default fallback to "New York."
- **Concurrency**: Utilizes Goroutines to send ticket confirmations asynchronously.
- **User-Friendly Output**:
  - Displays booking information dynamically.
  - Lists attendees and their first names.
- **Scalable**: Supports up to 50 unique bookings.

---


## Prerequisites

1. Install [Go](https://golang.org/dl/) (version 1.16 or later recommended).
2. Clone or download this repository.
3. Ensure your `GOPATH` is set up correctly.

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/meenakshi3151/Ticketi.git
   
