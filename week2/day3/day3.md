## Jenkins

![jenkins](https://bugfender.com/wp-content/uploads/2017/11/jenkins-1-3-jenkins.jpg)
Jenkins is a self-contained, open source automation server which can be used to automate all sorts of tasks related to building, testing, and delivering or deploying software.

Jenkins can be installed through native system packages, Docker, or even run standalone by any machine with a Java Runtime Environment (JRE) installed.

### Installation (Linux)

#### Hardware recommendations:

- Hardware: see the [Hardware Recommendations](https://www.jenkins.io/doc/book/scaling/hardware-recommendations) page

#### Software requirements:

- Java: see the [Java Requirements](https://www.jenkins.io/doc/administration/requirements/java) page. Jenkins requires Java in order to run, yet certain distributions don’t include this by default and [some Java versions are incompatible](https://www.jenkins.io/doc/administration/requirements/java/) with Jenkins.
- Web browser: see the [Web Browser Compatibility](https://www.jenkins.io/doc/administration/requirements/web-browsers) page

#### Commands to install (LTS release)

```sh
curl -fsSL https://pkg.jenkins.io/debian-stable/jenkins.io.key | sudo tee \
  /usr/share/keyrings/jenkins-keyring.asc > /dev/null

echo deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc] \
  https://pkg.jenkins.io/debian-stable binary/ | sudo tee \
  /etc/apt/sources.list.d/jenkins.list > /dev/null

sudo apt-get update

sudo apt-get install jenkins
```

#### Start Jenkins

Register the Jenkins service with the command:

```sh
sudo systemctl daemon-reload
```

You can start the Jenkins service with the command:

```sh
sudo systemctl start jenkins
```

You can check the status of the Jenkins service using the command:

```sh
sudo systemctl status jenkins
```

If everything has been set up correctly, you should see an output like this:

```sh
Loaded: loaded (/etc/rc.d/init.d/jenkins; generated)
Active: active (running) since Tue 20222-02-02 16:19:01 +03; 1min ago
```

#### Post-installation setup wizard

This setup wizard takes you through a few quick "one-off" steps to unlock Jenkins, customize it with plugins and create the first administrator user through which you can continue accessing Jenkins.

#### Unlocking Jenkins

When you first access a new Jenkins instance, you are asked to unlock it using an automatically-generated password.

Browse to `http://localhost:8080` (or whichever port you configured for Jenkins when installing it).
![unlock](https://www.jenkins.io/doc/book/resources/tutorials/setup-jenkins-01-unlock-jenkins-page.jpg)
The command: `sudo cat /var/lib/jenkins/secrets/initialAdminPassword will` print the password at console and paste the password in the above text field.
On the Unlock Jenkins page, paste this password into the Administrator password field and click Continue.
Notes:

#### Customizing Jenkins with plugins

After unlocking Jenkins, the Customize Jenkins page appears. Here you can install any number of useful plugins as part of your initial setup.

Click one of the two options shown:

- <b>Install suggested plugins</b> - to install the recommended set of plugins, which are based on most common use cases.

- <b>Select plugins to install</b> - to choose which set of plugins to initially install. When you first access the plugin selection page, the suggested plugins are selected by default.

### Must read resources

- [Intro to Jenkins](https://www.lambdatest.com/blog/what-is-jenkins/)
- [Installing Jenkins for all platforms ](https://www.jenkins.io/doc/book/installing/)
- [Jenkins: Go](https://plugins.jenkins.io/golang/)
- [Jenkins: Role-based Authorization Strategy](https://plugins.jenkins.io/role-strategy/)
