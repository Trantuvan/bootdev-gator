# bootdev-gator

RSS feed aggreGator 🐊. It's a CLI tool that allows users to:

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

`Learning Goals`

- Learn how to integrate a Go application with a PostgreSQL database
- Practice using your SQL skills to query and migrate a database (using [sqlc](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html) and [goose](https://github.com/pressly/goose), two lightweight tools for typesafe SQL in Go)
- Learn how to write a long-running service that continuously fetches new posts from RSS feeds and stores them in the database

## Installation

- The aggreGator 🐊 requires Golang and Postgres as dependencies.
- `.gatorconfig.json` will be needed as config file. Create this in home director with the structure

```json
{
  "db_url": "connectionString",
  "current_user_name": "userName"
}
```

- To install the aggreGator 🐊 runs. After this `bootdev-gator` appears as a command in the shell.

```bash
go install
```

Some commands to run:

- `bootdev-gator register <user_name>`
- `bootdev-gator addfeed <feed_name> <feed_url>`
- `bootdev-gator agg <duration_per_fetch>`
- `bootdev-gator browse <take> <skip>`
