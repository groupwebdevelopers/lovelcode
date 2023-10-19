### Backnd with GoFiber

## Info
Content-Type: application/json
base url: /api/v1/...


## Routes
GET /api/v1/home

# auth
POST /signup (json with email, password, name, family(optional))
    password lenght bigger than 8
    email example test@test.com
    name must letter without !@#$%^&*()_+-;'"
POST /signin (json with email, password)

# Project Doing Request
POST /pdr/create        (admin required)    json(title, description, suggestedPrice(optional))
GET /pdr/get/:id        (admin required)
GET /pdr/get-all        (admin required)
PUT /pdr/edit/:id       (admin required)    json(title, description, suggestedPrice(optional))
DELETE /pdr/delete/:id  (admin required)

# Plan filter f value
POST /admin/plan/create                       json(name, price)
POST /admin/plan/create-features/:planId      json(name, price, isHave)
POST /admin/upload/upload-plan-image/:planId  (Content-Type: multipart/form-data) (with FormData js object) (without base url)
PUT /admin/plan/edit/:planId                  json(name, price)
PUT	/admin/plan/edit-feature/:featureId       json(name, price, isHave)
GET	/admin/plan/get-all-plans
GET	/admin/plan/get-plan/:planId
GET	/admin/plan/get-all-features/:planId
GET	/admin/plan/get-feature/:featureId
GET /admin/plan/get-all-plans-and-features
DELETE /admin/plan/delete-plan/:planId
DELETE /admin/plan/delete-feature/:featureId