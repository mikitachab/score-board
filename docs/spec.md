# score-board specification

## Functional requirements

Main features:

- app should provide possibility to add new game score
- app should provide posibility to view all previous games scores
- app should work al least for [7 Wonders board](https://boardgamegeek.com/boardgame/68448/7-wonders) game
- app should show score as a table

Additioanl Features:  

- Access Restrictions/Login

## Non-functional requirements

- Back-end
  - Golang's [net/http](https://golang.org/pkg/net/http/) package
  - Possibly [Go kit](https://gokit.io/) helpers
  - Database: [PostgreSQL](https://www.postgresql.org/) + [ORM GORM](https://gorm.io/index.html)
- Front-end
  - Golangs [html/template](https://golang.org/pkg/html/template/) package
  - [Bootsrap](https://getbootstrap.com/) for styling
  - Possibly injection of vanilla JS

- Deploy options
  - Golang app on [Heroku](https://devcenter.heroku.com/articles/getting-started-with-go) (FREE)

## Mode of Work

- Changes delivery via GitHub PR, al leat one approve from other member (4 eyes rule)
- Tasks Management via [Trello](https://trello.com/)
- Tests with Golangs [testing](https://golang.org/pkg/testing/) package
- Lint with [golint](https://godoc.org/golang.org/x/lint/golint)
- CI with [GitHub Actions](https://github.com/features/actions)
