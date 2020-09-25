# Bands CLI

This repository contains the CLI for [Bands](https://bands.sh).
Bands.sh lets you create payment pages through a simple yaml file.

This repository is also used for issue tracking and feature requests.

## Installation

`bands` has been tested on Mac, Linux, and Windows WSL.

```shell
curl -fsSL 'https://bands-sh.s3.amazonaws.com/get-bands.sh' | sh
```

## Usage

### Create a sample yaml, lesson.yaml.

```
---
deployment_name: react_lessons_live

gateway:
  type: stripe

customer_portal: true

products:
  - id: react_lessons_live
    name: Online 1-on-1 React Lessons
    description: Online 1-on-1 React Lessons with World Class Developers. 
    images: []
    metadata:
      bullets:
        - Access to a professional React developer
        - Practice for interviews, or code together with a mentor
        - Cover basic and advanced JavaScript and React topics depending on your level
    type: service

plans:
  - product: react_lessons_live
    nickname: Once a week
    metadata:
      bullets:
        - Dedicated mentor
        - Go from a beginner to solid React developer in 4 months
        - Follow our proven and tested learning path, or decide on your own
        - Great for long term projects
      images: []
    usage_type: licensed
    interval: monthly
    currency: usd
    amount: '299.00'

  - product: react_lessons_live
    nickname: One time session
    metadata:
      bullets:
        - StackOverflow is great, but a pro helping you is better
        - Great for beginners or experts
        - Cover a wide range of topics
        - Prep for interviews with a mock interviewer
      images: []
    currency: usd
    amount: '99.00'
```

### Create a Bands.sh account.

`bands init --email <email>`

### Create a payment page.

`bands up -f lesson.yaml`

### Archive the payment page.

`bands down -f lesson.yaml`

## Contributing

Contributions are very much welcome!
