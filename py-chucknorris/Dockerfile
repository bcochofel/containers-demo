FROM python:3.6.1-alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /project
ADD . /project
RUN pip install -r requirements.txt
CMD ["python","main.py"]
