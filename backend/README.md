### Backnd with GoFiber

## Routes
GET /api/v1/home

# auth
POST /api/v1/signup (json with email, password, number, name, family)
    password lenght bigger than 8
    email example test@test.com
    number int
    name must letter without !@#$%^&*()_+-
POST /api/v1/signin (json with email, username, password) (one of email and username is required)

# Project Doing Request
POST /pdr/create (auth required) (json with title, description, suggestedPrice)
GET /pdr/get/:id (auth required)
GET /pdr/get-all (auth required)
POST /pdr/edit/:id (auth required) (json with title, description, suggestedPrice)

# Plan
POST /plan/create
POST /plan/create-features/:planId
POST /plan/upload-image/:planId
PUT /plan/edit/:planId
PUT	/plan/edit-feature/:featureId
PUT	/plan/edit-image/:planId
GET	/plan/get-all-plans
GET	/plan/get-plan/:planId
GET	/plan/get-all-features/:planId
GET	/plan/get-feature/:featureId
GET /plan/get-all-plans-and-features
DELETE /plan/delete-plan/:planId
DELETE /plan/delete-feature/:featureId