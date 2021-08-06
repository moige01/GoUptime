<!--
 Copyright (c) 2021 Moisés González

 This software is released under the MIT License.
 https://opensource.org/licenses/MIT
-->

# GoUptime

## Table of Contents

- [About](#about)
- [Contributing](#contributing)
    - [DataPopulation](#data_population)
    - [Handlers](#handlers)
- [Flow](#flow)
- [Prerequisites](#prerequisites)

## About <a name="about"></a>

GoUptime lets you verify if your favorites page are Up and notify you if there are something wrong with them!

## Contributing <a name="contributing"></a>

We are using a [Priority Queue](https://en.wikipedia.org/wiki/Priority_queue) to handle a pages list with their respective priority.
You can see the implementation right [here](./priority_queue.go).

The `UptimeChecker` structure ([here](./uptime_checker.go)) is your tools to check each page and do something with the result of.
It use the `PriorityQueue` type to iterate over all pages ordered by their priority (descending order). On each iterate, it just do a HTTP GET Request against the
URL provided and execute a method passing the URL, HTTP Status Code, HTTP Status Message and and `sync.WaitGroup`. The WaitGroup is used
to communicate to the main thread that all pages were handled since we're using go routines.

In essence, the project is extensible. For that we created two interfaces: `DataPopulation` and `Handlers`, [here](./data_population.go) and [here](./handlers.go) respective.

#### DataPopulation <a name="data_population">

This interface **MUST** be implemented if you want to create your own way to populate a list of pages. This interface only need to implement a single "method":
`Dispatch` that receive a `map[string]HandlerFunc` (hash list of handlers functions as value with their identifers as key) and returns a pointer to a `PriorityQueue`.
The source of your data runs by your own. Whatever your data source are, you **MUST** give a list of: page URL, priority and a handler identifer, e.g:

```python
[
    ["www.google.com", 100, "LOG"],
    ["www.alibaba.com", 20, "DISCORD"],
    ...
]
```

You will need to use the `Page` struct and the `Node` struct to construct your `PriorityQueue`. There is a built-in called `CSVDataPopulation` ([here](./csv_data_population.go)). It can help you to get familiarized.

NOTE: You have an sample CSV [here](./csv/pages.example.csv)

#### Handlers <a name="handlers"></a>

This interface **MUST** be implemented if you want to create your own way to handle the result of the `UptimeChecker`. In each iteration, `UptimeChecker` will call the `HandlerFunc` associated with the target page. You only need to implement the `Dispatch` method, receiving the Page URL, HTTP Status Code, HTTP Status Message and and pointer to a `sync.WaitGroup`. The first thing that you need to do is declare a `defer` to the `Done` method of the `WaitGroup`. This is **MANDATORY**. No exceptions. If you don't do that, the process will freeze until you force their exit.

Handlers let you do something with the result of the request against the target page, e.g., send email notification, send WhatsApp message, Discord message, etc.
There are three handlers built-in: `StdOutHandler` (useful for testing), `LoggerHandler` (write to a log file) and `DiscordHandler` (send message to a Discord channel), [here](./stdout_handler.go), [here](./log_handler.go) and [here](./discord_handler.go) respective.

Each handler **MUST** be registered in the `main` function. The `RegisterHandler` method of `Handlers` struct will do the work. You job is to give to you function an unique identifer as key and a pointer to your `HandlerFunc` as value. You can see the [main](./main.go) file to see how that work.

## Flow <a name="flow"></a>

We use [gocron](https://github.com/go-co-op/gocron) to make a unix CRON like behavior. So, the first step in `main` function is register all of the `HandlerFunc`s available. Then, we create a time frequency for our `UptimeChecker`, for now is executed every five minutes. On each call, we initialize our `UptimeChecker` with the `Init` method, passing our `DataPopulation` and our `Handlers` as arguments. Next, the `VerifyStatus` method is executed, which iterate over our populated pages and executing their `HandlerFunc`.

## Prerequisites <a name="prerequisites"></a>

* Golang (tested on go1.16.6 linux/amd64)

## Usage

To use with the `CSVDataPopulation` you need to set an environment var called `CSV_PATH` which contain the **absolute** path of the CSV file containing your populated pages with the format `url,priority,handler` (same as the CSV sample in this repo).

To use the Discord handler, you need to set two environment vars: `DISCORD_BOT_TOKEN` and `DISCORD_CHANNEL_ID`. Self explanatory.

```sh
$ CSV_PATH=$HOME/pages.csv DISCORD_BOT_TOKEN=XXX DISCORD_CHANNEL_ID=XXX ./gouptime
```
