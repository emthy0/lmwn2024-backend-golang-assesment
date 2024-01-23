# LMWN Assesment Covid-19 Summary API

## Overview

This project is part of Lineman Wongnai Junior Internship assesment

Given follwing instruction

>Please create a simple go project that satisfy the following requirements*
You're asked to write a simple JSON API to summarize COVID-19 stats using this public API, https://static.wongnai.com/devinterview/covid-cases.json.
>Your project must use Go, Go module, and Gin framework

>You create a JSON API at this endpoint /covid/summary

>The JSON API must count number of cases by provinces and age group

>There are 3 age groups: 0-30, 31-60, and 60+ if the case has no age data, please count as "N/A" group

>If the case has no province data, please count it as "N/A".

>We encourage you to write tests, which we will give you some extra score"

```json
--- sample response --
{
  "Province": {
    "Samut Sakhon": 3613,
    "Bangkok": 2774,
    "N/A": 6
  },
  "AgeGroup": {
    "0-30": 300,
    "31-60": 150,
    "61+": 250,
    "N/A": 4
  }
}
```

## Demo

This is demo using over engineering deployment on my personal kuberetes cluster.

API: [`https://lmwn-covid-summary.kuranasaki.work/covid/summary`](https://lmwn-covid-summary.kuranasaki.work/covid/summary)

Docs: [`https://lmwn-covid-summary.kuranasaki.work/docs/index.html`](https://lmwn-covid-summary.kuranasaki.work/docs/index.html)

Metrics: [`https://lmwn-covid-summary.kuranasaki.work/metrics`](https://lmwn-covid-summary.kuranasaki.work/metrics)

## Usage



### Install Dependency

` go mod tidy `

### Start develop

I prefer using air package to enable hot reload, I have create air configuration for both Mac, Linux, Windows, feel free to use with single command.

`make run`

API: [`http://localhost:8080/covid/summary`](http://localhost:8080/covid/summary)

Docs: [`http://localhost:8080/docs/index.html`](http://localhost:8080/docs/index.html)

### Test

`make test`

### Build

`make build`

### Clean

`make clean`

