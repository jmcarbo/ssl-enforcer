TAG=`cat version`
container:
	env GOOS=linux GOARCH=amd64 go build -o bin/enforcer main.go
	docker build -t goincremental/enforcer .
	docker tag -f goincremental/enforcer eu.gcr.io/gi-harbour/enforcer
	docker tag -f goincremental/enforcer eu.gcr.io/gi-harbour/enforcer:$(TAG)

publish:
	git tag v$(TAG)
	git push origin --tags
	gcloud docker push eu.gcr.io/gi-harbour/enforcer

rc:
	kubectl create -f rc.yml

svc:
	kubectl create -f svc.yml

edge-svc:
	kubectl create -f edge-svc.yml

edge-rc:
	kubectl create -f edge-rc.yml
