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


# Plan

GET	/plan/get-all-plans

GET	/plan/get-plan/:planId

GET	/plan/get-all-features/:planId

GET	/plan/get-feature/:featureId

GET /plan/get-all-plans-and-features

POST /admin/plan/create                       json(name, price)

POST /admin/plan/create-features/:planId      json(name, price, isHave)

POST /admin/upload/plan/image/:planId  (Content-Type: multipart/form-data) (with FormData js object) (without base 
url)

PUT /admin/plan/edit/:planId                  json(name, price)

PUT	/admin/plan/edit-feature/:featureId       json(name, price, isHave)

DELETE /admin/plan/delete-plan/:planId

DELETE /admin/plan/delete-feature/:featureId


# Member

GET /member/get-all

GET /member/get/:memberId

POST /admin/member/create

Put /admin/member/edit/:memberId

Delete /admin/member/delete/:memberId


# Article
GET /article/get-all/:page (page is number)

GET /article/get/:articleTitleUrl (articleTitleUrl is in json)

POST /admin/article/create              josn(title, body, tags, shortDesc) (tags splited with | example: 'test|art')
PUT /admin/article/edit/:articleId      josn(title, body, tags, shortDesc)
DELETE /admin/article/delete/:articleId

