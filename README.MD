# Curzor
Curzor is a basic vulnerable web application written in Golang, part of the Shadowspace cyber range. The application itself is intended for educational purposes only. It can be used as an attack entrypoint in case it's exposed to external networks and not properly encapsulated. Use with caution!

## Dependencies
You need the following dependencies in order to run the application:
 * Golang

## Usage
### On-system 
* Run `go build .` inside the directory of the project.
* Execute the application with `./curzor`.

or

* Directly run the application with `go run .` inside the directory of the project.

The server runs on port `8080`.

### Docker
* Execute `docker build . -t shadowspace/curzor` in order to build the container;
* Run the container in daemon mode with `sudo docker run -p 80:8080 shadowspace/curzor`.

Docker will bind the container's port `8080` to the host's port `80` on `0.0.0.0`.

## Vulnerabilities
The following vulnerabilities are present in Curzor:
* Stored Cross-Site Scripting (XSS);
* Remote Code Execution (RCE);
* Server-Side Request Forgery (SSRF);
* Template Injection;
* Use of Hardcoded Credentials;
* Unrestricted File Upload;
* Identification and Authentication Failure;
* Broken Access Control.

The vulnerabilities in the application have been encountered in the wild before (in the same or similar form).

## TODO
* Implement mechanism that allows dynamic configuration of the web application.

# Notes
* The application and part of the range are currently utilized in the Active Information Security course in Sofia University, Bulgaria.

