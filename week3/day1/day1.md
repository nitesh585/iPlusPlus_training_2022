## Use case for mini project

Have to pick any one of the given use cases:

- Online shopping cart
- Healthcare management system
- Train reservation system

### Train reservation system

Application which provides the features to search and book trains according to source, destination and date of their journey.

#### Functional Requirements:

<b>user management</b>

- user can register, login and logout
- maintain history
  - booked , canceled, failed tickets
  - ticket refunds
- update profile
  - password and basic personal details
  - add aadhaar card with KYC

<b>search module</b>

- check for the train availability on the basis of source, destination stations and date
- if direct route not found then show different available connected path from source to  
  destination
- add filters of classes, seat priority, etc.
- user should get list of available trains with further details:
  - seats available, RAC & waiting list count
  - arrival time and departure time
  - estimated journey time
  - train schedule
- user can update the search query

<b>booking engine</b>

- book ticket for selected train
- user must logged in before booking
- get the details of passengers (age, name and seat priority) and give a review page to user
- after reviewing, redirect to the payment method.
- cancel ticket after booking
  - if cancellation is after chart preparation or as per law, then refund after deduction otherwise not
  - user can cancel whole ticket or cancel ticket for one of the passenger
- give invoice or ticket in the end of successful booking on email or sms as per the subscribed service

<b>notification module</b>

- user should get alerts about tickets on subscribed services eg. email or sms
- should give alert to user about his/her ticket status - alert if the ticket got confirmed from RAC or WL

<b>payment method</b>

- integrate one of payment gateways
- user must logged in and selected at least one seat
- pay for the tickets
- cancel pay anytime before last step of payment
