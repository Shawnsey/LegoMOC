# LegoMOC
Application that allows lego creators the ability to post MOCs, plans, and parts lists. Users can also purchase MOCs direct from the store and have pick a bricks sent to there house.

Going to use Golang for the backend
Postgres as the db
go-chi for webservice framework
Auth0 for jwt auth

### End Goals

Create UI where users can do the following:
- Sign in to account: buyer, creator
- Add a creation with images, instructions, description, parts list
- Purchase creations
- visualize order prior to payment

Use S3/ block storage to house Images, instructions under a creators id and creation id

Automatically create purchase order to Lego store pick a brick API

Containerize everything, ease of spin up with docker-compose. Long term build manifests for kubernetes

Utilize some User mgmt service for housing user info either Keycloak or Auth0

use jwt for authentication/authorization of apis

