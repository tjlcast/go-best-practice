docker build . -t jialtang/greeter:v1.0.0

kind load docker-image jialtang/greeter:v1.0.0 --name my-cluster
