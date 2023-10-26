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

GET	/admin/plan/get-all-plans

GET	/admin/plan/get-plan/:planId

GET	/plan/get-all-features/:planId

GET	/admin/plan/get-feature/:featureId

GET /plan/get-all-plans-and-features

GET /plan/get-featured

POST /admin/plan/create                       json(name, price)

POST /admin/plan/create-features/:planId      json(name, price, isHave)

POST /admin/upload/plan/image/:planId  (Content-Type: multipart/form-data) (with FormData js object) (without /api/v1)

PUT /admin/plan/edit/:planId                  json(name, price)

PUT	/admin/plan/edit-feature/:featureId       json(name, price, isHave)

DELETE /admin/plan/delete-plan/:planId

DELETE /admin/plan/delete-feature/:featureId


# Member

GET /member/get-all

GET /member/get/:memberId

POST /admin/member/create

POST /admin/upload/member/image/:memberId (without /api/v1)

Put /admin/member/edit/:memberId

Delete /admin/member/delete/:memberId


# Article
GET /article/get-all/:page (page is number)

GET /article/get/:articleTitleUrl (articleTitleUrl is in json)

GET /article/get-featured

POST /admin/article/create              josn(title, body, tags, shortDesc) (tags splited with | example: 'test|art')

POST /admin/upload/article/image/:articleId (without /api/v1)

PUT /admin/article/edit/:articleId      josn(title, body, tags, shortDesc)

DELETE /admin/article/delete/:articleId


# Work Smaple

GET /work-sample/get-all/:page (page is number)

GET /work-sample/get-featured

GET /admin/works-sample/get/:workSampleId

POST /admin/work-sample/create              josn(title, imagePath, siteUrl, description, isFeatured)

POST /admin/upload/work-smaple/image/:workSampleId  (without /api/v1)

PUT /admin/article/edit/:articleId          josn(title, imagePath, siteUrl, description, isFeatured)

DELETE /admin/article/delete/:articleId


# Site Features
GET /site-features/get-all


# Comments

GET /comment/get-all-for-article/:articleTitleUrl

POST /comment/create json(body)

PUT /comment/edit/:id json(body)

DELETE /comment/delete/:id

# Contactus

GET /contactus/get-all-for-user

GET /contactus/get/:contactusTitle

GET /admin/contactus/get-all

POST /contactus/create json(title, body)

PUT /contactus/edit/:titleUrl json(title, body)

DELETE /contactus/delete/:titleUrl

# Customer

GET /customer/get-all

GET /customer/get-feautred

GET /admin/customer/get/:customerId

POST /admin/customer/create json(name, siteUrl, isFeatured)

POST /admin/upload/customer/image/:customerId (without /api/v1)

PUT /admin/customer/edit/:customerId json(name, siteUrl, isFeatured)

DELETE /admin/customer/delete/:customerId