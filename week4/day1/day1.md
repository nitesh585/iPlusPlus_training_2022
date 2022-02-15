## BDD (Behaviour Driven development)

Behaviour Driven Development (BDD) is a synthesis and refinement of practices stemming from Test Driven Development (TDD) and Acceptance Test Driven Development (ATDD). BDD augments TDD and ATDD with the following tactics:

- apply the “[Five Why’s](https://en.wikipedia.org/wiki/Five_whys)” principle to each proposed [user story](<https://www.agilealliance.org/glossary/user-stories/#q=~(infinite~false~filters~(postType~(~'page~'post~'aa_book~'aa_event_session~'aa_experience_report~'aa_glossary~'aa_research_paper~'aa_video)~tags~(~'user*20stories))~searchTerm~'~sort~false~sortDirection~'asc~page~1)>), so that its purpose is clearly related to business outcomes
- thinking “from the outside in”, in other words implement only those behaviors which contribute most directly to these business outcomes, so as to minimize waste
- describe behaviors in a single notation which is directly accessible to domain experts, testers and developers, so as to improve communication
- apply these techniques all the way down to the lowest levels of abstraction of the software, paying particular attention to the distribution of behavior, so that evolution remains cheap

### Example of BDD

In order to ensure the working of Login Functionality, we are developing acceptance test cases on the basis of BDD.

> **Feature**: Login Function
> To enter in the System
> User must be able to
> Access software when login is successful
>
> **Scenario**: Login
> **Given** User has its Email
> **And** Password
> **When** User enters the correct Email and Password
> **Then** It should be logged in
>
> **Scenario**: Unsuccessful Login
> **When** User enters either wrong Email or Password
> **Then** It should be reverse back on the login page with an error message

Several BDD Tools are in use for different platforms and programming languages. They are −

- Cucumber (Ruby framework)
- SpecFlow (.NET framework)
- Behave (Python framework)
- JBehave (Java framework)
- JBehave Web (Java framework with Selenium integration)
- Lettuce (Python framework)
- Concordion (Java framework)
- Behat (PHP framework)
- Kahlan (PHP framework)
- DaSpec (JavaScript framework)
- Jasmine (JavaScript framework)
- Cucumber-js (JavaScript framework)
- Squish GUI Tester (BDD GUI Testing Tool for JavaScript, Python, Perl, Ruby and Tcl)
- Spock (Groovy framework)
- Yadda (Gherkin language support for frameworks such as Jasmine (JavaScript framework))

## Go module

A module is a collection of Go packages stored in a file tree with a go.mod file at its root. The go.mod file defines the module’s module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build. Each dependency requirement is written as a module path and a specific semantic version.

### Create Go module

Let us now see how to create modules in Golang. To create a module first create a directory and then go inside it using the following commands:

```bash
mkdir go_modules
cd go_modules
```

To initialize the current directory as the root of the module that will allow us to manage dependencies, use the following command:

```bash
go mod init go_modules
```

As we are working outside the `$GOPATH/src` module, we need to explicitly specify the name of the module during initialization.

We can now check if the `go.mod` file is created and if present, the contents of it.

Next step is to create a simple Go file with the following code:

```go
// file inside the current module
package hello

import("fmt")

func initialiser() string {

       fmt.Printf("In hello package. \n")

       // returns the current module
       // and the package name
       return_string := "Module : go_modules."
       return return_string

}
```

## Must read resources

- [BDD: Behaviour-Driven Development](https://cucumber.io/blog/bdd/getting-started-with-bdd-part-1/)
- [Gherkin Reference](https://cucumber.io/docs/gherkin/reference/)
- [Cucumber for Go lang](https://github.com/cucumber/godog)
- [Using Go Modules](https://go.dev/blog/using-go-modules)
- [Migrating to Go Modules](https://go.dev/blog/migrating-to-go-modules)
- [Publishing Go Modules](https://go.dev/blog/publishing-go-modules)
- [Go Modules: v2 and Beyond](https://go.dev/blog/v2-go-modules)
- [Keeping Your Modules Compatible](https://go.dev/blog/module-compatibility)
