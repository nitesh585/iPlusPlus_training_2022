# Software Methodologies and terminologies

| Functional requirements                                                                                                                                                                                                                                   | Non-Functional requirements                                                                                                                                                                                                                                                        |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| A functional requirement defines a system or its component.                                                                                                                                                                                               | A non-functional requirement defines the quality attribute of a software system.                                                                                                                                                                                                   |
| It specifies “What should the software system do?”                                                                                                                                                                                                        | It places constraints on “How should the software system fulfill the functional requirements?”                                                                                                                                                                                     |
| Functional requirement is specified by User.                                                                                                                                                                                                              | Non-functional requirement is specified by technical peoples e.g. Architect, Technical leaders and software developers.                                                                                                                                                            |
| It is mandatory.                                                                                                                                                                                                                                          | It is not mandatory.                                                                                                                                                                                                                                                               |
| It is captured in use case.                                                                                                                                                                                                                               | It is captured as a quality attribute.                                                                                                                                                                                                                                             |
| Defined at a component level.                                                                                                                                                                                                                             | Applied to a system as a whole.                                                                                                                                                                                                                                                    |
| Helps you verify the functionality of the software.                                                                                                                                                                                                       | Helps you to verify the performance of the software.                                                                                                                                                                                                                               |
| Functional Testing like System, Integration, End to End, API testing, etc are done.                                                                                                                                                                       | Non-Functional Testing like Performance, Stress, Usability, Security testing, etc are done.                                                                                                                                                                                        |
| Usually easy to define.                                                                                                                                                                                                                                   | Usually more difficult to define.                                                                                                                                                                                                                                                  |
| Example: <b>1)</b> Authentication of user whenever he/she logs into the system. <b>2)</b> System shutdown in case of a cyber attack. <b>3)</b> A Verification email is sent to user whenever he/she registers for the first time on some software system. | Example: <b>1)</b> Emails should be sent with a latency of no greater than 12 hours from such an activity. <b>2)</b> The processing of each request should be done within 10 seconds <b>3)</b> The site should load in 3 seconds when the number of simultaneous users are > 10000 |

---

### WaterFall Model:

Waterfall approach was first SDLC Model to be used widely in Software Engineering to ensure success of the project. In "The Waterfall" approach, the whole process of software development is divided into separate phases. In this Waterfall model, typically, the outcome of one phase acts as the input for the next phase sequentially. [learn more](https://www.tutorialspoint.com/sdlc/sdlc_waterfall_model.htm)
<img src="https://asperbrothers.com/wp-content/uploads/2019/08/Zrzut-ekranu-2019-08-27-o-15.55.16.png" />

### V model:

V-Model also referred to as the Verification and Validation Model. In this, each phase of SDLC must complete before the next phase starts. It follows a sequential design process same as the waterfall model. Testing of the device is planned in parallel with a corresponding stage of development. [learn more](https://www.javatpoint.com/software-engineering-v-model)
![VModel](https://i.stack.imgur.com/GxMAg.png)

### Agile Model:

agile software development methodologies are all about delivering small pieces of working software quickly to improve customer satisfaction. These methodologies use adaptive approaches and teamwork to focus on continuous improvement. Usually, agile software development consists of small, self-organizing teams of software developers and business representatives regularly meeting in-person throughout the software development life cycle. Agile favors a lightweight approach to software documentation and embraces—rather than resists—changes at any stage of the life cycle. [learn more](https://www.javatpoint.com/agile-methodology)
![Agile](https://ncube-digest.com/wp-content/uploads/2019/08/Agile-Methodology.png)

---

## DDs in SDE

### TDD (Test Driven Development) :

TDD is an iterative development process. Each iteration starts with a set of tests written for a new piece of functionality. These tests are supposed to fail during the start of iteration as there will be no application code corresponding to the tests. In the next phase of the iteration, Application code is written with an intention to pass all the tests written earlier in the iteration. Once the application code is ready tests are run.
Any failures in the test run are marked and more Application code is written/re-factored to make these tests pass. Once application code is added/re-factored the tests are run again. This cycle keeps on happening until all the tests pass. Once all the tests pass we can be sure that all the features for which tests were written have been developed. [learn more](https://www.guru99.com/test-driven-development.html)
![tdd](https://www.guru99.com/images/8-2016/081216_0811_TestDrivenD1.png)

### BDD (Behavior-driven development) :

BDD is using examples at multiple levels to create a shared understanding and surface uncertainty to deliver software that matter.
Behavior Driven Development focuses on −

- Providing a shared process and shared tools promoting communication to the software developers, business analysts and stakeholders to collaborate on software development, with the aim of delivering product with business value.
- What a system should do and not on how it should be implemented.
- Providing better readability and visibility.
- Verifying not only the working of the software but also that it meets the customer’s expectations.[learn more](https://www.tutorialspoint.com/behavior_driven_development/behavior_driven_development_introduction.htm)

### DDD (Domain-driven development):

A domain model offers several benefits some of which are:

- It helps the team create a common model, between the business and IT stakeholders in the company, that the team can use to communicate about the business requirements, data entities, and process models.
- The model is modular, extensible and easy to maintain as the design reflects the business model.
- It improves the reusability and testability of the business domain objects.

[learn more](https://www.infoq.com/articles/ddd-in-practice/#:~:text=Domain%20Driven%20Design%20and%20Development%20In%20Practice%201,Testing%2FContinuous%20Integration.%20...%204%20Sample%20Application%20Design.%20)

### MDD (Model-driven development):

Model-Driven Development (MDD) is an approach that uses models as a specification of software and transformations of those models to get the source code. Models are created before the source code is written or generated. MDD aims at speeding up the software development and making it more cost efficient, by using the models to visualise the code and, if used to raise the abstraction, the problem domain. MDD also separates implementation technology from the business logic of a program.
The idea behind MDD is that, it is possible to create models of a system that can then be transformed into a real thing. Systems are modelled at several levels of abstraction and from multiple perspectives and the created models are the primary artefacts of software development. [learn more](https://www.vttresearch.com/sites/default/files/pdf/workingpapers/2009/W114.pdf)

## SRE fundamentals: SLIs, SLAs and SLOs

The concept of SRE starts with the idea that metrics should be closely tied to business objectives. We use several essential tools—SLO, SLA and SLI—in SRE planning and practice. [learn more](https://cloud.google.com/blog/products/devops-sre/sre-fundamentals-slis-slas-and-slos)

### SLO (Service-Level Objective):

When we set out to define the terms of SRE, we wanted to set a precise numerical target for system availability. We term this target the availability Service-Level Objective (SLO) of our system. Any discussion we have in the future about whether the system is running sufficiently reliably and what design or architectural changes we should make to it must be framed in terms of our system continuing to meet this SLO.

### SLA (Service Level Agreement):

An SLA – Service-Level Agreement – is a written contract agreement between the quality team (provider) and development team (customer) to ensure everyone understands the engagement, service of work, and expectations from the provider. If specific conditions are not met there are consequences and often the quality of the product suffers.

### SLI (Service-Level Indicator):

When we evaluate whether our system has been running within SLO for the past week, we look at the SLI to get the service availability percentage. If it goes below the specified SLO, we have a problem and may need to make the system more available in some way, such as running a second instance of the service in a different city and load-balancing between the two. If you want to know how reliable your service is, you must be able to measure the rates of successful and unsuccessful queries as your SLIs.

## WCAG A, AA, AAA Accessibility Standards

Accessibility standards for the internet are set by the Web Content Accessibility Guidelines (WCAG). These are the guidelines for making the web accessible to everyone, regardless of their needs or if they require assistive technologies to use the web. [learn more](https://www.bing.com/search?q=a+aa+aaa+standards+&qs=n&form=QBRE&msbsrank=1_2__0&sp=-1&pq=a+aa+aaa+standards+&sc=2-19&sk=&cvid=683C36F80227449EB89676A5D372C22B)

- <b>A (Single A)</b> is viewed as the minimum level of requirement which all websites, apps, and electronic content such as documents should adhere to.
- <b>AA (Double A)</b> is viewed as the acceptable level of accessibility for many online services, which should work with most assistive technology which is now widely available on both desktop and mobile devices, or which can be purchased as a third-party installation.
- <b>AAA (Triple A)</b> compliance is viewed as the gold standard level of accessibility, which provides everything for a complete accessible offering, including all the bells and whistles which make the difference between a very good experience and an excellent one.

### Availability: five 9s

Available for use means that it performs its agreed function successfully when required.[learn more](https://docs.aws.amazon.com/wellarchitected/latest/reliability-pillar/availability.html)

### OWASP Top Ten:

The OWASP Top 10 is a standard awareness document for developers and web application security. It represents a broad consensus about the most critical security risks to web applications. [learn more](https://owasp.org/www-project-top-ten/)

## Tools:

- [JIRA](https://www.atlassian.com/software/jira?gclid=a0779a3eb40415a1133cfbd11c71cb07&gclsrc=3p.ds&&adgroup=1307319557870123&campaign=380755286&creative=81707524366663&device=c&keyword=what%20is%20jira&ds_k=what+is+jira&matchtype=e&network=o&ds_kids=p54415850281&ds_e=MICROSOFT&ds_eid=700000001738795&ds_e1=MICROSOFT&msclkid=a0779a3eb40415a1133cfbd11c71cb07): It is a tool developed by Australian Company Atlassian. This software is used for bug tracking, issue tracking, and project management.
- [Trello](https://trello.com/): Trello is a popular, simple, and easy-to-use collaboration tool that enables you to organize projects and everything related to them into boards.
- [MeisterTask](https://www.meistertask.com/): The MeisterTask app is one of the most optimal solutions for monitoring ongoing work. It excels in organizing and tracking tasks in progress rather than larger and more complex projects.
- [Confluence](https://www.atlassian.com/software/confluence?gclid=6934a328ca031371fc73c2c762fa10f8&gclsrc=3p.ds&&adgroup=1301821999416486&campaign=380755062&creative=81363927004378&device=c&keyword=confulence&ds_k=confulence&matchtype=e&network=o&ds_kids=p54414063233&ds_e=MICROSOFT&ds_eid=700000001721838&ds_e1=MICROSOFT&msclkid=6934a328ca031371fc73c2c762fa10f8): Confluence is basically a team collaboration application or software that gives a platform for all teams working on a project to work together and share updates & information effectively. This can reduce a lot of communication gaps on various matters related to the SDLC.
- [Swagger](https://swagger.io/): Swagger is the standard way of documenting the Standard APIs.
- [Real World PlantUML](https://real-world-plantuml.com/): Used to design UML diagrams.
- [Figma](https://www.figma.com/): Figma is a cloud-based design and prototyping tool for digital projects. It’s made so that users can collaborate on projects and work pretty much anywhere.

---

Thanks and credit to all resources creators.
