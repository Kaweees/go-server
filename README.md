<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<div align="left">

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]

</div>

<a href="https://github.com/nvm-sh/logos">
  <source media="(prefers-color-scheme: dark)" srcset="assets/img/go-logo-light-mode.svg">
  <img alt="Text changing depending on mode. Light: 'Rust Logo Light Mode' Dark: 'Go Logo Dark Mode'" src="assets/img/go-logo-dark-mode.svg" align="right" width="150">
</a>

<div align="left">
  <h1><em>~go-server</em></h1>
</div>

<!-- ABOUT THE PROJECT -->

A template for developing concurrent http servers in Go

### Built With

[![Go][Go-shield]][Go-url]
[![GitHub Actions][github-actions-shield]][github-actions-url]

<!-- GETTING STARTED -->

## Getting Started

To get a local copy of the project up and running on your machine, follow these simple steps:

### Prerequisites

Before attempting to build this project, make sure you have [Go](https://go.dev/doc/install) installed on your machine.

### Installation

To get a local copy of the project up and running on your machine, follow these simple steps:

1. Clone the repo
   ```sh
   git clone https://github.com/Kaweees/go-server.git
   cd go-server
   ```
2. Install the project dependencies 
   ```sh
   go mod tidy
   ```

<!-- USAGE EXAMPLES -->

## Usage

#### Available Commands

Here is a list of all the available commands:

| Script      | Description                                                                                   |
| ----------- | --------------------------------------------------------------------------------------------- |
| go mod tidy | Installs all the dependencies listed in `go.mod`.                                             |
| go run .    | Runs the application without building it.                                                     |
| go build .  | Builds the application into an executable file for different platforms using the Go compiler. |

_For more examples, please refer to the [Go Command Documentation](https://go.dev/doc/cmd)_

<!-- PROJECT FILE STRUCTURE -->

## Project Structure

```
. go-server/
├── logger.go                      - 
├── main.go                        -
├── README.md                      - you are here
├── request_handler.go             -
├── request_types.go               -
└── response_types.go              -
```

<!-- LICENSE -->
<!-- https://choosealicense.com/ -->

## License

go-server is distributed under the terms of the GNU General Public License v3.0, as I firmly believe that collaborating on free and open-source software fosters innovations that mutually and equitably beneficial to both collaborators and users alike. See [LICENSE](LICENSE) for details and more information.

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/Kaweees/go-server.svg?style=for-the-badge
[contributors-url]: https://github.com/Kaweees/go-server/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/Kaweees/go-server.svg?style=for-the-badge
[forks-url]: https://github.com/Kaweees/go-server/network/members
[stars-shield]: https://img.shields.io/github/stars/Kaweees/go-server.svg?style=for-the-badge
[stars-url]: https://github.com/Kaweees/go-server/stargazers

<!-- MARKDOWN SHIELD BADGES & LINKS -->
<!-- https://github.com/Ileriayo/markdown-badges -->
[Go-shield]: https://img.shields.io/badge/go-%23008080.svg?style=for-the-badge&logo=go&logoColor=00add8&labelColor=222222&color=00add8
[Go-url]: https://gohugo.io/
[github-actions-shield]: https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=2671E5&labelColor=222222&color=2671E5
[github-actions-url]: https://github.com/features/actions
