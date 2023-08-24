# About
This repository is for experimental use unleash as feature toggle in Go. We uses a package provided by Unleash as an open source feature toggle.

# How to use
First, we need to start the Unleash Server. The easiest way is by running the following script in your terminal.
```bash
git clone git@github.com:Unleash/unleash-docker.git
cd unleash-docker
docker compose up 
```
Open your browser and go to localhost:4242. You need to log into Unleash dashbord using this credential:
```
username : admin
password : unleash4all
```
Later, you will use this dashboard to create and configure your feature toggle.

```
The simple working-example as a playground to try feature toggle is provided here. Assuming that you havel already installed local Unleash server, you just need to run
```
git clone git@github.com:rachadiannovansyah/unleash-go-feature-toggle.git
cd feature-toggle
go get github.com/Unleash/unleash-client-go/v3
go run main.go
```
Open your browser and hit `localhost:8080/`. With feature toggle `toggle` is enabled, you will get `feature is enabled` printed on your browser. Otherwise, it will show `feature is disabled`.
# References
- Dive deeper to the available [documentations](https://docs.getunleash.io/) and [articles](https://www.getunleash.io/blog/feature-toggle-best-practices) by Unleash to uncover tips and best practices to implement feature toggle in your project.
- Academic research [paper](https://arxiv.org/pdf/1907.06157.pdf) by Mahdavi-Hezaveh, R. et.al. about the use of feature toggle in industry.