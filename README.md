# Bank application
Backend app for bank. With PostgreSQL storage, JWT authentication and HTTP API (see documentation below). Complete aplication for managing banking accounts.

---
## HTTP API docs
`/accounts` (GET)
returns all exsisting bank accounts in database

---
`/accounts` (POST)
creates new user bank account and returns JWT token for authentication

| Parameter | Required |  Description |
| ----------- | ----------- |----------- |
| firstName | Yes | (string) First name of user  |
| lastName | Yes | (string) Last name of user  |

---
`/accounts/{id}` (GET)
returns info about bank accounts in database

| Parameter | Required |  Description |
| ----------- | ----------- |----------- |
| id | Yes | (int) User ID. Needs cookie authentication via JWT token |
---
`/transfer` (POST)
makes transfer from current account to specified in body specified amount

| Parameter | Required |  Description |
| ----------- | ----------- |----------- |
| toAccount | Yes | (int) Account number |
| amount | Yes | (int) Amount |