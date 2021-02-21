# Server Architecture

## Intro

I don't actually think this code is bad, although it is very proof-of-conceptual. It's the structure I came up with for the rest Go API I planned to use for [no-manga](https://github.com/SpencerWhitehead7/no-manga). However, I ultimately switched to GQL, and now I need somewhere to preserve this because I did a lot of research and spent a long time on it and I don't want to just lose it.

## API

The API is a REST-ful, and organized into endpoints by the types of data (eg, manga, mangaka, magazines) the app needs and the functionality it performs (eg, search). Each resource has its own set of `handler`, `controller`, `repository` and `model`. The server is initialized `main`.

### `Main`

The server runs from `server/main.go`, which connects to the DB, creates a router, and sets up all the route `handlers`. Route `handlers` are passed their base route and the DB connection.

### `Handler`

`Handlers` instantiate their `controller` with its `repository` and handle requests for all their sub-routes. A request is handled by validating its params, handling errors if necessary, then invoking a `controller` method.

### `Controller`

`Controllers` perform all the endpoint's business logic, handling errors if necessary, and ultimately respond to the request. They are given `repositories`, which they use to perform any necessary DB calls.

### `Repository`

`Repositories` are given a connection the DB and do all the app's DB calls. They are used in `controllers`, and implement interfaces for querying the database for structs for `models`.

### `Models`

`Models` define structs representing DB rows.
