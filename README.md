<div align="center" id="readme-top">

  <h1>Task Tracker CLI</h1>

  <p>
    A simple CLI task tracker that helps you to track your tasks and manage your to-do list. </p>
 <p>Built with Go, this tool is my implementation of the <a href="https://roadmap.sh/projects/task-tracker">Task Tracker</a> challenge from <a href="https://roadmap.sh">roadmap.sh</a>.
  </p>

<!-- Badges -->
<p>
  <a href="https://github.com/vicjeremy/task-cli/graphs/contributors">
    <img src="https://img.shields.io/github/contributors/vicjeremy/task-cli" alt="contributors" />
  </a>
  <a href="">
    <img src="https://img.shields.io/github/last-commit/vicjeremy/task-cli" alt="last update" />
  </a>
  <a href="https://github.com/vicjeremy/task-cli/network/members">
    <img src="https://img.shields.io/github/forks/vicjeremy/task-cli" alt="forks" />
  </a>
  <a href="https://github.com/vicjeremy/task-cli/stargazers">
    <img src="https://img.shields.io/github/stars/vicjeremy/task-cli" alt="stars" />
  </a>
  <a href="https://github.com/vicjeremy/task-cli/issues/">
    <img src="https://img.shields.io/github/issues/vicjeremy/task-cli" alt="open issues" />
  </a>
  <a href="https://github.com/vicjeremy/task-cli/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/vicjeremy/task-cli.svg" alt="license" />
  </a>
</p>

<h4>
    <a href="https://github.com/vicjeremy/task-cli/issues/">Report Bug</a>
  <span> · </span>
    <a href="https://github.com/vicjeremy/task-cli/issues/">Request Feature</a>
  </h4>
</div>

<br />

<!-- Table of Contents -->

# :notebook_with_decorative_cover: Table of Contents

- [About the Project](#star2-about-the-project)
  - [Screenshots](#camera-screenshots)
  - [Tech Stack](#space_invader-tech-stack)
  - [Features](#dart-features)
- [Getting Started](#toolbox-getting-started)
  - [Prerequisites](#bangbang-prerequisites)
  - [Installation](#gear-installation)
  - [Run Locally](#running-run-locally)
- [Usage](#eyes-usage)]
- [Contributing](#wave-contributing)
- [License](#warning-license)
- [Contact](#handshake-contact)
- [Acknowledgements](#gem-acknowledgements)

<!-- About the Project -->

## :star2: About the Project

<!-- Screenshots -->

### :camera: Screenshots

<div align="center">
  <img src="img/example-test.png" style="width:600px;height:400px" alt="screenshot" />
</div>

<!-- TechStack -->

### :space_invader: Tech Stack

- <a href="https://golang.org" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="30" height="30"/>[![Go][Go]][Go-url]</a>

<!-- Features -->

### :dart: Features

- Add, Update, and Delete tasks
- Mark a task as in progress or done
- List all tasks
- List all tasks that are done
- List all tasks that are not done
- List all tasks that are in progress

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Getting Started -->

## :toolbox: Getting Started

<!-- Prerequisites -->

### :bangbang: Prerequisites\

This project uses Go as its main language. Make sure you have it installed on your machine. If not, you can install it from the official website [here](https://golang.org/).

```bash
  go version
```

<!-- Run Locally -->

### :running: Run Locally

Clone the project

```bash
  git clone https://github.com/vicjeremy/task-cli.git
```

Go to the project directory

```bash
  cd task-cli
```

build the project

```bash
  go build -o task-cli
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Usage -->

## :eyes: Usage

Use this space to tell a little more about your project and how it can be used. Show additional screenshots, code samples, demos or link to other resources.

```bash
# To see the list of available commands
task-cli --help, -h, help

# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Contributing -->

## :wave: Contributing

<a href="https://github.com/vicjeremy/task-cli/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=vicjeremy/task-cli" />
</a>

Contributions are always welcome!

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- License -->

## :warning: License

Distributed under the MIT License. See [LICENSE.txt](LICENSE.txt) for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Contact -->

## :handshake: Contact

Vic Jeremy - [@viccjeremy](https://instagram.com/viccjeremy) - [vicjeremyp@gmail.com](mailto:vicjeremyp@gmail.com)

Project Link: [https://github.com/vicjeremy/task-cli](https://github.com/vicjeremy/task-cli)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Acknowledgments -->

## :gem: Acknowledgements

- [Backend Project ideas from roadmap.sh](https://roadmap.sh/backend/projects)

Packages or Libraries used:

- [strings](https://pkg.go.dev/strings)
- [fmt](https://pkg.go.dev/fmt)
- [os](https://pkg.go.dev/os)
- [strconv](https://pkg.go.dev/strconv)
- [time](https://pkg.go.dev/time)
- [encoding/json](https://pkg.go.dev/encoding/json)
- [log](https://pkg.go.dev/log)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

[Go]: https://img.shields.io/badge/GOlang-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://golang.org/
