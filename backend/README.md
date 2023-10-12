### Backnd with GoFiber

## Routes
GET /api/v1/home

# auth
POST /api/v1/auth/signup
POST /api/v1/auth/signin

# Project Doing Request
POST /pdr/create (auth required) (json with title, description, suggestedPrice)
GET /pdr/get/:id (auth required)
GET /pdr/get-all (auth required)
POST /pdr/edit/:id (auth required) (json with title, description, suggestedPrice)