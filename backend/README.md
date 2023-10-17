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
POST /pdr/create        (admin required)    json(title, description, suggestedPrice)
GET /pdr/get/:id        (admin required)
GET /pdr/get-all        (admin required)
PUT /pdr/edit/:id       (admin required)    json(title, description, suggestedPrice)
DELETE /pdr/delete/:id  (admin required)

# Plan
POST /plan/create                       json(name, price)
POST /plan/create-features/:planId      json(name, price, isHave)
POST /upload/upload-plan-image/:planId
PUT /plan/edit/:planId                  json(name, price)
PUT	/plan/edit-feature/:featureId       json(name, price, isHave)
GET	/plan/get-all-plans
GET	/plan/get-plan/:planId
GET	/plan/get-all-features/:planId
GET	/plan/get-feature/:featureId
GET /plan/get-all-plans-and-features
DELETE /plan/delete-plan/:planId
DELETE /plan/delete-feature/:featureId