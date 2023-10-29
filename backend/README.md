# Backnd with GoFiber

## Info
Content-Type: application/json
base url: /api/v1/...


### Routes
GET /api/v1/


#### auth

POST /signup (json with email, password, name, family(optional))
    password lenght bigger than 8
    email example test@test.com
    name must letter without !@#$%^&*()_+-;'"

POST /signin (json with email, password)


#### Project Doing Request

POST /pdr/create        (admin required)    json(title, description, suggestedPrice(optional))

GET /pdr/get/:id        (admin required)

GET /pdr/get-all        (admin required)

PUT /pdr/edit/:id       (admin required)    json(title, description, suggestedPrice(optional))

DELETE /pdr/delete/:id  (admin required)


#### Plan

<b>Public:</b>

GET /plan/get-all/?page=1&pageLimit=20

For main page
GET /plan/get-featured

<b>Admin:</b>

GET	/admin/plan/get-all-plans

GET	/admin/plan/get-plan/:planId

GET	/admin/plan/get-feature/:featureId

GET	/plan/get-all-features/:planId

POST /admin/plan/create                       json(name, price)

POST /admin/plan/create-features/:planId      json(name, price, isHave)

POST /admin/upload/plan/image/:planId  (Content-Type: multipart/form-data) (with FormData js object) (without /api/v1)

PUT /admin/plan/edit/:planId                  json(name, price)

PUT	/admin/plan/edit-feature/:featureId       json(name, price, isHave)

DELETE /admin/plan/delete-plan/:planId

DELETE /admin/plan/delete-feature/:featureId


#### Member

<b>Public:</b>

GET /member/get-all

GET /member/get/:memberId

<b>Admin:</b>

POST /admin/member/create

POST /admin/upload/member/image/:memberId (without /api/v1)

Put /admin/member/edit/:memberId

Delete /admin/member/delete/:memberId


#### Blog

<b>Public:</b>

GET /blog/get-all/?page=1&pageLimit=20

GET /blog/get/:articleTitleUrl (articleTitleUrl is in json)

(For main page)
GET /blog/get-featured


<b>Admin:</b>

POST /admin/article/create              josn(title, body, tags, shortDesc) (tags splited with | example: 'test|art')

POST /admin/upload/blog/image/:articleId (without /api/v1)

PUT /admin/blog/edit/:articleId      josn(title, body, tags, shortDesc)

DELETE /admin/blog/delete/:articleId


#### Portfolio

<b>Public:</b>

(Get all portfolios)
GET /portfolio/get-all/?page=1&pageLimit=20

(For main page)
GET /portfolio/get-featured

(Get single portfolio)
GET /admin/portfolio/get/:portfolioId

<b>Admin:</b>

POST /admin/portfolio/create              josn(title, imagePath, siteUrl, description, isFeatured)

POST /admin/upload/portfolio/image/:portfolioId  (without /api/v1)

PUT /admin/portfolio/edit/:articleId          josn(title, imagePath, siteUrl, description, isFeatured)

DELETE /admin/portfolio/delete/:articleId


# Site Features
GET /mainpage/site-features/


# Comments

<b>Public:</b>

GET /comment/get-all-for-article/:articleTitleUrl

POST /comment/create json(body, commentAnswerID(0 if not answer and id of comment for answer it))

PUT /comment/edit/:id json(body, commentAnswerID(0 if not answer and id of comment for answer it))

DELETE /comment/delete/:id

# Contactus

<b>Public:</b>

POST /contactus/create json(title, body, email, number)

<b>Admin:</b>

GET /admin/contactus/get-all/?page=1&pageLimit=20

GET /admin/contactus/get/:contactusTitle

GET /admin/contactus/get-all/?page=1&pageLimit=20

DELETE /admin/contactus/delete/:titleUrl

# Customer

<b>Public:</b>

(Get all customers)
GET /customer/get-all/?page=1&pageLimit=20

(For main page)
GET /customer/get-feautred

<b>Admin:</b>

(Get single customer)
GET /admin/customer/get/:customerId

POST /admin/customer/create json(name, siteUrl, isFeatured)

POST /admin/upload/customer/image/:customerId (without /api/v1)

PUT /admin/customer/edit/:customerId json(name, siteUrl, isFeatured)

DELETE /admin/customer/delete/:customerId

# Statistics

<b>Public:</b>

GET /statistics/get-public

<b>Admin:</b>

GET /admin/statistic/get-all

POST /admin/statistic/create    json(name, name2, number, isPublic)

PUT /admin/statistic/edit/:statisticId json(name, name2, number, isPublic)

DELETE /admin/statistic/delete/:statisticId