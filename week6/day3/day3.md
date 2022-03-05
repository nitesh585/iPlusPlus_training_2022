## Sonarcube

![sonarqube-img](https://docs.sonarqube.org/latest/images/dev-cycle.png)
An automatic code review tool to detect bugs, vulnerabilities, and code smells in your code. It can integrate with your existing workflow to enable continuous code inspection across your project branches and pull requests.

### General Guidlines

Before importing test coverage and execution reports, you need to have the appropriate SonarScanner configured as part of your build pipeline. Then, complete the following steps:

- Set up your coverage tool to run as part of your build pipeline. Your coverage tool should be set up to run before the SonarScanner analysis.
- Configure the coverage tool so that location and format of the output report files matches what the SonarScanner expects.
- Configure the analysis parameters of the SonarScanner so that it can import the report files.

---

Practical: Setup and try to run the sonarqube

## Must read resources:

- [Sonarqube](https://www.sonarqube.org/)

- [Golang Sonarqube](https://docs.sonarqube.org/latest/analysis/languages/go/)
