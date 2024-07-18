#!/bin/bash

# File to store the token and company_id
TOKEN_FILE="token.txt"
COMPANY_ID_FILE="company_id.txt"

# Function to register the user
register() {
  echo "Registering user..."
  curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{
    "username": "gjohndoe",
    "password": "password123",
    "firstname": "John",
    "lastname": "Doe"
  }'
}

# Function to login the user and get the token
login() {
  echo "Logging in user..."
  LOGIN_RESPONSE=$(curl -s -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{
    "username": "gjohndoe",
    "password": "password123"
  }')

  TOKEN=$(echo $LOGIN_RESPONSE | jq -r .token)

  echo "Token obtained: $TOKEN"
  echo $TOKEN > $TOKEN_FILE
}

# Function to create a company
create_company() {
  if [[ ! -f $TOKEN_FILE ]]; then
    echo "Token file not found. Please login first."
    exit 1
  fi

  TOKEN=$(cat $TOKEN_FILE)

  echo "Creating company..."
  CREATE_COMPANY_RESPONSE=$(curl -s -X POST http://localhost:8080/company -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{
    "name": "MyCompany",
    "description": "This is a description of MyCompany. It can be up to 3000 characters long.",
    "amount_of_employees": 50,
    "registered": true,
    "type": "Corporations"
  }')

  COMPANY_ID=$(echo $CREATE_COMPANY_RESPONSE | jq -r .company_id)

  echo "Company created with ID: $COMPANY_ID"
  echo $COMPANY_ID > $COMPANY_ID_FILE
}

# Function to get company details
get_company() {
  if [[ ! -f $TOKEN_FILE ]] || [[ ! -f $COMPANY_ID_FILE ]]; then
    echo "Token or company ID file not found. Please create a company first."
    exit 1
  fi

  TOKEN=$(cat $TOKEN_FILE)
  COMPANY_ID=$(cat $COMPANY_ID_FILE)

  echo "Getting company details..."
  curl -X GET "http://localhost:8080/company?company_id=$COMPANY_ID" -H "Authorization: Bearer $TOKEN"
}

# Function to update company details
update_company() {
  if [[ ! -f $TOKEN_FILE ]] || [[ ! -f $COMPANY_ID_FILE ]]; then
    echo "Token or company ID file not found. Please create a company first."
    exit 1
  fi

  TOKEN=$(cat $TOKEN_FILE)
  COMPANY_ID=$(cat $COMPANY_ID_FILE)

  echo "Updating company..."
  curl -X PUT "http://localhost:8080/company?company_id=$COMPANY_ID" -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{
    "name": "Company",
    "description": "This is an updated description of MyCompany.",
    "amount_of_employees": 100,
    "registered": true,
    "type": "Corporations"
  }'
}

# Function to delete the company
delete_company() {
  if [[ ! -f $TOKEN_FILE ]] || [[ ! -f $COMPANY_ID_FILE ]]; then
    echo "Token or company ID file not found. Please create a company first."
    exit 1
  fi

  TOKEN=$(cat $TOKEN_FILE)
  COMPANY_ID=$(cat $COMPANY_ID_FILE)

  echo "Deleting company..."
  curl -X DELETE "http://localhost:8080/company?company_id=$COMPANY_ID" -H "Authorization: Bearer $TOKEN"
}

# Main function to handle the command-line arguments
main() {
  case "$1" in
    register)
      register
      ;;
    login)
      login
      ;;
    create_company)
      create_company
      ;;
    get_company)
      get_company
      ;;
    update_company)
      update_company
      ;;
    delete_company)
      delete_company
      ;;
    *)
      echo "Usage: $0 {register|login|create_company|get_company|update_company|delete_company}"
      exit 1
  esac
}

main "$@"
